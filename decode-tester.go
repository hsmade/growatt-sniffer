package main

import (
	"encoding/binary"
	"fmt"
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"log/slog"
	"os"
)

func main() {
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	decryptedData := decrypt.Decrypt(raw)
	//fmt.Printf("decrypted data: '%s' -> %02x\n", decryptedData, decryptedData)

	fmt.Print("decode as uint8: ")
	for i := 565; i < len(decryptedData); i += 1 {
		fmt.Printf("%d ", binary.BigEndian.Uint16([]byte{0x00, decryptedData[i]}))
	}
	fmt.Println()

	fmt.Print("decode as uint16: ")
	for i := 565; i < len(decryptedData)-1; i += 2 {
		fmt.Printf("%d(", binary.BigEndian.Uint16(decryptedData[i:i+2]))
		fmt.Printf("%d) ", int16(binary.BigEndian.Uint16(decryptedData[i:i+2])))
	}
	fmt.Println()

	fmt.Print("decode as uint32: ")
	for i := 565; i < len(decryptedData)-3; i += 4 {
		fmt.Printf("%d ", binary.BigEndian.Uint32(decryptedData[i:i+4]))
	}
	fmt.Println()

	//data := decode.Data{}
	//_ = decode.UnmarshalBinary(decryptedData, &data)
	//fmt.Printf("%+v\n", data)
}
