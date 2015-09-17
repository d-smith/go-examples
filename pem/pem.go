package main
import (
	"encoding/pem"
	"fmt"
)


const twoBlocks = `-----BEGIN Stuff1-----
c3R1ZmYgIGZyb20gYmxvY2sgMQ==
-----END Stuff1-----
-----BEGIN Stuff2-----
b2YgYmxvY2sgMiBzdHVmZiwgeWVz
-----END Stuff2-----`

func encodeAndDecode() {
	var b pem.Block
	b.Type = "MyStuff"
	b.Bytes = []byte("This is my stuff, yeah")

	encoded := pem.EncodeToMemory(&b)

	decoded, theRest := pem.Decode(encoded)

	fmt.Println("encode and decode")
	fmt.Printf("\tdecoded: type - %v\n", decoded.Type)
	fmt.Printf("\tdecoded bytes as string: %s\n", string(decoded.Bytes))
	fmt.Println("\tthe rest: ", theRest)
}

func decodeGarbage() {
	decoded, theRest := pem.Decode([]byte("this is not PEM encoded"))
	fmt.Println("decodeGarbage")
	fmt.Println("\tdecoded: ", decoded)
	fmt.Println("\tthe rest as string: ", string(theRest))
}

func decodeMultiBlock() {
	decoded, theRest := pem.Decode([]byte(twoBlocks))
	fmt.Println("decodeMultiblock")
	fmt.Printf("\tdecoded: type - %v\n", decoded.Type)
	fmt.Printf("\tdecoded bytes as string: %s\n", string(decoded.Bytes))
	fmt.Println("\tmore bytes: ", len(theRest) > 0)

	decoded, theRest = pem.Decode(theRest)
	fmt.Printf("\tthe rest decoded: type - %v\n", decoded.Type)
	fmt.Printf("\tthe rest decoded bytes as string: %s\n", string(decoded.Bytes))
	fmt.Println("\tmore bytes: ", len(theRest) > 0)
}


func main() {
	encodeAndDecode()
	decodeGarbage()
	decodeMultiBlock()
}