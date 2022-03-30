package informer

type (
	Kind interface {
		SetName(string)
		Name() string
		SetDescription(string)
		Description() string
	}

	kind struct {
		properties map[string]string
	}
)

func (k *kind) SetName(s string) {
	k.properties["name"] = s
}

func (k *kind) Name() string {
	return k.properties["name"]
}

func (k *kind) SetDescription(s string) {
	k.properties["description"] = s
}

func (k *kind) Description() string {
	return k.properties["description"]
}
