package gear

// Gear is the component in charge of carrying, wearing and using items for a character

type Gear interface {
}

func CreateGear() Gear {
	return &gear{}
}

type gear struct{}
