package utils

type Attributes interface {
	Attribute(string) string
	SetAttribute(string, string)
}

func CreateAttributes() Attributes {
	return &attributes{
		attributes: map[string]string{},
	}
}

type attributes struct {
	attributes map[string]string
}

func (b *attributes) SetAttribute(t, v string) {
	b.attributes[t] = v
}

func (b *attributes) Attribute(t string) string {
	return b.attributes[t]
}
