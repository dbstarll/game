package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
)

type Attributes struct {
	Hp            int     // 生命值
	HpPercentage  float64 // 生命值%
	Atk           int     // 攻击力
	AtkPercentage float64 // 攻击力%
	Def           int     // 防御力
	DefPercentage float64 // 防御力%

	CriticalRate   float64 // 暴击率%
	CriticalDamage float64 // 暴击伤害%

	EnergyRecharge   float64 // 元素充能效率%
	ElementalMastery int     // 元素精通

	HealingBonus         float64 // 治疗加成%
	IncomingHealingBonus float64 // 受治疗加成%

	CDReduction    float64 // 冷却缩减%
	ShieldStrength float64 // 护盾强效%

	ElementalDamageBonus map[elemental.Elemental]float64 // 元素伤害加成%
	ElementalResist      map[elemental.Elemental]float64 // 元素抗性%
	PhysicalDamageBonus  float64                         // 物理伤害加成%
	PhysicalResist       float64                         // 物理抗性%
}

type AttributeModifier func(attributes *Attributes) func()

func MergeAttributes(modifiers ...AttributeModifier) AttributeModifier {
	return func(attributes *Attributes) func() {
		size := len(modifiers)
		cancelList := make([]func(), size)
		for idx, modifier := range modifiers {
			cancelList[size-idx-1] = modifier(attributes)
		}
		return func() {
			for _, cancel := range cancelList {
				cancel()
			}
		}
	}
}

func BaseAttributes(baseHp, baseAtk, baseDef int) AttributeModifier {
	return func(attributes *Attributes) func() {
		oldHp, oldAtk, oldDef := attributes.Hp, attributes.Atk, attributes.Def
		attributes.Hp, attributes.Atk, attributes.Def = baseHp, baseAtk, baseDef
		return func() {
			attributes.Hp, attributes.Atk, attributes.Def = oldHp, oldAtk, oldDef
		}
	}
}

func AddAtk(add int) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.Atk += add
		return func() {
			attributes.Atk -= add
		}
	}
}

func AddAtkPercentage(add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.AtkPercentage += add
		return func() {
			attributes.AtkPercentage -= add
		}
	}
}

func AddCriticalRate(add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.CriticalRate += add
		return func() {
			attributes.CriticalRate -= add
		}
	}
}

func AddEnergyRecharge(add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.EnergyRecharge += add
		return func() {
			attributes.EnergyRecharge -= add
		}
	}
}

func (a *Attributes) Apply(modifiers ...AttributeModifier) func() {
	return MergeAttributes(modifiers...)(a)
}
