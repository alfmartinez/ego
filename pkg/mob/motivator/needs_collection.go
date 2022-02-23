package motivator

type NeedsCollection interface {
	AddNeed(Need, int)
	UpdateNeeds()
	TopNeed() string
	Provide(Need, int)
}

type needsCollection struct {
	needs map[string]NeedLevel
}

func CreateNeedsCollection() NeedsCollection {
	needs := make(map[string]NeedLevel)
	return &needsCollection{needs}
}

func (c *needsCollection) AddNeed(need Need, value int) {
	needLevel := CreateNeedLevel(need, value)
	c.needs[need.Name()] = needLevel
}

func (c *needsCollection) UpdateNeeds() {

}

func (c *needsCollection) TopNeed() string {
	return "curiosity"
}

func (c *needsCollection) Provide(need Need, value int) {

}
