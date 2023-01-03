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
		return nil, err
	}
	usr, err := us.storage.User().Create(ctx, user)
	if err != nil {
		return nil, err
	}
	usr.Sanitize()
	return usr, nil
}

func (us *UserService) ReadAll(ctx context.Context, skip int, limit int) ([]model.User, error) {
	usrs, err := us.storage.User().ReadAll(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
	for i := range usrs {
		usrs[i].Sanitize()
	}
	return usrs, nil
}

func (us *UserService) FindById(ctx context.Context, id uint) (*model.User, error) {
	user, err := us.storage.User().FindById(ctx, id)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	user.Sanitize()
	return user, nil
}

func (us *UserService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	oldUser, err := us.storage.User().FindById(ctx, user.ID)
	if err != nil {
		return user, err
	}
	if !oldUser.VerifyPassword(user.Password) {
		return user, utils.ErrIncorrectEmailOrPassword
	}
	if err := user.BeforeCreate(); err != nil {
		return user, err
	}
	user, err = us.storage.User().Update(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Sanitize()
	return user, nil
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
	usrs, err := us.storage.User().GetUsersByFilter(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	for i := range usrs {
		usrs[i].Sanitize()
	}
	return usrs, nil
}
