package result

type CarResult struct {
	CarID     int    `json:"car_id"`
	Name      string `json:"name"`
	Stock     uint   `json:"stock"`
	DailyRent int64  `json:"daily_rent"`
}
