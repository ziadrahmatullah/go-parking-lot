package parking

import "git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"



type Attendance struct{
	lot Lot
}

func NewAttendance(lot Lot) *Attendance{
	return &Attendance{lot}
}

func (a *Attendance) Park(car entity.Car) (entity.Ticket, error){
	return a.lot.Park(car)
}

func (a *Attendance) Unpark(ticket entity.Ticket) (entity.Car, error){
	return a.lot.Unpark(ticket)
}