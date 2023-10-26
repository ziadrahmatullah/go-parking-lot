package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/mocks"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	// "github.com/stretchr/testify/mock"
)

func TestNotify(t *testing.T) {
	t.Run("should call Notify",func(t *testing.T) {
		l := parking.NewLot(1)
		c := entity.NewCar("GOLANG")
		mockSubscriber := new(mocks.Subscriber)
		l.Subscribe(mockSubscriber)
		mockSubscriber.On("Notify",l)

		l.Park(*c)

		mockSubscriber.AssertNumberOfCalls(t, "Notify", 1)
	} )
}