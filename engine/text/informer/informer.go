package informer

type (
	Object interface {
		SetName(string)
		Name() string
		SetDescription(string)
		Kind() Kind
	}

	object struct {
		name        string
		description string
		kind        Kind
	}

	Kind interface {
	}

	kind struct {
		name    string
		plural  string
		matches []string
		always  []string
		usually []string
		canHave []string
		parent  string
	}
)

var (
	kinds = map[string]Kind{
		"object": &kind{
			name:   "object",
			plural: "objects",
			usually: []string{
				"singular-name not plural-named",
				"improper-named not proper-named",
				"not ambiguously plural",
			},
			matches: []string{"value", "sayable value"},
			canHave: []string{"printed name", "list grouping key", "printed plural name"},
			parent:  "",
		},
		"room": &kind{
			name:    "room",
			plural:  "rooms",
			usually: []string{"lighted not dark", "unvisited not visited"},
			canHave: []string{"description", "map region"},
			parent:  "object",
		},
		"thing": &kind{
			name:    "thing",
			plural:  "things",
			usually: []string{"unlit not lit", "inedible not edible", "portable not fixed in place", "described not undescribed", "unmarked for listing not marked for listing", "mentioned not unmentioned"},
			canHave: []string{"description", "initial appearance"},
			parent:  "object",
		},
	}
)

func CreateObject(kind string) Object {
	return &object{
		kind: kinds[kind],
	}
}

func (o *object) SetName(s string) {
	o.name = s
}

func (o *object) Name() string {
	return o.name
}

func (o *object) SetDescription(s string) {
	o.description = s
}

func (o *object) Kind() Kind {
	return o.kind
}
