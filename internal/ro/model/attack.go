package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
)

type Attack struct {
	weapon   weapon.Weapon
	magic    bool
	remote   bool
	skill    bool
	critical bool
	nature   nature.Nature
}

func (a *Attack) WithSkill() *Attack {
	a.skill = true
	return a
}

func (a *Attack) WithCritical() *Attack {
	a.critical = true
	return a
}

func (a *Attack) WithNature(nature nature.Nature) *Attack {
	a.nature = nature
	return a
}
