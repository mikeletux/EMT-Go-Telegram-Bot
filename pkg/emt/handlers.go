package emt

import (
	"fmt"
	"strconv"

	"github.com/mikeletux/goemt"
	"github.com/mikeletux/goemt/busemtmad"
)

type Handler struct {
	// Description of the handler
	Description string

	// Actual fuction that carries out the logic
	Handler func(goemt.IAPI, []string) (string, error)
}

func GetAllHandlers() map[string]Handler {
	return map[string]Handler{
		"bus_wait_times": {
			Description: "This command returns the bus waiting times given a stop ID.\n" +
				"\tArguments: [1]: ID from the stop to query",
			Handler: HandlerGetBusWaitingTimes,
		},
	}
}

func HandlerGetBusWaitingTimes(c goemt.IAPI, arguments []string) (string, error) {
	// Get arguments
	stopID, err := strconv.Atoi(arguments[0])
	if err != nil {
		return "", fmt.Errorf("unable to get argument stop ID")
	}

	//GetTimeArrivalBus func needs a struct to use it when post
	postData := busemtmad.PostInfoTimeArrival{
		CultureInfo:             "ES",
		TextStopRequired:        "Y",
		TextEstimationsRequired: "Y",
		TextIncidencesRequired:  "N",
	}

	busTimes, err := busemtmad.GetTimeArrivalBus(c, stopID, 0, postData)
	if err != nil {
		return "", fmt.Errorf("Couldn't retrieve bus waiting times. Contact bot administrator")
	}

	var response string
	for _, v := range busTimes {
		for _, arrive := range v.Arrive {
			response += fmt.Sprintf("%s[%s]%s arrives in %d mins\n\n", busFrontEmoji, arrive.Line, arrive.Destination, arrive.EstimateArrive/60)
		}
	}

	return response, nil
}
