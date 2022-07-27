package main

import (
	"bytes"
	"fmt"
	"github.com/iralance/go-gateway/demo/base/unpack/unpack"
)

func main() {
	bytesBuffer := bytes.NewBuffer([]byte{})

	if err := unpack.Encode(bytesBuffer, "hello world 0!!!"); err != nil {
		panic(err)
	}
	if err := unpack.Encode(bytesBuffer, "hello world 1!!!"); err != nil {
		panic(err)
	}

	for {
		if bt, err := unpack.Decode(bytesBuffer); err == nil {
			fmt.Println(string(bt))
			continue
		}
		break
	}
}
