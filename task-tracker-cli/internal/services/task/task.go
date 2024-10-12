package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	taskenums "github.com/anilonayy/go-examples/task-tracker-cli/internal/enums/task"
	"github.com/anilonayy/go-examples/task-tracker-cli/internal/models"
)

const fileName = "tasks.json"

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if string(data) == "" {
		return nil, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, errors.New("error parsing file data")
	}

	return tasks, nil
}

func CreateTasksFile() error {
	_, err := os.Create(fileName)
	if err != nil {
		return errors.New("error creating file: " + err.Error())
	}

	return nil
}

func AddTask(taskName string, tasks []models.Task) error {
	newTask := models.Task{
		ID:     len(tasks) + 1,
		Title:  taskName,
		Status: taskenums.Status(taskenums.StatusToDo),
	}

	tasks = append(tasks, newTask)

	return SaveTasks(tasks)
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.New("error marshalling data")
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return errors.New("error writing data to file")
	}

	return nil
}

func ListTasks(tasks []models.Task, status taskenums.Status) {
	fmt.Println("Tasks:")

	for _, task := range tasks {
		if status != taskenums.StatusUnknown && task.Status != status {
			continue
		}

		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, task.Status.ToString())
	}
}

func GetTask(id int, tasks []models.Task) *models.Task {
	for _, task := range tasks {
		if task.ID == id {
			return &task
		}
	}

	return nil
}

func UpdateTasks(tasks []models.Task, task *models.Task) error {
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = *task
		}
	}

	return SaveTasks(tasks)
}

func DeleteTask(id int, tasks []models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
