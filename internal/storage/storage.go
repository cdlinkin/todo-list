package storage

import "todo-list/internal/models"

type Storage struct {
	Tasks []models.Task
}

func NewStorage() *Storage {
	return &Storage{
		Tasks: make([]models.Task, 0),
	}
}
