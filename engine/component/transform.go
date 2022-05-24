package component

type Coords2D struct {
	X, Y float32
}

type Transform2D struct {
	Position Coords2D
	Rotation Coords2D
	Scale    Coords2D
}
