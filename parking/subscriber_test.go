package parking_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/mocks"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	// "github.com/stretchr/testify/mock"
)

func TestNotifyFull(t *testing.T) {
	t.Run("should call NotifyFull", func(t *testing.T) {
		l := parking.NewLot(1)
		c := entity.NewCar("GOLANG")
		mockSubscriber := new(mocks.Subscriber)
		l.Subscribe(mockSubscriber)
		mockSubscriber.On("NotifyFull", l)

		l.Park(*c)

		mockSubscriber.AssertNumberOfCalls(t, "NotifyFull", 1)
	})
}

func TestNotifyAvailable(t *testing.T) {
	t.Run("should call NotifyAvailable", func(t *testing.T) {
		l := parking.NewLot(1)
		c := entity.NewCar("GOLANG")
		mockSubscriber := new(mocks.Subscriber)
		l.Subscribe(mockSubscriber)
		mockSubscriber.On("NotifyFull", l)
		mockSubscriber.On("NotifyAvailable", l)

		ticket, _ := l.Park(*c)
		l.Unpark(ticket)
		_, _ = l.Park(*c)

		mockSubscriber.AssertNumberOfCalls(t, "NotifyFull", 2)
		// mockSubscriber.AssertExpectations(t) Untuk semua yang di setup terpanggil
		mockSubscriber.AssertNumberOfCalls(t, "NotifyAvailable", 1)
	})
}
