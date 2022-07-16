package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type AtuthService struct {
	userService    *UserService
	authRepository storage.AuthRepository
}

func NewAuthService(userService *UserService, authRepository storage.AuthRepository) *AtuthService {
	return &AtuthService{
		userService:    userService,
		authRepository: authRepository,
	}
}

func (as *AtuthService) SignUp(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}

func (as *AtuthService) SignIn(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}

func (as *AtuthService) CheckAuthorization(ctx context.Context, id string, jwtToken string) error {
	return nil
}

func (as *AtuthService) Reauthorize(ctx context.Context, id string, jwtToken string) error {
	return nil
}

func (as *AtuthService) LogOut(ctx context.Context, id string, jwtToken string) error {
	return nil
}
