package parking_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"github.com/stretchr/testify/assert"
)

func TestAttendancePark(t *testing.T) {
	t.Run("should return not nil when parking car", func(t *testing.T) {
		car := entity.Car{PlateNumber: "GOLANG"}

		parkingLot := parking.NewLot(2)
		attendance := parking.NewAttendance(*parkingLot)
		ticket, _ := attendance.Park(car)
		assert.NotNil(t, ticket)	
	})



	t.Run("should return ErrNoAvailablePosition when parking lot full", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "RUBY"}
		expected := constant.ErrUnrecognizedParkingTicket

		parkingLot := parking.NewLot(1)
		attendance := parking.NewAttendance(*parkingLot)
		attendance.Park(car1)
		_, err := attendance.Park(car2)
		errors.Is(expected,err)	
	})

	t.Run("should return ErrCarHasBeenParked when nput 2 same car", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "GOLANG"}
		expected := constant.ErrCarHasBeenParked

		parkingLot := parking.NewLot(2)
		attendance := parking.NewAttendance(*parkingLot)
		attendance.Park(car1)
		_, err := parkingLot.Park(car2)
		errors.Is(expected,err)	
	})
}

func TestAttendanceUnpark(t *testing.T) {
	t.Run("should return car when input ticket", func(t *testing.T) {
		expectedCar := entity.Car{PlateNumber: "GOLANG"}

		parkingLot := parking.NewLot(2)
		attendance := parking.NewAttendance(*parkingLot)
		ticket, _ := attendance.Park(expectedCar)
		car ,_:= attendance.Unpark(ticket)

		assert.Equal(t, expectedCar, car)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input unrecognize ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		ticket := entity.Ticket{ID: "1234"}
		expected := constant.ErrUnrecognizedParkingTicket

		parkingLot := parking.NewLot(1)
		attendance := parking.NewAttendance(*parkingLot)
		
		attendance.Park(car1)
		_, err := attendance.Unpark(ticket)
		errors.Is(expected,err)	
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input 2 same ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		parkingLot := parking.NewLot(1)
		expected := constant.ErrUnrecognizedParkingTicket

		ticket, _ := parkingLot.Park(car1)
		attendance := parking.NewAttendance(*parkingLot)
		attendance.Unpark(ticket)
		_, err := attendance.Unpark(ticket)
		errors.Is(expected,err)	
	})
}