package psql

type ChatRepository struct {
	storage *Storage
}

func NewChatRepository(storage *Storage) *ChatRepository {
	return &ChatRepository{
		storage: storage,
	}
}
