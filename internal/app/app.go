package app

import (
	"github.com/eagle7802/todo-list/config"
	storage2 "github.com/eagle7802/todo-list/internal/storage"
)

func Start() error {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		return err
	}
	storage, err := storage2.NewPostgresDB(cfg)
	if err != nil {
		return nil
	}
	return nil
}
