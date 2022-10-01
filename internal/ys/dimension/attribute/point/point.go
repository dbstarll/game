package point

type Point int

const (
	Hp                   Point = iota // 生命值
	HpPercentage                      // 生命值%
	Atk                               // 攻击力
	AtkPercentage                     // 攻击力%
	Def                               // 防御力
	DefPercentage                     // 防御力%
	ElementalMastery                  // 元素精通
	CriticalRate                      // 暴击率
	CriticalDamage                    // 暴击伤害
	HealingBonus                      // 治疗加成
	IncomingHealingBonus              // 受治疗加成
	EnergyRecharge                    // 元素充能效率
	CDReduction                       // 冷却缩减
	ShieldStrength                    // 护盾强效
	PyroDamageBonus                   // 火元素伤害加成
	PyroResist                        // 火元素抗性
	HydroDamageBonus                  // 水元素伤害加成
	HydroResist                       // 水元素抗性
	DendroDamageBonus                 // 草元素伤害加成
	DendroResist                      // 草元素抗性
	ElectroDamageBonus                // 雷元素伤害加成
	ElectroResist                     // 雷元素抗性
	AnemoDamageBonus                  // 风元素伤害加成
	AnemoResist                       // 风元素抗性
	CryoDamageBonus                   // 冰元素伤害加成
	CryoResist                        // 冰元素抗性
	GeoDamageBonus                    // 岩元素伤害加成
	GeoResist                         // 岩元素抗性
	PhysicalDamageBonus               // 物理伤害加成
	PhysicalResist                    // 物理抗性
)

var EntryPoints = []Point{
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
	IncomingHealingBonus,
	EnergyRecharge,
	CDReduction,
	ShieldStrength,
	PyroDamageBonus,
	PyroResist,
	HydroDamageBonus,
	HydroResist,
	DendroDamageBonus,
	DendroResist,
	ElectroDamageBonus,
	ElectroResist,
	AnemoDamageBonus,
	AnemoResist,
	CryoDamageBonus,
	CryoResist,
	GeoDamageBonus,
	GeoResist,
	PhysicalDamageBonus,
	PhysicalResist,
}

func (e Point) String() string {
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
	case IncomingHealingBonus:
		return "受治疗加成"
	case EnergyRecharge:
		return "元素充能效率"
	case CDReduction:
		return "冷却缩减"
	case ShieldStrength:
		return "护盾强效"
	case PyroDamageBonus:
		return "火元素伤害加成"
	case PyroResist:
		return "火元素抗性"
	case HydroDamageBonus:
		return "水元素伤害加成"
	case HydroResist:
		return "水元素抗性"
	case DendroDamageBonus:
		return "草元素伤害加成"
	case DendroResist:
		return "草元素抗性"
	case ElectroDamageBonus:
		return "雷元素伤害加成"
	case ElectroResist:
		return "雷元素抗性"
	case AnemoDamageBonus:
		return "风元素伤害加成"
	case AnemoResist:
		return "风元素抗性"
	case CryoDamageBonus:
		return "冰元素伤害加成"
	case CryoResist:
		return "冰元素抗性"
	case GeoDamageBonus:
		return "岩元素伤害加成"
	case GeoResist:
		return "岩元素抗性"
	case PhysicalDamageBonus:
		return "物理伤害加成"
	case PhysicalResist:
		return "物理抗性"
	default:
		if e < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
