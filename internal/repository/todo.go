package repository

import (
	"context"

	"github.com/arkurl/mygo-todo/internal/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	GetById(ctx context.Context, id int) (*model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Update(ctx context.Context, task *model.Todo) error
	Delete(ctx context.Context, id int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepository) GetById(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) List(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.WithContext(ctx).Order("id desc").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepository) Update(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}
