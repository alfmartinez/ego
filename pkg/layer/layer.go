package layer

type Layer uint

const (
	FAR_BACKGROUND Layer = iota
	BACKGROUND
	MIDDLEGROUND
	FOREGROUND
	CLOSE_FOREGROUND
)
