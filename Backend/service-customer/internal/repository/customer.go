package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetCustomerByID(ctx context.Context, id int64) (*model.Customer, error) {
	query := `SELECT id, first_name, last_name, phone_number, photo_url, customer_telegram_id, create_at, login, email 
			  FROM customers WHERE id = $1`
	var c model.Customer
	var phone, photo, tg sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.FirstName, &c.LastName, &phone, &photo, &tg, &c.CreateAt, &c.Login, &c.Email,
	)
	if err != nil {
		return nil, err
	}
	c.PhoneNumber = phone.String
	c.PhotoURL = photo.String
	c.CustomerTelegramID = tg.String
	return &c, nil
}
