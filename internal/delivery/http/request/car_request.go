package request

type UpdateCarStock struct {
	Stock int `json:"stock" form:"stock"`
}

type UpdateCar struct {
	Name      string `json:"name" form:"name"`
	Stock     uint   `json:"stock" form:"stock"`
	DailyRent int64  `json:"daily_rent" form:"daily_rent"`
}