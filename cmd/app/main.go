package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/internal/service"
	"todo-list/internal/storage"
)

func main() {
	storage := storage.NewStorage()
	svc := service.NewService(storage)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if checkInput := scanner.Scan(); !checkInput {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		input := strings.Fields(text)

		command := input[0]
		// description := strings.Join(input[2:], "")

		switch command {
		case "add":
			svc.AddTask(idGenerator(storage), input[1], input[2])
		case "list":
			fmt.Println(svc.ListTask())
		case "done":
			omg := input[1]
			id, err := strconv.Atoi(omg)
			if err != nil {
				fmt.Println("oopss")
			}
			svc.DoneTask(id)
		case "delete":
			omg := input[1]
			id, err := strconv.Atoi(omg)
			if err != nil {
				fmt.Println("oopss")
			}
			svc.DeleteTask(id)
		}
	}
}

func idGenerator(s *storage.Storage) int {
	return len(s.Tasks) + 1
}
