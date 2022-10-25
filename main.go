package main

import (
	"fmt"
	"strings"
)

func main(){

	var conferenceName  = "Go Conference"
	const conferenceTickets = 50 
	var remainingTickets uint = 50 
	var bookings[] string


	fmt.Printf("Welcome to %v Booking APP\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	for  {

		var firstName string 
		var lastName string 
		var email string
		var userTickets uint



		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		remainingTickets =  remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)


		fmt.Printf("Thank you %v %v for booking  %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickects remaining for %v\n", remainingTickets, conferenceName)

		firstNames := [] string{}

		for _, booking := range bookings { 
			var names = strings.Fields(booking) 
			firstNames = append(firstNames, names[0])
		} 

		fmt.Printf("These are all the bookings  %v \n", firstNames)

	}
}
