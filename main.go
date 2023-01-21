package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("confrenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome in out %v booking application!\n", conferenceName)
	fmt.Printf("Today, we have %v tickets and %v tickets is still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for his name

	userName = "Mateusz"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)

}