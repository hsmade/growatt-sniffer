package main

import (
	"flag"
	"fmt"
	"github.com/hsmade/growatt-sniffer/pkg/decode"
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"github.com/hsmade/growatt-sniffer/pkg/tzsp"
	"log"
	"log/slog"
	"os"
)

var (
	port    = flag.Int("address", 9900, "address to listen on")
	verbose = flag.Bool("verbose", false, "enable verbose logging")
)

func main() {
	flag.Parse()

	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	if *verbose {
		programLevel.Set(slog.LevelDebug)
	}

	packets := make(chan []byte, 1)
	go func() {
		slog.Info("starting tzsp listener")
		err := tzsp.Listen(*port, packets)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// handle the received packets
	for {
		select {
		case packet := <-packets:
			slog.Debug("received packet", "size", len(packet))
			if len(packet) == 0 {
				continue
			}

			if len(packet) < 100 {
				slog.Debug("received packet too small", "size", len(packet))
				continue
			}

			// parse packet data into struct
			data := decode.Data{}
			err := decode.UnmarshalBinary(decrypt.Decrypt(packet), &data)
			if err != nil {
				slog.Warn("failed to handle packet", "error", err.Error())
				continue
			}

			// do something with the data
			fmt.Printf("%+v\n", data)
		}
	}
}
