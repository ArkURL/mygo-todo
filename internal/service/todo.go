package service

import (
	"context"
	"errors"
	"strings"

	"github.com/arkurl/mygo-todo/internal/model"
	"github.com/arkurl/mygo-todo/internal/repository"
)

var ErrTodoTitleRequired = errors.New("todo title is required")

type CreateTodoInput struct {
	Title   string
	Content string
}

type TodoService interface {
	Create(ctx context.Context, input CreateTodoInput) (*model.Todo, error)
	GetById(ctx context.Context, id int) (*model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Delete(ctx context.Context, id int) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) Create(ctx context.Context, input CreateTodoInput) (*model.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, ErrTodoTitleRequired
	}

	todo := &model.Todo{
		Title:   title,
		Content: strings.TrimSpace(input.Content),
		Done:    false,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) GetById(ctx context.Context, id int) (*model.Todo, error) {
	return s.repo.GetById(ctx, id)
}

func (s *todoService) List(ctx context.Context) ([]model.Todo, error) {
	return s.repo.List(ctx)
}

func (s *todoService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
