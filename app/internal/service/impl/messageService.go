package impl

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/pkg/utils"
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
	chat, err := ms.storage.Chat().FindById(ctx, message.ChatID)
	if err != nil {
		return message, err
	}
	if !chat.CheckPermissions(message.UserID) {
		return message, utils.ErrNoPermissions
	}
	message, err = ms.storage.Message().Create(ctx, message)
	if err != nil {
		return message, err
	}
	message.Sanitize()
	return message, nil
}

func (ms *MessageService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Message, error) {
	messages, err := ms.storage.Message().ReadAll(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
	for i, _ := range messages {
		messages[i].Sanitize()
	}
	return messages, nil
}

func (ms *MessageService) FindById(ctx context.Context, id uint) (*model.Message, error) {
	message, err := ms.storage.Message().FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	message.Sanitize()
	return message, nil
}

func (ms *MessageService) Update(ctx context.Context, message *model.Message) (*model.Message, error) {
	message, err := ms.storage.Message().Update(ctx, message)
	if err != nil {
		return nil, err
	}
	message.Sanitize()
	return message, nil
}

func (ms *MessageService) Delete(ctx context.Context, message *model.Message) error {
	return ms.storage.Message().Delete(ctx, message)
}

func (ms *MessageService) FilterMessage(ctx context.Context, messageFilter *model.MessageFilter) ([]model.Message, error) {
	messages, err := ms.storage.Message().FilterMessage(ctx, messageFilter)
	if err != nil {
		return nil, err
	}
	for i, _ := range messages {
		messages[i].Sanitize()
	}
	return messages, nil
}

func (ms *MessageService) MarkAsRead(ctx context.Context, id uint, user_id uint) error {
	return ms.storage.Message().MarkAsRead(ctx, id, user_id)
}
