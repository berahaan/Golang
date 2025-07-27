package main

import "fmt"

func main() {
	println("Hello welcome to go conference here in 2023")
	// let us learn variables here
	// var name = "John doe"
	var firstName string
	var lastName string
	var numberOffTicket int
	var email string
	remainTickets := 100

	fmt.Println("Enter your first name please:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name please:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Emails please:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to buy please::")
	fmt.Scan(&numberOffTicket)

	fmt.Printf("Thank you %v %v for buying %v tickets you will receive emails confirmations at %v\n", firstName, lastName, numberOffTicket, email)
	remain := remainTickets - numberOffTicket
	fmt.Println("Remaining tickets are :", remain)

}
