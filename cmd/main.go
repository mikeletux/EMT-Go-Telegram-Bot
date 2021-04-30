package main

import (
	"fmt"
	"os"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/auth"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/bot"
)

func main() {
	// Get env vars
	telegramBotToken := os.Getenv("TELEGRAM_API_TOKEN")
	if len(telegramBotToken) == 0 {
		panic("please provide a token for the bot to work")
	}

	// Create allowed users for auth
	allowedUsers := []string{"MikeletuX"}

	// Create auth struct
	auth := auth.NewSimpleAuth(allowedUsers)

	// Set TelegramBot config
	config := bot.TelegramBotConfig{
		Token: telegramBotToken,
		Debug: false,
	}

	// Create TelegramBot
	bot, err := bot.NewTelegramBot(config, auth)
	if err != nil {
		panic(fmt.Sprintf("error when creating the Telegram bot - %s", err))
	}

	// Run the bot
	err = bot.Run()
}
