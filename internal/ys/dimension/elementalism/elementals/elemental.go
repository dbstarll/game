package elementals

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/states"
)

// 元素
type Elemental int

type reactionFactor struct {
	reaction reactions.Reaction
	factor   float64
	state    states.State
}
type stateFactor struct {
	reaction  reactions.Reaction
	factor    float64
	elemental Elemental
}

const (
	Physical Elemental = iota // 物理
	Fire                      // 火
	Water                     // 水
	Grass                     // 草
	Electric                  // 雷
	Wind                      // 风
	Ice                       // 冰
	Earth                     // 岩
)

var (
	Elementals = []Elemental{
		Physical,
		Fire,
		Water,
		Grass,
		Electric,
		Wind,
		Ice,
		Earth,
	}
	restraintMap = map[Elemental]map[Elemental]int{
		Fire:     {Ice: 2, Electric: 1, Wind: 1},
		Water:    {Fire: 2, Electric: 1, Wind: 1},
		Ice:      {Water: 2, Electric: 1, Wind: 1},
		Electric: {Wind: 1},
	}
	reactionMap = map[Elemental]map[Elemental]*reactionFactor{
		Fire: {
			Water:    &reactionFactor{reactions.Vaporize, 1.5, -1},
			Grass:    &reactionFactor{reactions.Burn, 0.25, states.Burn},
			Ice:      &reactionFactor{reactions.Melt, 2, -1},
			Electric: &reactionFactor{reactions.Overload, 2, -1},
			Wind:     &reactionFactor{reactions.Swirl, 0.6, -1},
			Earth:    &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
		},
		Water: {
			Fire:     &reactionFactor{reactions.Vaporize, 2, -1},
			Grass:    &reactionFactor{reactions.Bloom, 2, states.Bloom},
			Electric: &reactionFactor{reactions.ElectroCharged, 1.2, states.ElectroCharged},
			Wind:     &reactionFactor{reactions.Swirl, 0.6, -1},
			Ice:      &reactionFactor{reactions.Frozen, 0, states.Frozen},
			Earth:    &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
		},
		Grass: {
			Fire:     &reactionFactor{reactions.Burn, 0.25, states.Burn},
			Water:    &reactionFactor{reactions.Bloom, 2, states.Bloom},
			Electric: &reactionFactor{reactions.Catalyze, 0, states.Quicken},
		},
		Electric: {
			Fire:  &reactionFactor{reactions.Overload, 2, -1},
			Water: &reactionFactor{reactions.ElectroCharged, 1.2, states.ElectroCharged},
			Grass: &reactionFactor{reactions.Catalyze, 0, states.Quicken},
			Wind:  &reactionFactor{reactions.Swirl, 0.6, -1},
			Ice:   &reactionFactor{reactions.Superconduct, 0.5, states.Superconduct},
			Earth: &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
		},
		Wind: {
			Fire:     &reactionFactor{reactions.Swirl, 0.6, -1},
			Water:    &reactionFactor{reactions.Swirl, 0.6, -1},
			Electric: &reactionFactor{reactions.Swirl, 0.6, -1},
			Ice:      &reactionFactor{reactions.Swirl, 0.6, -1},
		},
		Ice: {
			Fire:     &reactionFactor{reactions.Melt, 1.5, -1},
			Water:    &reactionFactor{reactions.Frozen, 0, states.Frozen},
			Electric: &reactionFactor{reactions.Superconduct, 0.5, states.Superconduct},
			Wind:     &reactionFactor{reactions.Swirl, 0.6, -1},
			Earth:    &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
		},
		Earth: {
			Fire:     &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
			Water:    &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
			Electric: &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
			Ice:      &reactionFactor{reactions.Crystallize, 0, states.Crystallize},
		},
	}
	stateMap = map[Elemental]map[states.State]*stateFactor{
		Fire: {
			states.Bloom: &stateFactor{reactions.Burgeon, 3, Grass},
		},
		Grass: {
			states.Quicken: &stateFactor{reactions.Spread, 1.25, -1},
		},
		Electric: {
			states.Bloom:   &stateFactor{reactions.Hyperbloom, 3, Grass},
			states.Quicken: &stateFactor{reactions.Aggravate, 1.15, -1},
		},
	}
)

func (e Elemental) IsValid() bool {
	return e >= Physical && e <= Earth
}

// 判断是否可作为角色的神之眼属性
func (e Elemental) IsCharacter() bool {
	return e.IsValid() && e != Physical
}

// 是否可附魔
func (e Elemental) CanInfusion() bool {
	switch e {
	case Fire, Water, Electric, Ice:
		return true
	default:
		return false
	}
}

func (e Elemental) Name() string {
	if e.IsCharacter() {
		return fmt.Sprintf("%s元素", e)
	} else {
		return e.String()
	}
}

func (e Elemental) String() string {
	switch e {
	case Physical:
		return "物理"
	case Fire:
		return "火"
	case Water:
		return "水"
	case Grass:
		return "草"
	case Electric:
		return "雷"
	case Wind:
		return "风"
	case Ice:
		return "冰"
	case Earth:
		return "岩"
	default:
		return fmt.Sprintf("未知[%d]", e)
	}
}

func (e Elemental) Multiple() float64 {
	switch e {
	case Physical:
		return 1.875
	default:
		return 1.5
	}
}

// 元素克制与元素量消耗比率
func (e Elemental) restraint(elemental Elemental) int {
	if ratios, exist := restraintMap[e]; exist {
		if ratio, exist := ratios[elemental]; exist {
			return ratio
		}
	}
	return 0
}

// 多种附魔属性叠加后的最终附魔属性
func (e Elemental) Infusion(infusionElemental Elemental) Elemental {
	if !infusionElemental.CanInfusion() {
		return e
	} else if e.restraint(infusionElemental) > 0 {
		return e
	} else if infusionElemental.restraint(e) > 0 {
		return infusionElemental
	} else if e.CanInfusion() {
		return e
	} else {
		return infusionElemental
	}
}

func (e Elemental) Reaction(attached Elemental) *reactions.React {
	if rs, exist := reactionMap[e]; exist {
		if r, exist := rs[attached]; exist {
			return reactions.NewReactWithState(r.reaction, r.factor, r.state)
		}
	}
	return nil
}

func (e Elemental) StateReaction(attached states.State) *ReactWithElemental {
	if e.IsValid() && attached.IsValid() {
		if attached == states.Frozen {
			return &ReactWithElemental{Reaction: reactions.Shattered, Factor: 1.5, Elemental: Physical}
		}
		if rs, exist := stateMap[e]; exist {
			if r, exist := rs[attached]; exist {
				return &ReactWithElemental{Reaction: r.reaction, Factor: r.factor, Elemental: r.elemental}
			}
		}
	}
	return nil
}

type ReactWithElemental struct {
	Reaction  reactions.Reaction
	Factor    float64
	Elemental Elemental
}
