//go:build wireinject

package main

import (
	"api/config"
	"api/internal/database"
	"api/internal/handler"
	"api/internal/repository"
	"api/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type App struct {
	DB          *gorm.DB
	Config      *config.Config
	UserHandler *handler.UserHandler
}

func loadConfig() (*config.Config, error) {
	return config.Load("config/config.yaml")
}

func provideMySQLConfig(cfg *config.Config) config.MySQLConfig {
	return cfg.MySQL
}

func provideJWTConfig(cfg *config.Config) *config.JWTConfig {
	return &cfg.JWT
}

func InitializeApp() (*App, error) {
	wire.Build(
		loadConfig,
		provideMySQLConfig,
		provideJWTConfig,
		database.NewMySQL,
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
