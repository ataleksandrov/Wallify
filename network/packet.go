package network

import "github.com/google/gopacket/layers"

type Packet struct {
	header *layers.IPv4
}

func (p *Packet) MatchesSrcIP(srcIP string) bool {
	return p.header.SrcIP.String() == srcIP
}

func (p *Packet) MatchesTTL(ttl uint8) bool {
	return p.header.TTL == ttl
}
