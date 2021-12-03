package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"math"
)

type Character struct {
	nature  nature.Nature
	race    race.Race
	shape   shape.Shape
	level   Level
	quality Quality
	profits Profits
}

type CharacterModifier func(character *Character)

func NewCharacter(nature nature.Nature, race race.Race, shape shape.Shape, modifiers ...CharacterModifier) *Character {
	c := &Character{
		nature: nature,
		race:   race,
		shape:  shape,
	}
	for _, modifier := range modifiers {
		modifier(c)
	}
	return c
}

func Merge(modifiers ...CharacterModifier) CharacterModifier {
	return func(character *Character) {
		for _, modifier := range modifiers {
			modifier(character)
		}
	}
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
		character.profits.AddGains(magic, gains)
	}
}

func AddDamage(incr *Damage) CharacterModifier {
	return func(character *Character) {
		character.profits.AddDamage(incr)
	}
}

func AddNatureAttack(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureAttack(incr)
	}
}

func AddRaceDamage(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddRaceDamage(incr)
	}
}

func AddShapeDamage(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddShapeDamage(incr)
	}
}

func AddNatureDamage(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureDamage(incr)
	}
}

func DetectDefenceByPanel(expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) {
		character.detectDefenceByPanel(false, expectPhysicalPanel)
		character.detectDefenceByPanel(true, expectMagicalPanel)
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
	return c.profits.Attack(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) int {
	return c.profits.Defence(magic)
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) int {
	return c.QualityAttack(magic, remote) + c.EquipmentAttack(magic)
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) int {
	return c.QualityDefence(magic) + c.EquipmentDefence(magic)
}

func (c *Character) PanelDefence(magic bool) float64 {
	return float64(c.Defence(magic)) * (1 + c.profits.DefencePer(magic)/100)
}

func (c *Character) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	rate = c.profits.SkillDamageRate(target, magic, skillNature)
	rate *= skillNature.Restraint(target.nature) //*属性克制
	return
}

func (c *Character) detectDefenceByPanel(magic bool, expect float64) (optimumDefence int, optimumPanel float64) {
	for min, max, current := 0, 100000, c.profits.Defence(magic); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		c.profits.setDefence(magic, current)
		actual := c.PanelDefence(magic)

		if math.Abs(actual-expect) < math.Abs(optimumPanel-expect) {
			optimumDefence, optimumPanel = current, actual
		}
		if actual > expect {
			if max == current && max-min == 1 {
				break
			} else {
				max = current
			}
		} else if min == current && max-min == 1 {
			break
		} else {
			min = current
		}
	}
	c.profits.setDefence(magic, optimumDefence)
	fmt.Printf("detectDefenceByPanel[magic=%t]: optimumDefence=%d, optimumPanel=%f\n", magic, optimumDefence, optimumPanel)
	return
}
