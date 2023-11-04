package main

import (
	"fmt"
	"github.com/hsmade/growatt-sniffer/pkg/decode"
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"os"
)

type test struct {
	a int16
}

func main() {
	//a := test{100}
	//var bin_buf bytes.Buffer
	//binary.Write(&bin_buf, binary.LittleEndian, a)
	//fmt.Println(bin_buf.Bytes())

	raw, err := os.ReadFile("data-1699135081.010264.out")
	if err != nil {
		panic(err)
	}

	data := decode.Data{}
	_ = decode.UnmarshalBinary(decrypt.Decrypt(raw), &data)
	fmt.Printf("%+v\n", data)
}
