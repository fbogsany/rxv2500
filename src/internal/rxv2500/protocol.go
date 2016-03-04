package rxv2500

import (
	"fmt"
	"time"

	"github.com/pkg/term"
)

const (
	etx = '\x03'
	dc1 = '\x11'
	dc2 = '\x12'
)

type T struct {
	t   *term.Term
	err error
}

func Open(tty string) (*T, error) {
	t, err := term.Open(tty, term.RawMode, term.Speed(9600))
	if err != nil {
		return nil, err
	}
	r := &T{t: t}
	if r.start(); r.err != nil {
		return nil, r.err
	}
	return r, nil
}

func (t *T) ready() {
	timeout := 500 // milliseconds
	packet := []byte{dc1, byte(timeout >> 16), byte(timeout >> 8), byte(timeout), etx}
	_, t.err = t.t.Write(packet)
}

func (t *T) readConfiguration(timeout time.Duration) {
	packet := make([]byte, 260)
	count := 9
	count, t.err = t.t.Read(packet[:count])
	if t.err != nil {
		return
	}
	if count != 9 {
		// TODO
	}
	fmt.Printf("DC2: %x\n", packet[0])
	fmt.Printf("ModelID: %s\n", string(packet[1:6]))
	fmt.Printf("Version: %x\n", packet[6])
	count = (int(packet[7]) << 8) | int(packet[8])
	fmt.Printf("Length: %d\n", count)
	count, t.err = t.t.Read(packet[:count+3])
}

func (t *T) start() {
	for i := 0; i < 5; i++ {
		t.ready()
		t.readConfiguration(time.Second)
		if t.err == nil {
			return
		}
	}
}

func (t *T) Close() { _ = t.t.Close() }
