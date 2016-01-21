package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Bowery/prompt"
	"github.com/jessevdk/go-flags"
	"github.com/lunixbochs/go-keychain"
	"io"
	"os"
)

type CLI struct {
}

func NewCLI() *CLI {
	return &CLI{}
}

type CommandlineOptions struct {
	Decrypt  bool `short:"d" long:"dec"      description:"Decrypt mode"`
	Password bool `short:"p" long:"password" description:"Set the password in Keychain"`
	Version  bool `short:"v" long:"version"  description:"Show program's version number"`
}

func (cli *CLI) parseCmdopts() (*CommandlineOptions, []string, error) {
	opts := &CommandlineOptions{}

	parser := flags.NewParser(opts, flags.Default)
	parser.Name = "txtmsk"
	parser.Usage = "[-d] [-p] [-v]"
	args, err := parser.Parse()

	if err != nil {
		return nil, nil, err
	}

	return opts, args, nil
}

func (cli *CLI) Run() error {
	opts, _, err := cli.parseCmdopts()

	if err != nil {
		return nil
	}

	if opts.Version {
		fmt.Println("Version 0.2")
		return nil
	}

	pw, err := GetPassword()

	if opts.Password || err != nil {
		SetPassword()
		return nil
	}

	var text string

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text += sc.Text()
	}

	if opts.Decrypt {
		decrypted_text, err := GetDecryptedText(text, []byte(pw))

		if err != nil {
			return err
		}

		fmt.Println(decrypted_text)
		return nil
	}

	cipher_text, err := GetEncryptedText([]byte(text), []byte(pw))

	if err != nil {
		return err
	}

	fmt.Println(cipher_text)
	return nil
}

func SetPassword() {
	stdin := os.Stdin
	os.Stdin, _ = os.Open("/dev/tty")

	for {
		pw, err := prompt.Password("Set the password in Keychain: ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			break
		}

		if len(pw) > 32 {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", "password len 32 is over")
			continue
		}

		keychain.Remove("txtmsk", "")

		if err := keychain.Add("txtmsk", "", pw); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		break
	}

	os.Stdin = stdin
}

func GetPassword() (string, error) {
	pw, err := keychain.Find("txtmsk", "")

	if err != nil {
		return "", err
	}

	l := len(pw)
	m := 0

	switch {
	case (l < 16):
		m = 16
	case (l < 24):
		m = 24
	case (l < 32):
		m = 32
	case (l > 32):
		return "", errors.New("32 error")
	case (l == 0):
		return "", errors.New("0 error")
	}

	for i := l; i < m; i++ {
		pw += "*"
	}

	return pw, nil
}

func GetEncryptedText(plain_text []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipher_text := make([]byte, aes.BlockSize+len(plain_text))
	iv := cipher_text[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encrypt_stream := cipher.NewCTR(block, iv)
	encrypt_stream.XORKeyStream(cipher_text[aes.BlockSize:], plain_text)

	return base64.StdEncoding.EncodeToString(cipher_text), nil
}

func GetDecryptedText(plain_text string, key []byte) (string, error) {
	cipher_text, err := base64.StdEncoding.DecodeString(plain_text)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	decrypted_text := make([]byte, len(cipher_text[aes.BlockSize:]))
	decrypt_stream := cipher.NewCTR(block, cipher_text[:aes.BlockSize])
	decrypt_stream.XORKeyStream(decrypted_text, cipher_text[aes.BlockSize:])

	return string(decrypted_text), nil
}
