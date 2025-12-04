package repository

import (
	"context"
	"database/sql"
	"service-customer/internal/model"

	_ "github.com/lib/pq"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetStats(ctx context.Context, productID int64) (int, float64, error) {
	query := `
		SELECT count(*), COALESCE(AVG(rating), 0)
		FROM products_reviews
		WHERE product_id = $1
	`
	var count int
	var avgRating float64
	err := r.db.QueryRowContext(ctx, query, productID).Scan(&count, &avgRating)
	if err != nil {
		return 0, 0, err
	}
	return count, avgRating, nil
}

func (r *ReviewRepository) GetByProductID(ctx context.Context, productID int64, offset, limit int) ([]model.Review, int, error) {
	var total int
	countQuery := `SELECT count(*) FROM products_reviews WHERE product_id = $1`
	err := r.db.QueryRowContext(ctx, countQuery, productID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, customer_id, product_id, review, rating 
			  FROM products_reviews 
			  WHERE product_id = $1 
			  ORDER BY id DESC 
			  LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, productID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var reviews []model.Review
	for rows.Next() {
		var rev model.Review
		var reviewText sql.NullString
		if err := rows.Scan(&rev.ID, &rev.CustomerID, &rev.ProductID, &reviewText, &rev.Rating); err != nil {
			return nil, 0, err
		}
		rev.Review = reviewText.String
		reviews = append(reviews, rev)
	}
	return reviews, total, nil
}
