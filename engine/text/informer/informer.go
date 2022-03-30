package informer

var (
	kinds = map[string]Kind{
		"object": &objectKind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
	}

	values = make(map[string]ValueKind)
)
