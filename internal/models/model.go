package models

import (
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	IsDone      bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(id int, title, description string) Task {
	return Task{
		ID: id,

		Title:       title,
		Description: description,
		IsDone:      false,

		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}

func (t *Task) Done() {
	done := time.Now()

	t.IsDone = true
	t.DoneAt = &done
}
