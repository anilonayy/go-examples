package command

import (
	"strings"
)

type Command int

const (
	CommandUnknown Command = iota
	CommandAdd
	CommandList
	CommandUpdate
	CommandDelete
	CommandMarkAsTodo
	CommandMarkAsInProg
	CommandMarkAsDone
)

var commands = map[string]Command{
	"add":              CommandAdd,
	"list":             CommandList,
	"update":           CommandUpdate,
	"delete":           CommandDelete,
	"mark-todo":        CommandMarkAsTodo,
	"mark-in-progress": CommandMarkAsInProg,
	"mark-done":        CommandMarkAsDone,
	"unknown":          CommandUnknown,
}

func (s Command) ToString() string {
	for k, v := range commands {
		if v == s {
			return k
		}
	}

	return "unknown"
}

func ParseCommand(input string) Command {
	if cmd, exists := commands[strings.ToLower(strings.TrimSpace(input))]; exists {
		return cmd
	}

	return CommandUnknown
}
