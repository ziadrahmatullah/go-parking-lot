package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

type Lot struct {
	field map[entity.Ticket]entity.Car
	cap   int
}

func NewLot(cap int) *Lot {
	field := make(map[entity.Ticket]entity.Car, 0)
	return &Lot{field: field, cap: cap}
}

func (l *Lot) Park(car entity.Car) (ticket entity.Ticket, err error) {
	if len(l.field) == l.cap {
		err = constant.ErrNoAvailablePosition
		return
	}
	if l.isCarAvailable(car) {
		err = constant.ErrCarHasBeenParked
		return
	}
	ticket = entity.NewTicket()
	l.field[ticket] = car
	return ticket, nil
}

func (l *Lot) Unpark(ticket entity.Ticket) (car entity.Car, err error) {
	car, ok := l.field[ticket]
	if !ok {
		err = constant.ErrUnrecognizedParkingTicket
		return
	}
	delete(l.field, ticket)
	return
}

func (l *Lot) isCarAvailable(car entity.Car) bool {
	for _, parkingCar := range l.field {
		if parkingCar.PlateNumber == car.PlateNumber {
			return true
		}
	}
	return false
}

func (l *Lot) isLotFull()bool{
	return len(l.field) == l.cap
}