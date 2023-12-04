package app

import (
	"flag"
	"github.com/IvanKonoplich/shortened_links_service/internal/infrastructure/repository/database_storage"
	in_memory "github.com/IvanKonoplich/shortened_links_service/internal/infrastructure/repository/in-memory"
	"github.com/IvanKonoplich/shortened_links_service/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	DatabasePassword      string = "POSTGRES_PASSWORD"
	Storage                      = "STORAGE"
	PostgresFlagParameter        = "postgres"
	InMemoryFlagParameter        = "in-memory"
)

func ConfigureService() *service.Service {
	getEnvVariables()

	newStorage := checkStorageFlags()

	newService := service.New(newStorage)

	return newService

}

func checkStorageFlags() service.Storage {
	var newStorage service.Storage

	if storageFlag := os.Getenv(Storage); storageFlag != "" {
		switch storageFlag {
		case InMemoryFlagParameter:
			newStorage = configureInMemoryStorage()
		case PostgresFlagParameter:
			newStorage = configurePostgresStorage()
		default:
			log.Fatalf("Задайте корректный параметр хранилища. Допустимые значения: %s, %s. Переданное значение: %s", PostgresFlagParameter, InMemoryFlagParameter, storageFlag)
		}
	} else {

		var (
			storageFlag string
		)
		flag.StringVar(&storageFlag, "s", "", "Какое хранилище использовать. Postgres или in-memory")
		flag.Parse()

		switch storageFlag {
		case InMemoryFlagParameter:
			newStorage = configureInMemoryStorage()
		case PostgresFlagParameter:
			newStorage = configurePostgresStorage()
		default:
			log.Fatalf("Задайте корректный параметр хралища. Допустимые значения: %s, %s. Переданное значение: %s", PostgresFlagParameter, InMemoryFlagParameter, storageFlag)
		}
	}
	return newStorage
}

func configurePostgresStorage() service.Storage {
	configDB := database_storage.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.DBName"),
		Password: os.Getenv(DatabasePassword),
		SSLMode:  viper.GetString("db.SSLMode"),
	}

	postgresDB, err := database_storage.OpenDBConnection(configDB)
	if err != nil {
		logrus.Fatalf("ошибка во время соединерния с БД: %s", err.Error())
	}

	postgresStorage := database_storage.NewDatabaseStorage(postgresDB)

	return postgresStorage
}

func configureInMemoryStorage() service.Storage {
	inMemoryStorage := in_memory.NewInMemoryStorage()

	return inMemoryStorage
}

func getEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf(".env файл не найден: %s", err.Error())
	}
}
