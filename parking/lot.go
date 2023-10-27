package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

type Lot struct {
	subscriberList []Subscriber
	field          map[entity.Ticket]entity.Car
	cap            int
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

	if len(l.field) == l.cap {
		l.notifierFull()
	}

	return ticket, nil
}

func (l *Lot) Unpark(ticket entity.Ticket) (car entity.Car, err error) {
	car, ok := l.field[ticket]
	if !ok {
		err = constant.ErrUnrecognizedParkingTicket
		return
	}
	if l.isLotFull(){
		l.notifierAvailable()
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

func (l *Lot) isLotFull() bool {
	return len(l.field) == l.cap
}

func (l *Lot) notifierFull() {
	for _, subscriber := range l.subscriberList {
		subscriber.NotifyFull(l)
	}
}

func (l *Lot) notifierAvailable() {
	for _, subscriber := range l.subscriberList {
		subscriber.NotifyAvailable(l)
	}
}

func (l *Lot) numberOfFreeSpace()int{
	return l.cap - len(l.field)
}

func (l *Lot) Subscribe(s Subscriber) {
	l.subscriberList = append(l.subscriberList, s)
}
