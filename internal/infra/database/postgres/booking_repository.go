package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mfmahendr/car-rental/internal/domain"
	"github.com/mfmahendr/car-rental/internal/domain/entities"
	"github.com/mfmahendr/car-rental/internal/infra/database"
)

type BookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{db}
}

func (r *BookingRepository) FindAll(ctx context.Context, page, limit *int) ([]entities.Booking, int64, error) {
	offset := getOffsetAndChangePageLimit(page, limit)

	queryString := `
		SELECT id, customer_id, car_id, start_rent, end_rent, total_cost, finished, COUNT(*) OVER() AS total_count
		FROM bookings
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`
	rows, err := getDB(ctx, r.db).Query(ctx, queryString, *limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var bookings []entities.Booking
	var total int64
	for rows.Next() {
		var b entities.Booking
		if err := rows.Scan(&b.BookingID, &b.CustomerID, &b.CarID, &b.StartRent, &b.EndRent, &b.TotalCost, &b.Finished, &total); err != nil {
			return nil, 0, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		bookings = append(bookings, b)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(bookings) == 0 || total == 0 {
		return nil, total, domain.ErrNotFound
	}

	return bookings, total, nil
}

func (r *BookingRepository) FindByID(ctx context.Context, id uint) (*entities.Booking, error) {
	q := `
		SELECT id, customer_id, car_id, start_rent, end_rent, total_cost, finished FROM bookings
		WHERE id=$1
	`
	row := getDB(ctx, r.db).QueryRow(ctx, q, id)
	b := new(entities.Booking)
	if err := row.Scan(&b.BookingID, &b.CustomerID, &b.CarID, &b.StartRent, &b.EndRent, &b.TotalCost, &b.Finished); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", database.ErrDBScan, err)
	}
	return b, nil
}

func (r *BookingRepository) GetBookingsByUserID(ctx context.Context, userID uint) ([]entities.Booking, int64, error) {
	q := `
		SELECT id, customer_id, car_id, start_rent, end_rent, total_cost, finished, COUNT(*) OVER() AS total_count
		FROM bookings
		WHERE customer_id = $1
		ORDER BY id ASC
	`
	rows, err := getDB(ctx, r.db).Query(ctx, q, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var bookings []entities.Booking
	var total int64
	for rows.Next() {
		var b entities.Booking
		if err := rows.Scan(&b.BookingID, &b.CustomerID, &b.CarID, &b.StartRent, &b.EndRent, &b.TotalCost, &b.Finished, &total); err != nil {
			return nil, 0, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		bookings = append(bookings, b)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(bookings) == 0 || total == 0 {
		return nil, 0, domain.ErrNotFound
	}

	return bookings, total, nil
}

func (r *BookingRepository) GetBookingsByCarID(ctx context.Context, carID uint) ([]entities.Booking, int64, error) {
	q := `
		SELECT id, customer_id, car_id, start_rent, end_rent, total_cost, finished, COUNT(*) OVER() AS total_count
		FROM bookings
		WHERE car_id = $1
		ORDER BY id ASC
	`
	rows, err := getDB(ctx, r.db).Query(ctx, q, carID)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var bookings []entities.Booking
	var total int64
	for rows.Next() {
		var b entities.Booking
		if err := rows.Scan(&b.BookingID, &b.CustomerID, &b.CarID, &b.StartRent, &b.EndRent, &b.TotalCost, &b.Finished, &total); err != nil {
			return nil, 0, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		bookings = append(bookings, b)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(bookings) == 0 || total == 0 {
		return nil, 0, domain.ErrNotFound
	}

	return bookings, total, nil
}

func (r *BookingRepository) Create(ctx context.Context, newBooking *entities.Booking) error {
	q := `
	INSERT INTO bookings (customer_id, car_id, start_rent, end_rent, total_cost, finished)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	err := getDB(ctx, r.db).QueryRow(ctx, q, newBooking.CustomerID, newBooking.CarID, newBooking.StartRent, newBooking.EndRent, newBooking.TotalCost, newBooking.Finished).Scan(&newBooking.CarID)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}

func (r *BookingRepository) Update(ctx context.Context, id uint, updatedBooking entities.Booking) error {
	const q = `
		UPDATE bookings 
		SET customer_id = $1, car_id = $2, start_rent = $3, end_rent = $4, total_cost = $5, finished  = $6
		WHERE id = $7
	`
	result, err := getDB(ctx, r.db).Exec(ctx, q, updatedBooking.CustomerID, updatedBooking.CarID, updatedBooking.StartRent, updatedBooking.EndRent, updatedBooking.TotalCost, updatedBooking.Finished, updatedBooking.BookingID)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *BookingRepository) UpdateRentDate(ctx context.Context, id uint, startRent, endRent time.Time) error {
	const q = `
		UPDATE bookings 
		SET start_rent = $1, end_rent = $2
		WHERE id = $3
	`
	result, err := getDB(ctx, r.db).Exec(ctx, q, startRent, endRent, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *BookingRepository) Delete(ctx context.Context, id uint) error {
	const q = `DELETE FROM bookings WHERE id = $1`
	_, err := getDB(ctx, r.db).Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}
