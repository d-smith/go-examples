package actor

const mailboxSize = 100

type actor struct {
	mailbox   chan Message
	receiveFn func(chan Message)
}

type Message struct {
	Content interface{}
	Sender  chan interface{}
}

func NewActor(receiveFn func(chan Message)) *actor {
	return &actor{
		mailbox:   make(chan Message, mailboxSize),
		receiveFn: receiveFn,
	}
}

func (a *actor) Send(msgContent interface{}, replyMailbox chan interface{}) {
	message := Message{msgContent, replyMailbox}
	a.mailbox <- message
}

func (a *actor) Run() {
	for {
		a.receiveFn(a.mailbox)
	}
}
