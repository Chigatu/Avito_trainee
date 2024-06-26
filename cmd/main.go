package main

import (
	"github.com/higatu/todo-app"
	"github.com/higatu/todo-app/pkg/handler"
	"github.com/higatu/todo-app/pkg/repository"
	"github.com/higatu/todo-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host: viper.GetString("db.host"), 
		Port: viper.GetString("db.port"), 
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname: viper.GetString("db.dbname"), 
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil{
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
