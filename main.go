package main

import (
	"bufio"
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/entity"
	"git.garena.com/sea-labs-id/batch-04/shared-projects/go-parking-lot/parking"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	menu := "Parking Lot\n" +
		"1. Setup\n" +
		"2. Park\n" +
		"3. Un Park\n" +
		"4. Exit"

	for !exit {
		fmt.Println(menu)
		input := promptInput(scanner, "input menu: ")

		p := parking.NewLot(2)

		switch input {
		case "1":
			capacities := promptInput(scanner, "input parking lot capacities: ")
			fmt.Println("menu 1 selected: ", capacities)
		case "2":
			ticket , err := p.Park(entity.Car{PlateNumber: "GOLANG"})
			fmt.Print(ticket, err)
		case "4":
			exit = true
		default:
			fmt.Println("invalid menu")
		}
	}
}
