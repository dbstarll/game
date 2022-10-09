package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
)

type Action struct {
	mode      attackMode.AttackMode
	dmg       float64
	name      string
	elemental elemental.Elemental
}

func NewAction(mode attackMode.AttackMode, dmg float64, elemental elemental.Elemental, name string) *Action {
	return &Action{
		mode:      mode,
		dmg:       dmg,
		elemental: elemental,
		name:      name,
	}
}

func (a *Action) String() string {
	if a.elemental < 0 {
		return fmt.Sprintf("%s[%s][物理][技能倍率: %+v%%]", a.name, a.mode, a.dmg)
	} else {
		return fmt.Sprintf("%s[%s][%s元素][技能倍率: %+v%%]", a.name, a.mode, a.elemental, a.dmg)
	}
}
