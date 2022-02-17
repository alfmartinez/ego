package renderer

type NullRenderer struct {
}

func (r *NullRenderer) Init() {}

func (r *NullRenderer) Start() {}

func (r *NullRenderer) Render(s Renderable) {

}
