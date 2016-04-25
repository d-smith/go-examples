This sample provides a simple event sourcing implementation. This includes an
aggregate (technically an entity as it is not a composite), commands and 
events, event routing and recording, and a method to restore entity state
from a set of events.

This example is slightly non-traditional in that it introduces an event stream
between the aggregate and the event store. The idea is there could be a durable/
reliable transport between the aggregate and the event store, so the events
get flush to the stream, then eventually written to the event store.

The use case is a single app where writes are made, and many readers who want
to build state from the event store.