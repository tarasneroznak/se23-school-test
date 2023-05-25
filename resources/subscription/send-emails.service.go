package subscription

import (
	file_storage "main/databases/file-storage"
	email_provider "main/services/email-provider"
)

func sendEmails() error {
	emails, err := file_storage.ReadFile("emails")
	if err != nil {
		return err
	}
	err = email_provider.SendEmail("", emails, "", "")
	if err != nil {
		return err
	}
	return nil
}
