package emt

import (
	"fmt"

	"github.com/mikeletux/goemt"
	"github.com/mikeletux/goemt/busemtmad"
)

type Emt interface {
	// GetBusWaitingTimes must return a formated string about bus waiting times
	GetBusWaitingTimes(stopID int) string
	Logout() error
}

type goEMT struct {
	Client *goemt.APIClient
}

func NewGoEMT(config goemt.ClientConfig) (Emt, error) {
	goEMT := goEMT{}
	c, err := goemt.Connect(config)
	if err != nil {
		return nil, err
	}
	goEMT.Client = c
	return &goEMT, nil

}

func (e *goEMT) GetBusWaitingTimes(stopID int) string {
	//GetTimeArrivalBus func needs a struct to use it when post
	postData := busemtmad.PostInfoTimeArrival{
		CultureInfo:             "ES",
		TextStopRequired:        "Y",
		TextEstimationsRequired: "Y",
		TextIncidencesRequired:  "N",
	}

	busTimes, err := busemtmad.GetTimeArrivalBus(e.Client, stopID, 0, postData)
	if err != nil {
		return "Couldn't retrieve bus waiting times. Contact bot administrator"
	}

	var response string
	for _, v := range busTimes {
		for _, arrive := range v.Arrive {
			response += fmt.Sprintf("Bus number %s will arrive in %d seconds and heads to %s\n\n", arrive.Line, arrive.EstimateArrive, arrive.Destination)
		}

	}

	return response
}

func (e *goEMT) Logout() error {
	e.Client.Logout()
	return nil
}
