package parking

import (
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/util"
)

type Attendance struct {
	lot          []*Lot
	availableLot []*Lot
	parkStyle    ParkStyle
}

func NewAttendance(lot []*Lot) *Attendance {
	newAttendance := &Attendance{lot, make([]*Lot, 0), &FirstAvailableStyle{}}
	for _, lt := range lot {
		newAttendance.availableLot = append(newAttendance.availableLot, lt)
		lt.Subscribe(newAttendance)
	}
	return newAttendance
}

func (a *Attendance) ChangeStyle(parkStyle ParkStyle) {
	a.parkStyle = parkStyle
}

func (a *Attendance) Park(car entity.Car) (ticket entity.Ticket, err error) {
	if len(a.availableLot) == 0 {
		err = constant.ErrNoAvailablePosition
		return
	}
	if a.isCarAvailable(car) {
		err = constant.ErrCarHasBeenParked
		return
	}
	a.parkStyle.ImplementStyle(a.availableLot)
	return a.availableLot[0].Park(car)
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

func (a *Attendance) Notify(lot *Lot, message string) {
	switch message {
	case "full":
		a.notifyFull(lot)
	case "available":
		a.notifyAvailable(lot)
	}
}

func (a *Attendance) notifyFull(lot *Lot) {
	for i, lt := range a.availableLot {
		if lt == lot {
			a.availableLot = util.DeleteElement[*Lot](a.availableLot, i)
			break
		}
	}
}

func (a *Attendance) notifyAvailable(lot *Lot) {
	a.availableLot = append(a.availableLot, lot)
}

func (a *Attendance) LotStatus() []*Lot{
	return a.lot
}


