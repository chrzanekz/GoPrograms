package main

import (
	"fmt"
	"strings"
)

//constant values, first variables
var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = []string{}


func main() {
	

	// switch test
	city := "Bytow"

	switch city {
		case "Bytow":
			fmt.Printf("Welcome to Bytow conference\n")
		case "Gdansk", "Gdynia":
			fmt.Printf("Welcome to Tricity conference\n")
		default:
			fmt.Printf("you didnt choose right city")
	}

	//some info for user
	greetUsers()

	
	
	for {

		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		//basic calculations
		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, email)


			fmt.Printf("the whole slice is: %v \n", bookings)
			fmt.Printf("the first value of slice is: %v \n", bookings[0])
			fmt.Printf("the type of slice is: %T \n", bookings)
			fmt.Printf("the length of slice is: %v \n", len(bookings))

			
			
			//printing first names of customers
			firstNames := printFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			
			if remainingTickets == 0 {
				//ending loop
				fmt.Println("Our conference is booked out. Please come back next year.")
				break
			}
			// else checking for variables input
		} else {
			if !isValidName {
				fmt.Printf("First name or last name is incorrect (too short).\n")
			}
			if !isValidEmail {
				fmt.Printf("Check your e-mail address\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Check your ticket number, it's incorrect or we don't have this much tickets available.\n")
			}
		}
	}
}

//additional function to test
func greetUsers(){
	fmt.Printf("confrenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome in out %v booking application!\n", conferenceName)
	fmt.Printf("Today, we have %v tickets and %v tickets is still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames()[]string{
	//loop for creating first name slice
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//print results
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	//bool variables checking
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint){
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	//results on the screen
	fmt.Printf("Thank You %v %v for buying %v tickets, you'll receive an email at %v with confirmation.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}