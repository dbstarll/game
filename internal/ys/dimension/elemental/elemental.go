package elemental

import "github.com/dbstarll/game/internal/ys/dimension/attribute/point"

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
	restraint = map[Elemental]map[Elemental]int{
		Fire:     {Ice: 2, Electric: 1, Wind: 1},
		Water:    {Fire: 2, Electric: 1, Wind: 1},
		Ice:      {Water: 2, Electric: 1, Wind: 1},
		Electric: {Wind: 1},
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
	if ratios, exist := restraint[e]; exist {
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
		return point.PyroDamageBonus
	case Water:
		return point.HydroDamageBonus
	case Grass:
		return point.DendroDamageBonus
	case Electric:
		return point.ElectroDamageBonus
	case Wind:
		return point.AnemoDamageBonus
	case Ice:
		return point.CryoDamageBonus
	case Earth:
		return point.GeoDamageBonus
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
		return point.PyroResist
	case Water:
		return point.HydroResist
	case Grass:
		return point.DendroResist
	case Electric:
		return point.ElectroResist
	case Wind:
		return point.AnemoResist
	case Ice:
		return point.CryoResist
	case Earth:
		return point.GeoResist
	default:
		if e < 0 {
			return point.PhysicalResist
		} else {
			return -1
		}
	}
}
