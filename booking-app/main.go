package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var matchDay string = "Manchester United VS Chelsea"
var bookedTickets = make([]UserDetails, 0)
var remainingTickets uint = 50
var wg = sync.WaitGroup{}

type UserDetails struct {
	firstName   string
	secondName  string
	email       string
	userTickets uint
}

func main() {
	greetUser()

	firstName, secondName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, secondName, email, userTickets, remainingTickets)

	if isValidTicketNumber && isValidEmail && isValidName {
		bookTicket(firstName, secondName, email, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, secondName, email)

		firstNames := getFirstNames(bookedTickets)
		fmt.Println(">>> Bookings Updated: ", firstNames)
		fmt.Printf(">>> Bookings lists map: %v \n", bookedTickets)
	} else {
		fmt.Println(">>> Invalid Input")
		if !isValidName {
			fmt.Println(">>> Wrong Name Input")
		}
		if !isValidEmail {
			fmt.Println(">>> Wrong Email Input")
		}
		if !isValidTicketNumber {
			fmt.Println(">>> Wrong ticket number input")
		}
	}
	if remainingTickets == 0 {
		fmt.Println("The match is booked out, sorry")
	}

	wg.Wait()
}
func greetUser() {
	fmt.Println("Welcome to OLD TRAFFORD ticket system")
	fmt.Println("Todays match, 2:30 pm")
	fmt.Printf("%v \n", matchDay)
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var userTickets uint

	fmt.Printf("There are %v available tickets\n", remainingTickets)
	fmt.Println("Enter First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter Second Name: ")
	fmt.Scan(&secondName)
	fmt.Println("Enter Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, secondName, email, userTickets
}
func getFirstNames(bookedTickets []UserDetails) []string {
	firstNames := []string{}
	for _, booking := range bookedTickets {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func bookTicket(firstName string, secondName string, email string, userTickets uint) {
	var userDetails = UserDetails{
		firstName:   firstName,
		secondName:  secondName,
		email:       email,
		userTickets: userTickets,
	}
	bookedTickets = append(bookedTickets, userDetails)
	remainingTickets = remainingTickets - userTickets
	fmt.Println(">>> Ticket Booked")
}
func sendTicket(userTickets uint, firstName string, secondName string, email string) {
	time.Sleep(5 * time.Second)
	var ticketDetails = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, secondName)
	fmt.Println(">>>>>>>>>>>>")
	fmt.Printf("Sending ticket to email: %v \n %v", email, ticketDetails)
	fmt.Println(">>>>>>>>>>>>")
	wg.Done()
}
