package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
)

type Weapon struct {
	level      int
	weaponType weaponType.WeaponType
	base       Attributes
}

type WeaponModifier func(weapon *Weapon) func()

func BaseWeapon(level, baseAttack int, baseModifier AttributeModifier) WeaponModifier {
	return func(weapon *Weapon) func() {
		oldLevel := weapon.level
		weapon.level = level
		callback := MergeAttributes(AddAttack(baseAttack), baseModifier)(&weapon.base)
		return func() {
			callback()
			weapon.level = oldLevel
		}
	}
}

func NewWeapon(weaponType weaponType.WeaponType, modifiers ...WeaponModifier) *Weapon {
	w := &Weapon{
		level:      1,
		weaponType: weaponType,
	}
	for _, modifier := range modifiers {
		modifier(w)
	}
	return w
}
