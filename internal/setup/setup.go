package setup

import (
	"github.com/mfmahendr/car-rental/internal/application/usecase"
	"github.com/mfmahendr/car-rental/internal/delivery/http"
	"github.com/mfmahendr/car-rental/internal/delivery/http/middleware"
	"github.com/mfmahendr/car-rental/internal/delivery/http/route"
	"github.com/mfmahendr/car-rental/internal/infra/database/postgres"
)

func (app *Application) Setup() {
	// repositories
	carRepo := postgres.NewCarRepository(app.db)
	customerRepo := postgres.NewCustomerRepository(app.db)
	bookingRepo := postgres.NewBookingRepository(app.db)

	tractor := postgres.NewTransactor(app.db)

	// applications
	carUsecae := usecase.NewCarUsecase(app.val, carRepo)
	customerUsecae := usecase.NewCustomerUsecase(app.val, customerRepo)
	bookingUsecae := usecase.NewBookingUsecase(app.val, tractor, bookingRepo, carRepo, customerRepo)

	// controllers/handlers
	carController := http.NewCarController(carUsecae)
	customerController := http.NewCustomerController(customerUsecae)
	bookingController := http.NewBookingController(bookingUsecae)

	router := route.Router{
		App:                app.httpServer,
		CarController:      carController,
		CustomerController: customerController,
		BookingController:  bookingController,
	}

	router.App.Use(middleware.APIVersion(app.config.App.AppVersion))
	router.App.Use(middleware.Logger())

	router.RegisterCarRoutes()
	router.RegisterCustomerRoutes()
	router.RegisterBookingRoutes()
}
