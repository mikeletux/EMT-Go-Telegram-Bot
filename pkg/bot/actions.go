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

type botActions struct {
	Actions map[string]*botAction

	Emt emt.Emt
}

func NewBotActions(emt emt.Emt) *botActions {
	return &botActions{
		Emt: emt,
		Actions: map[string]*botAction{
			"about": {
				Command:     "about",
				Description: "About EMT Telegram bot",
				Handler:     GetAbout,
			},
			"bus_waiting_times": {
				Command:     "bus_waiting_times",
				Description: "This command returns all bus waiting times given a stop ID.",
				Arguments: map[int]string{
					1: "Bus stop ID to query",
				},
				Handler: GetAllBusWaitingTimes,
			},
		},
	}
}

func (b *botActions) PerformAction(command string, args []string) (string, error) {
	// Check if command was help
	if command == "help" {
		return b.printAllHelp(), nil
	}

	// TO-DO write about

	// Check if action exists
	action, ok := b.Actions[command]
	if !ok {
		return "", fmt.Errorf("command not found")
	}

	// Check argument length is ok
	if action.Arguments != nil && len(args) != len(action.Arguments) {
		return "", fmt.Errorf("number of arguments incorrect. Please check help")
	}

	returnedData, err := action.Handler(b.Emt, args)
	if err != nil {
		return "", err
	}

	return returnedData, nil
}

func (b *botActions) printAllHelp() string {
	var help string
	help += fmt.Sprintf("%s Help from EMT Telegram Bot %s\n\n", busSideEmoji, busSideEmoji)

	keys := make([]string, 0, len(b.Actions))
	for k := range b.Actions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		help += printHelp(b.Actions[k]) + "\n"
	}

	return help
}

func printHelp(action *botAction) string {
	var temp string
	if len(action.Arguments) > 0 {
		for i := 0; i < len(action.Arguments); i++ {
			temp += fmt.Sprintf("\t - \\[%d] %s", i+1, action.Arguments[i+1])
			if i < len(action.Arguments)-1 {
				temp += "\n"
			}
		}
	} else {
		temp = "None\n"
	}

	return fmt.Sprintf("Command: */%s*\n"+
		" - *Description*: %s\n"+
		" - *Arguments*:\n %s", action.Command, action.Description, temp)
}
