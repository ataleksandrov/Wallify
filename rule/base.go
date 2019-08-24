package rule

import (
	"github.com/Wallify/network"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type baseRule struct {
	layer gopacket.Layer
}

func (r *baseRule) Matches(packet gopacket.Packet) bool {
	layer := packet.Layer(r.layer.LayerType())
	if layer != nil {
		if h, ok := layer.(*layers.Ethernet); ok {
			f := network.Frame{h}
			return f.MatchesSrcMAC()

		} else if h, ok := layer.(*layers.IPv4); ok {
			p := network.Packet{h}
			return p.MatchesSrcIP() || p.MatchesTTL()

		} else if h, ok := layer.(*layers.TCP); ok {
			s := network.Segment{h}
			return s.MatchesDstPort() || s.MatchesURG()
		}

		return false
	}
	return false
}

func (r *baseRule) Apply(packet gopacket.Packet) {
	// nothing to do here
}
