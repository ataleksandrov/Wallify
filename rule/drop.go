package rule

import (
	environment "github.com/Wallify/env"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type DropRule struct {
	layer gopacket.Layer
	c     *environment.DropConfig
}

func NewDropRule(c *environment.DropConfig) *DropRule {
	return &DropRule{
		layer: resolveLayer(c.Base.Layer),
		c:     c,
	}
}

func (r *DropRule) Matches(packet gopacket.Packet) bool {
	layer := packet.Layer(r.layer.LayerType())
	if layer != nil {
		if p, ok := layer.(*layers.Ethernet); ok {

		} else if p, ok := layer.(*layers.IPv4); ok {

		} else if p, ok := layer.(*layers.TCP); ok {

		}
		return false
	}
	return false
}

func (r *DropRule) Apply(packet gopacket.Packet) {
	// nothing to do... dropping packet
}
