package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"time"
)

// 生命值
func AddHp(add int) attr.AttributeModifier {
	return attr.New(point.Hp, float64(add)).Accumulation()
}

// 生命值%
func AddHpPercentage(add float64) attr.AttributeModifier {
	return attr.New(point.HpPercentage, add).Accumulation()
}

// 攻击力
func AddAtk(add int) attr.AttributeModifier {
	return attr.New(point.Atk, float64(add)).Accumulation()
}

// 攻击力%
func AddAtkPercentage(add float64) attr.AttributeModifier {
	return attr.New(point.AtkPercentage, add).Accumulation()
}

// 防御力
func AddDef(add int) attr.AttributeModifier {
	return attr.New(point.Def, float64(add)).Accumulation()
}

// 防御力%
func AddDefPercentage(add float64) attr.AttributeModifier {
	return attr.New(point.DefPercentage, add).Accumulation()
}

// 元素精通
func AddElementalMastery(add float64) attr.AttributeModifier {
	return attr.New(point.ElementalMastery, add).Accumulation()
}

// 暴击率
func AddCriticalRate(add float64) attr.AttributeModifier {
	return attr.New(point.CriticalRate, add).Accumulation()
}

// 暴击伤害
func AddCriticalDamage(add float64) attr.AttributeModifier {
	return attr.New(point.CriticalDamage, add).Accumulation()
}

// 治疗加成
func AddHealingBonus(add float64) attr.AttributeModifier {
	return attr.New(point.HealingBonus, add).Accumulation()
}

// 受治疗加成
func AddIncomingHealingBonus(add float64) attr.AttributeModifier {
	return attr.New(point.IncomingHealingBonus, add).Accumulation()
}

// 元素充能效率
func AddEnergyRecharge(add float64) attr.AttributeModifier {
	return attr.New(point.EnergyRecharge, add).Accumulation()
}

// 冷却缩减
func AddCDReduction(add float64) attr.AttributeModifier {
	return attr.New(point.CDReduction, add).Accumulation()
}

// 护盾强效
func AddShieldStrength(add float64) attr.AttributeModifier {
	return attr.New(point.ShieldStrength, add).Accumulation()
}

// 伤害加成
func AddDamageBonus(add float64) attr.AttributeModifier {
	return attr.New(point.DamageBonus, add).Accumulation()
}

// 受到的伤害加成
func AddIncomingDamageBonus(add float64) attr.AttributeModifier {
	return attr.New(point.IncomingDamageBonus, add).Accumulation()
}

// 无视防御
func AddIgnoreDefence(add float64) attr.AttributeModifier {
	return attr.New(point.IgnoreDefence, add).Accumulation()
}

// 防御减免
func AddDefenceReduction(add float64) attr.AttributeModifier {
	return attr.New(point.DefenceReduction, add).Accumulation()
}

// 元素/物理伤害加成
func AddElementalDamageBonus(add float64, es ...elementals.Elemental) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, e := range es {
		modifiers = append(modifiers, attr.AddElementalDamageBonus(e, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 全元素/物理抗性
func AddAllElementalResist(add float64) attr.AttributeModifier {
	return AddElementalResist(add, elementals.Elementals...)
}

// 元素/物理抗性
func AddElementalResist(add float64, es ...elementals.Elemental) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, e := range es {
		modifiers = append(modifiers, attr.AddElementalResist(e, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 元素影响下增伤
func AddElementalAttachedDamageBonus(add float64, es ...elementals.Elemental) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, e := range es {
		modifiers = append(modifiers, attr.AddElementalAttachedDamageBonus(e, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 元素反应系数提高/元素反应伤害提升
func AddReactionDamageBonus(add float64, rs ...reactions.Reaction) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, r := range rs {
		modifiers = append(modifiers, attr.AddReactionDamageBonus(r, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 攻击模式伤害加成
func AddAttackDamageBonus(add float64, rs ...attackMode.AttackMode) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, r := range rs {
		modifiers = append(modifiers, attr.AddAttackDamageBonus(r, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 攻击模式的技能倍率加成
func AddAttackFactorBonus(add float64, rs ...attackMode.AttackMode) attr.AttributeModifier {
	var modifiers []attr.AttributeModifier
	for _, r := range rs {
		modifiers = append(modifiers, attr.AddAttackFactorBonus(r, add))
	}
	return attr.MergeAttributes(modifiers...)
}

func Superposition(times int, duration, interval time.Duration, modifier attr.AttributeModifier) attr.AttributeModifier {
	modifiers := make([]attr.AttributeModifier, times)
	for i := 0; i < times; i++ {
		modifiers[i] = modifier
	}
	return attr.MergeAttributes(modifiers...)
}
