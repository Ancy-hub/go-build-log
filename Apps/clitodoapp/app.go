package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("CLI-TO-DO-APP")
	fileName := "tasks.json"
	tasks := loadTasks(fileName)

	for {
		fmt.Println(`
		1. Add Task
		2. View Tasks
		3. Mark Done
		4. Exit

		Enter your choice: 
	`)
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			newTask := addTask(tasks)
			tasks = append(tasks, newTask)
			saveTasks(tasks, fileName)
		case 2:
			viewTasks(tasks)
		case 3:
			var id int
			fmt.Print("Enter the task id to be marked done: ")
			fmt.Scanln(&id)
			if markDone(tasks, id) {
				saveTasks(tasks, fileName)
			}
		case 4:
			saveTasks(tasks, fileName)
			return
		default:
			fmt.Println("Enter a valid choice")
		}
	}
}


type Task struct {
	ID    int
	Title string
	Done  bool
}



func loadTasks(file string) []Task {
	data, err := os.ReadFile(file)
	if err != nil {
		return []Task{}
	}
	var tasksList []Task
	if err := json.Unmarshal(data, &tasksList); err != nil {
		return []Task{}
	}
	return tasksList
}

func addTask(tasks []Task) Task {
	fmt.Println("Lets add a task!")
	var newTask Task
	newTask.ID = getNextTaskID(tasks)
	fmt.Print("Enter the title of the task: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	newTask.Title = strings.TrimSpace(title)
	newTask.Done = false
	return newTask
}

func getNextTaskID(tasks []Task) int {
	max := 0
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}
	return max + 1
}

func viewTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✔]"
		}
		fmt.Printf("%d %s %s\n", task.ID, status, task.Title)
	}
}

func markDone(tasks []Task, id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			fmt.Printf("Task '%s' marked as complete!\n", tasks[i].Title)
			return true
		}
	}
	fmt.Println("No such task found!")
	return false
}

func saveTasks(tasks []Task, filename string) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	if err := os.WriteFile(filename, data, 0600); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}
