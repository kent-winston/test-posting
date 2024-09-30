package tools

import "net/mail"

func CheckEmailValidity(input string) bool {
	var (
		isValid bool = true
	)

	_, err := mail.ParseAddress(input)
	if err != nil {
		isValid = false
	}

	return isValid
}
