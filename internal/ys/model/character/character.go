package character

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/artifacts"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"github.com/dbstarll/game/internal/ys/model/weapon"
	"github.com/pkg/errors"
)

type Character struct {
	star            int
	elemental       elementals.Elemental
	weaponType      weaponType.WeaponType
	level           int
	base            *attr.Attributes
	talents         *talent.Talents
	talentsTemplate *talent.TalentsTemplate
	weapon          *weapon.Weapon
	artifacts       map[position.Position]*artifacts.Artifacts
	attached        *attr.Attributes
}

type Modifier func(character *Character) func()

func Base(level, baseHp, baseAtk, baseDef int, baseModifier attr.AttributeModifier) Modifier {
	return func(character *Character) func() {
		oldLevel := character.level
		character.level = level
		callback := attr.MergeAttributes(buff.AddHp(baseHp), buff.AddAtk(baseAtk), buff.AddDef(baseDef), baseModifier)(character.base)
		return func() {
			callback()
			character.level = oldLevel
		}
	}
}

func TalentsTemplateModifier(talentsTemplate *talent.TalentsTemplate) Modifier {
	return func(character *Character) func() {
		oldTemplate := character.talentsTemplate
		character.talentsTemplate = talentsTemplate
		return func() {
			character.talentsTemplate = oldTemplate
		}
	}
}

func New(star int, elemental elementals.Elemental, weaponType weaponType.WeaponType, modifiers ...Modifier) *Character {
	c := &Character{
		star:       star,
		elemental:  elemental,
		weaponType: weaponType,
		level:      1,
		base:       attr.NewAttributes(buff.AddCriticalRate(5), buff.AddCriticalDamage(50), buff.AddEnergyRecharge(100)),
		artifacts:  make(map[position.Position]*artifacts.Artifacts),
		attached:   attr.NewAttributes(),
	}
	for _, modifier := range modifiers {
		modifier(c)
	}
	return c
}

func (c *Character) Level() int {
	return c.level
}

func (c *Character) BaseAttr(point point.Point) float64 {
	return c.base.Get(point)
}

func (c *Character) WeaponAttr(point point.Point) float64 {
	return c.weapon.Get(point)
}

func (c *Character) Weapon(newWeapon *weapon.Weapon) (*weapon.Weapon, error) {
	if oldWeapon := c.weapon; newWeapon == nil {
		// 卸下武器
		c.weapon = nil
		return oldWeapon, nil
	} else if c.weaponType != newWeapon.Type() {
		return nil, errors.Errorf("不能装备此类型的武器: %s, 需要: %s", newWeapon.Type(), c.weaponType)
	} else {
		// 替换武器
		c.weapon = newWeapon
		return oldWeapon, nil
	}
}

func (c *Character) GetWeapon() *weapon.Weapon {
	return c.weapon
}

func (c *Character) Artifacts(newArtifacts *artifacts.Artifacts) *artifacts.Artifacts {
	if newArtifacts == nil {
		return nil
	}
	position := newArtifacts.Position()
	oldArtifacts, _ := c.artifacts[position]
	c.artifacts[position] = newArtifacts
	return oldArtifacts
}

func (c *Character) Talents(normal, skill, burst int) *Character {
	c.talents = c.talentsTemplate.Talents(normal, skill, burst)
	return c
}

func (c *Character) Apply(modifiers ...attr.AttributeModifier) func() {
	return attr.MergeAttributes(modifiers...)(c.attached)
}

func (c *Character) GetActions() *action.Actions {
	return c.talents.DMGs(c.weaponType, c.elemental)
}

func (c *Character) basicAttributes() *attr.Attributes {
	basic := attr.NewAttributes()
	c.base.Accumulation(false)(basic)
	c.weapon.AccumulationBase()(basic)
	return basic
}

func (c *Character) FinalAttributes() *attr.Attributes {
	final := attr.NewAttributes()
	c.basicAttributes().Accumulation(false)(final)
	final.Clear(point.Hp, point.Atk, point.Def)
	c.weapon.AccumulationRefine()(final)
	for _, artifacts := range c.artifacts {
		artifacts.Accumulation(false)(final)
	}
	c.attached.Accumulation(false)(final)
	return final
}

func (c *Character) String() string {
	return fmt.Sprintf("%s\n", c.base)
}

func (c *Character) Evaluate(replaceArtifacts ...*artifacts.Artifacts) map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	for _, artifact := range c.artifacts {
		for n, m := range artifact.Evaluate(replaceArtifacts...) {
			detects[n] = m
		}
	}
	for n, m := range c.weapon.Evaluate() {
		detects[n] = m
	}
	return detects
}
