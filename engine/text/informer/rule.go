package informer

import "github.com/alfmartinez/ego/engine/text/grammar"

type SemanticRule interface {
	Name() string
	Matches(*grammar.Statement) bool
	Execute(*grammar.Statement, Semantix)
}

func CreateSemanticRule(name string, match func(*grammar.Statement) bool, exec func(*grammar.Statement, Semantix)) SemanticRule {
	return &semanticRule{
		name:  name,
		match: match,
		exec:  exec,
	}
}

type semanticRule struct {
	name  string
	match func(*grammar.Statement) bool
	exec  func(*grammar.Statement, Semantix)
}

func (r *semanticRule) Name() string {
	return r.name
}

func (r *semanticRule) Matches(s *grammar.Statement) bool {
	return r.match(s)
}

func (r *semanticRule) Execute(s *grammar.Statement, a Semantix) {
	r.exec(s, a)
}
