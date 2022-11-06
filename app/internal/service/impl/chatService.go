package impl

import (
	"context"
	"fmt"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/storage"
)

type ChatService struct {
	storage storage.Storage
}

func NewChatService(storage storage.Storage) *ChatService {
	return &ChatService{
		storage: storage,
	}
}

func (cs *ChatService) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	if err := chat.BeforeCreate(); err != nil {
		return chat, err
	}
	fmt.Println(chat)
	chat, err := cs.storage.Chat().Create(ctx, chat)
	if err != nil {
		return chat, err
	}
	fmt.Println(chat)
	fmt.Println(chat)
	return chat, nil
}

func (cs *ChatService) ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error) {
	return cs.storage.Chat().ReadAll(ctx, skip, limit)
}

func (cs *ChatService) FindById(ctx context.Context, id uint) (*model.Chat, error) {
	chat, err := cs.storage.Chat().FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (cs *ChatService) Update(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	if err := chat.BeforeCreate(); err != nil {
		return chat, err
	}
	chat, err := cs.storage.Chat().Update(ctx, chat)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func (cs *ChatService) Delete(ctx context.Context, id uint) error {
	return cs.storage.Chat().Delete(ctx, id)
}

func (cs *ChatService) FindByUserId(ctx context.Context, id uint) ([]model.Chat, error) {
	chats, err := cs.storage.Chat().FindByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (cs *ChatService) FilterChat(ctx context.Context, chatFilter *model.ChatFilter) ([]model.Chat, error) {
	return cs.storage.Chat().FilterChat(ctx, chatFilter)
}
