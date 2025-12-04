package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"

	_ "github.com/lib/pq"
)

type VendorRepository struct {
	db *sql.DB
}

func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{db: db}
}

func (r *VendorRepository) GetVendorByID(ctx context.Context, id int64) (*model.Vendor, error) {
	query := `SELECT id, first_name, last_name, phone_number, photo_url, vendor_telegram_id, create_at, login, email 
			  FROM vendors WHERE id = $1`
	var v model.Vendor
	var phone, photo, tg sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&v.ID, &v.FirstName, &v.LastName, &phone, &photo, &tg, &v.CreateAt, &v.Login, &v.Email,
	)
	if err != nil {
		return nil, err
	}
	v.PhoneNumber = phone.String
	v.PhotoURL = photo.String
	v.VendorTelegramID = tg.String
	return &v, nil
}
