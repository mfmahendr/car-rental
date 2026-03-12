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

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) FindAll(ctx context.Context, page, limit *int) ([]entities.Customer, int64, error) {
	offset := getOffsetAndChangePageLimit(page, limit)

	queryString := "SELECT id, name, nik, phone_number, COUNT(*) OVER() AS total_count FROM customer ORDER BY id ASC LIMIT $1 OFFSET $2"
	rows, err := getDB(ctx, r.db).Query(ctx, queryString, *limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var customers []entities.Customer
	var total int64
	for rows.Next() {
		var c entities.Customer
		if err := rows.Scan(&c.CustomerID, &c.Name, &c.NIK, &c.PhoneNumber, &total); err != nil {
			return nil, 0, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		customers = append(customers, c)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(customers) == 0 || total == 0 {
		return nil, total, domain.ErrNotFound
	}

	return customers, total, nil
}

func (r *CustomerRepository) FindByIDs(ctx context.Context, ids []uint) ([]entities.Customer, error) {
	if len(ids) == 0 {
		return nil, domain.ErrNotFound
	}

	placeholders, args := buildInClausePlaceholdersAndArgs(ids)

	q := fmt.Sprintf("SELECT id, name, nik, phone_number FROM customers WHERE id IN (%s)", strings.Join(placeholders, ","))
	rows, err := getDB(ctx, r.db).Query(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		var c entities.Customer
		if err := rows.Scan(&c.CustomerID, &c.Name, &c.NIK, &c.PhoneNumber); err != nil {
			return nil, fmt.Errorf("%s: %w", database.ErrDBScan, err)
		}
		customers = append(customers, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", database.ErrDBIterating, err)
	}

	if len(customers) == 0 {
		return nil, domain.ErrNotFound
	}

	return customers, nil
}

func (r *CustomerRepository) FindByID(ctx context.Context, id uint) (*entities.Customer, error) {
	q := "SELECT id, name, nik, phone_number FROM customers WHERE id=$1"
	row := getDB(ctx, r.db).QueryRow(ctx, q, id)
	c := new(entities.Customer)
	if err := row.Scan(&c.CustomerID, &c.Name, &c.NIK, &c.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", database.ErrDBScan, err)
	}
	return c, nil
}

func (r *CustomerRepository) Create(ctx context.Context, newCustomer *entities.Customer) error {
	q := `
	INSERT INTO customers (name, nik, phone_number)
	VALUES ($1, $2, $3)
	RETURNING id`
	err := getDB(ctx, r.db).QueryRow(ctx, q, newCustomer.Name, newCustomer.NIK, newCustomer.PhoneNumber).Scan(&newCustomer.CustomerID)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}

func (r *CustomerRepository) Update(ctx context.Context, id uint, updatedCustomer entities.Customer) error {
	const q = `
		UPDATE customers 
		SET name = $1, nik = $2, phone_number = $3
		WHERE id = $4
	`
	result, err := getDB(ctx, r.db).Exec(ctx, q, updatedCustomer.Name, updatedCustomer.NIK, updatedCustomer.PhoneNumber, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *CustomerRepository) Delete(ctx context.Context, id uint) error {
	const q = `DELETE FROM customers WHERE id = $1`
	_, err := getDB(ctx, r.db).Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %w", database.ErrDBOperation, err)
	}
	return nil
}
