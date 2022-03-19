package module

func init() {
	RegisterModuleFactory("gravity", func() Module {
		return &gravity{}
	})
}

type gravity struct {
}
