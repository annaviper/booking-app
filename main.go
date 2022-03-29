package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string // slice

	// Welcome message
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for remainingTickets > 0 {
		// Inputs
		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Println("What's your name?")
		fmt.Scan(&firstName)

		fmt.Println("What's your surname?")
		fmt.Scan(&lastName)

		fmt.Println("What's your email?")
		fmt.Scan(&userEmail)

		fmt.Println("How many tickets do you want to purchase?")
		fmt.Scan(&userTickets)

		// Tickets logic
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			fullName := firstName + " " + lastName
			bookings = append(bookings, fullName)

			// Output
			fmt.Printf("User %v booked %v ticket(s). Confirmation email sent to %v.\n", fullName, userTickets, userEmail)
			fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

			// range loop
			firstNames := []string{}
			for _, fullName := range bookings {
				var nameFields = strings.Fields(fullName)
				firstNames = append(firstNames, nameFields[0])
			}
			fmt.Printf("These are all the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out.")
				break
			}
		} else {
			fmt.Printf("There are only %v tickets remaining.\n", remainingTickets)
			continue
		}
	}
}
