package instruction

type ApiClient interface {
	Call(realm byte, action byte, data ...byte) int
}
