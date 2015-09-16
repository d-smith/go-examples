package main
import (
	"encoding/pem"
	"os"
)


func main() {
	var b pem.Block
	b.Type = "MyStuff"
	b.Bytes = []byte("This is my stuff, yeah")

	pem.Encode(os.Stdout, &b)
}