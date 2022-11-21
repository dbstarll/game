package talent

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/action"
	"time"
)

type NormalAttack struct {
	name    string
	lv      int
	hits    []float64
	charged ChargedAttack
	plunge  PlungeAttack
}

type ChargedAttack struct {
	cyclic   float64
	final    float64
	hits     []float64
	stamina  int
	duration time.Duration
}

type PlungeAttack struct {
	dmg  float64
	low  float64
	high float64
}

func BaseNormalAttack(name string, maxLv, chargedStamina int) *NormalAttack {
	return &NormalAttack{name: name, lv: maxLv, charged: ChargedAttack{stamina: chargedStamina}}
}

func BaseNormalAttackByCyclic(name string, maxLv, chargedStamina int, chargedDuration time.Duration) *NormalAttack {
	return &NormalAttack{name: name, lv: maxLv, charged: ChargedAttack{stamina: chargedStamina, duration: chargedDuration}}
}

func LevelNormalAttack(lv int, hits []float64, plungeDmg float64, chargedHits ...float64) *NormalAttack {
	return &NormalAttack{
		lv:      lv,
		hits:    hits,
		charged: ChargedAttack{hits: chargedHits},
		plunge:  PlungeAttack{dmg: plungeDmg, low: plungeDmg * 2, high: plungeDmg * 2.5},
	}
}

func LevelNormalAttackByCyclic(lv int, hits []float64, plungeDmg, chargedCyclic, chargedFinal float64) *NormalAttack {
	return &NormalAttack{
		lv:      lv,
		hits:    hits,
		charged: ChargedAttack{cyclic: chargedCyclic, final: chargedFinal},
		plunge:  PlungeAttack{dmg: plungeDmg, low: plungeDmg * 2, high: plungeDmg * 2.5},
	}
}

func (a *NormalAttack) DMGs(weaponType weaponType.WeaponType, elemental elementals.Elemental) *action.Actions {
	actions, hitElemental := action.NewActions(), attackMode.NormalAttack.Elemental(weaponType, elemental)
	for idx, dmg := range a.hits {
		actions.Add(action.New(a.lv, attackMode.NormalAttack, dmg, hitElemental, fmt.Sprintf("%s·%d段", a.name, idx+1)))
	}
	actions.AddAll(a.charged.DMGs(a.lv, a.name, attackMode.ChargedAttack.Elemental(weaponType, elemental)))
	actions.AddAll(a.plunge.DMGs(a.lv, a.name, attackMode.PlungeAttack.Elemental(weaponType, elemental)))
	return actions
}

func (a *ChargedAttack) DMGs(lv int, name string, elemental elementals.Elemental) *action.Actions {
	actions := action.NewActions()
	if a.cyclic > 0 {
		actions.Add(action.New(lv, attackMode.ChargedAttack, a.cyclic, elemental, fmt.Sprintf("%s·重击持续", name)))
	}
	if a.final > 0 {
		actions.Add(action.New(lv, attackMode.ChargedAttack, a.final, elemental, fmt.Sprintf("%s·重击终结", name)))
	}
	hits := 0.0
	for _, hit := range a.hits {
		hits += hit
	}
	if hits > 0 {
		actions.Add(action.New(lv, attackMode.ChargedAttack, hits, elemental, fmt.Sprintf("%s·重击伤害", name)))
	}
	return actions
}

func (a *PlungeAttack) DMGs(lv int, name string, elemental elementals.Elemental) *action.Actions {
	actions := action.NewActions()
	if a.dmg > 0 {
		actions.Add(action.New(lv, attackMode.PlungeAttack, a.dmg, elemental, fmt.Sprintf("%s·下坠期间", name)))
	}
	if a.low > 0 {
		actions.Add(action.New(lv, attackMode.PlungeAttack, a.low, elemental, fmt.Sprintf("%s·低空坠地冲击", name)))
	}
	if a.high > 0 {
		actions.Add(action.New(lv, attackMode.PlungeAttack, a.high, elemental, fmt.Sprintf("%s·高空坠地冲击", name)))
	}
	return actions
}
