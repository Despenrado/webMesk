package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Despenrado/webMesk/internal/service/impl"
	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/storage/psql"
	"github.com/Despenrado/webMesk/internal/transport/restapi"
	"github.com/Despenrado/webMesk/internal/utils"
	pkgutils "github.com/Despenrado/webMesk/pkg/utils"
)

var (
	configPath string
)

func init() {
	os.Setenv("TZ", "UTC")
	flag.StringVar(&configPath, "config", "configs/default.yaml", "path to config file")
}

func main() {
	flag.Parse()

	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Panicln(err.Error())
	}

	ctx := context.TODO()
	logger := pkgutils.NewLogger()

	db, err := initDataBase(&config.PostgreSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// db := teststore.NewStorage()
	service := impl.NewService(db)

	uh := restapi.NewUserHandler(service)
	ch := restapi.NewChatHandler(service)
	mh := restapi.NewMessageHandler(service)
	ah := restapi.NewAuthHandler(service)

	srv := restapi.NewServer(ctx, config.RestAPIServer.Port, nil, logger, uh, mh, ch, ah)
	srv.InitDefaultEndpoints("restapi")
	err = srv.Run()
	if err != nil {
		log.Panicln(err)
	}
}

func initDataBase(config *utils.PostgreSQL) (storage.Storage, error) {
	db, err := psql.NewConnection(config.PSQLToString(), config.AutoMigration)
	storage := psql.NewStorage(db)
	return storage, err
}
