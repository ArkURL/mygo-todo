package model

import "time"

type Todo struct {
	ID        uint
	Title     string
	Content   string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Todo) TableName() string {
	return "todo"
}
