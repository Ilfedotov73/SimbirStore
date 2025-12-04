package repository

import (
	"context"
	"database/sql"
	"fmt"
	"service-customer/internal/model"
	"strings"

	_ "github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (*model.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.photo_url, p.create_at, p.characteristics, p.product_rating, vp.vendor_id
		FROM products p
		LEFT JOIN vendors_products vp ON p.id = vp.product_id
		WHERE p.id = $1
	`
	var p model.Product
	var photoUrl sql.NullString
	var chars sql.NullString
	var vendorID sql.NullInt64

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.Name, &p.Price, &photoUrl, &p.CreateAt, &chars, &p.ProductRating, &vendorID,
	)
	if err != nil {
		return nil, err
	}
	p.PhotoURL = photoUrl.String
	p.Characteristics = chars.String
	p.VendorID = vendorID.Int64

	return &p, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context, offset, limit int, minPrice, maxPrice float64, vendorID int64, queryParam, sortParam string) ([]model.Product, int, error) {
	var args []any
	baseQuery := `
		FROM products p
		JOIN vendors_products vp ON p.id = vp.product_id
	`
	whereClauses := []string{"1=1"}
	argID := 1

	if minPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("p.price >= $%d", argID))
		args = append(args, minPrice)
		argID++
	}
	if maxPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("p.price <= $%d", argID))
		args = append(args, maxPrice)
		argID++
	}
	if vendorID > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("vp.vendor_id = $%d", argID))
		args = append(args, vendorID)
		argID++
	}
	if queryParam != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(p.name ILIKE $%d OR p.characteristics ILIKE $%d)", argID, argID))
		args = append(args, "%"+queryParam+"%")
		argID++
	}

	whereSQL := strings.Join(whereClauses, " AND ")

	countQuery := "SELECT count(*) " + baseQuery + " WHERE " + whereSQL
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	orderBy := "ORDER BY p.id ASC"
	if sortParam != "" {
		parts := strings.Split(sortParam, ".")
		if len(parts) == 2 {
			field := parts[0]
			dir := parts[1]
			validFields := map[string]string{
				"price":  "p.price",
				"rating": "p.product_rating",
				"name":   "p.name",
			}
			if dbField, ok := validFields[field]; ok {
				if strings.ToUpper(dir) == "DESC" {
					orderBy = fmt.Sprintf("ORDER BY %s DESC", dbField)
				} else {
					orderBy = fmt.Sprintf("ORDER BY %s ASC", dbField)
				}
			}
		}
	}

	limitOffset := fmt.Sprintf("LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, limit, offset)

	selectQuery := `
		SELECT p.id, p.name, p.price, p.photo_url, p.create_at, p.characteristics, p.product_rating, vp.vendor_id 
	` + baseQuery + " WHERE " + whereSQL + " " + orderBy + " " + limitOffset

	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		var photoUrl sql.NullString
		var chars sql.NullString

		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &photoUrl, &p.CreateAt, &chars, &p.ProductRating, &p.VendorID); err != nil {
			return nil, 0, err
		}
		p.PhotoURL = photoUrl.String
		p.Characteristics = chars.String
		products = append(products, p)
	}

	return products, total, nil
}
