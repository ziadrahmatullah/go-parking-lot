package parking_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {
	t.Run("should return not nil when parking car", func(t *testing.T) {
		car := entity.Car{PlateNumber: "GOLANG"}

		parkingLot := parking.NewLot(2)
		ticket, _ := parkingLot.Park(car)
		assert.NotNil(t, ticket)	
	})



	t.Run("should return ErrNoAvailablePosition when parking lot full", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "RUBY"}
		expected := constant.ErrUnrecognizedParkingTicket

		parkingLot := parking.NewLot(1)
		parkingLot.Park(car1)
		_, err := parkingLot.Park(car2)
		errors.Is(expected,err)	
	})

	t.Run("should return ErrCarHasBeenParked when nput 2 same car", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		car2 := entity.Car{PlateNumber: "GOLANG"}
		expected := constant.ErrCarHasBeenParked

		parkingLot := parking.NewLot(2)
		parkingLot.Park(car1)
		_, err := parkingLot.Park(car2)
		errors.Is(expected,err)	
	})
}

func TestUnpark(t *testing.T) {
	t.Run("should return car when input ticket", func(t *testing.T) {
		expectedCar := entity.Car{PlateNumber: "GOLANG"}
		parkingLot := parking.NewLot(2)

		ticket, _ := parkingLot.Park(expectedCar)
		car ,_:= parkingLot.Unpark(ticket)

		assert.Equal(t, expectedCar, car)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input unrecognize ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		ticket := entity.Ticket{ID: "1234"}
		parkingLot := parking.NewLot(1)

		expected := constant.ErrUnrecognizedParkingTicket

		parkingLot.Park(car1)
		_, err := parkingLot.Unpark(ticket)
		errors.Is(expected,err)	
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input 2 same ticket", func(t *testing.T) {
		car1 := entity.Car{PlateNumber: "GOLANG"}
		parkingLot := parking.NewLot(1)
		expected := constant.ErrUnrecognizedParkingTicket

		ticket, _ := parkingLot.Park(car1)
		parkingLot.Unpark(ticket)
		_, err := parkingLot.Unpark(ticket)
		errors.Is(expected,err)	
	})
}