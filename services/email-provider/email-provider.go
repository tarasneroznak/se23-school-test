package email_provider

import (
	"fmt"
)

func SendEmail(from string, to []string, subject string, body string) error {
	for _, email := range to {
		fmt.Printf("sending email from: %v to: %v", from, email)
	}
	return nil
}
