package sniffing

import (
	"fmt"
	"log"
	"time"

	environment "github.com/Wallify/env"
	"github.com/Wallify/rule"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type sniffer struct {
	c      *environment.HostConfig
	rules  []rule.Rule
	handle *pcap.Handle
}

func newSniffer(hc *environment.HostConfig, r []rule.Rule) *sniffer {
	return &sniffer{
		c:     hc,
		rules: r,
	}
}

func (s *sniffer) Run() {
	var err error

	c := s.c
	s.handle, err = pcap.OpenLive(c.Device, c.Snaplen, c.Promisc, -1*time.Second)
	if err != nil {
		log.Fatal(fmt.Errorf("could not listen on device: %s error: %s", c.Device, err))
	}
	defer s.handle.Close()

	packetSource := gopacket.NewPacketSource(s.handle, s.handle.LinkType())

	for packet := range packetSource.Packets() {
		for _, r := range s.rules {
			if r.Matches(packet) {
				r.Apply(packet)
				break
			}
		}
	}

}
