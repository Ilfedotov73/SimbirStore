package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"
	"time"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) GetByCustomerID(ctx context.Context, customerID int64, from, to time.Time, offset, limit int) ([]model.Notification, int, error) {
	var total int
	countQuery := `SELECT count(*) FROM notices WHERE entity_id = $1 AND create_at BETWEEN $2 AND $3`
	err := r.db.QueryRow(countQuery, customerID, from, to).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, text, entity_id, create_at 
			  FROM notices 
			  WHERE entity_id = $1 AND create_at BETWEEN $2 AND $3
			  ORDER BY create_at DESC 
			  LIMIT $4 OFFSET $5`

	rows, err := r.db.QueryContext(ctx, query, customerID, from, to, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Notification
	for rows.Next() {
		var n model.Notification
		if err := rows.Scan(&n.ID, &n.Text, &n.EntityID, &n.CreateAt); err != nil {
			return nil, 0, err
		}
		items = append(items, n)
	}
	return items, total, nil
}
