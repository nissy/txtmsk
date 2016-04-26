package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	ApplicationName = "txtmsk"
	Version         = "1.0.3"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli, err := NewCLI()

	if err != nil {
		return 0
	}

	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return 1
	}

	return 0
}
