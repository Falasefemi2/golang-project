package main

import (
	"fmt"
	"os"

	"example.com/todo/list"
)

func main() {
	todoList := &list.TodoList{}

	list.LoadTasks(todoList)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter task description: ")
			var desc string
			fmt.Scanln(&desc)
			list.AddTodo(todoList, desc)
		case 2:
			list.ViewTask(todoList)
		case 3:
			list.ViewTask(todoList)
			fmt.Print("Enter task ID to mark as completed")
			var id int
			fmt.Scanln(&id)
			list.MarkTaskCompleted(todoList, id)
		case 4:
			list.ViewTask(todoList)
			fmt.Print("Enter task ID to remove: ")
			var id int
			fmt.Scanln(&id)
			list.RemoveTask(todoList, id)
		case 5:
			fmt.Println("Exiting...")
			list.Save(todoList)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
