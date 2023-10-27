package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestAttendancePark(t *testing.T) {
	t.Run("should return not nil when parking car", func(t *testing.T) {
		car := entity.Car{PlateNumber: "GOLANG"}
		var lots []*parking.Lot
		parkingLot := parking.NewLot(2)
		lots = append(lots, parkingLot)
		attendance := parking.NewAttendance(lots, 2)

		ticket, _ := attendance.Park(car)

		assert.NotNil(t, ticket)
	})

	t.Run("should return ErrNoAvailablePosition when all parking lot full", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "RUBY"}
		car3 := entity.Car{PlateNumber: "CPP"}
		expected := constant.ErrNoAvailablePosition
		var lots []*parking.Lot
		parkingLot1 := parking.NewLot(1)
		parkingLot2 := parking.NewLot(1)
		lots = append(lots, parkingLot1)
		lots = append(lots, parkingLot2)
		attendance := parking.NewAttendance(lots, 2)

		attendance.Park(car1)
		attendance.Park(car2)
		_, err := attendance.Park(car3)

		assert.ErrorIs(t, expected, err)
	})

	t.Run("should return ErrCarHasBeenParked when input 2 same car", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "GOLANG"}
		expected := constant.ErrCarHasBeenParked
		var lots []*parking.Lot
		parkingLot := parking.NewLot(2)
		lots = append(lots, parkingLot)
		attendance := parking.NewAttendance(lots, 2)

		attendance.Park(car1)
		_, err := attendance.Park(car2)

		assert.ErrorIs(t, expected, err)
	})

	t.Run("should return ErrCarHasBeenParked when car already in lot 2", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "GOLANG"}
		expected := constant.ErrCarHasBeenParked
		var lots []*parking.Lot
		parkingLot1 := parking.NewLot(2)
		parkingLot2 := parking.NewLot(2)
		lots = append(lots, parkingLot1)
		lots = append(lots, parkingLot2)
		attendance := parking.NewAttendance(lots, 2)
		parkingLot2.Park(car1)

		attendance.Park(car1)
		_, err := attendance.Park(car2)

		assert.ErrorIs(t, expected, err)
	})
}

func TestAttendanceUnpark(t *testing.T) {
	t.Run("should return car when input ticket", func(t *testing.T) {
		expectedCar := entity.Car{PlateNumber: "GOLANG"}

		var lots []*parking.Lot
		parkingLot := parking.NewLot(2)
		lots = append(lots, parkingLot)
		attendance := parking.NewAttendance(lots, 2)
		ticket, _ := attendance.Park(expectedCar)

		car, _ := attendance.Unpark(ticket)

		assert.Equal(t, expectedCar, car)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input unrecognize ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		ticket := entity.Ticket{ID: "1234"}
		expected := constant.ErrUnrecognizedParkingTicket
		var lots []*parking.Lot
		parkingLot := parking.NewLot(1)
		lots = append(lots, parkingLot)
		attendance := parking.NewAttendance(lots, 2)

		attendance.Park(car1)
		_, err := attendance.Unpark(ticket)

		assert.ErrorIs(t, expected, err)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input 2 same ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		expected := constant.ErrUnrecognizedParkingTicket
		var lots []*parking.Lot
		parkingLot := parking.NewLot(1)
		lots = append(lots, parkingLot)
		attendance := parking.NewAttendance(lots, 2)
		ticket, _ := attendance.Park(car1)

		attendance.Unpark(ticket)
		_, err := attendance.Unpark(ticket)

		assert.ErrorIs(t, expected, err)
	})
}

func TestChangeStyle(t *testing.T) {
	t.Run("should return car from HighestCapacityStyle", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		var lots []*parking.Lot
		parkingLot1 := parking.NewLot(2)
		parkingLot2 := parking.NewLot(4)
		parkingLot3 := parking.NewLot(3)
		lots = append(lots, parkingLot1)
		lots = append(lots, parkingLot2)
		lots = append(lots, parkingLot3)
		attendance := parking.NewAttendance(lots, 3)
		attendance.ChangeStyle(2)
		ticket, _ := attendance.Park(*car1)

		car, _ := parkingLot2.Unpark(ticket)

		assert.Equal(t, car1, &car)
	})

	t.Run("should return car from HighestNumberOfFreeSpaceStyle", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		car2 := entity.NewCar("CPP")
		car3 := entity.NewCar("RUBY")
		var lots []*parking.Lot
		parkingLot1 := parking.NewLot(2)
		parkingLot2 := parking.NewLot(3)
		lots = append(lots, parkingLot1)
		lots = append(lots, parkingLot2)
		attendance := parking.NewAttendance(lots, 2)
		attendance.ChangeStyle(3)
		attendance.Park(*car1)
		attendance.Park(*car2)
		ticket, _ := attendance.Park(*car3)

		carUnpark, _ := parkingLot1.Unpark(ticket)

		assert.Equal(t, car3, &carUnpark)
	})
}