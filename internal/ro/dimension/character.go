package dimension

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

type Character struct {
	_nature          nature.Nature
	_race            race.Race
	_shape           shape.Shape
	level            Level
	quality          Quality
	equipmentProfits Profits
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
