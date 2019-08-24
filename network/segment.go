package network

import "github.com/google/gopacket/layers"

type Segment struct {
	header *layers.TCP
}

func (s *Segment) MatchesDstPort(dstPort uint16) bool {
	return s.header.DstPort == layers.TCPPort(dstPort)
}

func (s *Segment) MatchesURG(urg bool) bool {
	return s.header.URG == urg
}
