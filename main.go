package main

import (
	"fmt"
	"github.com/mp-singh/base64Encode/Encode"
	"os"
)

func main() {
	encode := Encode.Encoding{Input: os.Args[1]}
	encode.Base64Encode()
	fmt.Println(encode.Encoded)
}
