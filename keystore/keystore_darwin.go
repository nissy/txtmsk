package keystore

import (
	"fmt"
	"os"

	"github.com/Bowery/prompt"
	"github.com/lunixbochs/go-keychain"
)

type KeyStore struct {
	name string
}

func New(name string) *KeyStore {
	return &KeyStore{
		name: name,
	}
}

func (key *KeyStore) Set() (string, error) {
	stdin := os.Stdin
	os.Stdin, _ = os.Open("/dev/tty")

	for {
		pw, err := prompt.Password("Set the password in keychain: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n", ErrPassordLenOver.Error())
			continue
		}

		keychain.Remove(key.name, "")

		if err := keychain.Add(key.name, "", pw); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", ErrNotSetPassword
}

func (key *KeyStore) Get() (string, error) {
	pw, err := keychain.Find(key.name, "")

	if err != nil {
		return "", err
	}

	return pw, nil
}
