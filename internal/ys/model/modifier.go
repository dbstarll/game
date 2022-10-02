package model

import "github.com/dbstarll/game/internal/ys/dimension/attribute/point"

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

func AddHp(add int) AttributeModifier {
	return NewAttribute(point.Hp, float64(add)).Accumulation()
}

func AddHpPercentage(add float64) AttributeModifier {
	return NewAttribute(point.HpPercentage, add).Accumulation()
}

func AddAtk(add int) AttributeModifier {
	return NewAttribute(point.Atk, float64(add)).Accumulation()
}

func AddAtkPercentage(add float64) AttributeModifier {
	return NewAttribute(point.AtkPercentage, add).Accumulation()
}

func AddDef(add int) AttributeModifier {
	return NewAttribute(point.Def, float64(add)).Accumulation()
}

func AddDefPercentage(add float64) AttributeModifier {
	return NewAttribute(point.DefPercentage, add).Accumulation()
}

func AddCriticalRate(add float64) AttributeModifier {
	return NewAttribute(point.CriticalRate, add).Accumulation()
}

func AddCriticalDamage(add float64) AttributeModifier {
	return NewAttribute(point.CriticalDamage, add).Accumulation()
}

func AddElementalMastery(add float64) AttributeModifier {
	return NewAttribute(point.ElementalMastery, add).Accumulation()
}

func AddEnergyRecharge(add float64) AttributeModifier {
	return NewAttribute(point.EnergyRecharge, add).Accumulation()
}

func AddShieldStrength(add float64) AttributeModifier {
	return NewAttribute(point.ShieldStrength, add).Accumulation()
}
