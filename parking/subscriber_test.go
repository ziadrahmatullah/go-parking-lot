package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/mocks"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
)

func TestNotify(t *testing.T) {
	t.Run("should call notifyFull", func(t *testing.T) {
		l := parking.NewLot(1)
		c := entity.NewCar("GOLANG")
		mockSubscriber := new(mocks.Subscriber)
		l.Subscribe(mockSubscriber)
		mockSubscriber.On("Notify", l, "full")

		l.Park(*c)

		mockSubscriber.AssertNumberOfCalls(t, "Notify", 1)
	})

	t.Run("should call notifyAvailable", func(t *testing.T) {
		l := parking.NewLot(1)
		c := entity.NewCar("GOLANG")
		mockSubscriber := new(mocks.Subscriber)
		l.Subscribe(mockSubscriber)
		mockSubscriber.On("Notify", l, "full")
		mockSubscriber.On("Notify", l, "available")

		ticket, _ := l.Park(*c)
		l.Unpark(ticket)
		_, _ = l.Park(*c)

		mockSubscriber.AssertNumberOfCalls(t, "Notify", 3)
	})
}