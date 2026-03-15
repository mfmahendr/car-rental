package input

type CreateCarInput struct {
	Name      string `json:"name" validate:"required,min=3,carname"`
	Stock     uint   `json:"stock" validate:"required,min=1"`
	DailyRent int64  `json:"daily_rent" validate:"required,min=0"`
}

type UpdateCarInput struct {
	Name      string `json:"name" validate:"omitempty,min=3,carname"`
	Stock     uint   `json:"stock" validate:"omitempty,min=1"`
	DailyRent int64  `json:"daily_rent" validate:"omitempty,min=0"`
}
