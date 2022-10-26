package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"github.com/dbstarll/game/internal/ys/model/weapon"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var (
	CharacterFactory草主 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(4, elementals.Grass, weaponType.Sword,
			BaseCharacter(90, 10875, 216, 683, buff.AddAtkPercentage(24)),
			TalentsTemplateModifier(NewTalentsTemplate(
				&NormalAttack{name: "异邦草翦", lv: 1, charged: ChargedAttack{stamina: 20}},
				&ElementalSkill{name: "草缘剑", lv: 13, cd: time.Second * 8},
				&ElementalBurst{name: "偃草若化", lv: 1, infusionDuration: time.Second * 12, cd: time.Second * 20, energyCost: 80}).
				addNormalAttacks(
					&NormalAttack{lv: 1, hits: []float64{44.5, 43.4, 53.0, 58.3, 70.8}, charged: ChargedAttack{hits: []float64{55.9, 60.8}}, plunge: PlungeAttack{63.9, 128, 160}},
				).
				addElementalSkills(
					&ElementalSkill{lv: 1, dmgs: map[string]float64{"技能伤害": 230}},
					&ElementalSkill{lv: 2, dmgs: map[string]float64{"技能伤害": 248}},
					&ElementalSkill{lv: 3, dmgs: map[string]float64{"技能伤害": 265}},
					&ElementalSkill{lv: 4, dmgs: map[string]float64{"技能伤害": 288}},
					&ElementalSkill{lv: 5, dmgs: map[string]float64{"技能伤害": 305}},
					&ElementalSkill{lv: 6, dmgs: map[string]float64{"技能伤害": 323}},
					&ElementalSkill{lv: 7, dmgs: map[string]float64{"技能伤害": 346}},
					&ElementalSkill{lv: 8, dmgs: map[string]float64{"技能伤害": 369}},
					&ElementalSkill{lv: 9, dmgs: map[string]float64{"技能伤害": 392}},
					&ElementalSkill{lv: 10, dmgs: map[string]float64{"技能伤害": 415}},
					&ElementalSkill{lv: 11, dmgs: map[string]float64{"技能伤害": 438}},
					&ElementalSkill{lv: 12, dmgs: map[string]float64{"技能伤害": 461}},
					&ElementalSkill{lv: 13, dmgs: map[string]float64{"技能伤害": 490}},
				).
				addElementalBursts(
					&ElementalBurst{lv: 1, dmgs: map[string]float64{"草灯莲攻击伤害": 80.2, "激烈爆发伤害": 400.8}},
				).check()),
		).Talents(normal, skill, burst)
	}
	CharacterFactory久岐忍 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(4, elementals.Electric, weaponType.Sword,
			BaseCharacter(90, 12289, 213, 751, buff.AddHpPercentage(24)),
			TalentsTemplateModifier(NewTalentsTemplate(
				&NormalAttack{name: "忍流飞刃斩", lv: 11, charged: ChargedAttack{stamina: 20}},
				&ElementalSkill{name: "越祓雷草之轮", lv: 13, duration: time.Second * 12, cd: time.Second * 15},
				&ElementalBurst{name: "御咏鸣神刈山祭", lv: 13, cd: time.Second * 15, energyCost: 60}).
				addNormalAttacks(
					&NormalAttack{lv: 1, hits: []float64{48.8, 44.5, 59.3, 76.1}, charged: ChargedAttack{hits: []float64{55.6, 66.8}}, plunge: PlungeAttack{63.9, 128, 160}},
					&NormalAttack{lv: 2, hits: []float64{52.7, 48.2, 64.2, 82.3}, charged: ChargedAttack{hits: []float64{60.2, 72.2}}, plunge: PlungeAttack{69.1, 138, 173}},
					&NormalAttack{lv: 3, hits: []float64{56.7, 51.8, 69.0, 88.5}, charged: ChargedAttack{hits: []float64{64.7, 77.6}}, plunge: PlungeAttack{74.3, 149, 186}},
					&NormalAttack{lv: 4, hits: []float64{62.4, 57.0, 75.9, 97.4}, charged: ChargedAttack{hits: []float64{71.2, 85.4}}, plunge: PlungeAttack{81.8, 164, 204}},
					&NormalAttack{lv: 5, hits: []float64{66.3, 60.6, 80.7, 103.5}, charged: ChargedAttack{hits: []float64{75.7, 90.8}}, plunge: PlungeAttack{87.0, 174, 217}},
					&NormalAttack{lv: 6, hits: []float64{70.9, 64.8, 86.3, 110.6}, charged: ChargedAttack{hits: []float64{80.9, 97.0}}, plunge: PlungeAttack{92.9, 186, 232}},
					&NormalAttack{lv: 7, hits: []float64{77.1, 70.4, 93.8, 120.4}, charged: ChargedAttack{hits: []float64{88.0, 105.6}}, plunge: PlungeAttack{101.1, 202, 253}},
					&NormalAttack{lv: 8, hits: []float64{83.3, 76.1, 101.4, 130.1}, charged: ChargedAttack{hits: []float64{95.1, 114.1}}, plunge: PlungeAttack{109.3, 219, 273}},
					&NormalAttack{lv: 9, hits: []float64{89.6, 81.8, 109.0, 139.8}, charged: ChargedAttack{hits: []float64{102.2, 122.7}}, plunge: PlungeAttack{117.5, 235, 293}},
					&NormalAttack{lv: 10, hits: []float64{96.4, 88.1, 117.3, 150.5}, charged: ChargedAttack{hits: []float64{110.0, 132.0}}, plunge: PlungeAttack{126.4, 253, 316}},
					&NormalAttack{lv: 11, hits: []float64{103.2, 94.3, 125.6, 161.1}, charged: ChargedAttack{hits: []float64{117.7, 141.3}}, plunge: PlungeAttack{135.3, 271, 338}},
				).
				addElementalSkills(
					&ElementalSkill{lv: 1, dmgs: map[string]float64{"技能伤害": 76, "越祓草轮伤害": 25.2}, curePercentage: 3.0, cure: 289},
					&ElementalSkill{lv: 2, dmgs: map[string]float64{"技能伤害": 81, "越祓草轮伤害": 27.1}, curePercentage: 3.2, cure: 318},
					&ElementalSkill{lv: 3, dmgs: map[string]float64{"技能伤害": 87, "越祓草轮伤害": 29.0}, curePercentage: 3.5, cure: 349},
					&ElementalSkill{lv: 4, dmgs: map[string]float64{"技能伤害": 95, "越祓草轮伤害": 31.6}, curePercentage: 3.8, cure: 383},
					&ElementalSkill{lv: 5, dmgs: map[string]float64{"技能伤害": 100, "越祓草轮伤害": 33.4}, curePercentage: 4.0, cure: 419},
					&ElementalSkill{lv: 6, dmgs: map[string]float64{"技能伤害": 106, "越祓草轮伤害": 35.3}, curePercentage: 4.2, cure: 457},
					&ElementalSkill{lv: 7, dmgs: map[string]float64{"技能伤害": 114, "越祓草轮伤害": 37.9}, curePercentage: 4.5, cure: 498},
					&ElementalSkill{lv: 8, dmgs: map[string]float64{"技能伤害": 121, "越祓草轮伤害": 40.4}, curePercentage: 4.8, cure: 542},
					&ElementalSkill{lv: 9, dmgs: map[string]float64{"技能伤害": 129, "越祓草轮伤害": 42.9}, curePercentage: 5.1, cure: 587},
					&ElementalSkill{lv: 10, dmgs: map[string]float64{"技能伤害": 136, "越祓草轮伤害": 45.4}, curePercentage: 5.4, cure: 636},
					&ElementalSkill{lv: 11, dmgs: map[string]float64{"技能伤害": 144, "越祓草轮伤害": 48.0}, curePercentage: 5.7, cure: 686},
					&ElementalSkill{lv: 12, dmgs: map[string]float64{"技能伤害": 151, "越祓草轮伤害": 50.5}, curePercentage: 6.0, cure: 739},
					&ElementalSkill{lv: 13, dmgs: map[string]float64{"技能伤害": 161, "越祓草轮伤害": 53.6}, curePercentage: 6.4, cure: 795},
				).
				addElementalBursts(
					&ElementalBurst{lv: 1, dmgs: map[string]float64{"技能伤害": 3.6, "总伤害": 25.2, "半血总伤害": 43.3}},
					&ElementalBurst{lv: 2, dmgs: map[string]float64{"技能伤害": 3.9, "总伤害": 27.1, "半血总伤害": 46.5}},
					&ElementalBurst{lv: 3, dmgs: map[string]float64{"技能伤害": 4.1, "总伤害": 29.0, "半血总伤害": 49.8}},
					&ElementalBurst{lv: 4, dmgs: map[string]float64{"技能伤害": 4.5, "总伤害": 31.5, "半血总伤害": 54.0}},
					&ElementalBurst{lv: 5, dmgs: map[string]float64{"技能伤害": 4.8, "总伤害": 33.4, "半血总伤害": 57.3}},
					&ElementalBurst{lv: 6, dmgs: map[string]float64{"技能伤害": 5.0, "总伤害": 35.3, "半血总伤害": 60.6}},
					&ElementalBurst{lv: 7, dmgs: map[string]float64{"技能伤害": 5.4, "总伤害": 37.9, "半血总伤害": 64.9}},
					&ElementalBurst{lv: 8, dmgs: map[string]float64{"技能伤害": 5.8, "总伤害": 40.4, "半血总伤害": 69.2}},
					&ElementalBurst{lv: 9, dmgs: map[string]float64{"技能伤害": 6.1, "总伤害": 42.9, "半血总伤害": 73.5}},
					&ElementalBurst{lv: 10, dmgs: map[string]float64{"技能伤害": 6.5, "总伤害": 45.4, "半血总伤害": 77.9}},
					&ElementalBurst{lv: 11, dmgs: map[string]float64{"技能伤害": 6.8, "总伤害": 47.9, "半血总伤害": 82.1}},
					&ElementalBurst{lv: 12, dmgs: map[string]float64{"技能伤害": 7.2, "总伤害": 50.5, "半血总伤害": 86.5}},
					&ElementalBurst{lv: 13, dmgs: map[string]float64{"技能伤害": 7.7, "总伤害": 53.6, "半血总伤害": 91.9}},
				).check()),
		).Talents(normal, skill, burst)
	}
	CharacterFactory迪卢克 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(5, elementals.Fire, weaponType.Claymore,
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
	elemental       elementals.Elemental
	weaponType      weaponType.WeaponType
	level           int
	base            *attr.Attributes
	talents         Talents
	talentsTemplate *TalentsTemplate
	weapon          *weapon.Weapon
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

func NewCharacter(star int, elemental elementals.Elemental, weaponType weaponType.WeaponType, modifiers ...CharacterModifier) *Character {
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

func (c *Character) GetActions() *action.Actions {
	return c.talents.DMGs(c.weaponType, c.elemental)
}

func (c *Character) basicAttributes() *attr.Attributes {
	basic := attr.NewAttributes()
	c.base.Accumulation(false)(basic)
	c.weapon.AccumulationBase()(basic)
	return basic
}

func (c *Character) finalAttributes() *attr.Attributes {
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

func (c *Character) Calculate(enemy *enemy.Enemy, action *action.Action, infusionElemental elementals.Elemental) *Calculator {
	return NewCalculator(c, enemy, action, infusionElemental)
}

func (c *Character) Evaluate() map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	for _, artifact := range c.artifacts {
		zap.S().Debugf("%s", artifact)
		detects[artifact.name] = attr.NewCharacterModifier(artifact.Accumulation(true))
		if artifactDetects := artifact.Evaluate(); len(artifactDetects) > 0 {
			for n, m := range artifactDetects {
				detects[fmt.Sprintf("%s - %s", artifact.name, n)] = m
			}
		}
		weaponBase, weaponRefine := attr.NewAttributes(), attr.NewAttributes()
		c.weapon.AccumulationBase()(weaponBase)
		c.weapon.AccumulationRefine()(weaponRefine)
		detects[fmt.Sprintf("%s - 基础", c.weapon.Name())] = attr.NewCharacterModifier(weaponBase.Accumulation(true))
		detects[fmt.Sprintf("%s - 精炼", c.weapon.Name())] = attr.NewCharacterModifier(weaponRefine.Accumulation(true))
		detects[c.weapon.Name()] = attr.NewCharacterModifier(attr.MergeAttributes(weaponBase.Accumulation(true), weaponRefine.Accumulation(true)))
	}
	return detects
}
