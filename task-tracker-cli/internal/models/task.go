package models

import (
	"github.com/anilonayy/go-examples/task-tracker-cli/internal/enums/task"
)

type Task struct {
	ID     int         `json:"id"`
	Title  string      `json:"title"`
	Status task.Status `json:"status"`
}
