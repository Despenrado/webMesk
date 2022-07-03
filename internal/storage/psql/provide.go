package psql

import (
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/utils"
	"go.uber.org/fx"
)

var Module fx.Option = fx.Provide(ProvidePSQLstorage)

func ProvidePSQLstorage(
	config *utils.Config,
) (storage.Storage, error) {
	db, err := NewConnection(config.PostgreSQL)
	if err != nil {
		return nil, err
	}
	return NewStorage(db), nil
}
