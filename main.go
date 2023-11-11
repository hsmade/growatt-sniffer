package main

import (
	"flag"
	"fmt"
	"github.com/hsmade/growatt-sniffer/pkg/decode"
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"github.com/rs/tzsp"
	"log"
	"log/slog"
	"net"
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

	// -------
	//raw, err := os.ReadFile("data-1699135081.012552.out")
	//if err != nil {
	//	panic(err)
	//}
	//
	//data := decode.Data{}
	//_ = decode.UnmarshalBinary(decrypt.Decrypt(raw), &data)
	//fmt.Printf("%+v\n", data)
	// -------

	// setup listener for tzsp stream
	addr := net.UDPAddr{
		Port: *port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	// read tzsp stream and send the received packets to the channel
	packets := make(chan []byte, 1)
	go func() {
		buf := make([]byte, 65535)
		for {
			bytes, _, err := conn.ReadFrom(buf)
			if err != nil {
				slog.Error("reading from tzsp stream", "error", err)
				continue
			}
			packet, err := tzsp.Parse(buf[:bytes])
			if err != nil {
				slog.Error("parsing tzsp stream", "error", err)
				continue
			}
			packets <- packet.Data
		}
	}()

	// handle the received packets
	for {
		select {
		case packet := <-packets:
			slog.Debug("received packet", "size", len(packet))
			if len(packet) < 93 {
				slog.Warn("received packet too small", "size", len(packet))
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
