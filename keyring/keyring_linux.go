package keyring

import (
	"errors"
	"fmt"
	"os"

	"github.com/Bowery/prompt"
	"github.com/jsipprell/keyctl"
)

type Keyring struct {
	Name string
}

func New(name string) *Keyring {
	return &Keyring{
		Name: name,
	}
}

func (key *Keyring) Set() (string, error) {
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

		if _, err := keyring.Add(key.Name, []byte(pw)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", errors.New("No set password")
}

func (key *Keyring) Get() (string, error) {
	session, err := keyctl.SessionKeyring()

	session.Id()

	if err != nil {
		return "", err
	}

	sessionKey, err := session.Search(key.Name)

	if err != nil {
		return "", err
	}

	pw, err := sessionKey.Get()

	if err != nil {
		return "", err
	}

	return string(pw), nil
}
