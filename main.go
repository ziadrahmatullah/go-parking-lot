package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
)

var attendance parking.Attendance

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func Setup() {
	scanner := bufio.NewScanner(os.Stdin)
	var lots []*parking.Lot
	fmt.Println(len(lots))
	capacities := promptInput(scanner, "input parking lot capacities: ")
	capLots := strings.Split(capacities, ",")
	for _, cap := range capLots {
		cap, _ := strconv.Atoi(cap)
		lots = append(lots, parking.NewLot(cap))
	}
	attendance = *parking.NewAttendance(lots)
}

func Park() {
	scanner := bufio.NewScanner(os.Stdin)
	plateNumber := promptInput(scanner, "input plate number: ")
	ticket, err := attendance.Park(*entity.NewCar(plateNumber))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Car parked with ticket id %s\n", ticket.ID)
	}
}

func Unpark() {
	scanner := bufio.NewScanner(os.Stdin)
	ticket := promptInput(scanner, "input ticket id: #")
	car, err := attendance.Unpark(entity.Ticket{ID: ticket})
	if err != nil {
		switch err{

		}
		fmt.Println(err)
	} else {
		fmt.Printf("Car %s successfully unparked!\n", car.PlateNumber)
	}
}

// Parking Lot Status:
// Lot #1: 4 space left
// #5678 B123M
// #3456 B898M

// Lot #2: 0 space left
// #5678 B123M

func ParkingLotStatus() {
	fmt.Println("Parking Lot Status:")
	for i, lt := range attendance.LotStatus() {
		fmt.Printf("Lot #%d: %d space left\n", i+1, lt.NumberOfFreeSpace())
		for ticket, car := range lt.FieldStatus(){
			fmt.Printf("#%s %s\n", ticket.ID, car.PlateNumber)
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	menu := "------------------\n" +
		"Parking Lot\n" +
		"1. Setup\n" +
		"2. Park\n" +
		"3. UnPark\n" +
		"4. Parking Lot Status\n" +
		"5. Exit\n"

	for !exit {
		fmt.Println(menu)
		input := promptInput(scanner, "input menu: ")

		switch input {
		case "1":
			Setup()
		case "2":
			Park()
		case "3":
			Unpark()
		case "4":
			ParkingLotStatus()
		case "5":
			exit = true
		default:
			fmt.Println("invalid menu")
		}
	}
}
