package elementals

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
)

// 元素
type Elemental int

type reactionFactor struct {
	reaction reactions.Reaction
	factor   float64
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
			Water:    &reactionFactor{reactions.Vaporize, 1.5},
			Grass:    &reactionFactor{reactions.Burn, 0.25},
			Ice:      &reactionFactor{reactions.Melt, 2},
			Electric: &reactionFactor{reactions.Overload, 2},
			Wind:     &reactionFactor{reactions.Swirl, 0.6},
		},
		Water: {
			Fire:     &reactionFactor{reactions.Vaporize, 2},
			Grass:    &reactionFactor{reactions.Bloom, 2},
			Electric: &reactionFactor{reactions.ElectroCharged, 1.2},
			Wind:     &reactionFactor{reactions.Swirl, 0.6},
			Ice:      &reactionFactor{reactions.Frozen, 0},
		},
		Grass: {
			Fire:  &reactionFactor{reactions.Burn, 0.25},
			Water: &reactionFactor{reactions.Bloom, 2},
		},
		Electric: {
			Fire:  &reactionFactor{reactions.Overload, 2},
			Water: &reactionFactor{reactions.ElectroCharged, 1.2},
			Wind:  &reactionFactor{reactions.Swirl, 0.6},
			Ice:   &reactionFactor{reactions.Superconduct, 0.5},
			//Grass: &reactionFactor{reaction.Hyperbloom, 3},
		},
		Wind: {
			Fire:     &reactionFactor{reactions.Swirl, 0.6},
			Water:    &reactionFactor{reactions.Swirl, 0.6},
			Electric: &reactionFactor{reactions.Swirl, 0.6},
			Ice:      &reactionFactor{reactions.Swirl, 0.6},
		},
		Ice: {
			Fire:     &reactionFactor{reactions.Melt, 1.5},
			Water:    &reactionFactor{reactions.Frozen, 0},
			Electric: &reactionFactor{reactions.Superconduct, 0.5},
			Wind:     &reactionFactor{reactions.Swirl, 0.6},
		},
		Earth: {
			Fire:     &reactionFactor{reactions.Crystallize, 1},
			Water:    &reactionFactor{reactions.Crystallize, 1},
			Electric: &reactionFactor{reactions.Crystallize, 1},
			Ice:      &reactionFactor{reactions.Crystallize, 1},
		},
	}
)

// 判断是否可作为角色的神之眼属性
func (e Elemental) IsCharacter() bool {
	return e > Physical && e <= Earth
}

func (e Elemental) Name() string {
	if e > Physical && e <= Earth {
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
		return "未知"
	}
}

// 元素克制与元素量消耗比率
func (e Elemental) Restraint(elemental Elemental) int {
	if ratios, exist := restraintMap[e]; exist {
		if ratio, exist := ratios[elemental]; exist {
			return ratio
		}
	}
	return 0
}

// 多种附魔属性叠加后的最终附魔属性
func (e Elemental) Infusion(infusionElemental Elemental) Elemental {
	if e.Restraint(infusionElemental) > 0 {
		return e
	} else if infusionElemental.Restraint(e) > 0 {
		return infusionElemental
	} else if e <= Physical {
		return infusionElemental
	} else {
		return e
	}
}

func (e Elemental) Reaction(attached Elemental) *reactions.React {
	if rs, exist := reactionMap[e]; exist {
		if r, exist := rs[attached]; exist {
			return reactions.NewReact(r.reaction, r.factor)
		}
	}
	return nil
}
