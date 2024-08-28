package handlers

import (
	"net/mail"
	"regexp"
)

func isValidPhoneNumber(phoneNumber string) bool {
	// Updated regex pattern for international phone numbers
	const phoneRegexPattern = `^\+\d{1,15}$`

	re := regexp.MustCompile(phoneRegexPattern)
	return re.MatchString(phoneNumber)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
