package main

import (
	"fmt"
	//"strings"
)

// Conference booking application
func main() {
	var conferenceName = "Kubecon"
	// Another way to use vars
	// conferenceName := "Kubecon"
	const conferenceTickets = 100
	var remainingTickets = 100

	// Arrays
	var bookings [50]string

	fmt.Println("Welcome to", conferenceName, "conference")
	fmt.Println("We have a total of", conferenceTickets, "tickets and now we are left with", remainingTickets, "only")
	// Another method
	// fmt.Printf("We have a total of %v tickets and now we are left with %v only", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int // uint for +ve
	// "&" is pointer/special var

	fmt.Println("First name: ")
	fmt.Scan(&firstName)
	fmt.Println("Last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Email ID: ")
	fmt.Scan(&email)
	fmt.Println("How many Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - int(userTickets)
	// Arrays
	// bookings[0] = firstName + " " + lastName
	// Slices
	//bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole array: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your tickets within 1hr on %v", firstName, lastName, userTickets, email)
	fmt.Printf("Now we are left with %v tickets", remainingTickets)

	// %T is the type of the variable you are using

	// Slices
	// What if we don't know the size of array, when creating it or 1 user may book all tickets and having an array with size 50 with only 1 element inside is not efficient
	// Slice is basically an abstraction of an array, they are flexible arrays.

	// Loops
	// Infinite
	// for {
	// }
	// for-each loop
	//firstNames := []string{}
	//for _, booking := range bookings{
	// If you are not using "index," var then use _,
	//	var names = strings.Fields(booking)
	//	firstNames = append(firstNames, names[0])
	//}
	//fmt.Println(bookings)

	// If-else, else-if conditions
	if userTickets > remainingTickets {
		fmt.Println("All tickets are sold")
		// break and continue statement
	}
}
