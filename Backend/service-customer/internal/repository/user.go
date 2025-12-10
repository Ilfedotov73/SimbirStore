package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, first_name, last_name, phone_number, photo_url, user_telegram_id, create_at, login, email, role
			  FROM users WHERE id = $1`
	var u model.User
	var phone, photo, tg sql.NullString
	var role int8

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID, &u.FirstName, &u.LastName, &phone, &photo, &tg, &u.CreateAt, &u.Login, &u.Email, &role,
	)
	if err != nil {
		return nil, err
	}
	u.Role = model.UserRole(role)
	u.PhoneNumber = phone.String
	u.PhotoURL = photo.String
	u.UserTelegramID = tg.String
	return &u, nil
}
