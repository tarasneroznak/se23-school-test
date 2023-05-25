package subscription

import (
	"errors"
	file_storage "main/databases/file-storage"

	"github.com/ilyakaznacheev/cleanenv"
)

type SubscribeConfig struct {
	FileNameEmail string `env:"EMAIL_STORAGE_FILEMANE" env-default:"emails"`
}

func subscribe(email string) error {
	var cfg SubscribeConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return err
	}
	emails, err := file_storage.ReadFile(cfg.FileNameEmail)
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
	return file_storage.AppendToFile(cfg.FileNameEmail, email)
}
