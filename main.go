package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome in out %v booking application!\n", conferenceName)
	fmt.Printf("Today, we have %v tickets and %v tickets is still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}