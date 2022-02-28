package motivator

import "sort"

type NeedsCollection interface {
	AddNeed(Need, int)
	UpdateNeeds()
	TopNeed() string
	Provide(Need, int, int)
}

type needsCollection struct {
	needs map[string]NeedLevel
}

type sorter struct {
	name     string
	priority int
	value    int
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
	for _, x := range c.needs {
		x.Update()
	}
}

func (c *needsCollection) TopNeed() string {
	values := make([]sorter, 0, len(c.needs))
	for _, x := range c.needs {
		if x.Value() < 50 {
			values = append(values, sorter{x.Name(), x.Priority(), x.Value()})
		}
	}

	if len(values) == 0 {
		return "none"
	}

	sort.Slice(values, func(i, j int) bool {
		a, b := values[i], values[j]
		if a.value == b.value {
			return a.priority > b.priority
		} else {
			return a.value > b.value
		}
	})

	return values[0].name
}

func (c *needsCollection) Provide(need Need, value int, duration int) {
	c.needs[need.Name()].Provide(value, duration)
}
