package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/auth"
)

type TelegramBotConfig struct {
	// Token is the token retrieved by Telegram BotFather
	Token string

	// Debug sets debug output on/off
	Debug bool
}

type TelegramBot struct {
	// Bot is be the struct in charge of speaking with Telegram servers
	Bot *tgbotapi.BotAPI

	// Auth is the struct that implements authentication
	Auth auth.Auth

	// Client for EMT
}

func NewTelegramBot(config TelegramBotConfig, auth auth.Auth) (*TelegramBot, error) {
	telegramBot := &TelegramBot{}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = config.Debug

	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	telegramBot.Bot = bot
	telegramBot.Auth = auth

	return telegramBot, nil
}

func (b *TelegramBot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.Bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("couldn't retrieve updates channel - %s", err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if ok := b.Auth.CheckUser(update.Message.From.UserName); !ok {
			continue // ignore non allowed users
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		b.Bot.Send(msg)
	}
	return nil //This will never be reached
}
