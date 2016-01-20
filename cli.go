package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/Songmu/prompter"
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
	Password bool `short:"p" long:"password" description:"Set password"`
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
		fmt.Println("Version 0.1")
		return nil
	}

	if opts.Password {
		password := prompter.Password("Set Password")
		keychain.Add("txtmsk", "", password)
		return nil
	}

	var text string

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text += sc.Text()
	}

	key, err := keychain.Find("txtmsk", "")

	if err != nil {
		return nil
	}

	if opts.Decrypt {
		decrypted_text, err := GetDecryptedText(text, []byte(key))

		if err != nil {
			return err
		}

		fmt.Println(decrypted_text)
		return nil
	}

	cipher_text, err := GetEncryptText([]byte(text), []byte(key))

	if err != nil {
		return err
	}

	fmt.Println(cipher_text)
	return nil
}

func GetEncryptText(plain_text []byte, key []byte) (string, error) {
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
