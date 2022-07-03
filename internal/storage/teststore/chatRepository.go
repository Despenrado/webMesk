package teststore

import (
	"context"

	"github.com/Despenrado/webMesk/internal/model"
)

type ChatRepository struct {
	chatDB  map[uint]*model.Chat
	storage *Storage
}

func NewChatRepository(db map[uint]*model.Chat, storage *Storage) *ChatRepository {
	return &ChatRepository{
		chatDB:  db,
		storage: storage,
	}
}

func (cr *ChatRepository) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatRepository) ReadAll(ctx context.Context, skip int, limit int) ([]model.Chat, error) {
	return nil, nil
}

func (cr *ChatRepository) FindById(ctx context.Context, id uint) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatRepository) Update(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return nil, nil
}

func (cr *ChatRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (cr *ChatRepository) FindByUserId(ctx context.Context, id uint) ([]model.Chat, error) {
	return nil, nil
}

func (cr *ChatRepository) FilterChat(ctx context.Context, chatFilter *model.ChatFilter) ([]model.Chat, error) {
	return nil, nil
}
