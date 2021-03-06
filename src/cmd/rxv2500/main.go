package main

import (
	"flag"
	"fmt"

	"internal/rxv2500"
)

func main() {
	tty := flag.String("tty", "/dev/ttyUSB0", "communication channel with the RX-v2500")
	flag.Parse()

	run(*tty)
}

func run(tty string) {
	t, err := rxv2500.Open(tty)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	t.Close()
}
