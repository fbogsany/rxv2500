package main

import (
	"flag"
	"fmt"

	"github.com/pkg/term"
)

func main() {
	tty := flag.String("tty", "/dev/ttyUSB0", "communication channel with the RX-v2500")
	flag.Parse()

	run(*tty)
}

func run(tty string) {
	t, err := term.Open(tty, term.RawMode)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	_ = t.Close()
}