package entry

import "github.com/dbstarll/game/internal/ys/dimension/artifacts/position"

// 圣遗物词条
type Entry int

const (
	Hp                  Entry = iota // 生命值
	HpPercentage                     // 生命值%
	Atk                              // 攻击力
	AtkPercentage                    // 攻击力%
	Def                              // 防御力
	DefPercentage                    // 防御力%
	ElementalMastery                 // 元素精通
	CriticalRate                     // 暴击率
	CriticalDamage                   // 暴击伤害
	HealingBonus                     // 治疗加成
	EnergyRecharge                   // 元素充能效率
	PhysicalDamageBonus              // 物理伤害加成
	FireDamageBonus                  // 火元素伤害加成
	WaterDamageBonus                 // 水元素伤害加成
	GrassDamageBonus                 // 草元素伤害加成
	ElectricDamageBonus              // 雷元素伤害加成
	WindDamageBonus                  // 风元素伤害加成
	IceDamageBonus                   // 冰元素伤害加成
	EarthDamageBonus                 // 岩元素伤害加成
)

var (
	Entries = []Entry{
		Hp,
		HpPercentage,
		Atk,
		AtkPercentage,
		Def,
		DefPercentage,
		ElementalMastery,
		CriticalRate,
		CriticalDamage,
		HealingBonus,
		EnergyRecharge,
		PhysicalDamageBonus,
		FireDamageBonus,
		WaterDamageBonus,
		GrassDamageBonus,
		ElectricDamageBonus,
		WindDamageBonus,
		IceDamageBonus,
		EarthDamageBonus,
	}

	primaryEntries = map[position.Position]map[Entry]bool{
		position.FlowerOfLife: {Hp: true},
		position.PlumeOfDeath: {Atk: true},
		position.SandsOfEon: {
			HpPercentage: true, AtkPercentage: true, DefPercentage: true, ElementalMastery: true,
			EnergyRecharge: true,
		},
		position.GobletOfEonothem: {
			HpPercentage: true, AtkPercentage: true, DefPercentage: true, ElementalMastery: true,
			PhysicalDamageBonus: true, FireDamageBonus: true, WaterDamageBonus: true, GrassDamageBonus: true,
			ElectricDamageBonus: true, WindDamageBonus: true, IceDamageBonus: true, EarthDamageBonus: true,
		},
		position.CircletOfLogos: {
			HpPercentage: true, AtkPercentage: true, DefPercentage: true, ElementalMastery: true,
			CriticalRate: true, CriticalDamage: true, HealingBonus: true,
		},
	}

	secondaryEntries = map[Entry]float64{
		Hp: 2, HpPercentage: 1,
		Atk: 2, AtkPercentage: 1,
		Def: 1, DefPercentage: 1,
		CriticalRate: 1, CriticalDamage: 1,
		ElementalMastery: 1, EnergyRecharge: 1,
	}
)

func (e Entry) String() string {
	switch e {
	case Hp:
		return "生命值"
	case HpPercentage:
		return "生命值%"
	case Atk:
		return "攻击力"
	case AtkPercentage:
		return "攻击力%"
	case Def:
		return "防御力"
	case DefPercentage:
		return "防御力%"
	case ElementalMastery:
		return "元素精通"
	case CriticalRate:
		return "暴击率"
	case CriticalDamage:
		return "暴击伤害"
	case HealingBonus:
		return "治疗加成"
	case EnergyRecharge:
		return "元素充能效率"
	case PhysicalDamageBonus:
		return "物理伤害加成"
	case FireDamageBonus:
		return "火元素伤害加成"
	case WaterDamageBonus:
		return "水元素伤害加成"
	case GrassDamageBonus:
		return "草元素伤害加成"
	case ElectricDamageBonus:
		return "雷元素伤害加成"
	case WindDamageBonus:
		return "风元素伤害加成"
	case IceDamageBonus:
		return "冰元素伤害加成"
	case EarthDamageBonus:
		return "岩元素伤害加成"
	default:
		return "未知"
	}
}

func (e Entry) Primary(pos position.Position) bool {
	if entries, exist := primaryEntries[pos]; exist {
		if ok, exist := entries[e]; exist {
			return ok
		}
	}
	return false
}

func (e Entry) Secondary() (float64, bool) {
	if rate, exist := secondaryEntries[e]; exist {
		return rate, true
	} else {
		return 0, false
	}
}
