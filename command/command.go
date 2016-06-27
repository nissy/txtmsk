package command

import (
	"flag"
	"fmt"
	"os"
)

var (
	decrypt  = flag.Bool("d", false, "decrypt mode")
	password = flag.Bool("p", false, "set password")
	version  = flag.Bool("v", false, "show version and exit")
	help     = flag.Bool("h", false, "this help")
)

type Command struct {
	Decrypt  bool
	Password bool
	Version  bool
	Help     bool
	Args     []string
}

func New() *Command {
	flag.Parse()

	return &Command{
		Decrypt:  *decrypt,
		Password: *password,
		Version:  *version,
		Help:     *help,
		Args:     flag.Args(),
	}
}

func (cmd *Command) ShowHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] text\n", os.Args[0])
	flag.PrintDefaults()
}
