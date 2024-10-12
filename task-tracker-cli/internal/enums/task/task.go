package task

import (
	"strings"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusDone
	StatusInProgress
	StatusToDo
)

var statuses = map[string]Status{
	"todo":        StatusToDo,
	"done":        StatusDone,
	"in-progress": StatusInProgress,
	"unknown":     StatusUnknown,
}

func (s Status) ToString() string {
	for k, v := range statuses {
		if v == s {
			return k
		}
	}

	return "unknown"
}

func ParseStatus(input string) Status {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "todo":
		return StatusToDo
	case "done":
		return StatusDone
	case "in-progress":
		return StatusInProgress
	default:
		return StatusUnknown
	}
}
