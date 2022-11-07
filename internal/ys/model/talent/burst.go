package talent

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model/action"
	"time"
)

type ElementalBurst struct {
	name             string
	lv               int
	dmgs             map[string]float64
	cd               time.Duration // 冷却时间
	infusionDuration time.Duration // 附魔持续时间
	energyCost       int           // 元素能量
}

func BaseElementalBurst(name string, maxLv, energyCost int, cd, duration time.Duration) *ElementalBurst {
	return &ElementalBurst{name: name, lv: maxLv, energyCost: energyCost, cd: cd, infusionDuration: duration}
}

func LevelElementalBurst(lv int, dmgs map[string]float64) *ElementalBurst {
	return &ElementalBurst{lv: lv, dmgs: dmgs}
}

func (a *ElementalBurst) DMGs(elemental elementals.Elemental) *action.Actions {
	actions := action.NewActions()
	for name, dmg := range a.dmgs {
		actions.Add(action.New(attackMode.ElementalBurst, dmg, elemental, fmt.Sprintf("%s·%s", a.name, name)))
	}
	return actions
}
