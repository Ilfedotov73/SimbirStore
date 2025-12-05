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

type VendorRepositorySuite struct {
	suite.Suite
	db   *sql.DB
	repo *repository.VendorRepository
	ctx  context.Context
}

func (s *VendorRepositorySuite) SetupSuite() {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5431/test_db?sslmode=disable"
	}

	var err error
	s.db, err = sql.Open("postgres", dsn)
	s.Require().NoError(err, "Failed to open DB connection")

	err = s.db.Ping()
	s.Require().NoError(err, "Failed to ping DB")

	s.repo = repository.NewVendorRepository(s.db)
	s.ctx = context.Background()
}

func (s *VendorRepositorySuite) TearDownSuite() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *VendorRepositorySuite) TearDownTest() {
	_, err := s.db.Exec(`TRUNCATE TABLE vendors RESTART IDENTITY CASCADE`)
	s.Require().NoError(err)
}

func (s *VendorRepositorySuite) createVendorOnDB(id int64, firstName, lastName, login, email, password, phone, photoURL, telegramID string) {
	_, err := s.db.Exec(`
		INSERT INTO vendors (
			id, first_name, last_name, phone_number, photo_url, vendor_telegram_id, create_at, login, email, password
		)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7, $8, $9) 
		ON CONFLICT (id) DO NOTHING`,
		id, firstName, lastName, phone, photoURL, telegramID, login, email, password)
	s.Require().NoError(err, fmt.Sprintf("Failed to insert vendor %d", id))
}

func (s *VendorRepositorySuite) TestGetVendorByID() {
	s.createVendorOnDB(101,
		"Ivan", "Petrov", "ivan_p", "ivan@test.com", "hash123",
		"79001234567", "http://photo.url/101.jpg", "@ivan_tg")

	s.createVendorOnDB(102,
		"Maria", "Sidorova", "masha_s", "masha@test.com", "hash456",
		"", "", "")

	var createAt time.Time
	err := s.db.QueryRow("SELECT create_at FROM vendors WHERE id = $1", 101).Scan(&createAt)
	s.Require().NoError(err)
	expectedTime := normalizeTime(createAt)

	tests := []struct {
		name        string
		inputID     int64
		expectedErr error
		verify      func(*testing.T, *model.Vendor)
	}{
		{
			name:        "Success: Get existing vendor with all optional fields filled",
			inputID:     101,
			expectedErr: nil,
			verify: func(t *testing.T, v *model.Vendor) {
				assert.Equal(t, int64(101), v.ID)
				assert.Equal(t, "Ivan", v.FirstName)
				assert.Equal(t, "Petrov", v.LastName)
				assert.Equal(t, "ivan_p", v.Login)
				assert.Equal(t, "ivan@test.com", v.Email)
				assert.Equal(t, "79001234567", v.PhoneNumber)
				assert.Equal(t, "http://photo.url/101.jpg", v.PhotoURL)
				assert.Equal(t, "@ivan_tg", v.VendorTelegramID)
				assert.Equal(t, expectedTime, normalizeTime(v.CreateAt))
			},
		},
		{
			name:        "Success: Get existing vendor with optional fields being NULL/empty",
			inputID:     102,
			expectedErr: nil,
			verify: func(t *testing.T, v *model.Vendor) {
				assert.Equal(t, int64(102), v.ID)
				assert.Equal(t, "Maria", v.FirstName)
				assert.Equal(t, "Sidorova", v.LastName)
				assert.Equal(t, "masha_s", v.Login)
				assert.Equal(t, "masha@test.com", v.Email)

				assert.Empty(t, v.PhoneNumber, "PhoneNumber should be empty string for NULL value")
				assert.Empty(t, v.PhotoURL, "PhotoURL should be empty string for NULL value")
				assert.Empty(t, v.VendorTelegramID, "VendorTelegramID should be empty string for NULL value")
			},
		},
		{
			name:        "Error: Vendor not found should return sql.ErrNoRows",
			inputID:     999,
			expectedErr: sql.ErrNoRows,
			verify:      nil,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			v, err := s.repo.GetVendorByID(s.ctx, tc.inputID)

			if tc.expectedErr != nil {
				s.ErrorIs(err, tc.expectedErr)
				s.Nil(v)
			} else {
				s.NoError(err)
				s.NotNil(v)
				if tc.verify != nil {
					tc.verify(s.T(), v)
				}
			}
		})
	}
}

func TestVendorRepositorySuite(t *testing.T) {
	suite.Run(t, new(VendorRepositorySuite))
}
