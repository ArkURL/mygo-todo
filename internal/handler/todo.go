package handler

import (
	"net/http"
	"strconv"

	"github.com/arkurl/mygo-todo/internal/response"
	"github.com/arkurl/mygo-todo/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

type CreateTodoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, "invalid request")
		return
	}

	_, err := h.service.Create(c.Request.Context(), service.CreateTodoInput{
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Created(c, gin.H{})
}

func (h *TodoHandler) List(c *gin.Context) {
	todos, err := h.service.List((c.Request.Context()))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, "internal error")
		return
	}

	response.Success(c, gin.H{
		"todos": todos,
	})
}

func (h *TodoHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, "invalid request")
		return
	}

	todo, err := h.service.GetById(c.Request.Context(), id)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"todo": todo,
	})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeInvalidRequest, "invalid request")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{})
}
