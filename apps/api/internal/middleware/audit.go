package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/audit"
)

type AuditMiddleware struct {
	logger *audit.Logger
}

func NewAuditMiddleware(logger *audit.Logger) *AuditMiddleware {
	return &AuditMiddleware{logger: logger}
}

func (m *AuditMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipAudit(c.Request.URL.Path) {
			c.Next()
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
			m.logAction(c, bodyBytes)
		}
	}
}

func (m *AuditMiddleware) logAction(c *gin.Context, bodyBytes []byte) {
	ctx := c.Request.Context()

	action, entityType, entityID := extractActionAndEntity(c)
	if action == "" {
		return
	}

	var newVal any
	if len(bodyBytes) > 0 {
		json.Unmarshal(bodyBytes, &newVal)
		newVal = sanitizeSensitiveData(newVal)
	}

	var userID *uuid.UUID
	if uid, exists := c.Get("user_id"); exists {
		if id, ok := uid.(uuid.UUID); ok {
			userID = &id
		}
	}

	if entityID == "" && userID != nil {
		entityID = userID.String()
	}

	m.logger.Log(ctx, userID, action, entityType, entityID, nil, newVal, c.ClientIP(), c.Request.UserAgent())
}

var sensitiveFields = []string{"password", "token", "secret", "api_key", "apikey", "access_token", "refresh_token", "credential"}

func sanitizeSensitiveData(data any) any {
	if data == nil {
		return nil
	}

	switch v := data.(type) {
	case map[string]any:
		result := make(map[string]any)
		for key, val := range v {
			if isSensitiveField(key) {
				result[key] = "[REDACTED]"
			} else {
				result[key] = sanitizeSensitiveData(val)
			}
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, val := range v {
			result[i] = sanitizeSensitiveData(val)
		}
		return result
	default:
		return v
	}
}

func isSensitiveField(field string) bool {
	lowerField := strings.ToLower(field)
	for _, sensitive := range sensitiveFields {
		if lowerField == sensitive || strings.Contains(lowerField, sensitive) {
			return true
		}
	}
	return false
}

func extractActionAndEntity(c *gin.Context) (audit.Action, string, string) {
	method := c.Request.Method
	path := c.Request.URL.Path

	patterns := []struct {
		pattern    *regexp.Regexp
		action     audit.Action
		entityType string
		entityID   int
	}{
		{regexp.MustCompile(`^/api/v1/workspaces/([^/]+)/members/([^/]+)$`), audit.ActionDelete, "workspace_member", 2},
		{regexp.MustCompile(`^/api/v1/workspaces/([^/]+)/members$`), audit.ActionCreate, "workspace_member", 0},
		{regexp.MustCompile(`^/api/v1/workspaces/([^/]+)$`), audit.ActionUpdate, "workspace", 1},
		{regexp.MustCompile(`^/api/v1/workspaces/([^/]+)$`), audit.ActionDelete, "workspace", 1},
		{regexp.MustCompile(`^/api/v1/workspaces$`), audit.ActionCreate, "workspace", 0},
		{regexp.MustCompile(`^/api/v1/users/profile$`), audit.ActionUpdate, "user", 0},
		{regexp.MustCompile(`^/api/v1/auth/login$`), audit.ActionLogin, "user", 0},
		{regexp.MustCompile(`^/api/v1/auth/logout$`), audit.ActionLogout, "user", 0},
	}

	for _, p := range patterns {
		if matches := p.pattern.FindStringSubmatch(path); matches != nil {
			if p.action == audit.ActionUpdate && method != "PUT" && method != "PATCH" {
				continue
			}
			if p.action == audit.ActionDelete && method != "DELETE" {
				continue
			}
			if p.action == audit.ActionCreate && method != "POST" {
				continue
			}

			entityID := ""
			if p.entityID > 0 && p.entityID < len(matches) {
				entityID = matches[p.entityID]
			}

			return p.action, p.entityType, entityID
		}
	}

	return "", "", ""
}

func shouldSkipAudit(path string) bool {
	skipped := []string{
		"/health",
		"/metrics",
		"/swagger",
	}

	for _, prefix := range skipped {
		if len(path) >= len(prefix) && path[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}
