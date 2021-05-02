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

	// config is kept in case bot needs to reconnect to EMT services because of token expiration
	config goemt.ClientConfig
}

func NewGoEMT(config goemt.ClientConfig) (Emt, error) {
	goEMT := goEMT{}
	c, err := goemt.Connect(config)
	if err != nil {
		return nil, err
	}
	goEMT.Client = c
	goEMT.Handlers = GetAllHandlers()
	goEMT.config = config

	return &goEMT, nil
}

func (e *goEMT) reconnect() error {
	c, err := goemt.Connect(e.config)
	if err != nil {
		return err
	}
	e.Client = c
	return nil
}

func (e *goEMT) PerformAction(command string, arguments string) (string, error) {
	// Check if help is chosen
	if command == "help" {
		helpStr := fmt.Sprint("Help from EMT Telegram bot\n--------\n")
		for k, v := range e.Handlers {
			helpStr += fmt.Sprintf("/%s - %s\n--------\n", k, v.Description)
		}
		return helpStr, nil
	}

	// Check if token from API is expired
	if e.Client.IsTokenExpired() {
		e.reconnect()
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
