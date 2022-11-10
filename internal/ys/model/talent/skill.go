package talent

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model/action"
	"time"
)

type ElementalSkill struct {
	name           string
	lv             int
	dmgs           map[string]float64
	cure           int
	curePercentage float64
	duration       time.Duration // 持续时间
	cd             time.Duration // 冷却时间
	interval       time.Duration // 间隔时间
}

func BaseElementalSkill(name string, maxLv int, cd, duration time.Duration) *ElementalSkill {
	return &ElementalSkill{name: name, lv: maxLv, cd: cd, duration: duration}
}

func BaseElementalSkillWithInterval(name string, maxLv int, cd, duration, interval time.Duration) *ElementalSkill {
	return &ElementalSkill{name: name, lv: maxLv, cd: cd, duration: duration, interval: interval}
}

func LevelElementalSkill(lv int, dmgs map[string]float64) *ElementalSkill {
	return &ElementalSkill{lv: lv, dmgs: dmgs}
}

func LevelElementalSkillWithCure(lv int, dmgs map[string]float64, curePercentage float64, cure int) *ElementalSkill {
	return &ElementalSkill{lv: lv, dmgs: dmgs, curePercentage: curePercentage, cure: cure}
}

func (a *ElementalSkill) DMGs(elemental elementals.Elemental) *action.Actions {
	actions := action.NewActions()
	for name, dmg := range a.dmgs {
		actions.Add(action.New(attackMode.ElementalSkill, dmg, elemental, fmt.Sprintf("%s·%s", a.name, name)))
	}
	return actions
}
