package bot

import (
	"fmt"
	"sort"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

type botAction struct {
	Command     string
	Description string
	Arguments   map[int]string
	Handler     func(emt.Emt, []string) (string, error)
}

type botActionsMap map[string]*botAction

func NewBotActions() botActionsMap {
	return map[string]*botAction{
		"bus_waiting_times": &botAction{
			Command:     "bus_waiting_times",
			Description: "This command returns all bus waiting times given a stop ID.",
			Arguments: map[int]string{
				1: "Bus stop ID to query",
			},
			Handler: GetAllBusWaitingTimes,
		},
	}
}

func (b botActionsMap) PerformAction(command string, args []string) (string, error) {
	// Check if command was help
	if command == "help" {
		return b.PrintAllHelp(), nil
	}
	// Check if action exists
	action, ok := b[command]
	if !ok {
		return "", fmt.Errorf("command not found")
	}

	// Check argument length is ok
	if len(args) != len(action.Arguments) {
		return "", fmt.Errorf("number of arguments incorrect. Please check help")
	}

	// TO-DO

	return "", nil
}

func (b botActionsMap) PrintAllHelp() string {
	var help string
	help += "Help from EMT Telegram Bot\n---\n"

	keys := make([]string, 0, len(b))
	for k := range b {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		help += printHelp(b[k]) + "\n---"
	}

	return help
}

func printHelp(action *botAction) string {
	var temp string
	for i := 0; i < len(action.Arguments); i++ {
		temp += fmt.Sprintf("\t[%d] - %s", i+1, action.Arguments[i+1])
	}

	return fmt.Sprintf("Command: %s\n" +
		"Description: %s\n" +
		"Arguments:\n" + temp)
}
