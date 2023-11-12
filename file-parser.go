package main

import (
	"fmt"
	"github.com/hsmade/growatt-sniffer/pkg/decode"
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
	data := decode.Data{}
	_ = decode.UnmarshalBinary(decryptedData, &data)

	// do something with the data
	fmt.Printf("%+v\n", data)
}
