package importas

import (
	"errors"
	"fmt"
	"regexp"
)

type Config struct {
	RequiredAlias        map[string]string
	DisallowUnaliased    bool
	DisallowExtraAliases bool
}

type CompiledConfig struct {
	Config
	rules []*Rule
}

func (c Config) Compile() (*CompiledConfig, error) {
	rules := make([]*Rule, 0, len(c.RequiredAlias))
	for path, alias := range c.RequiredAlias {
		reg, err := regexp.Compile(fmt.Sprintf("^%s$", path))
		if err != nil {
			return nil, err
		}

		rules = append(rules, &Rule{
			Regexp: reg,
			Alias:  alias,
		})
	}

	return &CompiledConfig{
		Config: c,
		rules:  rules,
	}, nil
}

func (c *CompiledConfig) AliasFor(path string) (string, bool) {
	rule := c.findRule(path)
	if rule == nil {
		return "", false
	}

	alias, err := rule.aliasFor(path)
	if err != nil {
		return "", false
	}

	return alias, true
}

func (c *CompiledConfig) findRule(path string) *Rule {
	for _, rule := range c.rules {
		if rule.Regexp.MatchString(path) {
			return rule
		}
	}

	return nil
}

type Rule struct {
	Alias  string
	Regexp *regexp.Regexp
}

func (r *Rule) aliasFor(path string) (string, error) {
	str := r.Regexp.FindString(path)
	if len(str) > 0 {
		return r.Regexp.ReplaceAllString(str, r.Alias), nil
	}

	return "", errors.New("mismatch rule")
}
