package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

type Character struct {
	nature           nature.Nature
	race             race.Race
	shape            shape.Shape
	level            Level
	quality          Quality
	equipmentProfits Profits
}

type CharacterModifier func(character *Character)

func NewCharacter(nature nature.Nature, race race.Race, shape shape.Shape, modifiers ...CharacterModifier) *Character {
	c := &Character{
		nature: nature,
		race:   race,
		shape:  shape,
	}
	for _, f := range modifiers {
		f(c)
	}
	return c
}

func AddQuality(quality *Quality) CharacterModifier {
	return func(character *Character) {
		character.quality.Add(quality)
	}
}

func AddLevel(level *Level) CharacterModifier {
	return func(character *Character) {
		character.level.Add(level)
	}
}

func AddGains(magic bool, gains *Gains) CharacterModifier {
	return func(character *Character) {
		character.equipmentProfits.Add(magic, gains)
	}
}

//素质攻击
func (c *Character) QualityAttack(magic, remote bool) int {
	return c.quality.Attack(magic, remote)
}

//素质防御
func (c *Character) QualityDefence(magic bool) int {
	return c.quality.Defence(magic)
}

//装备攻击
func (c *Character) EquipmentAttack(magic bool) int {
	return c.equipmentProfits.Attack(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) int {
	return c.equipmentProfits.Defence(magic)
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) int {
	return c.QualityAttack(magic, remote) + c.EquipmentAttack(magic)
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) int {
	return c.QualityDefence(magic) + c.EquipmentDefence(magic)
}
