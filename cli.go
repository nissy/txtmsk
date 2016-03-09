package main

import (
	"bufio"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/ngc224/mask"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

type CLI struct {
}

func NewCLI() *CLI {
	return &CLI{}
}

type CommandlineOptions struct {
	Decrypt  bool `short:"d" long:"decrypt"  description:"Decrypt mode"`
	Password bool `short:"p" long:"password" description:"Set the password"`
	Version  bool `short:"v" long:"version"  description:"Show program's version number"`
}

func (cli *CLI) parseCmdopts() (*CommandlineOptions, []string, error) {
	opts := &CommandlineOptions{}

	parser := flags.NewParser(opts, flags.Default)
	parser.Name = APPLICATION_NAME
	parser.Usage = "[-d] [-p] [-v] text"
	args, err := parser.Parse()

	if err != nil {
		return nil, nil, err
	}

	return opts, args, nil
}

func (cli *CLI) Run() error {
	opts, args, err := cli.parseCmdopts()

	if err != nil {
		return nil
	}

	if opts.Version {
		fmt.Println("Version " + VERSION)
		return nil
	}

	pw, err := GetPassword()

	if opts.Password || err != nil {
		pw, err = SetPassword()

		if err != nil {
			return err
		}
	}

	var text string

	if len(args) > 0 {
		text = args[0]
	} else {
		if terminal.IsTerminal(0) {
			return nil
		}

		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			text += sc.Text() + "\n"
		}

		text = strings.TrimRight(text, "\n")
	}

	m, err := mask.NewMask(pw)

	if err != nil {
		return err
	}

	if opts.Decrypt {
		decrypted_text, err := m.Decrypt(text)

		if err != nil {
			return err
		}

		fmt.Println(decrypted_text)
		return nil
	}

	cipher_text, err := m.Encrypt(text)

	if err != nil {
		return err
	}

	fmt.Println(cipher_text)
	return nil
}
