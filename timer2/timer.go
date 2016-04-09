package timer2

import (
	"fmt"
	"log"
	"time"
)

//Functions of the timer are accessed using commands
type command struct {
	opcode string
	args   []interface{}
}

//Opcodes for commands
const (
	killOp     = "kill"
	stopOp     = "stop"
	durationOp = "duration"
)

//EndToEnd timer is an opaque data type handed out to timer consumers. It exposes
//several methods, but allows direct access to the data members only from a goroutine
//spawned to manage the timer state.
type EndToEndTimer struct {
	name     string
	start    time.Time
	c        chan command
	r        chan interface{}
	duration time.Duration
	err      error
}

func (t *EndToEndTimer) handleStop(cmd command) {
	t.duration = time.Now().Sub(t.start)
	if len(cmd.args) > 0 {
		theErr, ok := cmd.args[0].(error)
		if ok {
			t.err = theErr
		} else {
			t.err = nil
		}
	}
}

func (t *EndToEndTimer) handleDuration() {
	t.r <- t.duration
}

//handleTimerOps is the internal go routine responsible for accessing the timer internals
func (t *EndToEndTimer) handleTimerOps() {
	for {
		cmd := <-t.c
		log.Println("handle command", cmd.opcode)
		switch cmd.opcode {
		case killOp:
			return
		case stopOp:
			t.handleStop(cmd)
		case durationOp:
			t.handleDuration()
		default:
			fmt.Println("command", cmd.opcode)
		}
	}
}

func NewEndToEndTimer(name string) *EndToEndTimer {
	e2e := &EndToEndTimer{
		name:  name,
		start: time.Now(),
		c:     make(chan command),
		r:     make(chan interface{}),
	}

	go e2e.handleTimerOps()

	return e2e
}

func (t *EndToEndTimer) Stop(err error) {
	t.c <- command{
		opcode: stopOp,
		args:   []interface{}{err},
	}
}

func (t *EndToEndTimer) Duration() time.Duration {
	t.c <- command{opcode: durationOp}
	r := <-t.r
	d, ok := r.(time.Duration)

	if ok {
		return d
	} else {
		fmt.Println("Was not ablt to coerce restult to a Duration")
		return 0 * time.Millisecond
	}
}

func (t *EndToEndTimer) Kill() {
	t.c <- command{opcode: killOp}
}
