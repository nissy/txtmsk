package command

import (
	"flag"
	"fmt"
	"os"
)

var (
	unmask   = flag.Bool("u", false, "unmask mode")
	password = flag.Bool("p", false, "set password")
	trim     = flag.Bool("t", false, "trim inline tags (unmask mode only)")
	version  = flag.Bool("v", false, "show version and exit")
	help     = flag.Bool("h", false, "this help")
)

type Command struct {
	UnMask   bool
	Password bool
	Trim     bool
	Version  bool
	Help     bool
	Args     []string
}

func New() *Command {
	flag.Parse()

	return &Command{
		UnMask:   *unmask,
		Password: *password,
		Trim:     *trim,
		Version:  *version,
		Help:     *help,
		Args:     flag.Args(),
	}
}

func (cmd *Command) ShowHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] textfile\n", os.Args[0])
	flag.PrintDefaults()
}
