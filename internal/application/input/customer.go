package input

type CreateCustomerInput struct {
	Name        string `json:"name" validate:"required,min=3,carname"`
	NIK         string `json:"nik" validate:"required,len=13,nik"`
	PhoneNumber string `json:"phone_number" validate:"required,phone"`
}

type UpdateCustomerInput struct {
	Name        string `json:"name" validate:"omitempty,min=3,carname"`
	NIK         string `json:"nik" validate:"omitempty,len=13,nik"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,phone"`
}
