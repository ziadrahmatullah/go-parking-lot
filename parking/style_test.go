package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	"github.com/stretchr/testify/assert"
)

func TestImplementStyle(t *testing.T) {
	t.Run("should do nothing when FirstAvailableStyle", func(t *testing.T) {
		lot1 := parking.NewLot(1)
		lot2 := parking.NewLot(2)
		lots := []*parking.Lot{lot1, lot2}
		parkingStyle := parking.FirstAvailableStyle{}
		expected := lots

		parkingStyle.ImplementStyle(lots)

		assert.Equal(t, expected, lots)
	})

	t.Run("should sort lots basd highest capacity", func(t *testing.T) {
		lot1 := parking.NewLot(1)
		lot2 := parking.NewLot(2)
		lots := []*parking.Lot{lot1, lot2}
		parkingStyle := parking.HighestCapacityStyle{}
		expected := lot2

		parkingStyle.ImplementStyle(lots)

		assert.Equal(t, expected, lots[0])
	})

	t.Run("should sort lots basd highest free space", func(t *testing.T) {
		lot1 := parking.NewLot(2)
		lot2 := parking.NewLot(2)
		lots := []*parking.Lot{lot1, lot2}
		lot1.Park(*entity.NewCar("GOLANG"))
		parkingStyle := parking.HighestNumberOfFreeSpaceStyle{}
		expected := lot2

		parkingStyle.ImplementStyle(lots)

		assert.Equal(t, expected, lots[0])
	})
}
