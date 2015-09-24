package main
import (
	"os"
	"io"
	"log"
	"bytes"
	"fmt"
	"encoding/hex"
)


func main() {
	var b bytes.Buffer
	dumper := hex.Dumper(&b)
	n, err := io.Copy(dumper, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(b.String())
	fmt.Println("copied", n, "bytes")
}