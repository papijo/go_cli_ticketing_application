package main

import (
	"booking-app/config"
	"booking-app/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mail.v2"
)

// Application Variables
var conference_name = "Go Conference"

const conference_tickets int = 50

var remaining_tickets uint = 50

var bookings = make([]UserData, 0)

//Structs

type UserData struct {
	first_name        string
	last_name         string
	email             string
	phonenumber       string
	number_of_tickets uint
}

// Creating a wait group (group of threads that the main thread should wait for before termininating)
var wg = sync.WaitGroup{}

func main() {

	//Load Environment Variables
	envErr := config.LoadEnvVariables()

	if envErr != nil {
		log.Fatal("Error loading environment variable file: ", envErr)
	}

	//Conditional For Loops

	//Welcome Users to conference
	greet_users()

	//Get Input from Users
	first_name, last_name, email, phonenumber, user_tickets := get_user_input()

	//Validate User Input
	is_valid_name, is_valid_email, is_valid_ticket_number := helper.Validate_user_input(first_name, last_name, email, user_tickets, remaining_tickets)

	if is_valid_name && is_valid_email && is_valid_ticket_number {

		//Book Ticket
		book_ticket(user_tickets, first_name, last_name, phonenumber, email)

		//add a thread/goroutine to a wait group
		wg.Add(2)
		//Send Ticket to User's email
		go send_ticket(user_tickets, first_name, last_name, email) //go starts a new go-routine or application thread
		go send_sms(phonenumber, user_tickets)
		//Call function to print first names
		first_names := get_first_names()
		fmt.Printf("These are all our bookings: %v\n", first_names)

		var no_tickets_remaining bool = remaining_tickets == 0

		if no_tickets_remaining {
			// End Program
			fmt.Println("Our conference is booked out. Come back next year")
			//break //Breaks the Loop
		}
	} else {
		if !is_valid_name {
			fmt.Println("First Name or Last Name is too short")
		}

		if !is_valid_email {
			fmt.Println("Email Address is in wrong format")
		}

		if !is_valid_ticket_number {
			fmt.Println("Number of tickets you entered is invalid")
		}
		fmt.Println("Your input data is invalid, Try again")

	}

	wg.Wait()

}

//Application Helper Functions

// Greet Users on Entry to Application
func greet_users() {
	fmt.Printf("Welcome to our %v. \n", conference_name)
	fmt.Printf("Welcome to %v booking application.\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conference_tickets, remaining_tickets)

}

// Get First Names of Users
func get_first_names() []string {
	first_names := []string{}
	for _, booking := range bookings { //_ signifies a blank identifier

		first_names = append(first_names, booking.first_name)
	}
	return first_names
}

// Get User Input
func get_user_input() (string, string, string, string, uint) {
	var first_name string
	var last_name string
	var phonenumber string
	var email string
	var user_tickets uint
	// var city string

	// // Ask users to ask for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&first_name)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&last_name)
	fmt.Println("Enter your Phone Number: ")
	fmt.Scan(&phonenumber)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&user_tickets)

	return first_name, last_name, email, phonenumber, user_tickets
}

func book_ticket(user_tickets uint, first_name string, last_name string, phonenumber string, email string) {
	remaining_tickets = remaining_tickets - user_tickets

	//Create a map for a user
	// var user_data = make(map[string]string) //`make` keyword can be used to make structures
	// user_data["first_name"] = first_name
	// user_data["last_name"] = last_name
	// user_data["email"] = email
	// user_data["number_of_tickets"] = strconv.FormatUint(uint64(user_tickets), 10)

	//Create a struct to hold user data
	user_data := UserData{
		first_name:        first_name,
		last_name:         last_name,
		email:             email,
		phonenumber:       phonenumber,
		number_of_tickets: user_tickets,
	}

	bookings = append(bookings, user_data)
	fmt.Printf("List of bookings is %v.\n", bookings)

	fmt.Printf("User: %v booked %v tickets.\n", first_name, user_tickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", first_name, last_name, user_tickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conference_name)
}

// Extend CLI Application
// Send Email Function
func send_email(subject, body, to string) error {

	var email string = os.Getenv("SMTP_EMAIL_SENDER")
	var password string = os.Getenv("SMTP_EMAIL_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("SMTP_EMAIL_PORT"))
	var email_url string = os.Getenv("SMTP_EMAIL_URL")

	if err != nil {
		log.Fatalln("Check Port Data Type")
	}

	from := email

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer(email_url, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	fmt.Printf("Email sent to %s\n", to)
	return nil
}

// Send ticket in Email
func send_ticket(user_tickets uint, first_name string, last_name string, email string) {
	time.Sleep(10 * time.Second)

	subject := "Thank you for purchasing Go Conference Tickets"
	var body string
	if user_tickets > 1 {
		body = fmt.Sprintf("You purchased %v tickets for %v %v", user_tickets, first_name, last_name)
	} else {
		body = fmt.Sprintf("You purchased %v ticket for %v %v", user_tickets, first_name, last_name)
	}

	fmt.Println("########################")
	fmt.Printf("Sending ticket to email address %v. \n", email)
	fmt.Println("########################")

	send_email(subject, body, email)
	wg.Done()

}

// Send Ticket Notification in SMS
func send_sms(phonenumber string, user_tickets uint) error {
	url := os.Getenv("TERMII_URL")
	apiKey := os.Getenv("TERMII_API_KEY")
	account := os.Getenv("TERMII_ACCOUNT")

	data := map[string]interface{}{
		"to":      "234" + phonenumber[1:],
		"from":    account,
		"sms":     fmt.Sprintf("You purchased %v number of tickets for Go Conference.", user_tickets),
		"type":    "plain",
		"api_key": apiKey,
		"channel": "generic",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error encoding JSON data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	fmt.Println("Response:", response)

	if response["code"] == "ok" {
		fmt.Println("SMS Sent")
	}

	wg.Done()
	return nil
}
