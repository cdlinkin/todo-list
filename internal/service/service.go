package service

import (
	"time"
	"todo-list/internal/errors"
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

func (s *Service) AddTask(id int, title, description string) error {
	if title == "" {
		return errors.ErrTitleIsEmpty
	}

	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			return errors.ErrTaskAlreadyExists
		}
	}

	task := models.NewTask(id, title, description)

	s.Storage.Tasks = append(s.Storage.Tasks, task)
	return nil
}

func (s *Service) ListTask() *[]models.Task {
	return &s.Storage.Tasks
}

func (s *Service) DoneTask(id int) error {
	for i := range s.Storage.Tasks {

		if id == s.Storage.Tasks[i].ID {

			if s.Storage.Tasks[i].IsDone == true {
				return errors.ErrTaskAlreadyDone
			}

			done := time.Now()
			s.Storage.Tasks[i].IsDone = true
			s.Storage.Tasks[i].DoneAt = &done

			return nil
		}
	}
	return nil
}

func (s *Service) DeleteTask(id int) error {
	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			s.Storage.Tasks = append(s.Storage.Tasks[:i], s.Storage.Tasks[i+1:]...)
			break
		}
	}

	return errors.ErrTaskIsNotFound
}

func IdGenerator(s *storage.Storage) int {
	return len(s.Tasks) + 1
}
