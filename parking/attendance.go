package parking

import (
	"fmt"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)



type Attendance struct{
	lot []*Lot
	capLot int
}

func NewAttendance(lot []*Lot, cap int) *Attendance{
	return &Attendance{lot, cap}
}

func (a *Attendance) Park(car entity.Car) (ticket entity.Ticket,err error){
	if a.isCarAvailable(car){
		err = constant.ErrCarHasBeenParked
		return
	}
	for _, lt := range a.lot {
		if !lt.isLotFull() {
			return lt.Park(car)
		}
	}
	err = constant.ErrNoAvailablePosition
	return 
}

func (a *Attendance) Unpark(ticket entity.Ticket) (car entity.Car,err error){
	for _, lt := range a.lot {
		if _, ok := lt.field[ticket]; ok{
			return lt.Unpark(ticket)
		}
	}
	err = constant.ErrUnrecognizedParkingTicket
	return 
}

func (a *Attendance) isCarAvailable(car entity.Car) bool{
	for _, lt := range a.lot{
		if lt.isCarAvailable(car){
			return true
		}
	}
	return false
}

func (a *Attendance) Notify(lot *Lot) {
	fmt.Println("lot penuh:", lot)
}
