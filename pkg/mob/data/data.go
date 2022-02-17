package data

type Data struct {
	name string
}

func CreateMobData(name string) *Data {
	return &Data{name}
}

func (m Data) Name() string {
	return m.name
}
