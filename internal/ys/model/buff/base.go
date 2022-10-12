package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

// 生命值
func AddHp(add int) attr.AttributeModifier {
	return attr.NewAttribute(point.Hp, float64(add)).Accumulation()
}

// 生命值%
func AddHpPercentage(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.HpPercentage, add).Accumulation()
}

// 攻击力
func AddAtk(add int) attr.AttributeModifier {
	return attr.NewAttribute(point.Atk, float64(add)).Accumulation()
}

// 攻击力%
func AddAtkPercentage(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.AtkPercentage, add).Accumulation()
}

// 防御力
func AddDef(add int) attr.AttributeModifier {
	return attr.NewAttribute(point.Def, float64(add)).Accumulation()
}

// 防御力%
func AddDefPercentage(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.DefPercentage, add).Accumulation()
}

// 元素精通
func AddElementalMastery(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.ElementalMastery, add).Accumulation()
}

// 暴击率
func AddCriticalRate(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.CriticalRate, add).Accumulation()
}

// 暴击伤害
func AddCriticalDamage(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.CriticalDamage, add).Accumulation()
}

// 治疗加成
func AddHealingBonus(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.HealingBonus, add).Accumulation()
}

// 受治疗加成
func AddIncomingHealingBonus(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.IncomingHealingBonus, add).Accumulation()
}

// 元素充能效率
func AddEnergyRecharge(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.EnergyRecharge, add).Accumulation()
}

// 冷却缩减
func AddCDReduction(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.CDReduction, add).Accumulation()
}

// 护盾强效
func AddShieldStrength(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.ShieldStrength, add).Accumulation()
}

// 单个元素伤害加成 or 物理伤害加成
func AddElementalDamageBonus(e elemental.Elemental, add float64) attr.AttributeModifier {
	return attr.NewAttribute(e.DamageBonusPoint(), add).Accumulation()
}

// 所有元素伤害加成 and 物理伤害加成
func AddAllElementalResist(add float64) attr.AttributeModifier {
	modifiers := []attr.AttributeModifier{attr.NewAttribute(point.PhysicalResist, add).Accumulation()}
	for _, e := range elemental.Elementals {
		modifiers = append(modifiers, AddElementalResist(e, add))
	}
	return attr.MergeAttributes(modifiers...)
}

// 单个元素抗性 or 物理抗性
func AddElementalResist(e elemental.Elemental, add float64) attr.AttributeModifier {
	return attr.NewAttribute(e.ResistPoint(), add).Accumulation()
}

// 伤害加成
func AddDamageBonus(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.DamageBonus, add).Accumulation()
}

// 受到的伤害加成
func AddIncomingDamageBonus(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.IncomingDamageBonus, add).Accumulation()
}

// 无视防御
func AddIgnoreDefence(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.IgnoreDefence, add).Accumulation()
}

// 防御减免
func AddDefenceReduction(add float64) attr.AttributeModifier {
	return attr.NewAttribute(point.DefenceReduction, add).Accumulation()
}

// 单个攻击模式的伤害加成
func AddAttackDamageBonus(m attackMode.AttackMode, add float64) attr.AttributeModifier {
	return attr.NewAttribute(m.DamageBonusPoint(), add).Accumulation()
}
