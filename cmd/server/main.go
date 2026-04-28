package main

import (
	"log"

	"github.com/arkurl/mygo-todo/internal/config"
	"github.com/arkurl/mygo-todo/internal/database"
	"github.com/arkurl/mygo-todo/internal/handler"
	"github.com/arkurl/mygo-todo/internal/logger"
	"github.com/arkurl/mygo-todo/internal/model"
	"github.com/arkurl/mygo-todo/internal/repository"
	"github.com/arkurl/mygo-todo/internal/router"
	"github.com/arkurl/mygo-todo/internal/service"
	"go.uber.org/zap"
)

func main() {
	config.Init()

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	db, err := database.NewPostgres(config.Conf.Database.DSN())

	if err != nil {
		logger.L().Fatal("connect database failed", zap.Error(err))
	}
	defer database.Close(db)

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		logger.L().Fatal("migrate database failed", zap.Error(err))
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r := router.New(todoHandler)

	if err := r.Run(config.Conf.Server.PORT()); err != nil {
		logger.L().Fatal("server stopped with error", zap.Error(err))
	}

}
