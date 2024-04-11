package list

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoList struct {
	Todo []Task `json:"tasks"`
}

func AddTodo(todo *TodoList, desc string) {
	taskID := len(todo.Todo) + 1
	newTask := Task{ID: taskID, Description: desc, Completed: false}
	todo.Todo = append(todo.Todo, newTask)
	fmt.Println("Task added successfully")
}

func ViewTask(todo *TodoList) {
	fmt.Println("Tasks:")
	for _, todo := range todo.Todo {
		status := "Incomplete"
		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", todo.ID, todo.Description, status)
	}
}

func MarkTaskCompleted(todo *TodoList, id int) {
	for i, task := range todo.Todo {
		if task.ID == id {
			todo.Todo[i].Completed = true
			fmt.Println("Task marked as complete")
			return
		}
		fmt.Println("Task not found")
	}
}

func RemoveTask(todo *TodoList, id int) {
	for i, task := range todo.Todo {
		if task.ID == id {
			todo.Todo = append(todo.Todo[:i], todo.Todo[i+1:]...)
			fmt.Println("Task removed")
			return
		}
	}
	fmt.Println("Task not found")
}

func Save(todo *TodoList) error {
	file, err := json.MarshalIndent(todo, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", file, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Tasks saved successfully")
	return nil
}

func LoadTasks(todo *TodoList) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, todo)
}
