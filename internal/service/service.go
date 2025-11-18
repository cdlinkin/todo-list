package service

import (
	"time"
	"todo-list/internal/models"
	"todo-list/internal/storage"
)

type Service struct {
	Storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Storage: storage,
	}
}

func (s *Service) AddTask(id int, title, description string) {
	task := models.NewTask(id, title, description)

	s.Storage.Tasks = append(s.Storage.Tasks, task)
}

func (s *Service) ListTask() *[]models.Task {
	return &s.Storage.Tasks
}

func (s *Service) DoneTask(id int) {
	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			done := time.Now()

			s.Storage.Tasks[i].IsDone = true
			s.Storage.Tasks[i].DoneAt = &done
		}
	}
}

func (s *Service) DeleteTask(id int) {
	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			s.Storage.Tasks = append(s.Storage.Tasks[:i], s.Storage.Tasks[i+1:]...)
			break
		}
	}
}
