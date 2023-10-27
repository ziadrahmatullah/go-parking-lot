package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

type Attendance struct {
	lot    []*Lot
	capLot int
	availableLot []*Lot
}

func NewAttendance(lot []*Lot, cap int) *Attendance {
	newAttendance := &Attendance{lot, cap, lot}
	for _, lt := range lot {
		lt.Subscribe(newAttendance)
	}
	return newAttendance
}

func (a *Attendance) Park(car entity.Car) (ticket entity.Ticket, err error) {
	// TODO: make is availablelot empty for error below
	
	if a.isCarAvailable(car) {
		err = constant.ErrCarHasBeenParked
		return
	}
	for _, lt := range a.availableLot{
		return lt.Park(car)
	}
	err = constant.ErrNoAvailablePosition
	return
}

func (a *Attendance) Unpark(ticket entity.Ticket) (car entity.Car, err error) {
	for _, lt := range a.lot {
		if _, ok := lt.field[ticket]; ok {
			return lt.Unpark(ticket)
		}
	}
	err = constant.ErrUnrecognizedParkingTicket
	return
}

func (a *Attendance) isCarAvailable(car entity.Car) bool {
	for _, lt := range a.lot {
		if lt.isCarAvailable(car) {
			return true
		}
	}
	return false
}

func (a *Attendance) NotifyFull(lot *Lot) {
	for i, lt := range a.availableLot {
		if lt == lot{
			a.availableLot = deleteElement(a.availableLot, i)
			break
		}
	}
}

func (a *Attendance) NotifyAvailable(lot *Lot) {
	a.availableLot = append(a.availableLot, lot)
}

func deleteElement(slice []*Lot, index int) []*Lot {
	return append(slice[:index], slice[index+1:]...)
}

//TODO: Satuin Notify
//TODO: Gunakan Iota

