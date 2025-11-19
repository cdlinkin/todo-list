package models

import (
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(id int, title, description string) Task {
	return Task{
		ID: id,

		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
