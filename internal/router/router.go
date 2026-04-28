package router

import (
	"github.com/arkurl/mygo-todo/internal/handler"
	"github.com/gin-gonic/gin"
)

func New(todoHandler *handler.TodoHandler) *gin.Engine {
	r := gin.Default()

	RegisterTodoRoutes(r, todoHandler)

	return r
}

func RegisterTodoRoutes(r gin.IRouter, h *handler.TodoHandler) {
	todos := r.Group("/todos")

	todos.POST("", h.Create)
	todos.GET("", h.List)
	todos.GET("/:id", h.GetById)
	todos.DELETE("/:id", h.Delete)
}
