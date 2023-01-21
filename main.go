package main

import "fmt"

func main() {
	//constant values, first variables
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	//some info for user
	fmt.Printf("confrenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome in out %v booking application!\n", conferenceName)
	fmt.Printf("Today, we have %v tickets and %v tickets is still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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
	bookings[0] = firstName + " " + lastName

	fmt.Printf("the whole array is: %v \n", bookings)
	fmt.Printf("the first value of array is: %v \n", bookings[0])
	fmt.Printf("the type of array is: %T \n", bookings)
	fmt.Printf("the length of array is: %v \n", len(bookings))

	//results on the screen
	fmt.Printf("Thank You %v %v for buying %v tickets, you'll receive an email at %v with confirmation.\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}