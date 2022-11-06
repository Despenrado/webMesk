package impl

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	putils "github.com/Despenrado/webMesk/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

type AtuthService struct {
	userService    *UserService
	authRepository storage.AuthRepository
	jwtConfig      *utils.JWTConfig
}

func NewAuthService(userService *UserService, cacheStorage storage.CacheStorage, config *utils.Config) *AtuthService {
	return &AtuthService{
		userService:    userService,
		authRepository: cacheStorage.Auth(),
		jwtConfig:      config.JWTConfig,
	}
}

func (as *AtuthService) SignUp(ctx context.Context, user *model.User) (string, error) {
	user, err := as.userService.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return as.login(ctx, user)
}

func (as *AtuthService) SignIn(ctx context.Context, user *model.User) (*model.User, string, error) {
	dbUser, err := as.userService.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, "", err
	}
	if !dbUser.VerifyPassword(user.Password) {
		return nil, "", putils.ErrIncorrectEmailOrPassword
	}
	dbUser.Sanitize()
	token, err := as.login(ctx, dbUser)
	return dbUser, token, err
}

func (as *AtuthService) ValidAuthorization(ctx context.Context, jwtToken string) (context.Context, error) {
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.jwtConfig.JWTKey), nil
	})
	if err != nil {
		return ctx, err
	}
	if !tkn.Valid {
		return ctx, jwt.ErrSignatureInvalid
	}
	redisUser, err := as.authRepository.FindById(ctx, claims.Id)
	if err != nil {
		return ctx, err
	}
	if redisUser.ID == claims.Id && redisUser.JWTToken == jwtToken {
		return context.WithValue(ctx, "user_id", claims.Id), nil
	}
	return ctx, putils.ErrUnauthorized
}

func (as *AtuthService) Reauthorize(ctx context.Context, id string, jwtToken string) error {

	return nil
}

func (as *AtuthService) LogOut(ctx context.Context, jwtToken string) error {
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.jwtConfig.JWTKey), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return jwt.ErrSignatureInvalid
	}
	if err := as.authRepository.DeleteByUserId(ctx, claims.Id); err != nil {
		return err
	}
	return nil
}

func (as *AtuthService) login(ctx context.Context, user *model.User) (string, error) {
	tokenString, err := as.createToken(user)
	if err != nil {
		return "", err
	}
	userAuth := &model.UserAuth{
		ID:        strconv.FormatUint(uint64(user.ID), 10),
		SessionId: "",
		JWTToken:  tokenString,
	}
	if err := as.authRepository.Set(ctx, userAuth, as.jwtConfig.TknExpires); err != nil {
		return "", err
	}
	return tokenString, nil
}

func (as *AtuthService) createToken(user *model.User) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(as.jwtConfig.TknExpires).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("jwtkey:", as.jwtConfig.JWTKey)
	log.Println("claims:", claims)
	tokenString, err := jwtToken.SignedString([]byte(as.jwtConfig.JWTKey))
	if err != nil {
		log.Println("jwtkey error")
		return "", err
	}
	return tokenString, nil
}
