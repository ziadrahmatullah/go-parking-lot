package parking

import (
	"sort"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
)

const (
	FirstAvailable int = iota + 1
	HighestCapacityStyle
	HighestNumberOfFreeSpaceStyle
)

type Attendance struct {
	lot          []*Lot
	capLot       int
	availableLot []*Lot
	parkStyle    int
}

func NewAttendance(lot []*Lot, cap int) *Attendance {
	newAttendance := &Attendance{lot, cap, lot, 1}
	for _, lt := range lot {
		lt.Subscribe(newAttendance)
	}
	return newAttendance
}

func (a *Attendance) ChangeStyle(input int) {
	switch input {
	case HighestCapacityStyle:
		a.parkStyle = HighestCapacityStyle
	case HighestNumberOfFreeSpaceStyle:
		a.parkStyle = HighestNumberOfFreeSpaceStyle
	}
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
	if a.parkStyle == HighestCapacityStyle {
		a.highestCapacityStyle()
	} else if a.parkStyle == HighestNumberOfFreeSpaceStyle {
		a.highestNumberOfFreeSpaceStyle()
	}
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

func (a *Attendance) NotifyFull(lot *Lot) {
	for i, lt := range a.availableLot {
		if lt == lot {
			a.availableLot = deleteElement(a.availableLot, i)
			break
		}
	}
}

func (a *Attendance) NotifyAvailable(lot *Lot) {
	a.availableLot = append(a.availableLot, lot)
}

func (a *Attendance) highestCapacityStyle() {
	sort.Slice(a.availableLot, func(i, j int) bool {
		return a.availableLot[i].cap > a.availableLot[j].cap
	})
}

func (a *Attendance) highestNumberOfFreeSpaceStyle() {
	sort.Slice(a.availableLot, func(i, j int) bool {
		return a.availableLot[i].numberOfFreeSpace() > a.availableLot[j].numberOfFreeSpace()
	})
}

func deleteElement(slice []*Lot, index int) []*Lot {
	return append(slice[:index], slice[index+1:]...)
}

//TODO: Satuin Notify
//TODO: Gunakan Iota
