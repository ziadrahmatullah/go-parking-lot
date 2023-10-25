package parking

import "git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"

type Lot struct{
	field map[entity.Ticket]entity.Car
}

func NewLot() *Lot{
	field := make(map[entity.Ticket]entity.Car)
	return &Lot{field}
}

func (pl *Lot) Park(car entity.Car) entity.Ticket{
	T := entity.NewTicket()
	pl.field[T] = car
	return T
}

func (pl *Lot) Unpark(ticket entity.Ticket) entity.Car{
	car := pl.field[ticket]
	delete(pl.field, ticket)
	return car
}