## Protobuf Sample

Via https://jacobmartins.com/2016/05/24/practical-golang-using-protobuffs/?utm_source=golangweekly&utm_medium=email

To generate the client structure:

<pre>
protoc --go_out=. clientStructure.proto
</pre>

Start the server, then run the client:

<pre>
go run server.go clientStructure.pb.go
go run client.go clientStructure.pb.go
</pre>



## Protobuf - Mac Setup

Might need to install xcode and do this first:

<pre>
sudo xcode-select --install
</pre>

Download the C++ Tarball, then

<pre>
./configure
make
make check
sudo make install
which protoc
protoc --version
</pre>

Then, for go support:

<pre>
go get -u github.com/golang/protobuf/protoc-gen-go
</pre>
