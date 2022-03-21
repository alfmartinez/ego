package layer

type Layer uint

const (
	NONE Layer = iota
	FAR_BACKGROUND
	BACKGROUND
	MIDDLEGROUND
	FOREGROUND
	CLOSE_FOREGROUND
)
