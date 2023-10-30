package entity

type LotStatus struct {
	TicketCar map[Ticket]Car
	FreeSpace int
}