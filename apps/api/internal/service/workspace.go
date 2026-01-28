package service

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/store"
	pkgstorage "github.com/lukabrkovic/artemis/pkg/storage"
)

var (
	ErrWorkspaceNotFound = errors.New("workspace not found")
	ErrForbidden         = errors.New("forbidden")
	ErrInvalidRole       = errors.New("invalid role")
)

type CreateWorkspaceInput struct {
	Name      string
	AvatarURL *string
}

type Workspace interface {
	CreateWorkspace(ctx context.Context, userID uuid.UUID, input CreateWorkspaceInput) (*store.Workspace, error)
	GetWorkspace(ctx context.Context, userID, workspaceID uuid.UUID) (*store.Workspace, error)
	GetMyWorkspaces(ctx context.Context, userID uuid.UUID, filters store.FilterParams) (*store.PaginatedResponse[store.WorkspaceWithRole], error)
	UpdateWorkspace(ctx context.Context, userID, workspaceID uuid.UUID, name string) (*store.Workspace, error)
	DeleteWorkspace(ctx context.Context, userID, workspaceID uuid.UUID) error
	AddMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID, role string) (*store.WorkspaceMember, error)
	AddMemberByEmail(ctx context.Context, requesterID, workspaceID uuid.UUID, email, role string) (*store.WorkspaceMember, error)
	RemoveMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID) error
	GetMembers(ctx context.Context, userID, workspaceID uuid.UUID, filters store.FilterParams) (*store.PaginatedResponse[store.WorkspaceMember], error)
	UploadAvatar(ctx context.Context, userID, workspaceID uuid.UUID, reader io.Reader, size int64, contentType string) (string, error)
}

type WorkspaceService struct {
	store   *store.Store
	storage pkgstorage.Provider
}

func NewWorkspaceService(store *store.Store, storage pkgstorage.Provider) *WorkspaceService {
	return &WorkspaceService{store: store, storage: storage}
}

var _ Workspace = (*WorkspaceService)(nil)

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, userID uuid.UUID, input CreateWorkspaceInput) (*store.Workspace, error) {
	var workspace *store.Workspace
	err := s.store.ExecTx(ctx, func(tx *store.Store) error {
		var err error
		workspace, err = tx.Workspaces.CreateWorkspace(ctx, store.CreateWorkspaceParams{
			Name:      input.Name,
			AvatarURL: input.AvatarURL,
		})
		if err != nil {
			return err
		}

		return tx.Workspaces.AddWorkspaceMember(ctx, workspace.ID, userID, "owner")
	})

	return workspace, err
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, userID, workspaceID uuid.UUID) (*store.Workspace, error) {
	_, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	return s.store.Workspaces.GetWorkspaceByID(ctx, workspaceID)
}

func (s *WorkspaceService) GetMyWorkspaces(ctx context.Context, userID uuid.UUID, filters store.FilterParams) (*store.PaginatedResponse[store.WorkspaceWithRole], error) {
	workspaces, total, err := s.store.Workspaces.GetUserWorkspaces(ctx, userID, filters)
	if err != nil {
		return nil, err
	}

	return &store.PaginatedResponse[store.WorkspaceWithRole]{
		Data: workspaces,
		Pagination: store.PaginationInfo{
			Limit:  filters.Limit,
			Offset: filters.Offset,
			Total:  total,
		},
		Filters: store.FilterInfo{
			SortBy: filters.SortBy,
			Order:  filters.Order,
			Search: filters.Search,
		},
	}, nil
}

func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, userID, workspaceID uuid.UUID, name string) (*store.Workspace, error) {
	role, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	if role != "owner" && role != "admin" {
		return nil, ErrForbidden
	}

	return s.store.Workspaces.UpdateWorkspace(ctx, workspaceID, name)
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, userID, workspaceID uuid.UUID) error {
	role, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return ErrForbidden
		}
		return err
	}

	if role != "owner" {
		return ErrForbidden
	}

	return s.store.Workspaces.DeleteWorkspace(ctx, workspaceID)
}

func (s *WorkspaceService) AddMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID, role string) (*store.WorkspaceMember, error) {
	requesterRole, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, requesterID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	if requesterRole != "owner" && requesterRole != "admin" {
		return nil, ErrForbidden
	}

	if role != "admin" && role != "member" {
		return nil, ErrInvalidRole
	}

	err = s.store.Workspaces.AddWorkspaceMember(ctx, workspaceID, targetUserID, role)
	if err != nil {
		return nil, err
	}

	user, err := s.store.Users.GetUserByID(ctx, targetUserID)
	if err != nil {
		return nil, err
	}

	member := &store.WorkspaceMember{
		WorkspaceID: workspaceID,
		UserID:      targetUserID,
		Role:        role,
		JoinedAt:    user.CreatedAt,
		Name:        user.Name,
		Email:       *user.Email,
		AvatarURL:   user.AvatarURL,
	}

	return member, nil
}

func (s *WorkspaceService) AddMemberByEmail(ctx context.Context, requesterID, workspaceID uuid.UUID, email, role string) (*store.WorkspaceMember, error) {
	user, err := s.store.Users.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			return nil, store.ErrUserNotFound
		}
		return nil, err
	}

	return s.AddMember(ctx, requesterID, workspaceID, user.ID, role)
}

func (s *WorkspaceService) RemoveMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID) error {
	requesterRole, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, requesterID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return ErrForbidden
		}
		return err
	}

	if requesterID == targetUserID {
		if requesterRole == "owner" {
			return errors.New("owner cannot leave workspace, delete it instead")
		}
	} else {
		if requesterRole == "member" {
			return ErrForbidden
		}

		targetRole, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, targetUserID)
		if err != nil {
			return err
		}

		if requesterRole == "admin" && (targetRole == "admin" || targetRole == "owner") {
			return ErrForbidden
		}
	}

	return s.store.Workspaces.RemoveWorkspaceMember(ctx, workspaceID, targetUserID)
}

func (s *WorkspaceService) GetMembers(ctx context.Context, userID, workspaceID uuid.UUID, filters store.FilterParams) (*store.PaginatedResponse[store.WorkspaceMember], error) {
	_, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	members, total, err := s.store.Workspaces.GetWorkspaceMembers(ctx, workspaceID, filters)
	if err != nil {
		return nil, err
	}

	return &store.PaginatedResponse[store.WorkspaceMember]{
		Data: members,
		Pagination: store.PaginationInfo{
			Limit:  filters.Limit,
			Offset: filters.Offset,
			Total:  total,
		},
		Filters: store.FilterInfo{
			SortBy: filters.SortBy,
			Order:  filters.Order,
			Search: filters.Search,
		},
	}, nil
}

func (s *WorkspaceService) UploadAvatar(ctx context.Context, userID, workspaceID uuid.UUID, reader io.Reader, size int64, contentType string) (string, error) {
	if err := store.CheckContext(ctx); err != nil {
		return "", err
	}

	role, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return "", ErrForbidden
		}
		return "", err
	}

	if role != "owner" && role != "admin" {
		return "", ErrForbidden
	}

	// Upload to storage first
	avatarURL, err := s.storage.UploadAvatar(ctx, workspaceID.String(), reader, size, contentType)
	if err != nil {
		return "", err
	}

	// Update database - if this fails, we need to clean up the uploaded file
	_, err = s.store.Workspaces.UpdateWorkspaceAvatar(ctx, workspaceID, avatarURL)
	if err != nil {
		// Best effort cleanup of orphaned file
		_ = s.storage.DeleteAvatar(ctx, workspaceID.String())
		return "", err
	}

	return avatarURL, nil
}
