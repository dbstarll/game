package attr

import "github.com/dbstarll/game/internal/ys/dimension/elemental"

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
