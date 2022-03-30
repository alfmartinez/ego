package informer

var (
	kinds = map[string]Kind{
		"object": &kind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
	}
)
