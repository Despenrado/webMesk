package psql_test

import (
	"os"
	"testing"

	"github.com/Despenrado/webMesk/internal/storage/psql"
	"github.com/Despenrado/webMesk/internal/utils"
)

var storageInt *psql.Storage

func TestMain(m *testing.M) {
	os.Setenv("TZ", "UTC")
	config := &utils.Config{
		PostgreSQL: &utils.PostgreSQLConfig{
			Master: struct {
				Host          string "yaml:\"host\""
				Port          string "yaml:\"port\""
				User          string "yaml:\"user\""
				DBName        string "yaml:\"dbname\""
				Password      string "yaml:\"password\""
				SSLMode       string "yaml:\"sslmode\""
				TimeZone      string "yaml:\"timeZone\""
				AutoMigration bool   "yaml:\"autoMigration\""
			}{
				Host:          "localhost",
				Port:          "5432",
				User:          "gorm",
				Password:      "gorm",
				DBName:        "webmesk",
				SSLMode:       "disable",
				TimeZone:      "Europe/Warsaw",
				AutoMigration: true,
			},
		},
	}
	db, err := psql.NewConnection(config.PostgreSQL)
	if err != nil {
		panic(err)
	}
	psqlStorage := psql.NewStorage(db)
	storageInt = psqlStorage
	m.Run()
	defer func() {
		// psqlStorage.DB.Migrator().DropTable(&model.User{}, &model.Chat{}, &model.Message{}, "user_chat")
	}()
}
