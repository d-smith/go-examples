package timer2
import (
	"time"
	"fmt"
	"log"
)

type command struct {
	opcode string
	args []interface{}
}



type EndToEndTimer struct {
	name string
	start time.Time
	c chan command
	r chan interface{}
	duration time.Duration
	err error
}

func (t *EndToEndTimer) handleTimerOps() {
	for {
		cmd := <- t.c
		log.Println("handle command",cmd.opcode)
		switch cmd.opcode {
		case "kill":
			return
		case "stop":
			t.duration = time.Now().Sub(t.start)
			if len(cmd.args) > 0 {
				theErr,ok := cmd.args[0].(error)
				if ok {
					t.err = theErr
				} else {
					t.err = nil
				}
			}
		case "duration":
			t.r <- t.duration
		default:
			fmt.Println("command", cmd.opcode)
		}
	}
}

func NewEndToEndTimer(name string) *EndToEndTimer {
	e2e := &EndToEndTimer{
		name: name,
		start: time.Now(),
		c: make(chan command),
		r: make(chan interface{}),
	}

	go e2e.handleTimerOps()

	return e2e
}

func (t *EndToEndTimer) Stop(err error) {
	t.c <- command{
		opcode: "stop",
		args: []interface{}{err},
	}
}

func (t *EndToEndTimer) Duration() time.Duration {
	t.c <- command{opcode:"duration"}
	r := <- t.r
	d, ok := r.(time.Duration)

	if ok {
		return d
	} else {
		fmt.Println("Was not ablt to coerce restult to a Duration")
		return 0 * time.Millisecond
	}
}

func (t *EndToEndTimer) Kill() {
	t.c <- command{opcode:"kill"}
}