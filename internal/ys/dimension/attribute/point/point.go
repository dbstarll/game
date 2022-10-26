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
	DamageBonus                       // 伤害加成
	IncomingDamageBonus               // 受到的伤害加成
	IgnoreDefence                     // 无视防御
	DefenceReduction                  // 防御减免
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
	DamageBonus,
	IncomingDamageBonus,
	IgnoreDefence,
	DefenceReduction,
}

func (p Point) IsPercentage() bool {
	switch p {
	case Hp, Atk, Def, ElementalMastery:
		return false
	default:
		return true
	}
}

func (p Point) Multiple() float64 {
	switch p {
	case Hp: // 生命值
		return 153.7
	case HpPercentage: // 生命值%
		return 1.5
	case Atk: // 攻击力
		return 10.005
	case AtkPercentage: // 攻击力%
		return 1.5
	case Def: // 防御力
		return 5.955
	case DefPercentage: // 防御力%
		return 1.875
	case ElementalMastery: // 元素精通
		return 5.996
	case CriticalRate: // 暴击率
		return 1
	case CriticalDamage: // 暴击伤害
		return 2
	case HealingBonus: // 治疗加成
		return 1.155
	//IncomingHealingBonus              // 受治疗加成
	case EnergyRecharge: // 元素充能效率
		return 5.0 / 3
	//CDReduction                       // 冷却缩减
	//ShieldStrength                    // 护盾强效
	//DamageBonus                       // 伤害加成
	//IncomingDamageBonus               // 受到的伤害加成
	//IgnoreDefence                     // 无视防御
	//DefenceReduction                  // 防御减免
	default:
		return 0
	}
}

func (p Point) String() string {
	switch p {
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
	case DamageBonus:
		return "伤害加成"
	case IncomingDamageBonus:
		return "受到的伤害加成"
	case IgnoreDefence:
		return "无视防御"
	case DefenceReduction:
		return "防御减免"
	default:
		if p < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
