package main

import (
	"flag"
	"log"

	"github.com/Despenrado/webMesk/internal/storage"
	"github.com/Despenrado/webMesk/internal/storage/psql"
	"github.com/Despenrado/webMesk/internal/utils"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/default.yaml", "path to config file")
}

func main() {
	defer log.Println("Server stopped")
	flag.Parse()
	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(config.PostgreSQL.PSQLToString())
	_, err = initDataBase(&config.PostgreSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func initDataBase(config *utils.PostgreSQL) (storage.Storage, error) {
	db, err := psql.NewConnection(config.PSQLToString(), config.AutoMigration)
	storage := psql.NewStorage(db)
	return storage, err
}
