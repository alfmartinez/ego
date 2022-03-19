package module

func init() {
	RegisterModuleFactory("simulate", func() Module {
		return &simulate{}
	})
}

type simulate struct {
}
