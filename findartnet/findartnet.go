package findartnet

import (
	"fmt"
	"net"
	"time"

	"github.com/jsimonetti/go-artnet"
	"github.com/jsimonetti/go-artnet/packet"
)

type artnetInfo struct {
	IP   net.IP
	Name string
}

func FindArtnet() []artnetInfo {

	dst := fmt.Sprintf("%s:%d", "255.255.255.255", packet.ArtNetPort)
	broadcastAddr, _ := net.ResolveUDPAddr("udp", dst)
	src := fmt.Sprintf("%s:%d", "192.168.11.219", packet.ArtNetPort)
	localAddr, _ := net.ResolveUDPAddr("udp", src)

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("error opening udp: %s\n", err)
		return nil
	}

	p := &packet.ArtPollPacket{}
	b, err := p.MarshalBinary()
	if err != nil {
		fmt.Printf("error marshalling packet: %s\n", err)
		return nil
	}

	n, err := conn.WriteTo(b, broadcastAddr)
	if err != nil {
		fmt.Printf("error writing packet: %s\n", err)
		return nil
	}
	fmt.Printf("packet sent, wrote %d bytes\n", n)

	// wait 5 seconds for a reply
	timer := time.NewTimer(5 * time.Second)

	recvAddr := make(chan artnetInfo)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, addr, err := conn.ReadFromUDP(buf) // First packet you read might be your own
			if err != nil {
				fmt.Printf("error reading packet: %s\n", err)
				continue
			}

			// Skip packets from the local machine
			if addr.IP.Equal(localAddr.IP) {
				continue
			}

			// Check if addr is valid and not nil
			if addr != nil && addr.IP != nil {
				p, _ := packet.Unmarshal(buf[:n])
				cf := artnet.ConfigFromArtPollReply(*p.(*packet.ArtPollReplyPacket))
				recvAddr <- artnetInfo{cf.IP, cf.Name}
			} else {
				fmt.Println("Received invalid address.")
			}
		}
	}()
	artnetAddrs := []artnetInfo{}
	for {
		select {
		case ip := <-recvAddr:
			artnetAddrs = append(artnetAddrs, ip)
		case <-timer.C:
			fmt.Printf("timeout reached\n")
			return artnetAddrs
		}
	}
}
