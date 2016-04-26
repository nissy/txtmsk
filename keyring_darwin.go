package main

import (
	"errors"
	"fmt"
	"github.com/Bowery/prompt"
	"github.com/lunixbochs/go-keychain"
	"os"
)

func SetPassword() (string, error) {
	stdin := os.Stdin
	os.Stdin, _ = os.Open("/dev/tty")

	for {
		pw, err := prompt.Password("Set the password in Keychain: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", "Password len 32 is over")
			continue
		}

		keychain.Remove(ApplicationName, "")

		if err := keychain.Add(ApplicationName, "", pw); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", errors.New("Error: No set password ")
}

func GetPassword() (string, error) {
	pw, err := keychain.Find(ApplicationName, "")

	if err != nil {
		return "", err
	}

	return pw, nil
}
