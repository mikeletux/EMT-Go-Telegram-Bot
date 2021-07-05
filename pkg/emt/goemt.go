package emt

import (
	"fmt"

	"github.com/mikeletux/goemt"
	"github.com/mikeletux/goemt/busemtmad"
)

type goEMT struct {
	// Client is the struct that will hold the EMT client
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

func (e *goEMT) GetAllBusWaitingTimes(stopID int) ([]busemtmad.TimeArrivalBus, error) {
	//GetTimeArrivalBus func needs a struct to use it when post
	postData := busemtmad.PostInfoTimeArrival{
		CultureInfo:             "ES",
		TextStopRequired:        "Y",
		TextEstimationsRequired: "Y",
		TextIncidencesRequired:  "N",
	}

	busTimes, err := busemtmad.GetTimeArrivalBus(e.Client, stopID, 0, postData)
	if err != nil {
		return nil, fmt.Errorf("error. Could not retrieve bus waiting time - %s", err)
	}

	return busTimes, nil
}

func (e *goEMT) Logout() error {
	e.Client.Logout()
	return nil
}
