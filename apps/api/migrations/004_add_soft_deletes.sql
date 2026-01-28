-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE sessions ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE workspaces ADD COLUMN deleted_at TIMESTAMPTZ;
ALTER TABLE workspace_members ADD COLUMN deleted_at TIMESTAMPTZ;

CREATE INDEX idx_users_deleted_at ON users(deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_sessions_deleted_at ON sessions(deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_workspaces_deleted_at ON workspaces(deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_workspace_members_deleted_at ON workspace_members(deleted_at) WHERE deleted_at IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_sessions_deleted_at;
DROP INDEX IF EXISTS idx_workspaces_deleted_at;
DROP INDEX IF EXISTS idx_workspace_members_deleted_at;

ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE sessions DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE workspaces DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE workspace_members DROP COLUMN IF EXISTS deleted_at;
-- +goose StatementEnd
