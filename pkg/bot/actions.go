package bot

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/auth"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botAction struct {
	Command     string
	Description string
	Arguments   map[int]string
	Handler     func(emt.Emt, *tgbotapi.Message) (string, error)
}

type botActions struct {
	Actions map[string]*botAction

	Emt emt.Emt

	Auth auth.Auth
}

func NewBotActions(emt emt.Emt, auth auth.Auth) *botActions {
	return &botActions{
		Emt:  emt,
		Auth: auth,
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

func (b *botActions) PerformAction(message *tgbotapi.Message) (string, error) {

	switch message.Command() {

	case "help":
		return b.printAllHelp(), nil

	case "register":
		err := b.Auth.Register(message.From.UserName)
		if err != nil {
			if _, ok := err.(auth.UserAlreadyExists); ok {
				return "", fmt.Errorf("the user %s is already registered in the  system. No need to register again", message.From.UserName)
			}
			return "", err
		}

		log.Printf("username %s added to the repository", message.From.UserName)
		return fmt.Sprintf("Username %s registered sucessfully", message.From.UserName), nil

	default:
		// Check if user is registered
		if err := b.Auth.CheckUser(message.From.UserName); err != nil {
			switch e := err.(type) {
			case auth.UserNotFoundError:
				log.Printf("User %s not authorized", e)
				return "", fmt.Errorf("user %s is not registered in the system. Please user */register* to sign up", message.From.UserName)
			}
		}
		// Check if action exists
		action, ok := b.Actions[message.Command()]
		if !ok {
			return "", fmt.Errorf("command not found")
		}

		// Check argument length is ok
		if action.Arguments != nil && len(strings.Split(message.CommandArguments(), " ")) != len(action.Arguments) {
			return "", fmt.Errorf("number of arguments incorrect. Please check */help*")
		}

		returnedData, err := action.Handler(b.Emt, message)
		if err != nil {
			return "", err
		}

		return returnedData, nil
	}
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
