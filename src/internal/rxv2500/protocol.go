package rxv2500

import (
	"time"

	"github.com/pkg/term"
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
