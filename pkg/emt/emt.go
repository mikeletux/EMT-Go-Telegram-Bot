package emt

import (
	"fmt"
	"strings"

	"github.com/mikeletux/goemt"
)

type Emt interface {
	PerformAction(command string, arguments string) (string, error)
	Logout() error
}

type goEMT struct {
	// Client is the struct that will hold the EMT client
	Client *goemt.APIClient

	// Handlers is where all action handlers will be defined
	Handlers map[string]Handler
}

func NewGoEMT(config goemt.ClientConfig) (Emt, error) {
	goEMT := goEMT{}
	c, err := goemt.Connect(config)
	if err != nil {
		return nil, err
	}
	goEMT.Client = c
	goEMT.Handlers = GetAllHandlers()

	return &goEMT, nil
}

func (e *goEMT) PerformAction(command string, arguments string) (string, error) {
	// Check if help
	if command == "help" {
		return "Help for EMT Telegram bot:", nil
	}

	// Get handler
	handler, ok := e.Handlers[command]
	if !ok {
		return "", fmt.Errorf("the option you chose doesn't exist")
	}

	// Split arguments string
	args := strings.Split(arguments, " ")

	return handler.Handler(e.Client, args)
}

func (e *goEMT) Logout() error {
	e.Client.Logout()
	return nil
}
