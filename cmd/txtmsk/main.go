package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ngc224/txtmsk"
	"github.com/ngc224/txtmsk/keystore"
	"github.com/ngc224/txtmsk/mask"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	applicationName = "txtmsk"
	version         = "1.2.1"
)

var (
	isUnMask   = flag.Bool("u", false, "unmask mode")
	isPassword = flag.Bool("p", false, "set password")
	isTrim     = flag.Bool("t", false, "trim inline tags (unmask mode only)")
	isVersion  = flag.Bool("v", false, "show version and exit")
	isHelp     = flag.Bool("h", false, "this help")
)

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
	flag.Parse()
	args := flag.Args()

	if *isVersion {
		fmt.Println("v" + version)
		return nil
	}

	if *isHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] textfile\n", os.Args[0])
		flag.PrintDefaults()
		return nil
	}

	key := keystore.New(applicationName)
	pw, err := key.Get()

	if *isPassword || err != nil {
		pw, err = key.Set()

		if err != nil {
			return err
		}
	}

	var fp io.Reader = os.Stdin

	if len(args) > 0 {
		f, err := os.Open(args[0])

		if err != nil {
			return err
		}

		defer f.Close()

		fp = f
	} else {
		if terminal.IsTerminal(0) {
			return nil
		}
	}

	var tBuf bytes.Buffer

	sc := bufio.NewScanner(fp)

	bufLen := bufio.MaxScanTokenSize
	sc.Buffer(make([]byte, bufLen, 1000*bufLen), 1000*bufLen)

	for sc.Scan() {
		tBuf.Write(sc.Bytes())
		tBuf.WriteString("\n")
	}

	if sc.Err() != nil {
		return sc.Err()
	}

	text := strings.TrimRight(tBuf.String(), "\n")

	m, err := mask.New(pw)

	if err != nil {
		return err
	}

	if *isUnMask {
		umText, err := txtmsk.TryUnMask(m, text)

		if err != nil {
			if err != mask.ErrNotUseText {
				fmt.Println(text)
			}

			return err
		}

		if *isTrim {
			umText = txtmsk.TrimInLineTag(umText)
		}

		fmt.Println(umText)

		if umText == text {
			return mask.ErrNotDecrypt
		}

		return nil
	}

	mText, err := txtmsk.TryMask(m, text)

	if err != nil {
		return err
	}

	fmt.Println(mText)
	return nil
}
