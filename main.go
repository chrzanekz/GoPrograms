package main

import (
	"fmt"
	"strings"
)

func main() {
	//constant values, first variables
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	//some info for user
	fmt.Printf("confrenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome in out %v booking application!\n", conferenceName)
	fmt.Printf("Today, we have %v tickets and %v tickets is still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		//declare user variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for his informations
		fmt.Println("Enter first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter email address:")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets do you want to buy:")
		fmt.Scan(&userTickets)

		//basic calculations
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("the whole slice is: %v \n", bookings)
		fmt.Printf("the first value of slice is: %v \n", bookings[0])
		fmt.Printf("the type of slice is: %T \n", bookings)
		fmt.Printf("the length of slice is: %v \n", len(bookings))

		//results on the screen
		fmt.Printf("Thank You %v %v for buying %v tickets, you'll receive an email at %v with confirmation.\n", firstName, lastName, email, userTickets)
		fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
		
		//loop for creating first name slice
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		//print results
		fmt.Printf("The first names clients of our bookings are: %v\n", firstNames)
	}
}