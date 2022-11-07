package talent

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/action"
)

type Talents struct {
	normalAttack   *NormalAttack
	elementalSkill *ElementalSkill
	elementalBurst *ElementalBurst
}

func (t *Talents) DMGs(weaponType weaponType.WeaponType, elemental elementals.Elemental) *action.Actions {
	actions := action.NewActions()
	if t != nil {
		actions.AddAll(t.normalAttack.DMGs(weaponType, elemental))
		actions.AddAll(t.elementalSkill.DMGs(attackMode.ElementalSkill.Elemental(weaponType, elemental)))
		actions.AddAll(t.elementalBurst.DMGs(attackMode.ElementalBurst.Elemental(weaponType, elemental)))
	}
	return actions
}
