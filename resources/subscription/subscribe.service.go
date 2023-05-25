package subscription

import (
	"errors"
	file_storage "main/databases/file-storage"
)

func subscribe(email string) error {
	emails, err := file_storage.ReadFile("emails")
	if err != nil {
		return err
	}
	emailsset := make(map[string]bool)
	for _, email := range emails {
		emailsset[email] = true
	}
	if emailsset[email] {
		return errors.New("email already exists")
	}
	file_storage.AppendToFile("emails", email)
	return nil
}
