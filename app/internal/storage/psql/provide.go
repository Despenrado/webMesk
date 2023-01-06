package psql

import (
	"time"

	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	"go.uber.org/fx"
)

var Module fx.Option = fx.Provide(ProvidePSQLStorage)

func ProvidePSQLStorage(
	config *utils.Config,
) (storage.Storage, error) {
	db, err := NewConnection(config.PostgreSQL)
	if err != nil {
		for err != nil {
			time.Sleep(5 * time.Second)
			db, err = NewConnection(config.PostgreSQL)
		}
	}
	return NewStorage(db), nil
}
