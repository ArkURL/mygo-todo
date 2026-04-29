package handler

import "github.com/arkurl/mygo-todo/internal/model"

type CreateTodoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type TodoResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func ToTodoResponse(todo *model.Todo) TodoResponse {
	return TodoResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Content: todo.Content,
		Done:    todo.Done,
	}
}

func ToTodoResponses(todos []model.Todo) []TodoResponse {
	items := make([]TodoResponse, 0, 10)
	for i := range todos {
		items = append(items, ToTodoResponse(&todos[i]))
	}
	return items
}

type UpdateTodoRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Done    *bool   `json:"done"`
}
