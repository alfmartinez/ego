package informer

type (
	Kind interface {
	}

	ObjectKind interface {
		Set(string, string)
		Get(string) string
	}

	ValueKind interface {
		SetValues([]string)
		Values() []string
		Name() string
	}

	objectKind struct {
		parent     Kind
		properties map[string]string
	}

	valueKind struct {
		name   string
		values []string
	}
)

func (k *objectKind) Set(property, value string) {
	k.properties[property] = value
}

func (k *objectKind) Get(property string) string {
	return k.properties[property]
}

func (v *valueKind) SetValues(values []string) {
	v.values = values
}

func (v *valueKind) Values() []string {
	return v.values
}

func (v *valueKind) Name() string {
	return v.name
}
