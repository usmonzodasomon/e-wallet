package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	wallet "github.com/usmonzodasomon/e-wallet"
	"github.com/usmonzodasomon/e-wallet/db"
	"github.com/usmonzodasomon/e-wallet/pkg/handler"
	"github.com/usmonzodasomon/e-wallet/pkg/repository"
	"github.com/usmonzodasomon/e-wallet/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Println("Error occured while init viper config: ", err)
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err)
		return
	}

	var dbConfig db.Config
	if err := viper.UnmarshalKey("db", &dbConfig); err != nil {
		log.Println("Error unmarshalling config: ", err)
		return
	}
	dbConfig.Password = os.Getenv("DB_PASSWORD")

	db.StartDbConnection(dbConfig)

	if err := db.Migrate(db.GetDBConn()); err != nil {
		log.Println("Error while migrating tables: ", err)
		return
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(wallet.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Println("Error occured while starting server: ", err)
			return
		}
	}()
	log.Println("Server started...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := db.CloseDbConnection(); err != nil {
		log.Println("Error occured on database connection closing: ", err)
		return
	}

	log.Println("Server closed...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println("Error server shutting down: ", err)
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
