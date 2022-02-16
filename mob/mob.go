package mob

import (
	"ego/behaviour"
	"ego/configuration"
)

type Mob struct {
	config     configuration.Mob
	Status     string
	behaviours map[string]behaviour.Behaviour
}

func New(config configuration.Mob) *Mob {
	mob := &Mob{config, "idle", make(map[string]behaviour.Behaviour)}
	return mob
}

func (mob *Mob) Init() {
	mob.Status = "idle"
}

func (mob *Mob) Tick() {
	behaviour := mob.behaviours[mob.Status]
	behaviour.Evaluate()
	if behaviour.IsOver() {
		behaviour.Reset()
		mob.Status = behaviour.Next()
	}

}

func (mob *Mob) addBehaviour(behaviour behaviour.Behaviour) {
	mob.behaviours[behaviour.GetName()] = behaviour
}
