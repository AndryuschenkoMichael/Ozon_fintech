package main

import (
	"Ozon_fintech"
	"Ozon_fintech/pkg/handler"
	"Ozon_fintech/pkg/repository"
	"Ozon_fintech/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Error with initialization config file: %s", err.Error())
	}

	repositories := repository.NewRepository()
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
