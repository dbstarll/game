package point

type Point int

const (
	Hp                        Point = iota // 生命值
	HpPercentage                           // 生命值%
	Atk                                    // 攻击力
	AtkPercentage                          // 攻击力%
	Def                                    // 防御力
	DefPercentage                          // 防御力%
	ElementalMastery                       // 元素精通
	CriticalRate                           // 暴击率
	CriticalDamage                         // 暴击伤害
	HealingBonus                           // 治疗加成
	IncomingHealingBonus                   // 受治疗加成
	EnergyRecharge                         // 元素充能效率
	CDReduction                            // 冷却缩减
	ShieldStrength                         // 护盾强效
	FireDamageBonus                        // 火元素伤害加成
	FireResist                             // 火元素抗性
	WaterDamageBonus                       // 水元素伤害加成
	WaterResist                            // 水元素抗性
	GrassDamageBonus                       // 草元素伤害加成
	GrassResist                            // 草元素抗性
	ElectricDamageBonus                    // 雷元素伤害加成
	ElectricResist                         // 雷元素抗性
	WindDamageBonus                        // 风元素伤害加成
	WindResist                             // 风元素抗性
	IceDamageBonus                         // 冰元素伤害加成
	IceResist                              // 冰元素抗性
	EarthDamageBonus                       // 岩元素伤害加成
	EarthResist                            // 岩元素抗性
	PhysicalDamageBonus                    // 物理伤害加成
	PhysicalResist                         // 物理抗性
	DamageBonus                            // 伤害加成
	IncomingDamageBonus                    // 受到的伤害加成
	IgnoreDefence                          // 无视防御
	DefenceReduction                       // 防御减免
	NormalAttackDamageBonus                // 普通攻击伤害加成
	ChargedAttackDamageBonus               // 重击伤害加成
	PlungeAttackDamageBonus                // 下坠攻击伤害加成
	ElementalSkillDamageBonus              // 元素战技伤害加成
	ElementalBurstDamageBonus              // 元素爆发伤害加成
	NormalAttackFactorBonus                // 普通攻击技能倍率加成
	ChargedAttackFactorBonus               // 重击技能倍率加成
	PlungeAttackFactorBonus                // 下坠攻击技能倍率加成
	ElementalSkillFactorBonus              // 元素战技技能倍率加成
	ElementalBurstFactorBonus              // 元素爆发技能倍率加成
)

var Points = []Point{
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
	FireDamageBonus,
	FireResist,
	WaterDamageBonus,
	WaterResist,
	GrassDamageBonus,
	GrassResist,
	ElectricDamageBonus,
	ElectricResist,
	WindDamageBonus,
	WindResist,
	IceDamageBonus,
	IceResist,
	EarthDamageBonus,
	EarthResist,
	PhysicalDamageBonus,
	PhysicalResist,
	DamageBonus,
	IncomingDamageBonus,
	IgnoreDefence,
	DefenceReduction,
	NormalAttackDamageBonus,
	ChargedAttackDamageBonus,
	PlungeAttackDamageBonus,
	ElementalSkillDamageBonus,
	ElementalBurstDamageBonus,
	NormalAttackFactorBonus,
	ChargedAttackFactorBonus,
	PlungeAttackFactorBonus,
	ElementalSkillFactorBonus,
	ElementalBurstFactorBonus,
}

func (e Point) IsPercentage() bool {
	switch e {
	case Hp, Atk, Def, ElementalMastery:
		return false
	default:
		return true
	}
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
	case FireDamageBonus:
		return "火元素伤害加成"
	case FireResist:
		return "火元素抗性"
	case WaterDamageBonus:
		return "水元素伤害加成"
	case WaterResist:
		return "水元素抗性"
	case GrassDamageBonus:
		return "草元素伤害加成"
	case GrassResist:
		return "草元素抗性"
	case ElectricDamageBonus:
		return "雷元素伤害加成"
	case ElectricResist:
		return "雷元素抗性"
	case WindDamageBonus:
		return "风元素伤害加成"
	case WindResist:
		return "风元素抗性"
	case IceDamageBonus:
		return "冰元素伤害加成"
	case IceResist:
		return "冰元素抗性"
	case EarthDamageBonus:
		return "岩元素伤害加成"
	case EarthResist:
		return "岩元素抗性"
	case PhysicalDamageBonus:
		return "物理伤害加成"
	case PhysicalResist:
		return "物理抗性"
	case DamageBonus:
		return "伤害加成"
	case IncomingDamageBonus:
		return "受到的伤害加成"
	case IgnoreDefence:
		return "无视防御"
	case DefenceReduction:
		return "防御减免"
	case NormalAttackDamageBonus:
		return "普通攻击伤害加成"
	case ChargedAttackDamageBonus:
		return "重击伤害加成"
	case PlungeAttackDamageBonus:
		return "下坠攻击伤害加成"
	case ElementalSkillDamageBonus:
		return "元素战技伤害加成"
	case ElementalBurstDamageBonus:
		return "元素爆发伤害加成"
	case NormalAttackFactorBonus:
		return "普通攻击技能倍率加成"
	case ChargedAttackFactorBonus:
		return "重击技能倍率加成"
	case PlungeAttackFactorBonus:
		return "下坠攻击技能倍率加成"
	case ElementalSkillFactorBonus:
		return "元素战技技能倍率加成"
	case ElementalBurstFactorBonus:
		return "元素爆发技能倍率加成"
	default:
		if e < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
