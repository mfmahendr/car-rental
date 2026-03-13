package result

import "github.com/mfmahendr/car-rental/internal/application/input"

type PageResult[T any] struct {
	Items      []T  `json:"items"`
	TotalItems int  `json:"total_items"`
	TotalPages int  `json:"total_pages"`
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

func NewPageResult[T any](items []T, totalItems int, in input.PaginationInput) PageResult[T] {
	totalPages := (totalItems + in.Size - 1) / in.Size

	return PageResult[T]{
		Items:      items,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Page:       in.Page,
		Size:       in.Size,
		HasNext:    in.Page < totalPages,
		HasPrev:    in.Page > 1,
	}
}
