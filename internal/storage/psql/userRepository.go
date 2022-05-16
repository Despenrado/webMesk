package psql

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
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
	res := ur.storage.db.WithContext(ctx).Create(user)
	if res.RowsAffected != 1 {
		return user, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return user, res.Error
}

func (ur *UserRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error) {
	users := []model.User{}
	res := ur.storage.db.WithContext(ctx).Limit(limit).Offset(skip).Find(&users)
	return users, res.Error
}

func (ur *UserRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.db.WithContext(ctx).First(user, id)
	if res.RowsAffected != 1 {
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return user, res.Error
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	res := ur.storage.db.WithContext(ctx).Save(user)
	if res.RowsAffected != 1 {
		return user, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return user, res.Error
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	res := ur.storage.db.WithContext(ctx).Delete(&model.User{ID: id})
	if res.RowsAffected != 1 {
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return res.Error
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.db.WithContext(ctx).Where(map[string]interface{}{"email": email}).First(user)
	if res.RowsAffected != 1 {
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return user, res.Error
}

func (ur *UserRepository) FindByUserName(ctx context.Context, userName string) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.db.WithContext(ctx).Where(map[string]interface{}{"user_name": userName}).First(user)
	if res.RowsAffected != 1 {
		return nil, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	return user, res.Error
}
