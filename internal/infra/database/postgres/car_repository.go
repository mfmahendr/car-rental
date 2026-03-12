package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mfmahendr/car-rental/internal/domain"
	"github.com/mfmahendr/car-rental/internal/domain/entities"
	"github.com/mfmahendr/car-rental/internal/infra/database"
)

type CarRepository struct {
	db *pgxpool.Pool
}

func NewCarRepository(db *pgxpool.Pool) *CarRepository {
	return &CarRepository{db}
}

func (r *CarRepository) FindAll(ctx context.Context, page, limit *int) ([]entities.Car, int64, error) {
	offset := getOffsetAndChangePageLimit(page, limit)

	q := "SELECT id, name, stock, daily_rent, COUNT(*) OVER() AS total_count FROM cars ORDER BY id ASC LIMIT $1 OFFSET $2"
	rows, err := getDB(ctx, r.db).Query(ctx, q, *limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var cars []entities.Car
	var total int64
	for rows.Next() {
		var c entities.Car
		if err := rows.Scan(&c.CarID, &c.Name, &c.Stock, &c.DailyRent, &total); err != nil {
			return nil, 0, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		cars = append(cars, c)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(cars) == 0 || total == 0 {
		return nil, total, domain.ErrNotFound
	}

	return cars, total, nil
}

func (r *CarRepository) FindByIDs(ctx context.Context, ids []uint) ([]entities.Car, error) {
	if len(ids) == 0 {
		return nil, domain.ErrNotFound
	}

	placeholders, args := buildInClausePlaceholdersAndArgs(ids)

	q := fmt.Sprintf("SELECT id, name, stock, daily_rent FROM cars WHERE id IN (%s)", strings.Join(placeholders, ","))
	rows, err := getDB(ctx, r.db).Query(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var c entities.Car
		if err := rows.Scan(&c.CarID, &c.Name, &c.Stock, &c.DailyRent); err != nil {
			return nil, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		cars = append(cars, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(cars) == 0 {
		return nil, domain.ErrNotFound
	}

	return cars, nil
}

func (r *CarRepository) FindByID(ctx context.Context, id uint) (*entities.Car, error) {
	q := "SELECT id, name, stock, daily_rent FROM cars WHERE id=$1 ORDER BY id ASC"
	row := getDB(ctx, r.db).QueryRow(ctx, q, id)
	c := new(entities.Car)
	if err := row.Scan(&c.CarID, &c.Name, &c.Stock, &c.DailyRent); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", database.ErrDBScan, err)
	}
	return c, nil
}

func (r *CarRepository) Create(ctx context.Context, newCar *entities.Car) error {
	q := `
	INSERT INTO cars (name, stock, daily_rent)
	VALUES ($1, $2, $3)
	RETURNING id`
	err := getDB(ctx, r.db).QueryRow(ctx, q, newCar.Name, newCar.Stock, newCar.DailyRent).Scan(&newCar.CarID)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}

func (r *CarRepository) Update(ctx context.Context, id uint, updatedCar entities.Car) error {
	const q = `
		UPDATE cars 
		SET name = $1, stock = $2, daily_rent = $3 
		WHERE id = $4
	`
	result, err := getDB(ctx, r.db).Exec(ctx, q, updatedCar.Name, updatedCar.Stock, updatedCar.DailyRent, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *CarRepository) UpdateStock(ctx context.Context, id int64, stockChanges int) error {
	const q = `
		UPDATE cars
		SET stock = (SELECT stock FROM cars WHERE id = $1) + $2
		WHERE id = $1
	`
	result, err := getDB(ctx, r.db).Exec(ctx, q, id, stockChanges)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *CarRepository) Delete(ctx context.Context, id uint) error {
	const q = `DELETE FROM cars WHERE id = $1`
	_, err := getDB(ctx, r.db).Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}
