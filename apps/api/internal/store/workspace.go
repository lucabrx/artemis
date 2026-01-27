package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrWorkspaceNotFound = errors.New("workspace not found")
	ErrAlreadyMember     = errors.New("user is already a member")
	ErrNotMember         = errors.New("user is not a member")
)

type Workspace struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	AvatarURL *string   `json:"avatar_url" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceID uuid.UUID `json:"workspace_id" db:"workspace_id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Role        string    `json:"role" db:"role"`
	JoinedAt    time.Time `json:"joined_at" db:"joined_at"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	AvatarURL   *string   `json:"avatar_url" db:"avatar_url"`
}

type WorkspaceWithRole struct {
	Workspace
	Role string `json:"role" db:"role"`
}

type CreateWorkspaceParams struct {
	Name      string
	AvatarURL *string
}

type WorkspaceRepository interface {
	CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (*Workspace, error)
	GetWorkspaceByID(ctx context.Context, id uuid.UUID) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, id uuid.UUID, name string) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, id uuid.UUID) error
	AddWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID, role string) error
	GetWorkspaceMembers(ctx context.Context, workspaceID uuid.UUID) ([]WorkspaceMember, error)
	RemoveWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID) error
	GetUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]WorkspaceWithRole, error)
	GetWorkspaceMemberRole(ctx context.Context, workspaceID, userID uuid.UUID) (string, error)
	UpdateWorkspaceAvatar(ctx context.Context, id uuid.UUID, avatarURL string) (*Workspace, error)
}

type workspaceRepository struct {
	db DBTX
}

func NewWorkspaceRepository(db DBTX) WorkspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (*Workspace, error) {
	workspace := &Workspace{}
	query := `
		INSERT INTO workspaces (name, avatar_url)
		VALUES ($1, $2)
		RETURNING id, name, avatar_url, created_at, updated_at
	`
	err := r.db.GetContext(ctx, workspace, query, arg.Name, arg.AvatarURL)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (r *workspaceRepository) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	var workspace Workspace
	query := `SELECT * FROM workspaces WHERE id = $1`
	err := r.db.GetContext(ctx, &workspace, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWorkspaceNotFound
		}
		return nil, err
	}
	return &workspace, nil
}

func (r *workspaceRepository) UpdateWorkspace(ctx context.Context, id uuid.UUID, name string) (*Workspace, error) {
	var workspace Workspace
	query := `
		UPDATE workspaces
		SET name = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, name, avatar_url, created_at, updated_at
	`
	err := r.db.GetContext(ctx, &workspace, query, name, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWorkspaceNotFound
		}
		return nil, err
	}
	return &workspace, nil
}

func (r *workspaceRepository) UpdateWorkspaceAvatar(ctx context.Context, id uuid.UUID, avatarURL string) (*Workspace, error) {
	var workspace Workspace
	query := `
		UPDATE workspaces
		SET avatar_url = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, name, avatar_url, created_at, updated_at
	`
	err := r.db.GetContext(ctx, &workspace, query, avatarURL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWorkspaceNotFound
		}
		return nil, err
	}
	return &workspace, nil
}

func (r *workspaceRepository) DeleteWorkspace(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM workspaces WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *workspaceRepository) AddWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID, role string) error {
	query := `
		INSERT INTO workspace_members (workspace_id, user_id, role)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, query, workspaceID, userID, role)
	return err
}

func (r *workspaceRepository) GetWorkspaceMembers(ctx context.Context, workspaceID uuid.UUID) ([]WorkspaceMember, error) {
	var members []WorkspaceMember
	query := `
		SELECT wm.*, u.name, u.email, u.avatar_url
		FROM workspace_members wm
		JOIN users u ON wm.user_id = u.id
		WHERE wm.workspace_id = $1
		ORDER BY wm.joined_at ASC
	`
	err := r.db.SelectContext(ctx, &members, query, workspaceID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *workspaceRepository) RemoveWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID) error {
	query := `DELETE FROM workspace_members WHERE workspace_id = $1 AND user_id = $2`
	result, err := r.db.ExecContext(ctx, query, workspaceID, userID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotMember
	}
	return nil
}

func (r *workspaceRepository) GetUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]WorkspaceWithRole, error) {
	var workspaces []WorkspaceWithRole
	query := `
		SELECT w.*, wm.role
		FROM workspaces w
		JOIN workspace_members wm ON w.id = wm.workspace_id
		WHERE wm.user_id = $1
		ORDER BY w.created_at DESC
	`
	err := r.db.SelectContext(ctx, &workspaces, query, userID)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (r *workspaceRepository) GetWorkspaceMemberRole(ctx context.Context, workspaceID, userID uuid.UUID) (string, error) {
	var role string
	query := `SELECT role FROM workspace_members WHERE workspace_id = $1 AND user_id = $2`
	err := r.db.GetContext(ctx, &role, query, workspaceID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotMember
		}
		return "", err
	}
	return role, nil
}
