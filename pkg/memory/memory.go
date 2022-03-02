package memory

type Memory interface {
	Interested
	RemembersPlaces
}

type memory struct {
	RemembersPlaces
	Interested
}

func CreateMemory() Memory {
	places := CreatePlaces()
	var i = CreateInterests()
	return &memory{places, i}
}
