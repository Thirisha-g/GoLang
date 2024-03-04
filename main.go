package main

import (
	"fmt"
	//"strconv"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTickets = 50
var bookings = make([]UserData, 0) //array type

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
	//isOptedInForNewsletter bool
}

func main() {
	greetings()

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("Enter first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTicket)

		//validation
		isValidName, isValidEmail, isValidTicketNumber := validateInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			//call book func
			bookTicket(int(userTicket), firstName, lastName, email)
			sendTicket(uint(userTicket), firstName, lastName, email)
			//call func print first name
			fName := getsFirstName()
			fmt.Printf("The first names of bookings are : %v\n", fName)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out.come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("first name or last name entered is short")
			}
			if !isValidEmail {
				fmt.Printf("email address doesn't contain '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetings() {
	fmt.Println("welcome to our boooking APP")
	fmt.Println("welcome to", conferenceName, "booking application")
	fmt.Println(conferenceTicket, " tickets are available and ", remainingTickets, " tickets are remaining")

	fmt.Println("Get your tickets here to attend")

}

func getsFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func validateInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= uint(remainingTickets)

	return isValidName, isValidEmail, isValidTicketNumber

}

func bookTicket(userTicket int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - int(userTicket)

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets are reaming for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("#######################################")
	fmt.Printf("Sending tickeet:\n %v \nto email address %v\n", tickets, email)
	fmt.Println("#######################################")

}

//city := "London"

// switch city{
// case "New York":

// case "Singapore" , "Hong Kong":

// case "London":

// case "Mexico City":

// case "Berlin":

// default:
// 	fmt.Print("No valid city selected")
// }
