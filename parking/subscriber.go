package parking

type Subscriber interface {
	Notify(lot *Lot, message string)
}
