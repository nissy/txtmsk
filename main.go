package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ngc224/txtmsk/command"
	"github.com/ngc224/txtmsk/keystore"
	"github.com/ngc224/txtmsk/mask"
	"golang.org/x/crypto/ssh/terminal"
)

const applicationName = "txtmsk"

func main() {
	os.Exit(exitcode(run()))
}

func exitcode(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return 1
	}

	return 0
}

func run() error {
	cmd := command.New()

	if cmd.Version {
		fmt.Println("v" + version)
		return nil
	}

	if cmd.Help {
		cmd.ShowHelp()
		return nil
	}

	key := keystore.New(applicationName)
	pw, err := key.Get()

	if cmd.Password || err != nil {
		pw, err = key.Set()

		if err != nil {
			return err
		}
	}

	var text string

	if len(cmd.Args) > 0 {
		text = cmd.Args[0]
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

	m, err := mask.New(pw)

	if err != nil {
		return err
	}

	if cmd.UnMask {
		umText, err := runUnMask(m, text)

		if err != nil {
			return err
		}

		if cmd.Trim {
			umText = trimInLineTag(umText)
		}

		fmt.Println(umText)
		return nil
	}

	mText, err := runMask(m, text)

	if err != nil {
		return err
	}

	fmt.Println(mText)
	return nil
}
