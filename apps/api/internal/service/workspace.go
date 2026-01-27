package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/store"
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
	GetMyWorkspaces(ctx context.Context, userID uuid.UUID) ([]store.WorkspaceWithRole, error)
	UpdateWorkspace(ctx context.Context, userID, workspaceID uuid.UUID, name string) (*store.Workspace, error)
	DeleteWorkspace(ctx context.Context, userID, workspaceID uuid.UUID) error
	AddMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID, role string) error
	RemoveMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID) error
	GetMembers(ctx context.Context, userID, workspaceID uuid.UUID) ([]store.WorkspaceMember, error)
}

type WorkspaceService struct {
	store *store.Store
}

func NewWorkspaceService(store *store.Store) *WorkspaceService {
	return &WorkspaceService{store: store}
}

var _ Workspace = (*WorkspaceService)(nil)

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, userID uuid.UUID, input CreateWorkspaceInput) (*store.Workspace, error) {
	// Transaction: Create Workspace -> Add Creator as Owner
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
	// Check membership first
	_, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	return s.store.Workspaces.GetWorkspaceByID(ctx, workspaceID)
}

func (s *WorkspaceService) GetMyWorkspaces(ctx context.Context, userID uuid.UUID) ([]store.WorkspaceWithRole, error) {
	return s.store.Workspaces.GetUserWorkspaces(ctx, userID)
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

func (s *WorkspaceService) AddMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID, role string) error {
	requesterRole, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, requesterID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return ErrForbidden
		}
		return err
	}

	if requesterRole != "owner" && requesterRole != "admin" {
		return ErrForbidden
	}

	if role != "admin" && role != "member" {
		return ErrInvalidRole
	}

	return s.store.Workspaces.AddWorkspaceMember(ctx, workspaceID, targetUserID, role)
}

func (s *WorkspaceService) RemoveMember(ctx context.Context, requesterID, workspaceID, targetUserID uuid.UUID) error {
	requesterRole, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, requesterID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return ErrForbidden
		}
		return err
	}

	// Owner can remove anyone. Admin can remove members. Users can leave (remove themselves).
	if requesterID == targetUserID {
		// Consenting to leave
		// Prevent owner from leaving without transferring ownership? For simplicity allow strictly if they delete workspace OR separate logic.
		// Usually owner cannot leave, must verify.
		if requesterRole == "owner" {
			// Owner cannot leave, must delete workspace or transfer.
			// For now, let's block owner leaving.
			return errors.New("owner cannot leave workspace, delete it instead")
		}
	} else {
		// Removing someone else
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

func (s *WorkspaceService) GetMembers(ctx context.Context, userID, workspaceID uuid.UUID) ([]store.WorkspaceMember, error) {
	_, err := s.store.Workspaces.GetWorkspaceMemberRole(ctx, workspaceID, userID)
	if err != nil {
		if errors.Is(err, store.ErrNotMember) {
			return nil, ErrForbidden
		}
		return nil, err
	}

	return s.store.Workspaces.GetWorkspaceMembers(ctx, workspaceID)
}
