package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type MessageService struct {
	storage storage.Storage
}

func NewMessageService(storage storage.Storage) *MessageService {
	return &MessageService{
		storage: storage,
	}
}

func (ms *MessageService) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	if err := message.BeforeCreate(); err != nil {
		return message, err
	}
	message, err := ms.storage.Message().Create(ctx, message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (ms *MessageService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	return ms.storage.Message().ReadAll(ctx, skip, limit)
}

func (ms *MessageService) FindById(ctx context.Context, id uint) (*model.Message, error) {
	message, err := ms.storage.Message().FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (ms *MessageService) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	message, err := ms.storage.Message().Update(ctx, message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (ms *MessageService) Delete(ctx context.Context, id uint) error {
	return ms.storage.Message().Delete(ctx, id)
}

func (ms *MessageService) FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error) {
	return ms.storage.Message().FilterMessage(ctx, messageFilter)
}
