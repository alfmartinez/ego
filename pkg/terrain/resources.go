package terrain

type Resource int

const (
	Health Resource = iota
	Food
	Water
)

type Resources interface {
	AddResource(Resource, uint)
	HasResource(Resource) bool
	Consume(Resource)
}

type resources struct {
	res map[Resource]uint
}

func CreateResources() Resources {
	res := make(map[Resource]uint)
	return &resources{res}
}

func (r *resources) AddResource(name Resource, quantity uint) {
	r.res[name] = quantity
}

func (r *resources) HasResource(name Resource) bool {
	res, ok := r.res[name]
	return ok && res > 0
}

func (r *resources) Consume(name Resource) {
	r.res[name] = r.res[name] - 1

}
