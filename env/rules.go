package environment

import (
	"fmt"
)

type rulesConfig struct {
	DropRules     []DropConfig     `mapstructure:"drop"`
	RedirectRules []RedirectConfig `mapstructure:"redirect"`
}

// Validate validates properties of rulesConfig. Implements Validator.
func (rc *rulesConfig) Validate() error {
	// validate all drop rules
	for _, r := range rc.DropRules {
		if err := r.Base.Validate(); err != nil {
			return err
		}
	}

	// validate all redirect rules
	for _, r := range rc.RedirectRules {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// RedirectConfig represents a redirect rule configuration.
type RedirectConfig struct {
	Base        *BaseConfig `mapstructure:"base"`
	Destination string      `mapstructure:"destination"`
}

// Validate validates properties of RedirectConfig. Implements Validator.
func (rc *RedirectConfig) Validate() error {
	if err := rc.Base.Validate(); err != nil {
		return err
	}

	if len(rc.Destination) == 0 {
		return fmt.Errorf("missing destination")
	}

	return nil
}

// DropConfig represents a drop rule configuration.
type DropConfig struct {
	Base *BaseConfig `mapstructure:"base"`
}

// BaseConfig represents a base configuration for a rule.
type BaseConfig struct {
	Layer string `mapstructure:"layer"`
	Parm  string `mapstructure:"parm"`
}

// Validate validates properties of BaseConfig. Implements Validator.
func (c *BaseConfig) Validate() error {
	if len(c.Layer) == 0 {
		return fmt.Errorf("missing layer")
	}
	if len(c.Parm) == 0 {
		return fmt.Errorf("missing parm")
	}
	return nil
}
