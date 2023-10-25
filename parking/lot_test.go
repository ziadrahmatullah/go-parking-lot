package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {
	t.Run("should return not nil when parking car", func(t *testing.T) {
		car := entity.Car{PlateNumber: "GOLANG"}

		parkingLot := parking.NewLot()
		ticket := parkingLot.Park(car)
		assert.NotNil(t, ticket)	
	})
}

func TestUnpark(t *testing.T) {
	t.Run("should return car when input ticket", func(t *testing.T) {
		car := entity.Car{PlateNumber: "GOLANG"}

		parkingLot := parking.NewLot()
		ticket := parkingLot.Park(car)
		expected := parkingLot.Unpark(ticket)

		assert.Equal(t, expected, car)
	})
}