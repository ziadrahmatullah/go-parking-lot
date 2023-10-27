package parking

type Subscriber interface{
    NotifyFull(lot *Lot)
    NotifyAvailable(lot *Lot)
}