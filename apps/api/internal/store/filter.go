package store

import (
	"fmt"
	"strings"
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type FilterParams struct {
	Limit  int32 `json:"limit" query:"limit"`
	Offset int32 `json:"offset" query:"offset"`

	SortBy string `json:"sort_by" query:"sort_by"`
	Order  string `json:"order" query:"order"`

	Search string `json:"search" query:"search"`

	Filters map[string]string `json:"filters" query:"filters"`
}

func DefaultFilter() FilterParams {
	return FilterParams{
		Limit:   20,
		Offset:  0,
		SortBy:  "created_at",
		Order:   "desc",
		Filters: make(map[string]string),
	}
}

func (f *FilterParams) Normalize() {
	if f.Limit <= 0 {
		f.Limit = 20
	}
	if f.Limit > 100 {
		f.Limit = 100
	}
	if f.Offset < 0 {
		f.Offset = 0
	}

	order := strings.ToLower(f.Order)
	if order != string(SortOrderAsc) && order != string(SortOrderDesc) {
		f.Order = "desc"
	} else {
		f.Order = order
	}

	f.SortBy = sanitizeSortColumn(f.SortBy)

	if f.Filters == nil {
		f.Filters = make(map[string]string)
	}
}

func (f *FilterParams) GetSortClause(defaultSortBy string) string {
	sortBy := f.SortBy
	if sortBy == "" {
		sortBy = defaultSortBy
	}
	return fmt.Sprintf("ORDER BY %s %s", sortBy, strings.ToUpper(f.Order))
}

func (f *FilterParams) GetPaginationClause() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", f.Limit, f.Offset)
}

func sanitizeSortColumn(column string) string {
	validColumns := map[string]bool{
		"id":         true,
		"email":      true,
		"name":       true,
		"created_at": true,
		"updated_at": true,

		"workspace_id": true,
		"user_id":      true,
		"role":         true,
		"joined_at":    true,

		"expires_at": true,
	}

	column = strings.ToLower(strings.TrimSpace(column))
	if validColumns[column] {
		return column
	}
	return "created_at"
}

func (f *FilterParams) HasSearch() bool {
	return strings.TrimSpace(f.Search) != ""
}

func (f *FilterParams) GetSearchPattern() string {
	return "%" + strings.ToLower(strings.TrimSpace(f.Search)) + "%"
}

// PaginatedResponse is a generic wrapper for paginated list responses
// @Description Generic paginated response containing data and pagination metadata
type PaginatedResponse[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
	Filters    FilterInfo     `json:"filters"`
}

// PaginationInfo contains pagination metadata for paginated responses
// @Description Pagination metadata including current page info and totals
type PaginationInfo struct {
	Limit         int32 `json:"limit" example:"20"`
	Offset        int32 `json:"offset" example:"0"`
	TotalRecords  int64 `json:"total_records" example:"100"`
	LastPage      int32 `json:"last_page" example:"5"`
	CurrentPage   int32 `json:"current_page" example:"1"`
	HasNextPage   bool  `json:"has_next_page" example:"true"`
	HasPrevPage   bool  `json:"has_prev_page" example:"false"`
}

// FilterInfo contains the active filter parameters used for the query
// @Description Active filter and sorting parameters
type FilterInfo struct {
	SortBy string `json:"sort_by" example:"created_at"`
	Order  string `json:"order" example:"desc"`
	Search string `json:"search,omitempty" example:"john"`
}

func BuildFilterResponse[T any](data []T, total int64, filters FilterParams) *PaginatedResponse[T] {
	// Calculate pagination metadata
	limit := filters.Limit
	if limit <= 0 {
		limit = 20
	}

	currentPage := int32(filters.Offset/limit) + 1
	lastPage := int32((total + int64(limit) - 1) / int64(limit))
	if lastPage < 1 {
		lastPage = 1
	}
	hasNextPage := filters.Offset+limit < int32(total)
	hasPrevPage := filters.Offset > 0

	return &PaginatedResponse[T]{
		Data: data,
		Pagination: PaginationInfo{
			Limit:        limit,
			Offset:       filters.Offset,
			TotalRecords: total,
			LastPage:     lastPage,
			CurrentPage:  currentPage,
			HasNextPage:  hasNextPage,
			HasPrevPage:  hasPrevPage,
		},
		Filters: FilterInfo{
			SortBy: filters.SortBy,
			Order:  filters.Order,
			Search: filters.Search,
		},
	}
}

// PaginatedSessionsResponse represents a paginated list of user sessions
// @Description Paginated response containing user session data
// swagger:model PaginatedSessionsResponse
type PaginatedSessionsResponse struct {
	// List of user sessions
	Data []Session `json:"data"`
	// Pagination metadata
	Pagination PaginationInfo `json:"pagination"`
	// Applied filters
	Filters FilterInfo `json:"filters"`
}

// PaginatedWorkspacesResponse represents a paginated list of workspaces
// @Description Paginated response containing workspace data with user roles
// swagger:model PaginatedWorkspacesResponse
type PaginatedWorkspacesResponse struct {
	// List of workspaces with user roles
	Data []WorkspaceWithRole `json:"data"`
	// Pagination metadata
	Pagination PaginationInfo `json:"pagination"`
	// Applied filters
	Filters FilterInfo `json:"filters"`
}

// PaginatedMembersResponse represents a paginated list of workspace members
// @Description Paginated response containing workspace member data
// swagger:model PaginatedMembersResponse
type PaginatedMembersResponse struct {
	// List of workspace members
	Data []WorkspaceMember `json:"data"`
	// Pagination metadata
	Pagination PaginationInfo `json:"pagination"`
	// Applied filters
	Filters FilterInfo `json:"filters"`
}
