package environment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Env represents an abstraction where the configuration of the Wallify will be loaded
type Env struct {
	HC *hostConfig  `mapstructure:"host"`
	RC *rulesConfig `mapstructure:"rules"`
}

// Validate validates that the configuration contains all mandatory properties. Implements Validator.
func (e *Env) Validate() error {
	validatable := []Validator{e.RC, e.HC}

	for _, item := range validatable {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// NewDefault creates new default environment
func NewDefault() *Env {
	v := newDefaultViper()

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Could not load configuration: %s", err))
	}

	e := &Env{}
	if err := v.Unmarshal(e); err != nil {
		log.Fatal(fmt.Sprintf("Could not unmarshal environment: %s", err))
	}

	return e
}

func newDefaultViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	return v
}
