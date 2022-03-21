package instruction

type ApiClient interface {
	Global(GlobalAction)
	Item(ItemAction, ...int)
}
