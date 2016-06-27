package keyring

import (
	"errors"
	"fmt"
	"os"

	"github.com/Bowery/prompt"
	"github.com/lunixbochs/go-keychain"
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
		pw, err := prompt.Password("Set the password in keychain: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", "Password len 32 is over")
			continue
		}

		keychain.Remove(key.Name, "")

		if err := keychain.Add(key.Name, "", pw); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", errors.New("No set password")
}

func (key *Keyring) Get() (string, error) {
	pw, err := keychain.Find(key.Name, "")

	if err != nil {
		return "", err
	}

	return pw, nil
}
