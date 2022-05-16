package restapi

import "github.com/Despenrado/webMesk/internal/service"

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
