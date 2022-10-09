package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
)

type Action struct {
	mode attackMode.AttackMode
	dmg  float64
	name string
}

func NewAction(mode attackMode.AttackMode, dmg float64, name string) *Action {
	return &Action{
		mode: mode,
		dmg:  dmg,
		name: name,
	}
}

func (a *Action) String() string {
	return fmt.Sprintf("%s[%s][技能倍率: %+v%%]", a.name, a.mode, a.dmg)
}
