# Task manager / ToDo-list CLI

This is a pet project. I have tried to implement a scalable architecture.

## Description

The application allows you to:
- add new tasks;
- mark completed tasks;
- remove tasks from the list;
- view the list of current tasks.

 The project is implemented in Golang. Storage is implemented using a slice. Error handling is in place, and bufio.Scanner is used for input processing.

## How to get started

1. Copy the repository to your computer:
```
git clone https://github.com/cdlinkin/todo-list.git
```
2. Commands for using the cli:
```
"- add {title} {description}: add a task"
"- list: view the list of tasks"
"- done {id}: mark a task as completed"
"- delete {id}: delete task"
"- help: the list of commands"
"- exit: exit the CLI"
```
