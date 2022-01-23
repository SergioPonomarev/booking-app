package main

import (
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	userEmail       string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		var firstName, lastName, userEmail, userTickets = getUserInput()

		var isValidName, isValidEmail, isValidTicketNumber = validateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, userEmail)

			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func bookTickets(
	userTickets uint,
	firstName string,
	lastName string,
	userEmail string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		userEmail:       userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName,
		lastName,
		userTickets,
		userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	//firstName = "Sergey"

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	//lastName = "Ponomarev"

	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)
	//userEmail = "s@p.com"

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	//userTickets = 3

	return firstName, lastName, userEmail, userTickets
}

func getFirstNames() []string {
	var result []string
	for _, booking := range bookings {
		result = append(result, booking.firstName)
	}

	return result
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
