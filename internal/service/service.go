package service

import (
	"fmt"
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
	if title == "" {
		fmt.Println("title пуст") // err
	}

	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			fmt.Println("такой id уже существует") // err
		}
	}

	task := models.NewTask(id, title, description)

	s.Storage.Tasks = append(s.Storage.Tasks, task)
}

func (s *Service) ListTask() *[]models.Task {
	return &s.Storage.Tasks
}

func (s *Service) DoneTask(id int) error {
	for i := range s.Storage.Tasks {
		if s.Storage.Tasks[i].IsDone == true {
			return fmt.Errorf("задача уже выполнена") // err
		}

		if id == s.Storage.Tasks[i].ID {
			done := time.Now()
			s.Storage.Tasks[i].IsDone = true
			s.Storage.Tasks[i].DoneAt = &done

			return nil
		}
	}
	return fmt.Errorf("задача уже выполнена")
}

func (s *Service) DeleteTask(id int) {
	for i := range s.Storage.Tasks {
		if id == s.Storage.Tasks[i].ID {
			s.Storage.Tasks = append(s.Storage.Tasks[:i], s.Storage.Tasks[i+1:]...)
			break
		}
	}

	fmt.Println("задача не найдена") // err
}

func IdGenerator(s *storage.Storage) int {
	return len(s.Tasks) + 1
}
