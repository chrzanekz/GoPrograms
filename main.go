package main

import (
	"fmt"
	"booking-app/helper" //own package
	"time"
	"sync"
)

//constant values, first variables
var conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)


type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	

	// switch test
	city := "Bytow"

	switch city {
		case "Bytow":
			fmt.Printf("Welcome to Bytow conference\n")
		case "Gdansk", "Gdynia":
			fmt.Printf("Welcome to Tricity conference\n")
		default:
			fmt.Printf("you didnt choose right city\n")
	}

	//some info for user
	greetUsers()

	
	
	for {

		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//basic calculations
		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)


			fmt.Printf("the whole slice is: %v \n", bookings)
			fmt.Printf("the first value of slice is: %v \n", bookings[0])
			fmt.Printf("the type of slice is: %T \n", bookings)
			fmt.Printf("the length of slice is: %v \n", len(bookings))

			
			
			//printing first names of customers
			firstNames := printFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			
			if remainingTickets == 0 {
				//ending loop
				fmt.Println("Our conference is booked out. Please come back next year.\n")
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
	wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)
	}
	//print results
	return firstNames
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

	// create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//results on the screen
	fmt.Printf("Thank You %v %v for buying %v tickets, you'll receive an email at %v with confirmation.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n##########################")
	fmt.Printf("Sending ticket:\n%v \nto email adress %v\n", ticket, email)
	fmt.Println("##########################\n")
	wg.Done() 
}