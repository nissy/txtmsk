// +build !darwin

package keystore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Bowery/prompt"
)

type KeyStore struct {
	name string
	file string
}

func New(name string) *KeyStore {
	return &KeyStore{
		name: name,
		file: filepath.Join(os.Getenv("HOME"), "."+name),
	}
}

func (key *KeyStore) Set() (string, error) {
	stdin := os.Stdin
	os.Stdin, _ = os.Open("/dev/tty")

	for {
		pw, err := prompt.Password("Set the password in file: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n", ErrPassordLenOver.Error())
			continue
		}

		if err := ioutil.WriteFile(key.file, []byte(pw), 0600); err != nil {
			return "", err
		}

		os.Stdin = stdin
		return pw, nil
	}

	return "", ErrNotSetPassword
}

func (key *KeyStore) Get() (string, error) {
	src, err := ioutil.ReadFile(key.file)
	pw := string(src)

	if err != nil || len(pw) == 0 {
		return "", err
	}

	return pw, nil
}
