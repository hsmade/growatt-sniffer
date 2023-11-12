package tzsp

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/rs/tzsp"
	"log/slog"
	"net"
)

func Listen(port int, packets chan []byte) error {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return fmt.Errorf("tzsp:Listener: set up listener: %w", err)
	}

	buf := make([]byte, 65535)
	for {
		for {
			bytes, _, err := conn.ReadFrom(buf)
			if err != nil {
				slog.Error("tzsp:Listener: reading from tzsp stream", "error", err)
				continue
			}
			payload, err := tzsp.Parse(buf[:bytes])
			if err != nil {
				slog.Error("tzsp:Listener: parsing tzsp stream", "error", err)
				continue
			}

			ethPacket := gopacket.NewPacket(payload.Data, layers.LayerTypeEthernet, gopacket.Default)
			tcpLayer := ethPacket.Layer(layers.LayerTypeTCP)
			if tcpLayer == nil {
				slog.Debug("tzsp:Listener: received a non-tcp packet")
				continue // not a tcp packet
			}
			tcp, _ := tcpLayer.(*layers.TCP)
			packets <- tcp.Payload
		}
	}
}
