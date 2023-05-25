package file_storage

import (
	"errors"
	"os"
	"strings"
)

func createFile(path string) error {
	err := os.WriteFile(path, []byte(""), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getFilepath(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename is empty")
	}
	return "./" + filename + ".txt", nil
}

func ReadFile(filename string) ([]string, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}
	fn, err := getFilepath(filename)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fn)
	if err := createFile(fn); err != nil {
		createFile(fn)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n"), nil
}

func AppendToFile(filename string, text string) error {
	fn, err := getFilepath(filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text + "\n"); err != nil {
		return err
	}
	return nil
}
