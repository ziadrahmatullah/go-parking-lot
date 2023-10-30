package parking

import "sort"

type ParkStyle interface {
	ImplementStyle([]*Lot)
}

type FirstAvailableStyle struct{}

func (fa *FirstAvailableStyle) ImplementStyle(lots []*Lot){
}

type HighestCapacityStyle struct{}

func (hc *HighestCapacityStyle) ImplementStyle(lots []*Lot){
	sort.Slice(lots, func(i, j int) bool {
		return lots[i].isHigherCapacityThan(lots[j])
	})
} 

type HighestNumberOfFreeSpaceStyle struct{}

func (hs *HighestNumberOfFreeSpaceStyle) ImplementStyle(lots []*Lot){
	sort.Slice(lots, func(i, j int) bool {
		return lots[i].isHigherSpaceThan(lots[j])
	})
} 