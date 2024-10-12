package command

import (
	"strings"
)

type Command int

const (
	CommandUnknown Command = iota
	CommandAdd
	CommandList
	CommandDelete
	CommandSummary
	CommandClear
)

var commands = map[string]Command{
	"add":     CommandAdd,
	"list":    CommandList,
	"delete":  CommandDelete,
	"summary": CommandSummary,
	"clear":   CommandClear,
	"unknown": CommandUnknown,
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
