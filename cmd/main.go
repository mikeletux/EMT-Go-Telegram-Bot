package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/auth"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/bot"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"

	"github.com/mikeletux/goemt"
)

func main() {
	// Get env vars
	telegramBotToken := os.Getenv("TELEGRAM_API_TOKEN")
	if len(telegramBotToken) == 0 {
		panic("please provide a token for the bot to work")
	}

	// Create allowed users for auth
	allowedUsers := strings.Split(os.Getenv("TELEGRAM_ALLOWED_USERS"), ",")

	// Create auth struct
	auth := auth.NewSimpleAuth(allowedUsers)

	// Create EMT
	configProtected := goemt.ClientConfig{
		Enpoint:   os.Getenv("EMT_ENDPOINT"), // in this case will be https://openapi.emtmadrid.es/v2
		XClientID: os.Getenv("EMT_XCLIENTID"),
		PassKey:   os.Getenv("EMT_PASSKEY"),
	}
	emt, err := emt.NewGoEMT(configProtected)
	if err != nil {
		panic(fmt.Sprintf("there was an issue when creating emt - %s", err))
	}

	// Set TelegramBot config
	config := bot.TelegramBotConfig{
		Token: telegramBotToken,
		Debug: false,
	}

	// Create TelegramBot
	bot, err := bot.NewTelegramBot(config, auth, emt)
	if err != nil {
		panic(fmt.Sprintf("error when creating the Telegram bot - %s", err))
	}

	// Run the bot
	err = bot.Run()
}
