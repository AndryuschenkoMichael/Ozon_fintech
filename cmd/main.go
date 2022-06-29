package main

import (
	"Ozon_fintech"
	"Ozon_fintech/pkg/handler"
	"Ozon_fintech/pkg/repository"
	"Ozon_fintech/pkg/service"
	"Ozon_fintech/pkg/storage"
	strgen "Ozon_fintech/pkg/string_generator"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Error with initialization config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error with initializing environment file: %s", err.Error())
	}

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//})
	//
	//if err != nil {
	//	log.Fatalf("Failed to initialize db: %s", err.Error())
	//}

	generators := strgen.NewStringGeneratorRandom(service.LengthLink)
	storages := storage.New(generators.GenerateString)
	repositories := repository.NewRepositoryStorage(storages)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(Ozon_fintech.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error in running server: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
