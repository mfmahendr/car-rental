package input

type PaginationInput struct {
    Page int `json:"page" validate:"min=1"`
    Size int `json:"size" validate:"min=1,max=100"`
}

func (r *PaginationInput) GetOffset() int {
	return (r.Page - 1) * r.Size
}