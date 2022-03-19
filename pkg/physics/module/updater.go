package module

func init() {
	RegisterModuleFactory("updater", func() Module {
		return &updater{}
	})
}

type updater struct {
}
