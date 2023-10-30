package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/constant"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
)

var attendance parking.Attendance

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
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
			capacities := promptInput(scanner, "input parking lot capacities: ")
			var lots []*parking.Lot
			capLots := strings.Split(capacities, ",")
			for _, cap := range capLots {
				cap, _ := strconv.Atoi(cap)
				lots = append(lots, parking.NewLot(cap))
			}
			attendance = *parking.NewAttendance(lots)
		case "2":
			plateNumber := promptInput(scanner, "input plate number: ")
			ticket, err := attendance.Park(*entity.NewCar(plateNumber))
			if err != nil {
				switch{
				case errors.Is(err, constant.ErrNoAvailablePosition):
					fmt.Println("Parking lot is full!")
				case errors.Is(err, constant.ErrCarHasBeenParked):
					fmt.Println("Car is parked!")
				default:
					fmt.Printf("Unexpected error: %s\n", err)
				}
			} else {
				fmt.Printf("Car parked with ticket id #%s\n", ticket.ID)
			}
		case "3":
			ticket := promptInput(scanner, "input ticket id: #")
			car, err := attendance.Unpark(entity.Ticket{ID: ticket})
			if err != nil {
				switch{
				case errors.Is(err, constant.ErrUnrecognizedParkingTicket):
					fmt.Printf("Unrecognize Ticket!")
				default:
					fmt.Printf("Unexpected error: %s\n", err)
				}
			} else {
				fmt.Printf("Car %s successfully unparked!\n", car.PlateNumber)
			}
		case "4":
			fmt.Println("Parking Lot Status:")
			for i, lt := range attendance.LotStatuses() {
				fmt.Printf("Lot #%d: %d space left\n", i+1, lt.FreeSpace)
				for ticket, car := range lt.TicketCar {
					fmt.Printf("#%s %s\n", ticket.ID, car.PlateNumber)
				}
				fmt.Println()
			}
		case "5":
			exit = true
		default:
			fmt.Println("invalid menu")
		}
	}
}
