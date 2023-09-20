package helper

import "strings"

// Validate User Input
func Validate_user_input(first_name string, last_name string, email string, user_tickets uint, remaining_tickets uint) (bool, bool, bool) {
	//Input Validation
	is_valid_name := len(first_name) >= 2 && len(last_name) >= 2
	is_valid_email := strings.Contains(email, "@")
	is_valid_ticket_number := user_tickets > 0 && user_tickets <= remaining_tickets

	return is_valid_name, is_valid_email, is_valid_ticket_number
}
