package elemental

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
)

// 元素类型
type Elemental int

const (
	Fire     Elemental = iota // 火(Pyro)
	Water                     // 水(Hydro)
	Grass                     // 草(Dendro)
	Electric                  // 雷(Electro)
	Wind                      // 风(Anemo)
	Ice                       // 冰(Cryo)
	Earth                     // 岩(Geo)
)

var (
	Elementals = []Elemental{
		Fire,
		Water,
		Grass,
		Electric,
		Wind,
		Ice,
		Earth,
	}
	restraints = map[Elemental]map[Elemental]int{
		Fire:     {Ice: 2, Electric: 1, Wind: 1},
		Water:    {Fire: 2, Electric: 1, Wind: 1},
		Ice:      {Water: 2, Electric: 1, Wind: 1},
		Electric: {Wind: 1},
	}
	reactions = map[Elemental]map[Elemental]*reaction.Factor{
		Fire: {
			Water:    reaction.NewFactor(reaction.Vaporize, 1.5),
			Grass:    reaction.NewFactor(reaction.Burn, 0.25),
			Ice:      reaction.NewFactor(reaction.Melt, 2),
			Electric: reaction.NewFactor(reaction.Overload, 2),
			Wind:     reaction.NewFactor(reaction.Swirl, 0.6),
		},
		Water: {
			Fire:     reaction.NewFactor(reaction.Vaporize, 2),
			Grass:    reaction.NewFactor(reaction.Bloom, 2),
			Electric: reaction.NewFactor(reaction.ElectroCharged, 1.2),
			Wind:     reaction.NewFactor(reaction.Swirl, 0.6),
			Ice:      reaction.NewFactor(reaction.Frozen, 0),
		},
		Grass: {
			Fire:  reaction.NewFactor(reaction.Burn, 0.25),
			Water: reaction.NewFactor(reaction.Bloom, 2),
		},
		Electric: {
			Fire:  reaction.NewFactor(reaction.Overload, 2),
			Water: reaction.NewFactor(reaction.ElectroCharged, 1.2),
			Wind:  reaction.NewFactor(reaction.Swirl, 0.6),
			Ice:   reaction.NewFactor(reaction.Superconduct, 0.5),
		},
		Wind: {
			Fire:     reaction.NewFactor(reaction.Swirl, 0.6),
			Water:    reaction.NewFactor(reaction.Swirl, 0.6),
			Electric: reaction.NewFactor(reaction.Swirl, 0.6),
			Ice:      reaction.NewFactor(reaction.Swirl, 0.6),
		},
		Ice: {
			Fire:     reaction.NewFactor(reaction.Melt, 1.5),
			Water:    reaction.NewFactor(reaction.Frozen, 0),
			Electric: reaction.NewFactor(reaction.Superconduct, 0.5),
			Wind:     reaction.NewFactor(reaction.Swirl, 0.6),
		},
		Earth: {
			Fire:     reaction.NewFactor(reaction.Crystallize, 1),
			Water:    reaction.NewFactor(reaction.Crystallize, 1),
			Electric: reaction.NewFactor(reaction.Crystallize, 1),
			Ice:      reaction.NewFactor(reaction.Crystallize, 1),
		},
	}
)

func (e Elemental) String() string {
	switch e {
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
		if e < -1 {
			return "不限"
		} else if e == -1 {
			return "物理"
		} else {
			return "未知"
		}
	}
}

func (e Elemental) Restraint(elemental Elemental) int {
	if ratios, exist := restraints[e]; exist {
		if ratio, exist := ratios[elemental]; exist {
			return ratio
		}
	}
	return 0
}

// 附魔优先级：
// 火雷附魔，火伤
// 火冰附魔，火伤
// 冰水附魔，冰伤
// 冰雷附魔，冰伤
// 水火附魔，水伤
// 水雷附魔，水伤
// 风元素附魔会被水火冰雷任何一种元素覆盖
func (e Elemental) Infusion(elemental Elemental) Elemental {
	if e.Restraint(elemental) > 0 {
		return e
	} else if elemental.Restraint(e) > 0 {
		return elemental
	} else if e < 0 {
		return elemental
	} else {
		return e
	}
}

func (e Elemental) DamageBonusPoint() point.Point {
	switch e {
	case Fire:
		return point.FireDamageBonus
	case Water:
		return point.WaterDamageBonus
	case Grass:
		return point.GrassDamageBonus
	case Electric:
		return point.ElectricDamageBonus
	case Wind:
		return point.WindDamageBonus
	case Ice:
		return point.IceDamageBonus
	case Earth:
		return point.EarthDamageBonus
	default:
		if e < 0 {
			return point.PhysicalDamageBonus
		} else {
			return -1
		}
	}
}

func (e Elemental) ResistPoint() point.Point {
	switch e {
	case Fire:
		return point.FireResist
	case Water:
		return point.WaterResist
	case Grass:
		return point.GrassResist
	case Electric:
		return point.ElectricResist
	case Wind:
		return point.WindResist
	case Ice:
		return point.IceResist
	case Earth:
		return point.EarthResist
	default:
		if e < 0 {
			return point.PhysicalResist
		} else {
			return -1
		}
	}
}

func (e Elemental) Reaction(attached Elemental) *reaction.Factor {
	if rs, exist := reactions[e]; exist {
		if r, exist := rs[attached]; exist {
			return r
		}
	}
	return nil
}
