package handlers

import (
	"net/mail"
	"regexp"
)

func isValidPhoneNumber(phoneNumber string) bool {
	// Regex pattern to match valid phone numbers
	const phoneRegexPattern = `^\+?[0-9]{1,3}?[-.\s]?(\(?[0-9]{1,4}?\))?[-.\s]?[0-9]{1,4}[-.\s]?[0-9]{1,9}$`

	re := regexp.MustCompile(phoneRegexPattern)
	return re.MatchString(phoneNumber)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
