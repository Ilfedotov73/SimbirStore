package service

import (
	"context"
	"service-customer/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=UserRepository --output=../../mocks --outpkg=mocks --with-expecter=true
type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
