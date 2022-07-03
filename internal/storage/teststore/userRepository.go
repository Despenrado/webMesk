package teststore

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
	"github.com/google/uuid"
)

type UserRepository struct {
	userDB  map[uint]*model.User
	storage *Storage
}

func NewUserRepository(db map[uint]*model.User, storage *Storage) *UserRepository {
	return &UserRepository{
		userDB:  db,
		storage: storage,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = uint(uuid.New().ID())
	ur.userDB[user.ID] = user
	user, ok := ur.userDB[user.ID]
	if !ok {
		return user, utils.ErrRowsNumberAffected(0)
	}
	return user, nil
}

func (ur *UserRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error) {
	users := []model.User{}
	for _, v := range ur.userDB {
		users = append(users, *v)
	}
	return users, nil
}

func (ur *UserRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	user, ok := ur.userDB[id]
	if !ok {
		return user, utils.ErrRowsNumberAffected(0)
	}
	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	ur.userDB[user.ID] = user
	user, ok := ur.userDB[user.ID]
	if !ok {
		return user, utils.ErrRowsNumberAffected(0)
	}
	return user, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	_, ok1 := ur.userDB[id]
	delete(ur.userDB, id)
	_, ok2 := ur.userDB[id]
	if ok1 == ok2 {
		return utils.ErrRowsNumberAffected(0)
	}
	return nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	for _, v := range ur.userDB {
		if v.Email == email {
			user = *v
			return &user, nil
		}
	}
	return nil, utils.ErrRowsNumberAffected(0)
}

func (ur *UserRepository) FindByUserName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	for _, v := range ur.userDB {
		if v.UserName == userName {
			user = *v
			return &user, nil
		}
	}
	return nil, utils.ErrRowsNumberAffected(0)
}

func (ur *UserRepository) FilterUser(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	return nil, nil
}
