package main

import (
	"flag"
	"os"

	"github.com/Despenrado/webMesk/internal/service/impl"
	"github.com/Despenrado/webMesk/internal/storage/psql"
	"github.com/Despenrado/webMesk/internal/storage/redis"
	"github.com/Despenrado/webMesk/internal/transport/restapi"
	"github.com/Despenrado/webMesk/internal/utils"
	pkgutils "github.com/Despenrado/webMesk/pkg/utils"

	"go.uber.org/fx"
)

func init() {
	os.Setenv("TZ", "UTC")
	flag.StringVar(&utils.ConfigPath, "config", "configs/default.yaml", "path to config file")
}

func main() {
	flag.Parse()
	fx.New(
		fx.Provide(
			pkgutils.NewLogger,
			impl.NewUserService,
			impl.NewChatService,
			impl.NewMessageService,
			impl.NewAuthService,
			restapi.NewUserHandler,
			restapi.NewChatHandler,
			restapi.NewMessageHandler,
			restapi.NewAuthHandler,
			utils.LoadConfig,
			restapi.NewServer,
		),
		redis.Module,
		psql.Module,
		impl.Module,
		fx.Invoke(
			restapi.RegisterHundlers,
		),
	).Run()

}

// func initDataBase(config *utils.PostgreSQL) (storage.Storage, error) {
// 	db, err := psql.NewConnection(config.PSQLToString(), config.AutoMigration)
// 	storage := psql.NewStorage(db)
// 	return storage, err
// }

// func mainStandard() {
// 	flag.Parse()

// 	config, err := utils.LoadConfig()
// 	if err != nil {
// 		log.Panicln(err.Error())
// 	}

// 	ctx := context.TODO()
// 	logger := pkgutils.NewLogger()

// 	db, err := initDataBase(&config.PostgreSQL)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	// db := teststore.NewStorage()
// 	service := impl.NewService(db)

// 	uh := restapi.NewUserHandler(service)
// 	ch := restapi.NewChatHandler(service)
// 	mh := restapi.NewMessageHandler(service)
// 	ah := restapi.NewAuthHandler(service)

// 	srv := restapi.NewServer(ctx, config.RestAPIServer.Port, nil, logger, uh, mh, ch, ah)
// 	restapi.InitDefaultEndpoints(srv)
// 	err = srv.Run()
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// }
