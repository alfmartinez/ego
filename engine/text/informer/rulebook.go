package informer

import (
	"strings"
)

type RuleBook interface {
	Name() string
	Matches(string) bool
	Applies(o Object) bool
}

func CreateRuleBook(name string, args ...string) RuleBook {
	var kind = ""
	if len(args) > 0 {
		kind = args[0]
	}

	return &rulebook{
		name: name,
		kind: kind,
	}
}

type rulebook struct {
	name, kind string
}

func (b *rulebook) Name() string {
	return b.name
}

func (b *rulebook) Matches(search string) bool {
	return strings.Contains(search, b.name)
}

func (b *rulebook) Applies(o Object) bool {
	return b.kind == "" || o.IsKind(b.kind)
}
