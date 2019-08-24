package network

import "github.com/google/gopacket/layers"

type Frame struct {
	header *layers.Ethernet
}

func (f *Frame) MatchesSrcMAC(srcMac string) bool {
	return f.header.SrcMAC.String() == srcMac
}
