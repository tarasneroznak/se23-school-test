package subscription

import (
	file_storage "main/databases/file-storage"
	email_provider "main/services/email-provider"

	"github.com/ilyakaznacheev/cleanenv"
)

type SendEmailsConfig struct {
	FromEmail string `env:"FROM_EMAIL"`
}

func sendEmails() error {
	emails, err := file_storage.ReadFile("emails")
	if err != nil {
		return err
	}
	var cfg SendEmailsConfig
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return err
	}
	err = email_provider.SendEmail(cfg.FromEmail, emails, "Weekly updates", "Dont miss out on our new content!")
	if err != nil {
		return err
	}
	return nil
}
