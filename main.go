package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome in out", conferenceName, "booking application!")
	fmt.Println("Today, we have", conferenceTickets, "tickets and", remainingTickets, "tickets is still available.")
	fmt.Println("Get your tickets here to attend")
}