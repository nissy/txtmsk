package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	APPLICATION_NAME = "txtmsk"
	VERSION          = "0.2"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli := NewCLI()
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return 1
	}

	return 0
}
