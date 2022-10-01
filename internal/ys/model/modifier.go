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

func AddAtk(add int) AttributeModifier {
	return NewAttribute(point.Atk, float64(add)).Accumulation()
}

func AddAtkPercentage(add float64) AttributeModifier {
	return NewAttribute(point.AtkPercentage, add).Accumulation()
}

func AddCriticalRate(add float64) AttributeModifier {
	return NewAttribute(point.CriticalRate, add).Accumulation()
}

func AddEnergyRecharge(add float64) AttributeModifier {
	return NewAttribute(point.EnergyRecharge, add).Accumulation()
}

func AddShieldStrength(add float64) AttributeModifier {
	return NewAttribute(point.ShieldStrength, add).Accumulation()
}
