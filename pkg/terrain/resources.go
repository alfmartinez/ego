package terrain

type Resources interface {
	AddResource(string, uint)
	HasResource(string) bool
	Consume(string)
}

type resources struct {
	res map[string]uint
}

func CreateResources() Resources {
	res := make(map[string]uint)
	return &resources{res}
}

func (r *resources) AddResource(name string, quantity uint) {
	r.res[name] = quantity
}

func (r *resources) HasResource(name string) bool {
	res, ok := r.res[name]
	return ok && res > 0
}

func (r *resources) Consume(name string) {
	r.res[name] = r.res[name] - 1

}
