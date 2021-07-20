package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

func GetAbout(emt emt.Emt, message *tgbotapi.Message) (string, error) {
	return fmt.Sprintf("%s This bot has been developed by MikeletuX.\n"+
		"%s API provided by mobilitylabs (https://mobilitylabs.emtmadrid.es)\n"+
		"%s Code available at https://github.com/mikeletux/EMT-Go-Telegram-Bot\n"+
		"2021", phone, key, paper), nil
}
