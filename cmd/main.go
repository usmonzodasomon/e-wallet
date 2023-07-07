package main

import (
	"os"

	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/usmonzodasomon/e-wallet/db"
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
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
