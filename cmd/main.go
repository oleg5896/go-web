package main

import (
	"os"

	"github.com/joho/godotenv"
	goweb "github.com/oleg5896/go-web"
	"github.com/oleg5896/go-web/pkg/handler"
	"github.com/oleg5896/go-web/pkg/repository"
	"github.com/oleg5896/go-web/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error init Config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Printf("err env: %s", err.Error())
	}

	db, err := repository.NewPgDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(goweb.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error start server: " + err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
