Generating code from the proto files

<pre>
protoc ./person.proto --go_out=plugins=grpc:.
</pre>

Cribbing from https://talks.golang.org/2015/gotham-grpc.slide#1

After code generation, we a golang interface for our service definition.

<pre>
type DirectoryServer interface {
	LookupPersonByName(context.Context, *NameRequest) (*Person, error)
}
</pre>

Next, implement the method is a server.



