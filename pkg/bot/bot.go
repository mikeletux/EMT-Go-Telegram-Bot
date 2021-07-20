package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot interface {
	GetUpdatesChan(tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
}
