package motivator

import "sort"

type NeedsCollection interface {
	AddNeed(Need, int)
	UpdateNeeds()
	TopNeed() Need
	Provide(Need, int, int)
}

type needsCollection struct {
	needs map[Need]NeedLevel
}

func CreateNeedsCollection() NeedsCollection {
	needs := make(map[Need]NeedLevel)
	return &needsCollection{needs}
}

func (c *needsCollection) AddNeed(need Need, value int) {
	needLevel := CreateNeedLevel(need, value)
	c.needs[need] = needLevel
}

func (c *needsCollection) UpdateNeeds() {
	for _, x := range c.needs {
		x.Update()
	}
}

func (c *needsCollection) TopNeed() Need {
	values := make([]NeedLevel, 0, len(c.needs))
	for _, x := range c.needs {
		if x.Value() < 50 {
			values = append(values, x)
		}
	}

	if len(values) == 0 {
		return None
	}

	sort.Slice(values, func(i, j int) bool {
		a, b := values[i], values[j]
		if a.Value() == b.Value() {
			return a.Need() > b.Need()
		} else {
			return a.Value() > b.Value()
		}
	})

	return values[0].Need()
}

func (c *needsCollection) Provide(need Need, value int, duration int) {
	increment := CreateLevelIncrement(value, duration)
	c.needs[need].AddIncrement(increment)
}
