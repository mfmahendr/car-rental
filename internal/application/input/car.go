package input

type CreateCarInput struct {
	Name      string `validate:"required,min=3,regex=^[a-zA-Z ]+$"`
	Stock     uint   `validate:"required,min=1"`
	DailyRent int64  `validate:"required,min=0"`
}

type UpdateCarInput struct {
	Name      string `validate:"omitempty,min=3,regex=^[a-zA-Z ]+$"`
	Stock     uint   `validate:"omitempty,min=1"`
	DailyRent int64  `validate:"omitempty,min=0"`
}
