package constant

import(
	"errors"
)

var (
	ErrUnrecognizedParkingTicket = errors.New("unrecognized parking ticket")
	ErrNoAvailablePosition = errors.New("no available position")
	ErrCarHasBeenParked = errors.New("car has been parked")
)