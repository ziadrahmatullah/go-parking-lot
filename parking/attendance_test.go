package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
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

		assert.ErrorIs(t,expected,err)	
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

		assert.ErrorIs(t,expected,err)	
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

		assert.ErrorIs(t,expected,err)	
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

		car ,_:= attendance.Unpark(ticket)

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

		assert.ErrorIs(t,expected,err)	
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

		assert.ErrorIs(t,expected,err)	
	})
}