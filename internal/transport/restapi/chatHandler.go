package restapi

import "github.com/Despenrado/webMesk/internal/service"

type ChatHandler struct {
	service service.Service
}

func NewChatHandler(service service.Service) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}
