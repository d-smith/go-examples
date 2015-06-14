This is a project exploring the idea of doing actor based concurrency
in Go Lang, built using the concurrency features of Go Lang.

First step was to define something that you can make that can receive messages
sent to it.

Next step: refactor into something using embedding so those who wish to
send items to an Actor don't have to supply the behavior (expressed in a
receive function) for the actor they are instantiating.
