package layer

type Layered interface {
	Layer() Layer
	SetLayer(Layer)
}

func CreateLayered() Layered {
	return &layered{}
}

type layered struct {
	layer Layer
}

func (p *layered) Layer() Layer {
	return p.layer
}

func (p *layered) SetLayer(l Layer) {
	p.layer = l
}
