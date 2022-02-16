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
	mob.Status = mob.behaviours[mob.Status].Evaluate()
}

func (mob *Mob) addBehaviour(behaviour behaviour.Behaviour) {
	mob.behaviours[behaviour.GetName()] = behaviour
}
