package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/pkg/errors"
	"time"
)

var (
	CharacterFactory迪卢克 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(5, elemental.Pyro, weaponType.Claymore,
			BaseCharacter(90, 12981, 335, 784, AddCriticalRate(19.2)),
			TalentsTemplateModifier(NewTalentsTemplate(
				&NormalAttack{lv: 11, charged: ChargedAttack{stamina: 40, duration: time.Second * 5}},
				&ElementalSkill{lv: 13, cd: time.Second * 10},
				&ElementalBurst{lv: 14, cd: time.Second * 12, infusionDuration: time.Second * 8, energyCost: 40}).
				addNormalAttacks(
					&NormalAttack{lv: 1, hits: []float64{89.7, 87.6, 98.8, 134}, charged: ChargedAttack{cyclic: 68.8, final: 125}, plunge: PlungeAttack{89.5, 179, 224}},
					&NormalAttack{lv: 2, hits: []float64{97.0, 94.8, 107, 145}, charged: ChargedAttack{cyclic: 74.4, final: 135}, plunge: PlungeAttack{96.8, 194, 242}},
					&NormalAttack{lv: 3, hits: []float64{104, 102, 115, 156}, charged: ChargedAttack{cyclic: 80, final: 145}, plunge: PlungeAttack{104, 208, 260}},
					&NormalAttack{lv: 4, hits: []float64{115, 112, 126, 171}, charged: ChargedAttack{cyclic: 88, final: 160}, plunge: PlungeAttack{114, 229, 286}},
					&NormalAttack{lv: 5, hits: []float64{122, 119, 134, 182}, charged: ChargedAttack{cyclic: 93.6, final: 170}, plunge: PlungeAttack{122, 243, 304}},
					&NormalAttack{lv: 6, hits: []float64{130, 127, 144, 195}, charged: ChargedAttack{cyclic: 100, final: 181}, plunge: PlungeAttack{130, 260, 325}},
					&NormalAttack{lv: 7, hits: []float64{142, 139, 156, 212}, charged: ChargedAttack{cyclic: 109, final: 197}, plunge: PlungeAttack{142, 283, 354}},
					&NormalAttack{lv: 8, hits: []float64{153, 150, 169, 229}, charged: ChargedAttack{cyclic: 118, final: 213}, plunge: PlungeAttack{153, 306, 382}},
					&NormalAttack{lv: 9, hits: []float64{165, 161, 182, 246}, charged: ChargedAttack{cyclic: 126, final: 229}, plunge: PlungeAttack{164, 329, 411}},
					&NormalAttack{lv: 10, hits: []float64{177, 173, 195, 265}, charged: ChargedAttack{cyclic: 136, final: 247}, plunge: PlungeAttack{177, 354, 442}},
					&NormalAttack{lv: 11, hits: []float64{192, 187, 211, 286}, charged: ChargedAttack{cyclic: 147, final: 266}, plunge: PlungeAttack{189, 379, 473}},
				).
				addElementalSkills(
					&ElementalSkill{lv: 1, DMGs: map[string]float64{"一段伤害": 94.4, "二段伤害": 97.6, "三段伤害": 129}},
					&ElementalSkill{lv: 2, DMGs: map[string]float64{"一段伤害": 101, "二段伤害": 105, "三段伤害": 138}},
					&ElementalSkill{lv: 3, DMGs: map[string]float64{"一段伤害": 109, "二段伤害": 112, "三段伤害": 148}},
					&ElementalSkill{lv: 4, DMGs: map[string]float64{"一段伤害": 118, "二段伤害": 122, "三段伤害": 161}},
					&ElementalSkill{lv: 5, DMGs: map[string]float64{"一段伤害": 125, "二段伤害": 129, "三段伤害": 171}},
					&ElementalSkill{lv: 6, DMGs: map[string]float64{"一段伤害": 132, "二段伤害": 137, "三段伤害": 180}},
					&ElementalSkill{lv: 7, DMGs: map[string]float64{"一段伤害": 142, "二段伤害": 146, "三段伤害": 193}},
					&ElementalSkill{lv: 8, DMGs: map[string]float64{"一段伤害": 151, "二段伤害": 156, "三段伤害": 206}},
					&ElementalSkill{lv: 9, DMGs: map[string]float64{"一段伤害": 160, "二段伤害": 166, "三段伤害": 219}},
					&ElementalSkill{lv: 10, DMGs: map[string]float64{"一段伤害": 170, "二段伤害": 176, "三段伤害": 232}},
					&ElementalSkill{lv: 11, DMGs: map[string]float64{"一段伤害": 179, "二段伤害": 185, "三段伤害": 245}},
					&ElementalSkill{lv: 12, DMGs: map[string]float64{"一段伤害": 189, "二段伤害": 195, "三段伤害": 258}},
					&ElementalSkill{lv: 13, DMGs: map[string]float64{"一段伤害": 201, "二段伤害": 207, "三段伤害": 274}},
				).
				addElementalBursts(
					&ElementalBurst{lv: 1, DMGs: map[string]float64{"斩击伤害": 204, "持续伤害": 60.0, "爆裂伤害": 204}},
					&ElementalBurst{lv: 2, DMGs: map[string]float64{"斩击伤害": 219, "持续伤害": 64.5, "爆裂伤害": 219}},
					&ElementalBurst{lv: 3, DMGs: map[string]float64{"斩击伤害": 235, "持续伤害": 69.0, "爆裂伤害": 235}},
					&ElementalBurst{lv: 4, DMGs: map[string]float64{"斩击伤害": 255, "持续伤害": 75.0, "爆裂伤害": 255}},
					&ElementalBurst{lv: 5, DMGs: map[string]float64{"斩击伤害": 270, "持续伤害": 79.5, "爆裂伤害": 270}},
					&ElementalBurst{lv: 6, DMGs: map[string]float64{"斩击伤害": 286, "持续伤害": 84.0, "爆裂伤害": 286}},
					&ElementalBurst{lv: 7, DMGs: map[string]float64{"斩击伤害": 306, "持续伤害": 90.0, "爆裂伤害": 306}},
					&ElementalBurst{lv: 8, DMGs: map[string]float64{"斩击伤害": 326, "持续伤害": 96.0, "爆裂伤害": 326}},
					&ElementalBurst{lv: 9, DMGs: map[string]float64{"斩击伤害": 347, "持续伤害": 102.0, "爆裂伤害": 347}},
					&ElementalBurst{lv: 10, DMGs: map[string]float64{"斩击伤害": 367, "持续伤害": 108.0, "爆裂伤害": 367}},
					&ElementalBurst{lv: 11, DMGs: map[string]float64{"斩击伤害": 388, "持续伤害": 114.0, "爆裂伤害": 388}},
					&ElementalBurst{lv: 12, DMGs: map[string]float64{"斩击伤害": 408, "持续伤害": 120.0, "爆裂伤害": 408}},
					&ElementalBurst{lv: 13, DMGs: map[string]float64{"斩击伤害": 434, "持续伤害": 128.0, "爆裂伤害": 434}},
					&ElementalBurst{lv: 14, DMGs: map[string]float64{"斩击伤害": 459, "持续伤害": 135.0, "爆裂伤害": 459}},
				).check()),
		).Talents(normal, skill, burst)
	}
)

//slashingDMG      float64       // 斩击伤害
//dot              float64       // 持续伤害
//explosionDMG     float64       // 爆裂伤害

type Character struct {
	star            int
	elemental       elemental.Elemental
	weaponType      weaponType.WeaponType
	level           int
	base            *Attributes
	talents         Talents
	talentsTemplate *TalentsTemplate
	weapon          *Weapon
	artifacts       map[position.Position]*Artifacts
	attached        *Attributes
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

func TalentsTemplateModifier(talentsTemplate *TalentsTemplate) CharacterModifier {
	return func(character *Character) func() {
		oldTemplate := character.talentsTemplate
		character.talentsTemplate = talentsTemplate
		return func() {
			character.talentsTemplate = oldTemplate
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

func (c *Character) Talents(normal, skill, burst int) *Character {
	c.talents.normalAttack = c.talentsTemplate.normalAttacks[normal]
	c.talents.elementalSkill = c.talentsTemplate.elementalSkills[skill]
	c.talents.elementalBurst = c.talentsTemplate.elementalBursts[burst]
	return c
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
