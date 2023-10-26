package main

import (
	"bufio"
	"fmt"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
	// "os"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var lots []*parking.Lot
	car := entity.NewCar("123")
	lot1 := parking.NewLot(1)
	lots = append(lots, lot1)
	attendant := parking.NewAttendance(lots, 2)
	attendant.Park(*car)

	// scanner := bufio.NewScanner(os.Stdin)
	// exit := false
	// menu := "Parking Lot\n" +
	// 	"1. Setup\n" +
	// 	"2. Park\n" +
	// 	"3. Un Park\n" +
	// 	"4. Exit"

	// for !exit {
	// 	fmt.Println(menu)
	// 	input := promptInput(scanner, "input menu: ")

	// 	switch input {
	// 	case "1":
	// 		capacities := promptInput(scanner, "input parking lot capacities: ")
	// 		fmt.Println("menu 1 selected: ", capacities)
	// 	case "4":
	// 		exit = true
	// 	default:
	// 		fmt.Println("invalid menu")
	// 	}
	// }
}
