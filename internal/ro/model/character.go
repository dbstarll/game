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

func NewCharacter(race race.Race, nature nature.Nature, shape shape.Shape, modifiers ...CharacterModifier) *Character {
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

func AddRaceResist(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddRaceResist(incr)
	}
}

func AddShapeDamage(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddShapeDamage(incr)
	}
}

func AddShapeResist(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddShapeResist(incr)
	}
}

func AddNatureDamage(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureDamage(incr)
	}
}

func AddNatureResist(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureResist(incr)
	}
}

func DetectAttackByPanel(remote bool, expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) {
		character.detectAttackByPanel(false, remote, expectPhysicalPanel)
		character.detectAttackByPanel(true, remote, expectMagicalPanel)
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

//装备攻击
func (c *Character) EquipmentAttack(magic bool) int {
	return c.profits.Attack(magic)
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) int {
	return c.QualityAttack(magic, remote) + c.EquipmentAttack(magic)
}

//面板攻击 = 攻击 * (1 + 攻击%)
func (c *Character) PanelAttack(magic, remote bool) float64 {
	return float64(c.Attack(magic, remote)) * (1 + c.profits.AttackPer(magic)/100)
}

//素质防御
func (c *Character) QualityDefence(magic bool) int {
	return c.quality.Defence(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) int {
	return c.profits.Defence(magic)
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) int {
	return c.QualityDefence(magic) + c.EquipmentDefence(magic)
}

//面板防御 = 防御 * (1 + 防御%)
func (c *Character) PanelDefence(magic bool) float64 {
	return float64(c.Defence(magic)) * (1 + c.profits.DefencePer(magic)/100)
}

func (c *Character) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	rate = c.profits.SkillDamageRate(target, magic, skillNature)
	rate *= 1 - target.profits.raceResist[c.race]/100 //*(1-种族减伤%)
	rate *= skillNature.Restraint(target.nature)      //*属性克制
	return
}

func (c *Character) detectAttackByPanel(magic, remote bool, expect float64) (optimumAttack int, optimumPanel float64) {
	for min, max, current := 0, 100000, c.profits.Attack(magic); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		c.profits.setAttack(magic, current)
		actual := c.PanelAttack(magic, remote)

		if math.Abs(actual-expect) < math.Abs(optimumPanel-expect) {
			optimumAttack, optimumPanel = current, actual
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
	c.profits.setAttack(magic, optimumAttack)
	fmt.Printf("detectAttackByPanel[magic=%t,remote=%t]: optimumAttack=%d, optimumPanel=%f\n", magic, remote, optimumAttack, optimumPanel)
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
