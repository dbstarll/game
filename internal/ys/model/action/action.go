package action

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
)

var (
	NopCallBack = func() {}
)

type Action struct {
	mode      attackMode.AttackMode
	dmg       float64
	name      string
	elemental elementals.Elemental
}

type Modifier func(action *Action) func()

func Infusion(elemental elementals.Elemental) Modifier {
	return func(action *Action) func() {
		switch action.Mode() {
		case attackMode.NormalAttack, attackMode.ChargedAttack, attackMode.PlungeAttack:
			if infusion := action.elemental.Infusion(elemental); infusion != action.elemental {
				old := action.elemental
				action.elemental = infusion
				return func() {
					action.elemental = old
				}
			} else {
				return NopCallBack
			}
		default:
			return NopCallBack
		}
	}
}

func New(mode attackMode.AttackMode, dmg float64, elemental elementals.Elemental, name string) *Action {
	return &Action{
		mode:      mode,
		dmg:       dmg,
		elemental: elemental,
		name:      name,
	}
}

func (a *Action) Apply(modifier Modifier) func() {
	return modifier(a)
}

func (a *Action) Mode() attackMode.AttackMode {
	return a.mode
}

func (a *Action) Name() string {
	return a.name
}

func (a *Action) DMG() float64 {
	return a.dmg
}

func (a *Action) Elemental() elementals.Elemental {
	return a.elemental
}

func (a *Action) String() string {
	return fmt.Sprintf("%s[%s][%s][技能倍率: %+v%%]", a.name, a.mode, a.elemental.Name(), a.dmg)
}
