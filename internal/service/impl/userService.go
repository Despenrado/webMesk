package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/pkg/utils"
)

type UserService struct {
	service *Service
	storage storage.Storage
}

func NewUserService(storage storage.Storage, service *Service) *UserService {
	return &UserService{
		service: service,
		storage: storage,
	}
}

func (us *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.BeforeCreate()
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
	user.BeforeCreate()
	if tmp, _ := us.storage.User().FindByEmail(ctx, user.Email); tmp != nil {
		return user, utils.ErrRecordAlreadyExists
	}
	if tmp, _ := us.storage.User().FindByUserName(ctx, user.UserName); tmp != nil {
		return user, utils.ErrRecordAlreadyExists
	}
	return us.storage.User().Update(ctx, user)
}

func (us *UserService) Delete(ctx context.Context, id uint) error {
	return us.storage.User().Delete(ctx, id)
}

func (us *UserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return us.storage.User().FindByEmail(ctx, email)
}
