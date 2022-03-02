package data

type Data interface {
	Name() string
}

type data struct {
	name string
}

func CreateData(name string) Data {
	return &data{name}
}

func (m data) Name() string {
	return m.name
}