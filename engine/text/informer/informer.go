package informer

import (
	"golang.org/x/exp/slices"
)

var (
	kinds = map[string]ObjectKind{
		"object": &objectKind{
			properties: map[string]string{
				"name":   "object",
				"plural": "objects",
			},
		},
		"action": &objectKind{
			properties: map[string]string{
				"name":   "action",
				"plural": "actions",
			},
		},
	}

	values = make(map[string]ValueKind)
)

func FindPropertyByValue(value string) ValueKind {
	for _, e := range values {
		if slices.Contains(e.Values(), value) {
			return e
		}
	}
	return nil
}
