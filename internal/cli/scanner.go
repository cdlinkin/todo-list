package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/internal/service"
	"todo-list/internal/storage"
)

func Run() { // demo
	storage := storage.NewStorage()
	services := service.NewService(storage)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if checkInput := scanner.Scan(); !checkInput {
			fmt.Println("Ошибка ввода!") // err
			return
		}

		text := scanner.Text()

		input := strings.Fields(text)

		if len(input) == 0 {
			fmt.Println("пустая команда") // err
			continue
		}

		command := input[0]

		switch command {
		case "add":
			if len(input) < 3 {
				fmt.Println("Недостаточно аргументов") // err
			} else {
				title := input[1]
				description := strings.Join(input[2:], " ")
				services.AddTask(service.IdGenerator(storage), title, description)
			}
		case "list":
			fmt.Println(services.ListTask())
		case "done":
			inputId := input[1]
			id, err := strconv.Atoi(inputId)
			if err != nil {
				fmt.Println("error") // err
			}
			services.DoneTask(id)
		case "delete":
			inputId := input[1]
			id, err := strconv.Atoi(inputId)
			if err != nil {
				fmt.Println("error") // err
			}
			services.DeleteTask(id)
		case "help":
			help()
			continue
		case "exit":
			fmt.Println("спасибо за внимание!")
			return
		default:
			fmt.Printf("Неизвестная команда: %s\n", input) // err
		}
	}
}

func help() {
	fmt.Println("commands to use:")
	fmt.Println("- add {title} {description}: доабвить задачу")
	fmt.Println("- list: просмотреть список задач")
	fmt.Println("- done {id}: Отметить выполнение задачи")
	fmt.Println("- delete {id}: Удалить задачу")
	fmt.Println("- help: Список команд")
	fmt.Println("- exit: Выйти из приложения")
}
