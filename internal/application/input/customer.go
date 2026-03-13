package input

type CreateCustomerInput struct {
	Name        string `validate:"required,min=3,regex=^[a-zA-Z ]+$"`
	NIK         string `validate:"required,len=13,regex=^\\d+$"`
	PhoneNumber string `validate:"required,regex=^(62|0)8\\d{8,10}$"`
}

type UpdateCustomerInput struct {
	Name        string `validate:"omitempty,min=3,regex=^[a-zA-Z ]+$"`
	NIK         string `validate:"omitempty,len=13,regex=^\\d+$"`
	PhoneNumber string `validate:"omitempty,regex=^(62|0)8\\d{8,10}$"`
}
