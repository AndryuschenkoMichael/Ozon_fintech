package main

import (
	"Ozon_fintech"
	"Ozon_fintech/pkg/handler"
	"Ozon_fintech/pkg/repository"
	"Ozon_fintech/pkg/service"
	"Ozon_fintech/pkg/storage"
	strgen "Ozon_fintech/pkg/string_generator"
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

var dbType = flag.String("db_type", "STORAGE", "type of db")

func main() {
	flag.Parse()

	if err := initConfigs(); err != nil {
		log.Fatalf("Error with initialization config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error with initializing environment file: %s", err.Error())
	}

	generators := strgen.NewStringGeneratorRandom(service.LengthLink)
	var repositories *repository.Repository

	switch *dbType {
	case "STORAGE":
		storages := storage.New(generators.GenerateString)
		repositories = repository.NewRepositoryStorage(storages)

	case "POSTGRES":
		db, err := repository.NewPostgresDB(repository.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		})

		if err != nil {
			log.Fatalf("Failed to initialize db: %s", err.Error())
		}

		repositories = repository.NewRepositoryPostgres(db, generators)

	default:
		log.Fatalf("undefined flag detected: %s", *dbType)
	}

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
