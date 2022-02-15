package main

import "log"

type Mob struct {
	config MobConfig
	status string
}

func newMob(config MobConfig) *Mob {
	mob := &Mob{config, "idle"}
	return mob
}

func (mob *Mob) addBehaviour(behaviour string) {
	log.Printf("Add behaviour %s to %s", behaviour, mob.config.Name)
}
