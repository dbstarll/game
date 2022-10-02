package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/pkg/errors"
)

var (
	CharacterFactory迪卢克 = func(constellation int) *Character {
		return NewCharacter(5, elemental.Pyro, weaponType.Claymore,
			BaseCharacter(90, 12981, 335, 784, AddCriticalRate(19.2)))
	}
)

type Character struct {
	star       int
	elemental  elemental.Elemental
	weaponType weaponType.WeaponType
	level      int
	base       *Attributes
	weapon     *Weapon
	artifacts  map[position.Position]*Artifacts
	attached   *Attributes
}

type CharacterModifier func(character *Character) func()

func BaseCharacter(level, baseHp, baseAtk, baseDef int, baseModifier AttributeModifier) CharacterModifier {
	return func(character *Character) func() {
		oldLevel := character.level
		character.level = level
		callback := MergeAttributes(AddHp(baseHp), AddAtk(baseAtk), AddDef(baseDef), baseModifier)(character.base)
		return func() {
			callback()
			character.level = oldLevel
		}
	}
}

func NewCharacter(star int, elemental elemental.Elemental, weaponType weaponType.WeaponType, modifiers ...CharacterModifier) *Character {
	c := &Character{
		star:       star,
		elemental:  elemental,
		weaponType: weaponType,
		level:      1,
		base:       NewAttributes(AddCriticalRate(5), AddCriticalDamage(50), AddEnergyRecharge(100)),
		artifacts:  make(map[position.Position]*Artifacts),
		attached:   NewAttributes(),
	}
	for _, modifier := range modifiers {
		modifier(c)
	}
	return c
}

func (c *Character) Weapon(newWeapon *Weapon) (*Weapon, error) {
	if oldWeapon := c.weapon; newWeapon == nil {
		// 卸下武器
		c.weapon = nil
		return oldWeapon, nil
	} else if c.weaponType != newWeapon.weaponType {
		return nil, errors.Errorf("不能装备此类型的武器: %s, 需要: %s", newWeapon.weaponType, c.weaponType)
	} else {
		// 替换武器
		c.weapon = newWeapon
		return oldWeapon, nil
	}
}

func (c *Character) Artifacts(newArtifacts *Artifacts) *Artifacts {
	if newArtifacts == nil {
		return nil
	}
	position := newArtifacts.position
	oldArtifacts, _ := c.artifacts[position]
	c.artifacts[position] = newArtifacts
	return oldArtifacts
}

func (c *Character) Apply(modifiers ...AttributeModifier) func() {
	return MergeAttributes(modifiers...)(c.attached)
}

func (c *Character) basicAttributes() *Attributes {
	basic := NewAttributes()
	c.base.Accumulation()(basic)
	c.weapon.AccumulationBase()(basic)
	return basic
}

func (c *Character) finalAttributes() *Attributes {
	final := NewAttributes()
	c.basicAttributes().Accumulation()(final)
	final.Clear(point.Hp, point.Atk, point.Def)
	c.weapon.AccumulationRefine()(final)
	for _, artifacts := range c.artifacts {
		artifacts.Accumulation()(final)
	}
	c.attached.Accumulation()(final)
	return final
}

// 基础区
func (c *Character) calculatorAttack() float64 {
	return 0
}

// 暴击区
func (c *Character) calculatorCritical() float64 {
	return 0
}

// 防御区
func (c *Character) calculatorDefense() float64 {
	return 0
}

// 基础伤害
func (c *Character) calculatorDamageBasic() float64 {
	return 0
}

func (c *Character) String() string {
	return fmt.Sprintf("%s\n", c.base)
}
