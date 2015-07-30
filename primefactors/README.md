To build the exe for docker packaging:

<pre>
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o pfservice ./cmd
</pre>
