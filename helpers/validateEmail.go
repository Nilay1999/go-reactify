package helpers

import "regexp"

func ValidateEmail(email string) bool {
	// Regular expression for email validation
	// This is a simple pattern, you can use a more comprehensive one if needed
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}
