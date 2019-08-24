package rule

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func resolveLayer(layer string) gopacket.Layer {
	var l gopacket.Layer

	switch layer {
	case "ethernet":
		l = &layers.Ethernet{}
	case "ipv4":
		l = &layers.IPv4{}
	case "tcp":
		l = &layers.TCP{}
	default:
		log.Fatal(fmt.Errorf("could not create rule: invalid layer name: %s", layer))
	}

	return l
}
