package main

import "strings"

func validateUserInput(
	firstName string,
	lastName string,
	userEmail string,
	userTickets uint,
	remainingTickets uint) (bool, bool, bool) {

	var isValidName = len(firstName) > 1 && len(lastName) > 1
	var isValidEmail = strings.Contains(userEmail, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
