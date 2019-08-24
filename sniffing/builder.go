package sniffing

import (
	environment "github.com/Wallify/env"
	"github.com/Wallify/rule"
)

type SnifferBuilder struct {
	env *environment.Env
}

func NewBuilder(env *environment.Env) (*SnifferBuilder, error) {
	if err := env.Validate(); err != nil {
		return nil, err
	}

	return &SnifferBuilder{
		env: env,
	}, nil
}

func (b *SnifferBuilder) Build() *sniffer {
	env := b.env.RC
	rules := make([]rule.Rule, len(env.DropRules)+len(env.RedirectRules)+1)

	for _, r := range env.DropRules {
		rules = append(rules, rule.NewDropRule(&r))
	}
	for _, r := range env.RedirectRules {
		rules = append(rules, rule.NewRedirectRule(&r))
	}

	return newSniffer(b.env.HC, rules)
}
