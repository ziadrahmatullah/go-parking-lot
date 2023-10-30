package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {
	t.Run("should return not nil when parking car", func(t *testing.T) {
		car := entity.NewCar("GOLANG")
		parkingLot := parking.NewLot(2)
		
		ticket, _ := parkingLot.Park(*car)

		assert.NotNil(t, ticket)
	})

	t.Run("should return ErrNoAvailablePosition when parking lot full", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		car2 := entity.NewCar("RUBY")
		expected := constant.ErrNoAvailablePosition
		parkingLot := parking.NewLot(1)

		parkingLot.Park(*car1)
		_, err := parkingLot.Park(*car2)

		assert.ErrorIs(t, expected, err)
	})

	t.Run("should return ErrCarHasBeenParked when nput 2 same car", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		car2 := entity.NewCar("GOLANG")
		expected := constant.ErrCarHasBeenParked

		parkingLot := parking.NewLot(2)
		parkingLot.Park(*car1)
		_, err := parkingLot.Park(*car2)

		assert.ErrorIs(t, expected, err)
	})
}

func TestUnpark(t *testing.T) {
	t.Run("should return car when input ticket", func(t *testing.T) {
		expectedCar := entity.NewCar("GOLANG")
		parkingLot := parking.NewLot(2)

		ticket, _ := parkingLot.Park(*expectedCar)
		car, _ := parkingLot.Unpark(ticket)

		assert.Equal(t, expectedCar, &car)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input unrecognize ticket", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		ticket := entity.Ticket{ID: "1234"}
		parkingLot := parking.NewLot(1)
		expected := constant.ErrUnrecognizedParkingTicket

		parkingLot.Park(*car1)
		_, err := parkingLot.Unpark(ticket)

		assert.ErrorIs(t, expected, err)
	})

	t.Run("should return ErrUnrecognizedParkingTicket when input 2 same ticket", func(t *testing.T) {
		car1 := entity.NewCar("GOLANG")
		parkingLot := parking.NewLot(1)
		expected := constant.ErrUnrecognizedParkingTicket

		ticket, _ := parkingLot.Park(*car1)
		parkingLot.Unpark(ticket)
		_, err := parkingLot.Unpark(ticket)

		assert.ErrorIs(t, expected, err)
	})
}

func TestLotStatus(t *testing.T) {
	t.Run("should return not nil when call LotStatus", func(t *testing.T) {
		parkingLot := parking.NewLot(2)
		
		lotStatus:= parkingLot.LotStatus()

		assert.NotNil(t, lotStatus)
	})
}

