package informer

var (
	kinds = map[string]Kind{
		"object": &kind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
		"room": &kind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
		"thing": &kind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
	}
)
