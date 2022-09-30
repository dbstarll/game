package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
)

var (
	WeaponFactory无工之剑 = func(refine int) *Weapon {
		return NewWeapon(5, weaponType.Claymore, BaseWeapon(90, 608, AddAtkPercentage(49.6)))
	}
	WeaponFactory原木刀 = func(refine int) *Weapon {
		return NewWeapon(4, weaponType.Sword, BaseWeapon(90, 565, AddEnergyRecharge(30.6)))
	}
)

type Weapon struct {
	star       int
	weaponType weaponType.WeaponType
	level      int
	base       Attributes
}

type WeaponModifier func(weapon *Weapon) func()

func BaseWeapon(level, baseAttack int, baseModifier AttributeModifier) WeaponModifier {
	return func(weapon *Weapon) func() {
		oldLevel := weapon.level
		weapon.level = level
		callback := MergeAttributes(AddAtk(baseAttack), baseModifier)(&weapon.base)
		return func() {
			callback()
			weapon.level = oldLevel
		}
	}
}

func NewWeapon(star int, weaponType weaponType.WeaponType, modifiers ...WeaponModifier) *Weapon {
	w := &Weapon{
		star:       star,
		weaponType: weaponType,
		level:      1,
	}
	for _, modifier := range modifiers {
		modifier(w)
	}
	return w
}
