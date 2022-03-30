package informer

type (
	Kind interface {
	}

	ObjectKind interface {
		SetName(string)
		Name() string
		SetDescription(string)
		Description() string
		SetProperty(string, string)
		GetProperty(string) string
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

func (k *objectKind) SetName(s string) {
	k.properties["name"] = s
}

func (k *objectKind) Name() string {
	return k.properties["name"]
}

func (k *objectKind) SetDescription(s string) {
	k.properties["description"] = s
}

func (k *objectKind) Description() string {
	return k.properties["description"]
}

func (k *objectKind) SetProperty(property, value string) {
	k.properties[property] = value
}

func (k *objectKind) GetProperty(property string) string {
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
