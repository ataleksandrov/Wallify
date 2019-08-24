package rule

import (
	environment "github.com/Wallify/env"
	"github.com/google/gopacket"
)

type RedirectRule struct {
	layer  gopacket.Layer
	config *environment.RedirectConfig
}

func NewRedirectRule(c *environment.RedirectConfig) *RedirectRule {
	return &RedirectRule{
		layer:  resolveLayer(c.Base.Layer),
		config: c,
	}
}

func (r *RedirectRule) Matches(gopacket.Packet) bool {
	return false
}

func (r *RedirectRule) Apply(packet gopacket.Packet) {

}
