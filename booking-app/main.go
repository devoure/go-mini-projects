package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var matchDay string = "Manchester United VS Chelsea"
var bookedTickets = make([]map[string]string, 0)
var remainingTickets uint = 50

func main() {
	for {
		greetUser()

		firstName, secondName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, secondName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidEmail && isValidName {

			bookTicket(firstName, secondName, email, userTickets)

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
			break
		}
	}
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
func getFirstNames(bookedTickets []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookedTickets {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}
func bookTicket(firstName string, secondName string, email string, userTickets uint) {
	var userDetails = make(map[string]string)
	userDetails["firstName"] = firstName
	userDetails["secondName"] = secondName
	userDetails["email"] = email
	userDetails["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookedTickets = append(bookedTickets, userDetails)
	remainingTickets = remainingTickets - userTickets
	fmt.Println(">>> Ticket Booked")
}
