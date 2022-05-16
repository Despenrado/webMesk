package restapi

import "github.com/Despenrado/webMesk/internal/service"

type MessageHandler struct {
	service service.Service
}

func NewMessageHandler(service service.Service) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}
