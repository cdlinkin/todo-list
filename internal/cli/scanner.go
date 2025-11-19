package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/internal/errors"
	"todo-list/internal/service"
	"todo-list/internal/storage"
)

func Run() error {
	storage := storage.NewStorage()
	services := service.NewService(storage)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if checkInput := scanner.Scan(); !checkInput {
			fmt.Println(errors.ErrInput)
		}

		text := scanner.Text()
		input := strings.Fields(text)

		if len(input) == 0 {
			fmt.Println(errors.ErrInput)
			continue
		}

		command := input[0]

		switch command {
		case "add":
			if len(input) < 3 {
				fmt.Println(errors.ErrNotEnoughAruments)
				continue
			}

			title := input[1]
			description := strings.Join(input[2:], " ")
			services.AddTask(service.IdGenerator(storage), title, description)
		case "list":
			if len(storage.Tasks) == 0 {
				fmt.Println(errors.ErrListIsEmpty)
				continue
			}

			services.ListTask()
		case "done":
			if len(input) != 2 {
				fmt.Println(errors.ErrNotEnoughAruments)
				continue
			}

			inputId := input[1]
			id, err := strconv.Atoi(inputId)
			if err != nil {
				fmt.Println(errors.ErrConvertStrconv)
			}
			services.DoneTask(id)
		case "delete":
			if len(input) != 2 {
				fmt.Println(errors.ErrNotEnoughAruments)
				continue
			}

			inputId := input[1]
			id, err := strconv.Atoi(inputId)
			if err != nil {
				fmt.Println(errors.ErrConvertStrconv)
			}
			services.DeleteTask(id)
		case "help":
			help()
		case "exit":
			fmt.Println("спасибо за внимание!")
			return nil
		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
		}
	}
}

func help() {
	fmt.Println("commands to use:")
	fmt.Println("- add {title} {description}: Добавить задачу")
	fmt.Println("- list: Просмотреть список задач")
	fmt.Println("- done {id}: Отметить выполнение задачи")
	fmt.Println("- delete {id}: Удалить задачу")
	fmt.Println("- help: Список команд")
	fmt.Println("- exit: Выйти из приложения")
}
