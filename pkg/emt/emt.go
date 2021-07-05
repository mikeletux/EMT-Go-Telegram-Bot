package emt

import "github.com/mikeletux/goemt/busemtmad"

type Emt interface {
	GetAllBusWaitingTimes(stopID int) ([]busemtmad.TimeArrivalBus, error)
	Logout() error
}
