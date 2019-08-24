package environment

import "fmt"

type HostConfig struct {
	Device  string `mapstructure:"device"`
	Snaplen int32  `mapstructure:"snaplen"`
	Promisc bool   `mapstructure:"promisc"`
}

func (hc *HostConfig) Validate() error {
	if len(hc.Device) == 0 {
		return fmt.Errorf("missing host device")
	}

	if hc.Snaplen == 0 {
		return fmt.Errorf("missing snaplen")
	}

	return nil
}
