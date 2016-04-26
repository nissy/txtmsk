package main

import (
	"bufio"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/ngc224/txtmsk/mask"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

type Command struct {
	Decrypt  bool `short:"d" long:"decrypt"  description:"Decrypt mode"`
	Password bool `short:"p" long:"password" description:"Set the password"`
	Version  bool `short:"v" long:"version"  description:"Show program's version number"`
}

type CLI struct {
	Command *Command
	Args    []string
}

var text string

func NewCLI() (*CLI, error) {
	opts, args, err := ParseCmdopts()

	if err != nil {
		return nil, err
	}

	return &CLI{
		Command: opts,
		Args:    args,
	}, nil
}

func ParseCmdopts() (*Command, []string, error) {
	opts := &Command{}

	parser := flags.NewParser(opts, flags.Default)
	parser.Name = ApplicationName
	parser.Usage = "[-d] [-p] [-h] [-v] TEXT"
	args, err := parser.Parse()

	if err != nil {
		return nil, nil, err
	}

	return opts, args, nil
}

func (cli *CLI) Run() error {
	if cli.Command.Version {
		fmt.Println("Version " + Version)
		return nil
	}

	pw, err := GetPassword()

	if cli.Command.Password || err != nil {
		pw, err = SetPassword()

		if err != nil {
			return err
		}
	}

	if len(cli.Args) > 0 {
		text = cli.Args[0]
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

	if cli.Command.Decrypt {
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
