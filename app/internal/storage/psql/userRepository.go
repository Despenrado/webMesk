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
	res := ur.storage.DB.WithContext(ctx).Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error) {
	users := []model.User{}
	res := ur.storage.DB.WithContext(ctx).Limit(limit).Offset(skip).Find(&users)
	return users, res.Error
}

func (ur *UserRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.DB.WithContext(ctx).First(user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	res := ur.storage.DB.WithContext(ctx).Save(user)
	if res.RowsAffected != 1 {
		return user, utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	tx := ur.storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	res := tx.WithContext(ctx).Exec("DELETE FROM user_chat WHERE user_id = ?", id)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	res = tx.WithContext(ctx).Delete(&model.User{ID: id})
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if res.RowsAffected != 1 {
		tx.Rollback()
		return utils.ErrRowsNumberAffected(int(res.RowsAffected))
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.DB.WithContext(ctx).Where(map[string]interface{}{"email": email}).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) FindByUserName(ctx context.Context, userName string) (*model.User, error) {
	user := &model.User{}
	res := ur.storage.DB.WithContext(ctx).Where(map[string]interface{}{"user_name": userName}).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) GetUsersByFilter(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	query := ur.storage.DB.WithContext(ctx)
	if userFilter.Email != "" {
		query = query.Where("email LIKE ?", "%"+userFilter.Email+"%")
	}
	if userFilter.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+userFilter.UserName+"%")
	}
	if userFilter.SessionId != "" {
		query = query.Where("sessionId LIKE ?", "%"+userFilter.SessionId+"%")
	}
	query = query.Offset(int(userFilter.Skip))
	if userFilter.Limit != 0 {
		query = query.Limit(int(userFilter.Limit))
	}
	users := []model.User{}
	res := query.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}
