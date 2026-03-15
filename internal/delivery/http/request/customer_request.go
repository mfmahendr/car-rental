package request

type UpdateCustomer struct {
	Name        string `json:"name" form:"name"`
	NIK         string `json:"nik" form:"nik"`
	PhoneNumber string `json:"phone_number" form:"phone"`
}