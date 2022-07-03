package psql

import (
	"context"
	"fmt"

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
	tx := ur.storage.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	fmt.Println("1")
	res := tx.WithContext(ctx).Exec("DELETE FROM user_chat WHERE user_id = ?", id)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	fmt.Println("2")
	res = tx.WithContext(ctx).Delete(&model.User{ID: id})
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	fmt.Println("3")
	if res.RowsAffected != 1 {
		tx.Rollback()
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	fmt.Println("4")
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
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
