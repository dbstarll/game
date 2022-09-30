package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/element"
)

type Attributes struct {
	Hp         int     // 生命值
	HpPer      float64 // 生命值%
	Attack     int     // 攻击力
	AttackPer  float64 // 攻击力%
	Defence    int     // 防御力
	DefencePer float64 // 防御力%

	Critical       float64 // 暴击率%
	CriticalDamage float64 // 暴击伤害%

	ElementCharge float64 // 元素充能效率%
	ElementMaster int     // 元素精通

	Cure  float64 // 治疗加成%
	Cured float64 // 受治疗加成%

	Cooling float64 // 冷却缩减%
	Shield  float64 // 护盾强效%

	ElementDamage  map[element.Element]float64 // 元素伤害加成%
	ElementResist  map[element.Element]float64 // 元素抗性%
	PhysicalDamage float64                     // 物理伤害加成%
	PhysicalResist float64                     // 物理抗性%
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

func BaseAttributes(baseHp, baseAttack, baseDefence int) AttributeModifier {
	return func(attributes *Attributes) func() {
		oldHp, oldAttack, oldDefence := attributes.Hp, attributes.Attack, attributes.Defence
		attributes.Hp, attributes.Attack, attributes.Defence = baseHp, baseAttack, baseDefence
		return func() {
			attributes.Hp, attributes.Attack, attributes.Defence = oldHp, oldAttack, oldDefence
		}
	}
}

func AddAttack(attack int) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.Attack += attack
		return func() {
			attributes.Attack -= attack
		}
	}
}

func AddAttackPer(attackPer float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.AttackPer += attackPer
		return func() {
			attributes.AttackPer -= attackPer
		}
	}
}

func AddCritical(critical float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		attributes.Critical += critical
		return func() {
			attributes.Critical -= critical
		}
	}
}

func (a *Attributes) Apply(modifiers ...AttributeModifier) func() {
	return MergeAttributes(modifiers...)(a)
}
