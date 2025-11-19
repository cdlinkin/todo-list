package service

import (
	"fmt"
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

func (s *Service) ListTask() {
	taskDone := "[✓]"
	taskUnDone := "[ ]"
	var status string
	for _, task := range s.Storage.Tasks {
		if task.IsDone {
			status = taskDone
		} else {
			status = taskUnDone
		}

		fmt.Printf("%s | %d |Задача: %s | Описание: %s | Создана: %s |\n", status, task.ID, task.Title, task.Description, task.CreatedAt.Format("2006-01-02 15:04"))
		if task.DoneAt != nil {
			fmt.Printf(" Время выполнения: %s |\n", task.DoneAt.Format("2006-01-02 15:04"))
		}
	}

}

func (s *Service) DoneTask(id int) error {
	found := false
	for i := range s.Storage.Tasks {

		if id == s.Storage.Tasks[i].ID {

			found = true

			done := time.Now()
			s.Storage.Tasks[i].IsDone = true
			s.Storage.Tasks[i].DoneAt = &done
		}

		if !found {
			return errors.ErrTaskIsNotFound
		}
	}

	return nil
}

func (s *Service) DeleteTask(id int) error {
	found := false
	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			s.Storage.Tasks = append(s.Storage.Tasks[:i], s.Storage.Tasks[i+1:]...)

			found = true
			break
		}
	}
	if !found {
		return errors.ErrTaskIsNotFound
	}

	return nil
}

func IdGenerator(s *storage.Storage) int {
	return len(s.Tasks) + 1
}
