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
	// Pagination
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

type PaginatedResponse[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
	Filters    FilterInfo     `json:"filters"`
}

type PaginationInfo struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Total  int64 `json:"total"`
}

type FilterInfo struct {
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
	Search string `json:"search,omitempty"`
}

func BuildFilterResponse[T any](data []T, total int64, filters FilterParams) PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Data: data,
		Pagination: PaginationInfo{
			Limit:  filters.Limit,
			Offset: filters.Offset,
			Total:  total,
		},
		Filters: FilterInfo{
			SortBy: filters.SortBy,
			Order:  filters.Order,
			Search: filters.Search,
		},
	}
}

// fu swagger
type PaginatedSessionsResponse struct {
	Data       []Session      `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
	Filters    FilterInfo     `json:"filters"`
}

type PaginatedWorkspacesResponse struct {
	Data       []WorkspaceWithRole `json:"data"`
	Pagination PaginationInfo      `json:"pagination"`
	Filters    FilterInfo          `json:"filters"`
}

type PaginatedMembersResponse struct {
	Data       []WorkspaceMember `json:"data"`
	Pagination PaginationInfo    `json:"pagination"`
	Filters    FilterInfo        `json:"filters"`
}
