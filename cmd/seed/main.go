package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mfmahendr/car-rental/internal/config"
	"github.com/mfmahendr/car-rental/internal/domain/entities"
	"github.com/mfmahendr/car-rental/internal/infra/database/postgres"
)

func main() {
	cfg := config.Load()

	pool, err := postgres.NewPool(cfg.Db)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `TRUNCATE TABLE bookings,cars,customers RESTART IDENTITY CASCADE`)
	if err != nil {
		log.Printf("Failed to truncate tables: %v", err)
		os.Exit(1)
	}

	for _, c := range cars {
		_, err := pool.Exec(context.Background(),
			"INSERT INTO cars (name, stock, daily_rent) VALUES ($1, $2, $3)",
			c.Name, c.Stock, c.DailyRent)
		if err != nil {
			log.Printf("Failed to add data: %v", err)
		}
	}
	for _, c := range customers {
		_, err := pool.Exec(context.Background(),
			"INSERT INTO customers (name, nik, phone_number) VALUES ($1, $2, $3)",
			c.Name, c.NIK, c.PhoneNumber)
		if err != nil {
			log.Printf("Failed to add data: %v", err)
		}
	}
	for _, c := range bookings {
		_, err := pool.Exec(context.Background(),
			"INSERT INTO bookings (customer_id, car_id, start_rent, end_rent, total_cost, finished) VALUES ($1, $2, $3, $4, $5, $6)",
			c.CustomerID, c.CarID, c.StartRent, c.EndRent, c.TotalCost, c.Finished)
		if err != nil {
			log.Printf("Failed to add data: %v", err)
		}
	}
	
	fmt.Println("Seeding done!")
}

var (
	cars = []entities.Car{
		{Name: "Toyota Camry", Stock: 2, DailyRent: 500000},
		{Name: "Toyota Avalon", Stock: 2, DailyRent: 500000},
		{Name: "Toyota Yaris", Stock: 2, DailyRent: 400000},
		{Name: "Toyota Agya", Stock: 2, DailyRent: 400000},
		{Name: "Toyota Fortuner", Stock: 1, DailyRent: 700000},
		{Name: "Toyota Rush", Stock: 1, DailyRent: 600000},
		{Name: "Toyota Hiace", Stock: 1, DailyRent: 1000000},
		{Name: "Honda Brio", Stock: 3, DailyRent: 500000},
		{Name: "Honda Civic", Stock: 1, DailyRent: 500000},
		{Name: "Honda Jazz", Stock: 1, DailyRent: 500000},
		{Name: "Honda Mobilio", Stock: 2, DailyRent: 700000},
		{Name: "Honda Amaze", Stock: 1, DailyRent: 700000},
		{Name: "Honda Breeze", Stock: 2, DailyRent: 700000},
		{Name: "Mitsubishi Pajero Sport", Stock: 5, DailyRent: 800000},
		{Name: "Mitsubishi Mirage", Stock: 3, DailyRent: 600000},
	}

	customers = []entities.Customer{
		{Name: "Wawan Hermawan", NIK: "3372093912739", PhoneNumber: "081237123682"},
		{Name: "Philip Walker", NIK: "3372093912785", PhoneNumber: "081237123683"},
		{Name: "Hugo Fleming", NIK: "3372093912800", PhoneNumber: "081237123684"},
		{Name: "Maximillian Mendez", NIK: "3372093912848", PhoneNumber: "081237123685"},
		{Name: "Felix Dixon", NIK: "3372093912851", PhoneNumber: "081237123686"},
		{Name: "Nicholas Riddle", NIK: "3372093912929", PhoneNumber: "081237123687"},
		{Name: "Stephen Wheeler", NIK: "3372093912976", PhoneNumber: "081237123688"},
		{Name: "Roy Brennan", NIK: "3372093913022", PhoneNumber: "081237123689"},
		{Name: "Eliza Le", NIK: "3372093913106", PhoneNumber: "081237123690"},
		{Name: "Jesse Taylor", NIK: "3372093913126", PhoneNumber: "081237123691"},
		{Name: "Damien Kaufman", NIK: "3372093913202", PhoneNumber: "081237123692"},
		{Name: "Ayesha Richardson", NIK: "3372093913257", PhoneNumber: "081237123693"},
		{Name: "Margaret Stokes", NIK: "3372093913262", PhoneNumber: "081237123694"},
		{Name: "Sara Livingston", NIK: "3372093913268", PhoneNumber: "081237123695"},
		{Name: "Callie Townsend", NIK: "3372093913281", PhoneNumber: "081237123696"},
		{Name: "Lilly Fischer", NIK: "3372093913325", PhoneNumber: "081237123697"},
		{Name: "Theresa Barton", NIK: "3372093913335", PhoneNumber: "081237123698"},
		{Name: "Mia Curtis", NIK: "3372093913343", PhoneNumber: "081237123699"},
		{Name: "Flora Barlow", NIK: "3372093913400", PhoneNumber: "081237123700"},
		{Name: "Vanessa Patton", NIK: "3372093913434", PhoneNumber: "081237123701"},
	}

	bookings = []entities.Booking{
		{CustomerID: 3, CarID: 2, StartRent: parseDate("1/1/2021"), EndRent: parseDate("1/2/2021"), TotalCost: 1_000_000, Finished: true},
		{CustomerID: 11, CarID: 2, StartRent: parseDate("1/10/2021"), EndRent: parseDate("1/11/2021"), TotalCost: 1_000_000, Finished: true},
		{CustomerID: 7, CarID: 1, StartRent: parseDate("1/12/2021"), EndRent: parseDate("1/14/2021"), TotalCost: 1_500_000, Finished: true},
		{CustomerID: 1, CarID: 15, StartRent: parseDate("1/14/2021"), EndRent: parseDate("1/16/2021"), TotalCost: 1_800_000, Finished: true},
		{CustomerID: 16, CarID: 7, StartRent: parseDate("1/29/2021"), EndRent: parseDate("1/29/2021"), TotalCost: 1_000_000, Finished: true},
		{CustomerID: 12, CarID: 14, StartRent: parseDate("2/16/2021"), EndRent: parseDate("2/16/2021"), TotalCost: 800_000, Finished: true},
		{CustomerID: 5, CarID: 9, StartRent: parseDate("2/20/2021"), EndRent: parseDate("2/22/2021"), TotalCost: 1_500_000, Finished: true},
		{CustomerID: 2, CarID: 8, StartRent: parseDate("3/30/2021"), EndRent: parseDate("3/30/2021"), TotalCost: 500_000, Finished: false},
	}
)

func parseDate(d string) time.Time {
	format := "1/11/2011"
	t, _ := time.Parse(format, d)
	return t
}
