package attr

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
)

var (
	NopCallBack          = func() {}
	NopAttributeModifier = func(_ *Attributes) func() {
		return NopCallBack
	}
)

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

// 单个元素伤害加成
func AddElementalDamageBonus(e elemental.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalDamageBonus(e, add)
	}
}

// 单个元素抗性
func AddElementalResist(e elemental.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalResist(e, add)
	}
}

// 单个元素影响下增伤
func AddElementalAttachedDamageBonus(e elemental.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalAttachedDamageBonus(e, add)
	}
}

// 单个元素反应系数提高/元素反应伤害提升
func AddReactionDamageBonus(r reaction.Reaction, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addReactionDamageBonus(r, add)
	}
}

// 单个攻击模式伤害加成
func AddAttackDamageBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackDamageBonus(r, add)
	}
}

// 攻击模式技能倍率加成
func AddAttackFactorBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackFactorBonus(r, add)
	}
}
