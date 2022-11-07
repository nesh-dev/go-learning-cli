package main

import (
	"fmt"

)

var conferenceName = "Go Conference"
const conferenceTickets = 50 
var remainingTickets uint = 50 
var bookings = make([]UserData, 0);  

type UserData struct { 
	firstName string
	lastName string 
	email string
	userTickets uint
}

func main() {

	greetUsers()

	for {
		

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets) 
		if isValidName && isValidEmail && isValidTicket {
		
			bookTicket(userTickets, firstName, lastName, email)
			sendTickets(userTickets, firstName, lastName, email)
		
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v \n", firstNames)

			if remainingTickets == 0 { 
				fmt.Println("Our conference is booked out come back next year")
				break
			}

		 } else { 

			if !isValidEmail { 
				fmt.Printf("email you entered is not correct \n")
			}

			if !isValidName { 
				fmt.Printf("first name or last name you entered is too short \n")
			}

			if !isValidTicket { 
				fmt.Printf(" Number of tickets entered is invalid \n")
			}
		}

	}
}

func validateUserInputs(firstName, lastName, email string, userTickets uint) {
	panic("unimplemented")
}


func greetUsers() { 
	fmt.Printf("Welcome to %v Booking APP\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() [] string{ 
	firstNames := [] string{}

	for _, booking := range bookings { 
		firstNames = append(firstNames, booking.firstName)
	} 

	return firstNames

}



func getUserInput() (string, string, string, uint){ 
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

	return firstName, lastName, email, userTickets
}


func bookTicket(userTickets uint, firstName string, lastName string, email string,){ 
	remainingTickets =  remainingTickets - userTickets

	var userData =  UserData { 
		firstName: firstName, 
		lastName: lastName, 
		email: email, 
		userTickets: userTickets,
	}



	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking  %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickects remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string,){ 
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)

	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")

}
