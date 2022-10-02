package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
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

func AddElementalDamageBonus(e elemental.Elemental, add float64) AttributeModifier {
	switch e {
	case elemental.Pyro:
		return NewAttribute(point.PyroDamageBonus, add).Accumulation()
	case elemental.Hydro:
		return NewAttribute(point.HydroDamageBonus, add).Accumulation()
	case elemental.Dendro:
		return NewAttribute(point.DendroDamageBonus, add).Accumulation()
	case elemental.Electro:
		return NewAttribute(point.ElectroDamageBonus, add).Accumulation()
	case elemental.Anemo:
		return NewAttribute(point.AnemoDamageBonus, add).Accumulation()
	case elemental.Cryo:
		return NewAttribute(point.CryoDamageBonus, add).Accumulation()
	case elemental.Geo:
		return NewAttribute(point.GeoDamageBonus, add).Accumulation()
	default:
		return NopAttributeModifier
	}
}

func AddDamageBonus(add float64) AttributeModifier {
	return NewAttribute(point.DamageBonus, add).Accumulation()
}

func AddIncomingDamageBonus(add float64) AttributeModifier {
	return NewAttribute(point.IncomingDamageBonus, add).Accumulation()
}
