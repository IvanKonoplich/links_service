package main

import (
	"github.com/IvanKonoplich/shortened_links_service/internal/app"
	"github.com/IvanKonoplich/shortened_links_service/pkg/shortened_links"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка во время чтения конфига:%s", err.Error())
	}

	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		logrus.Fatalf("Ошибка во время старта сервиса: %s", err.Error())
	}

	service := app.ConfigureService()
	grpcServer := grpc.NewServer()
	shortened_links.RegisterShortenedLinksServer(grpcServer, service)

	logrus.Info("Запуск сервера...")

	err = grpcServer.Serve(listener)
	if err != nil {
		logrus.Fatalf("Произошла критическая ошибка во время работы сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetDefault("port", "8000")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")

	return viper.ReadInConfig()
}
