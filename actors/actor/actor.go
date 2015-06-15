package actor

const mailboxSize = 100

type Actor struct {
	Mailbox   chan Message
	ReceiveFn func(chan Message)
}

type Message struct {
	Content interface{}
	Sender  chan interface{}
}

func (a *Actor) Send(msgContent interface{}, replyMailbox chan interface{}) {
	message := Message{msgContent, replyMailbox}
	a.Mailbox <- message
}

func (a *Actor) Run() {
	if a.ReceiveFn == nil {
		panic("No receive function defined for actor")
	}
	for {
		a.ReceiveFn(a.Mailbox)
	}
}
