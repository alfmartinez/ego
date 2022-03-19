package module

func init() {
	RegisterModuleFactory("collider", func() Module {
		return &collider{}
	})
}

type collider struct {
}
