package utils

import "strconv"

type Pagination struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

type PaginationResponse struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

func NewPagination(limitStr, pageStr string, defaultLimit, defaultPage int) Pagination {
	limit := defaultLimit
	page := defaultPage

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	offset := (page - 1) * limit

	return Pagination{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}
}
