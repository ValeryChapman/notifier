package api

import "net/mail"

func emailValidator(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

func emailsValidator(l []string) bool {
	for _, e := range l {
		if !emailValidator(e) {
			return false
		}
	}
	return true
}
