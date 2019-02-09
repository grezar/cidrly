package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"syscall"
)

func run() int {
	if terminal.IsTerminal(syscall.Stdin) {
		fmt.Fprintf(os.Stderr, "Only support pipeline usage\n")
		return 1
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read JSON from stdin: %v", err)
			return 1
		}
		NewSubnets(bytes).renderTree()
	}
	return 0
}

func main() {
	os.Exit(run())
}
