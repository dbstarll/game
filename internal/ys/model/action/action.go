package action

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

func New(mode attackMode.AttackMode, dmg float64, elemental elemental.Elemental, name string) *Action {
	return &Action{
		mode:      mode,
		dmg:       dmg,
		elemental: elemental,
		name:      name,
	}
}

func (a *Action) Mode() attackMode.AttackMode {
	return a.mode
}

func (a *Action) DMG() float64 {
	return a.dmg
}

func (a *Action) Elemental() elemental.Elemental {
	return a.elemental
}

func (a *Action) String() string {
	return fmt.Sprintf("%s[%s][%s][技能倍率: %+v%%]", a.name, a.mode, a.elemental.Name(), a.dmg)
}
