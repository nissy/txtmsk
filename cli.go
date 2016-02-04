package main

import (
	"bufio"
	"fmt"
	"github.com/Bowery/prompt"
	"github.com/jessevdk/go-flags"
	"github.com/lunixbochs/go-keychain"
	"github.com/ngc224/mask"
	"os"
)

type CLI struct {
}

func NewCLI() *CLI {
	return &CLI{}
}

type CommandlineOptions struct {
	Decrypt  bool `short:"d" long:"decrypt"  description:"Decrypt mode"`
	Password bool `short:"p" long:"password" description:"Set the password in Keychain"`
	Version  bool `short:"v" long:"version"  description:"Show program's version number"`
}

func (cli *CLI) parseCmdopts() (*CommandlineOptions, []string, error) {
	opts := &CommandlineOptions{}

	parser := flags.NewParser(opts, flags.Default)
	parser.Name = APPLICATION_NAME
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
		fmt.Println("Version " + VERSION)
		return nil
	}

	pw, err := GetPassword()

	if opts.Password || err != nil {
		SetPassword()
		return nil
	}

	mask := mask.NewMask(pw)

	var text string

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text += sc.Text() + "\n"
	}

	if opts.Decrypt {
		decrypted_text, err := mask.Decrypt(text)

		if err != nil {
			return err
		}

		fmt.Println(decrypted_text)
		return nil
	}

	cipher_text, err := mask.Encrypt(text)

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

		keychain.Remove(APPLICATION_NAME, "")

		if err := keychain.Add(APPLICATION_NAME, "", pw); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
			continue
		}

		break
	}

	os.Stdin = stdin
}

func GetPassword() (string, error) {
	pw, err := keychain.Find(APPLICATION_NAME, "")

	if err != nil {
		return "", err
	}

	return mask.GetKey(pw)
}
