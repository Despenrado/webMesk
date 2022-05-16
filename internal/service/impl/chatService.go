package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type ChatService struct {
	service *Service
	storage storage.Storage
}

func NewChatService(storage storage.Storage, service *Service) *ChatService {
	return &ChatService{
		service: service,
		storage: storage,
	}
}

func (cr *ChatService) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error) {
	return nil, nil
}

func (cr *ChatService) FindById(ctx context.Context, id uint) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatService) Update(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatService) Delete(ctx context.Context, id uint) error {
	return nil
}
