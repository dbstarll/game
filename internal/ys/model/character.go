package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/pkg/errors"
	"time"
)

var (
	CharacterFactory迪卢克 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(5, elemental.Fire, weaponType.Claymore,
			BaseCharacter(90, 12981, 335, 784, buff.AddCriticalRate(19.2)),
			TalentsTemplateModifier(NewTalentsTemplate(
				&NormalAttack{name: "淬炼之剑", lv: 11, charged: ChargedAttack{stamina: 40, duration: time.Second * 5}},
				&ElementalSkill{name: "逆焰之刃", lv: 13, cd: time.Second * 10},
				&ElementalBurst{name: "黎明", lv: 14, cd: time.Second * 12, infusionDuration: time.Second * 8, energyCost: 40}).
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
					&ElementalSkill{lv: 1, dmgs: map[string]float64{"1段": 94.4, "2段": 97.6, "3段": 129}},
					&ElementalSkill{lv: 2, dmgs: map[string]float64{"1段": 101, "2段": 105, "3段": 138}},
					&ElementalSkill{lv: 3, dmgs: map[string]float64{"1段": 109, "2段": 112, "3段": 148}},
					&ElementalSkill{lv: 4, dmgs: map[string]float64{"1段": 118, "2段": 122, "3段": 161}},
					&ElementalSkill{lv: 5, dmgs: map[string]float64{"1段": 125, "2段": 129, "3段": 171}},
					&ElementalSkill{lv: 6, dmgs: map[string]float64{"1段": 132, "2段": 137, "3段": 180}},
					&ElementalSkill{lv: 7, dmgs: map[string]float64{"1段": 142, "2段": 146, "3段": 193}},
					&ElementalSkill{lv: 8, dmgs: map[string]float64{"1段": 151, "2段": 156, "3段": 206}},
					&ElementalSkill{lv: 9, dmgs: map[string]float64{"1段": 160, "2段": 166, "3段": 219}},
					&ElementalSkill{lv: 10, dmgs: map[string]float64{"1段": 170, "2段": 176, "3段": 232}},
					&ElementalSkill{lv: 11, dmgs: map[string]float64{"1段": 179, "2段": 185, "3段": 245}},
					&ElementalSkill{lv: 12, dmgs: map[string]float64{"1段": 189, "2段": 195, "3段": 258}},
					&ElementalSkill{lv: 13, dmgs: map[string]float64{"1段": 201, "2段": 207, "3段": 274}},
				).
				addElementalBursts(
					&ElementalBurst{lv: 1, dmgs: map[string]float64{"斩击": 204, "持续": 60.0, "爆裂": 204}},
					&ElementalBurst{lv: 2, dmgs: map[string]float64{"斩击": 219, "持续": 64.5, "爆裂": 219}},
					&ElementalBurst{lv: 3, dmgs: map[string]float64{"斩击": 235, "持续": 69.0, "爆裂": 235}},
					&ElementalBurst{lv: 4, dmgs: map[string]float64{"斩击": 255, "持续": 75.0, "爆裂": 255}},
					&ElementalBurst{lv: 5, dmgs: map[string]float64{"斩击": 270, "持续": 79.5, "爆裂": 270}},
					&ElementalBurst{lv: 6, dmgs: map[string]float64{"斩击": 286, "持续": 84.0, "爆裂": 286}},
					&ElementalBurst{lv: 7, dmgs: map[string]float64{"斩击": 306, "持续": 90.0, "爆裂": 306}},
					&ElementalBurst{lv: 8, dmgs: map[string]float64{"斩击": 326, "持续": 96.0, "爆裂": 326}},
					&ElementalBurst{lv: 9, dmgs: map[string]float64{"斩击": 347, "持续": 102.0, "爆裂": 347}},
					&ElementalBurst{lv: 10, dmgs: map[string]float64{"斩击": 367, "持续": 108.0, "爆裂": 367}},
					&ElementalBurst{lv: 11, dmgs: map[string]float64{"斩击": 388, "持续": 114.0, "爆裂": 388}},
					&ElementalBurst{lv: 12, dmgs: map[string]float64{"斩击": 408, "持续": 120.0, "爆裂": 408}},
					&ElementalBurst{lv: 13, dmgs: map[string]float64{"斩击": 434, "持续": 128.0, "爆裂": 434}},
					&ElementalBurst{lv: 14, dmgs: map[string]float64{"斩击": 459, "持续": 135.0, "爆裂": 459}},
				).check()),
		).Talents(normal, skill, burst)
	}
)

type Character struct {
	star            int
	elemental       elemental.Elemental
	weaponType      weaponType.WeaponType
	level           int
	base            *attr.Attributes
	talents         Talents
	talentsTemplate *TalentsTemplate
	weapon          *Weapon
	artifacts       map[position.Position]*Artifacts
	attached        *attr.Attributes
}

type CharacterModifier func(character *Character) func()

func BaseCharacter(level, baseHp, baseAtk, baseDef int, baseModifier attr.AttributeModifier) CharacterModifier {
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
		base:       attr.NewAttributes(buff.AddCriticalRate(5), buff.AddCriticalDamage(50), buff.AddEnergyRecharge(100)),
		artifacts:  make(map[position.Position]*Artifacts),
		attached:   attr.NewAttributes(),
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

func (c *Character) Apply(modifiers ...attr.AttributeModifier) func() {
	return attr.MergeAttributes(modifiers...)(c.attached)
}

func (c *Character) GetActions() *Actions {
	return c.talents.DMGs(c.weaponType, c.elemental)
}

func (c *Character) basicAttributes() *attr.Attributes {
	basic := attr.NewAttributes()
	c.base.Accumulation()(basic)
	c.weapon.AccumulationBase()(basic)
	return basic
}

func (c *Character) finalAttributes() *attr.Attributes {
	final := attr.NewAttributes()
	c.basicAttributes().Accumulation()(final)
	final.Clear(point.Hp, point.Atk, point.Def)
	c.weapon.AccumulationRefine()(final)
	for _, artifacts := range c.artifacts {
		artifacts.Accumulation()(final)
	}
	c.attached.Accumulation()(final)
	return final
}

func (c *Character) String() string {
	return fmt.Sprintf("%s\n", c.base)
}

func (c *Character) Calculate(enemy *Enemy, action *Action, infusionElemental elemental.Elemental) *Calculator {
	return NewCalculator(c, enemy, action, infusionElemental)
}
