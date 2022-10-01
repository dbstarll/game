package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"time"
)

var (
	WeaponFactory无工之剑 = func(refine int) *Weapon {
		return NewWeapon(5, weaponType.Claymore, BaseWeapon(90, 608, AddAtkPercentage(49.6)),
			AddShieldStrength(float64(15+refine*5)),
			Superposition(5, time.Second*8, time.Millisecond*300, AddAtkPercentage(float64(3+refine))),
		)
	}
	WeaponFactory原木刀 = func(refine int) *Weapon {
		return NewWeapon(4, weaponType.Sword, BaseWeapon(90, 565, AddEnergyRecharge(30.6)))
	}
)

type Weapon struct {
	star            int
	weaponType      weaponType.WeaponType
	level           int
	base            Attributes
	refineModifiers []AttributeModifier
}

type WeaponModifier func(weapon *Weapon) func()

func BaseWeapon(level, baseAtk int, baseModifier AttributeModifier) WeaponModifier {
	return func(weapon *Weapon) func() {
		oldLevel := weapon.level
		weapon.level = level
		callback := MergeAttributes(AddAtk(baseAtk), baseModifier)(&weapon.base)
		return func() {
			callback()
			weapon.level = oldLevel
		}
	}
}

func NewWeapon(star int, weaponType weaponType.WeaponType, baseModifier WeaponModifier, refineModifiers ...AttributeModifier) *Weapon {
	w := &Weapon{
		star:            star,
		weaponType:      weaponType,
		level:           1,
		refineModifiers: refineModifiers,
	}
	baseModifier(w)
	return w
}

func (w *Weapon) AccumulationBase() AttributeModifier {
	if w == nil {
		return NopAttributeModifier
	} else {
		return w.base.Accumulation()
	}
}

func (w *Weapon) AccumulationRefine() AttributeModifier {
	if w == nil {
		return NopAttributeModifier
	} else {
		return MergeAttributes(w.refineModifiers...)
	}
}
