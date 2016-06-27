package keyring

import (
	"errors"
	"fmt"
	"os"

	"github.com/Bowery/prompt"
	"github.com/jsipprell/keyctl"
)

func SetPassword() (string, error) {
	stdin := os.Stdin
	os.Stdin, _ = os.Open("/dev/tty")

	for {
		pw, err := prompt.Password("Set the password in keyring: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", "Password len 32 is over")
			continue
		}

		keyring, err := keyctl.SessionKeyring()

		if err != nil {
			return "", err
		}

		if _, err := keyring.Add(ApplicationName, []byte(pw)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", errors.New("No set password")
}

func GetPassword() (string, error) {
	keyring, err := keyctl.SessionKeyring()

	keyring.Id()

	if err != nil {
		return "", err
	}

	key, err := keyring.Search(ApplicationName)

	if err != nil {
		return "", err
	}

	pw, err := key.Get()

	if err != nil {
		return "", err
	}

	return string(pw), nil
}
