package bot

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

func GetAllBusWaitingTimes(emt emt.Emt, message *tgbotapi.Message) (string, error) {
	// Get arguments
	args := strings.Split(message.CommandArguments(), " ")

	stopID, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("unable to get argument stop ID")
	}

	busTimes, err := emt.GetAllBusWaitingTimes(stopID)
	if err != nil {
		return "", err
	}

	var response string
	for _, v := range busTimes {
		for _, arrive := range v.Arrive {
			response += fmt.Sprintf("%s\\[%s]%s arrives in %d mins\n\n", busFrontEmoji, arrive.Line, arrive.Destination, arrive.EstimateArrive/60)
		}
	}

	return response, nil
}
