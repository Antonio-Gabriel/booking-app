package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickers int = 50
	var remainingTickets int = 50
	var bookings = []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!.\n", conferenceTickers, remainingTickets)
	fmt.Println("Get your tickers here to attend")

	for {
		var userName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name

		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter your tickets number: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",
				remainingTickets, userTickets)

			break
		}

		remainingTickets -= userTickets

		bookings = append(bookings, userName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)

		fmt.Printf(
			"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			userName, lastName, email, userTickets)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickers)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)

			firstNames = append(firstNames, names[0])
		}

		noTicketsRemaining := (remainingTickets == 0)
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back to the next year.")
			break
		}

		fmt.Printf("The first names of bookints are: %v\n", firstNames)
	}
}
