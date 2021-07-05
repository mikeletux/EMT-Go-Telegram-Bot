package bot

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot interface {
	GetUpdatesChan(UpdateConfig) (UpdatesChannel, error)
	Send(Chattable) (Message, error)
}
