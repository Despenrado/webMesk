package impl_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Despenrado/webMesk/internal/service/impl"
	"github.com/Despenrado/webMesk/internal/storage/psql"
	"github.com/Despenrado/webMesk/internal/utils"
	"github.com/bitcomplete/sqltestutil"
)

var storageInt *psql.Storage
var serviceInt *impl.Service

// func TestMain(m *testing.M) {
// 	os.Setenv("TZ", "UTC")
// 	config := &utils.Config{
// 		PostgreSQL: &utils.PostgreSQLConfig{
// 			Master: struct {
// 				Host          string "yaml:\"host\""
// 				Port          string "yaml:\"port\""
// 				User          string "yaml:\"user\""
// 				DBName        string "yaml:\"dbname\""
// 				Password      string "yaml:\"password\""
// 				SSLMode       string "yaml:\"sslmode\""
// 				TimeZone      string "yaml:\"timeZone\""
// 				AutoMigration bool   "yaml:\"autoMigration\""
// 			}{
// 				Host:          "localhost",
// 				Port:          "5432",
// 				User:          "gorm",
// 				Password:      "gorm",
// 				DBName:        "webmesk",
// 				SSLMode:       "disable",
// 				TimeZone:      "Europe/Warsaw",
// 				AutoMigration: true,
// 			},
// 		},
// 	}
// 	db, err := psql.NewConnection(config.PostgreSQL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	psqlStorage := psql.NewStorage(db)
// 	storageInt = psqlStorage
// 	m.Run()
// 	defer func() {
// 		// psqlStorage.DB.Migrator().DropTable(&model.User{}, &model.Chat{}, &model.Message{}, "user_chat")
// 	}()
// }

func TestMain(m *testing.M) {
	os.Setenv("TZ", "UTC")
	ctx := context.Background()
	pg, _ := sqltestutil.StartPostgresContainer(ctx, "12")
	fmt.Println(pg.ConnectionString())
	defer pg.Shutdown(ctx)
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
				Host:          "127.0.0.1",
				Port:          strings.Split(strings.Split(pg.ConnectionString(), ":")[3], "/")[0],
				User:          "pgtest",
				Password:      strings.Split(strings.Split(pg.ConnectionString(), ":")[2], "@")[0],
				DBName:        "pgtest",
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
	serviceInt = impl.NewService(
		storageInt,
		nil,
		impl.NewUserService(storageInt),
		impl.NewChatService(storageInt),
		impl.NewMessageService(storageInt),
		nil,
	)
	m.Run()
}
