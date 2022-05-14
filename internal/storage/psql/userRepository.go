package psql

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type UserRepository struct {
	storage *Storage
}

func NewUserRepository(storage *Storage) *UserRepository {
	return &UserRepository{
		storage: storage,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetAll(ctx context.Context) (*[]model.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint64) (*model.User, error) {
	return nil, nil
}
