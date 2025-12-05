package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"service-customer/internal/model"
	"service-customer/internal/repository"
)

type ProductRepositorySuite struct {
	suite.Suite
	db   *sql.DB
	repo *repository.ProductRepository
	ctx  context.Context
}

func (s *ProductRepositorySuite) SetupSuite() {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5431/test_db?sslmode=disable"
	}

	var err error
	s.db, err = sql.Open("postgres", dsn)
	s.Require().NoError(err, "Failed to open DB connection")

	err = s.db.Ping()
	s.Require().NoError(err, "Failed to ping DB")

	s.repo = repository.NewProductRepository(s.db)
	s.ctx = context.Background()
}

func (s *ProductRepositorySuite) TearDownSuite() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *ProductRepositorySuite) TearDownTest() {
	_, err := s.db.Exec(`TRUNCATE TABLE products, vendors_products RESTART IDENTITY CASCADE`)
	s.Require().NoError(err)
}

func (s *ProductRepositorySuite) setupBaseEnv() {
	vendorIDs := []int64{1, 2, 10, 20, 30}

	for _, id := range vendorIDs {
		firstName := fmt.Sprintf("FName%d", id)
		lastName := fmt.Sprintf("LName%d", id)
		login := fmt.Sprintf("vendor_login_%d", id)
		email := fmt.Sprintf("vendor%d@test.com", id)
		passwordHash := "dummy_secure_hash"

		_, err := s.db.Exec(`
			INSERT INTO vendors (
				id, first_name, last_name, login, email, password
			) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			ON CONFLICT (id) DO NOTHING`,
			id, firstName, lastName, login, email, passwordHash)
		s.Require().NoError(err, fmt.Sprintf("Failed to insert vendor %d", id))
	}
}

func (s *ProductRepositorySuite) createProductOnDB(id int64, name string, price float64, rating float64, chars string, photoURL string, vendorID int64) {
	_, err := s.db.Exec(`
		INSERT INTO products (id, name, price, photo_url, create_at, characteristics, product_rating)
		VALUES ($1, $2, $3, $4, NOW(), $5, $6) 
		ON CONFLICT (id) DO UPDATE SET 
			name = EXCLUDED.name, price = EXCLUDED.price, 
			characteristics = EXCLUDED.characteristics, product_rating = EXCLUDED.product_rating`,
		id, name, price, photoURL, chars, rating)
	s.Require().NoError(err, fmt.Sprintf("Failed to insert product %d", id))

	_, err = s.db.Exec(`
		INSERT INTO vendors_products (product_id, vendor_id)
		VALUES ($1, $2) 
		ON CONFLICT DO NOTHING`,
		id, vendorID)
	s.Require().NoError(err, fmt.Sprintf("Failed to link vendor %d to product %d", vendorID, id))
}

func normalizeTime(t time.Time) time.Time {
	return t.Truncate(time.Second)
}

func (s *ProductRepositorySuite) TestGetProductByID() {
	s.setupBaseEnv()
	s.createProductOnDB(101, "Test Laptop", 1500.00, 4.5, "Powerful, 16GB RAM", fmt.Sprintf("url_%d", 101), 1)
	s.createProductOnDB(102, "Simple Mouse", 15.00, 3.8, "", "", 2)

	var createAt time.Time
	err := s.db.QueryRow("SELECT create_at FROM products WHERE id = $1", 101).Scan(&createAt)
	s.Require().NoError(err)
	expectedTime := normalizeTime(createAt)

	tests := []struct {
		name        string
		inputID     int64
		expectedErr error
		verify      func(*testing.T, *model.Product)
	}{
		{
			name:        "Success: Get existing product with all fields filled",
			inputID:     101,
			expectedErr: nil,
			verify: func(t *testing.T, p *model.Product) {
				assert.Equal(t, int64(101), p.ID)
				assert.Equal(t, "Test Laptop", p.Name)
				assert.Equal(t, 1500.00, p.Price)
				assert.Equal(t, int64(1), p.VendorID)
				assert.Equal(t, 4.5, p.ProductRating)
				assert.Equal(t, "Powerful, 16GB RAM", p.Characteristics)
				assert.Equal(t, expectedTime, normalizeTime(p.CreateAt))
			},
		},
		{
			name:        "Success: Get existing product with optional fields being NULL/empty",
			inputID:     102,
			expectedErr: nil,
			verify: func(t *testing.T, p *model.Product) {
				assert.Equal(t, int64(102), p.ID)
				assert.Equal(t, 15.00, p.Price)
				assert.Empty(t, p.Characteristics, "Characteristics should be empty string for NULL/empty value")
				assert.Empty(t, p.PhotoURL, "PhotoURL should be empty string for NULL/empty value")
				assert.Equal(t, int64(2), p.VendorID)
			},
		},
		{
			name:        "Error: Product not found should return sql.ErrNoRows",
			inputID:     999,
			expectedErr: sql.ErrNoRows,
			verify:      nil,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			p, err := s.repo.GetProductByID(s.ctx, tc.inputID)

			if tc.expectedErr != nil {
				s.ErrorIs(err, tc.expectedErr)
				s.Nil(p)
			} else {
				s.NoError(err)
				s.NotNil(p)
				if tc.verify != nil {
					tc.verify(s.T(), p)
				}
			}
		})
	}
}

func (s *ProductRepositorySuite) TestGetProducts() {
	s.setupBaseEnv()
	s.createProductOnDB(1, "Apple iPhone 14", 800.00, 4.8, "Mobile phone, iOS", fmt.Sprintf("url_%d", 1), 10)
	s.createProductOnDB(2, "Samsung Galaxy S22", 650.00, 4.2, "Mobile phone, Android", fmt.Sprintf("url_%d", 2), 20)
	s.createProductOnDB(3, "MacBook Pro M2", 2100.00, 5.0, "Laptop, 16GB RAM", fmt.Sprintf("url_%d", 3), 10)
	s.createProductOnDB(4, "Dell XPS 13", 1200.00, 4.5, "Windows Laptop, i7", fmt.Sprintf("url_%d", 4), 30)
	s.createProductOnDB(5, "HP Printer", 150.00, 3.5, "Inkjet Printer", fmt.Sprintf("url_%d", 5), 20)

	tests := []struct {
		name          string
		offset        int
		limit         int
		minPrice      float64
		maxPrice      float64
		vendorID      int64
		queryParam    string
		sortParam     string
		expectedTotal int
		expectedIDs   []int64
		expectedErr   error
	}{
		{
			name:          "Success: No filters, default sort (ID ASC), pagination",
			offset:        0,
			limit:         3,
			expectedTotal: 5,
			expectedIDs:   []int64{1, 2, 3},
		},
		{
			name:          "Success: Pagination with offset",
			offset:        3,
			limit:         3,
			expectedTotal: 5,
			expectedIDs:   []int64{4, 5},
		},
		{
			name:          "Success: Filter by Min and Max Price",
			offset:        0,
			limit:         10,
			minPrice:      600.00,
			maxPrice:      1300.00,
			expectedTotal: 3,
			expectedIDs:   []int64{1, 2, 4},
		},
		{
			name:          "Success: Filter by VendorID",
			offset:        0,
			limit:         10,
			vendorID:      20,
			expectedTotal: 2,
			expectedIDs:   []int64{2, 5},
		},
		{
			name:          "Success: Search by Name (ILIKE, partial match)",
			offset:        0,
			limit:         10,
			queryParam:    "sam",
			expectedTotal: 1,
			expectedIDs:   []int64{2},
		},
		{
			name:          "Success: Search by Characteristics (ILIKE)",
			offset:        0,
			limit:         10,
			queryParam:    "android",
			expectedTotal: 1,
			expectedIDs:   []int64{2},
		},
		{
			name:          "Success: Search by query matching both name and characteristics",
			offset:        0,
			limit:         10,
			queryParam:    "lap",
			expectedTotal: 2,
			expectedIDs:   []int64{3, 4},
		},
		{
			name:          "Success: Sort by Price ASC",
			offset:        0,
			limit:         10,
			sortParam:     "price.asc",
			expectedTotal: 5,
			expectedIDs:   []int64{5, 2, 1, 4, 3},
		},
		{
			name:          "Success: Sort by Rating DESC",
			offset:        0,
			limit:         3,
			sortParam:     "rating.desc",
			expectedTotal: 5,
			expectedIDs:   []int64{3, 1, 4},
		},
		{
			name:          "Success: Combined Filters (Vendor 10) and Sort (Price DESC)",
			offset:        0,
			limit:         10,
			vendorID:      10,
			sortParam:     "price.desc",
			expectedTotal: 2,
			expectedIDs:   []int64{3, 1},
		},
		{
			name:          "Success: Invalid sort parameter defaults to ID ASC",
			offset:        0,
			limit:         2,
			sortParam:     "invalid.param",
			expectedTotal: 5,
			expectedIDs:   []int64{1, 2},
		},
		{
			name:          "Success: No results found",
			offset:        0,
			limit:         10,
			minPrice:      5000.00,
			expectedTotal: 0,
			expectedIDs:   []int64{},
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			products, total, err := s.repo.GetProducts(
				s.ctx,
				tc.offset,
				tc.limit,
				tc.minPrice,
				tc.maxPrice,
				tc.vendorID,
				tc.queryParam,
				tc.sortParam,
			)

			if tc.expectedErr != nil {
				s.ErrorIs(err, tc.expectedErr)
				s.Empty(products)
				s.Equal(0, total)
			} else {
				s.NoError(err)
				s.Equal(tc.expectedTotal, total, "Total count mismatch")
				s.Len(products, len(tc.expectedIDs), "Returned items count mismatch")

				actualIDs := make([]int64, len(products))
				for i, p := range products {
					actualIDs[i] = p.ID
				}
				s.Equal(tc.expectedIDs, actualIDs, "Returned product IDs are not in expected order or list")
			}
		})
	}
}

func TestProductRepositorySuite(t *testing.T) {
	suite.Run(t, new(ProductRepositorySuite))
}
