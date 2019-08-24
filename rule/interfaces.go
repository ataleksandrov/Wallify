package rule

import "github.com/google/gopacket"

type Rule interface {
	Matches(gopacket.Packet) bool
	Apply(gopacket.Packet)
}
