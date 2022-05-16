package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type MessageService struct {
	service *Service
	storage storage.Storage
}

func NewMessageService(storage storage.Storage, service *Service) *MessageService {
	return &MessageService{
		service: service,
		storage: storage,
	}
}

func (mr *MessageService) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	return nil, nil
}

func (mr *MessageService) FindById(ctx context.Context, id uint) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageService) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	return nil, nil
}

func (mr *MessageService) Delete(ctx context.Context, id uint) error {
	return nil
}
