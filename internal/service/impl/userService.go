package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/pkg/utils"
)

type UserService struct {
	storage storage.Storage
}

func NewUserService(storage storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := user.BeforeCreate(); err != nil {
		return user, err
	}
	if tmp, _ := us.storage.User().FindByEmail(ctx, user.Email); tmp != nil {
		return user, utils.ErrRecordAlreadyExists
	}
	if tmp, _ := us.storage.User().FindByUserName(ctx, user.UserName); tmp != nil {
		return user, utils.ErrRecordAlreadyExists
	}
	return us.storage.User().Create(ctx, user)
}

func (us *UserService) ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error) {
	return us.storage.User().ReadAll(ctx, skip, limit)
}

func (us *UserService) FindById(ctx context.Context, id uint) (*model.User, error) {
	return us.storage.User().FindById(ctx, id)
}

func (us *UserService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	oldUser, err := us.FindById(ctx, user.ID)
	if err != nil {
		return user, err
	}
	if !oldUser.VerifyPassword(user.Password) {
		return user, utils.ErrIncorrectEmailOrPassword
	}
	if err := user.BeforeCreate(); err != nil {
		return user, err
	}
	return us.storage.User().Update(ctx, user)
}

func (us *UserService) Delete(ctx context.Context, id uint) error {
	return us.storage.User().Delete(ctx, id)
}

func (us *UserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := us.storage.User().FindByEmail(ctx, email)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return user, nil
}

func (us *UserService) FilterUser(ctx context.Context, userFilter *model.UserFilter) ([]model.User, error) {
	return us.storage.User().FilterUser(ctx, userFilter)
}
