package main

import (
	"errors"
	"fmt"
	"strings"

	"os"
	"strconv"

	"github.com/anilonayy/go-examples/task-tracker-cli/internal/enums/command"
	taskenums "github.com/anilonayy/go-examples/task-tracker-cli/internal/enums/task"
	"github.com/anilonayy/go-examples/task-tracker-cli/internal/models"
	taskservice "github.com/anilonayy/go-examples/task-tracker-cli/internal/services/task"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	fileData, err := initializeTasksFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputCommand := command.ParseCommand(args[0])

	switch inputCommand {
	case command.CommandAdd:
		handleAddTask(args, fileData)
	case command.CommandList:
		handleListTasks(args, fileData)
	case command.CommandUpdate:
		handleUpdateTask(args, fileData)
	case command.CommandDelete:
		handleDeleteTask(args, fileData)
	case
		command.CommandMarkAsTodo,
		command.CommandMarkAsInProg,
		command.CommandMarkAsDone:

		handleMarkStatus(args, fileData)
	default:
		fmt.Println("Invalid command")
	}
}

func initializeTasksFile() ([]models.Task, error) {
	fileData, err := taskservice.GetAllTasks()
	if errors.Is(err, os.ErrNotExist) {
		return nil, taskservice.CreateTasksFile()
	}
	return fileData, err
}

func handleAddTask(args []string, fileData []models.Task) {
	if len(args) < 2 {
		fmt.Println("Please provide a task name")
		return
	}
	if err := taskservice.AddTask(args[1], fileData); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Task `%s` created successfully\n", args[1])
	}
}

func handleListTasks(args []string, fileData []models.Task) {
	var status = taskenums.StatusUnknown
	if len(args) >= 2 {
		status = taskenums.ParseStatus(args[1])
		if status == taskenums.StatusUnknown {
			fmt.Println("Please provide a valid status")
			return
		}
	}
	taskservice.ListTasks(fileData, taskenums.Status(status))
}

func handleUpdateTask(args []string, fileData []models.Task) {
	if len(args) < 3 {
		fmt.Println("Please provide a task ID and title")
		return
	}
	taskID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	task := taskservice.GetTask(taskID, fileData)
	if task == nil {
		fmt.Println("Task not found")
		return
	}

	task.Title = args[2]

	if err := taskservice.UpdateTasks(fileData, task); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("UPDATED: id:%d title:`%s` updated successfully\n", taskID, args[2])
	}
}

func handleDeleteTask(args []string, fileData []models.Task) {
	if len(args) < 2 {
		fmt.Println("Please provide a task ID")
		return
	}

	taskID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	if err := taskservice.DeleteTask(taskID, fileData); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Task with ID %d deleted successfully\n", taskID)
	}
}

func handleMarkStatus(args []string, fileData []models.Task) {
	if len(args) < 2 {
		fmt.Println("Please provide a task ID and status")
		return
	}

	taskID, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	task := taskservice.GetTask(taskID, fileData)
	if task == nil {
		fmt.Println("Task not found")
		return
	}

	operation := strings.Split(args[0], "mark-")[1]

	status := taskenums.ParseStatus(operation)
	if status == taskenums.StatusUnknown {
		fmt.Println("Please provide a valid status")
		return
	}

	task.Status = status

	if err := taskservice.UpdateTasks(fileData, task); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Task with ID %d marked as %s\n", taskID, status.ToString())
	}
}
