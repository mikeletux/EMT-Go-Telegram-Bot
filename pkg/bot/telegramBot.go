package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/auth"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

type TelegramBotConfig struct {
	// Token is the token retrieved by Telegram BotFather
	Token string

	// Debug sets debug output on/off
	Debug bool
}

type TelegramBot struct {
	// Bot is be the struct in charge of speaking with Telegram servers
	Bot Bot

	// Auth is the struct that implements authentication
	Auth auth.Auth

	actions *botActions
}

func NewTelegramBot(config TelegramBotConfig, auth auth.Auth, emt emt.Emt) (*TelegramBot, error) {
	telegramBot := &TelegramBot{}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = config.Debug

	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	telegramBot.Bot = bot
	telegramBot.Auth = auth
	telegramBot.actions = NewBotActions(emt, auth)

	return telegramBot, nil
}

func (b *TelegramBot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.Bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("couldn't retrieve updates channel - %s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			log.Printf("[INFO] user: %s - message: %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.ParseMode = tgbotapi.ModeMarkdown

				res, err := b.actions.PerformAction(update.Message)
				if err != nil {
					msg.Text = fmt.Sprintf("[Error] - %s", err)
				} else {
					msg.Text = res
				}
				b.Bot.Send(msg)
			}
		case <-c:
			log.Printf("[INFO] finishing gracefully telegram bot...")
			err := b.actions.Emt.Logout()
			if err != nil {
				return err
			}
			log.Printf("[INFO] EMT client closed successfully")
			return nil
		}
	}
}
