package main

import (
	"log"

	"github.com/arkurl/mygo-todo/internal/config"
	"github.com/arkurl/mygo-todo/internal/database"
	"github.com/arkurl/mygo-todo/internal/handler"
	"github.com/arkurl/mygo-todo/internal/model"
	"github.com/arkurl/mygo-todo/internal/repository"
	"github.com/arkurl/mygo-todo/internal/router"
	"github.com/arkurl/mygo-todo/internal/service"
)

func main() {
	config.Init()

	db, err := database.NewPostgres(config.Conf.Database.DSN())

	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(db)

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal(err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r := router.New(todoHandler)

	if err := r.Run(config.Conf.Server.PORT()); err != nil {
		log.Fatal(err)
	}

}
