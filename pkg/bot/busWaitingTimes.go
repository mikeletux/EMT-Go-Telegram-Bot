package bot

import (
	"fmt"
	"strconv"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

func GetAllBusWaitingTimes(emt emt.Emt, args []string) (string, error) {
	// Get arguments
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
