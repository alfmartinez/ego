package mob

import (
	"ego/behaviour"
	"ego/configuration"
	"log"
)

type Mob struct {
	config configuration.Mob
	status string
}

func New(config configuration.Mob) *Mob {
	mob := &Mob{config, "idle"}
	return mob
}

func (mob *Mob) addBehaviour(behaviour behaviour.Behaviour) {
	log.Printf("Add behaviour %+v to %s", behaviour, mob.config.Name)
}
