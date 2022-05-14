package psql

type MessageRepository struct {
	storage *Storage
}

func NewMessageRepository(storage *Storage) *MessageRepository {
	return &MessageRepository{
		storage: storage,
	}
}
