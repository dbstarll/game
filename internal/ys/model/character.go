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
	"github.com/dbstarll/game/internal/ys/model/talent"
	"github.com/dbstarll/game/internal/ys/model/weapon"
	"github.com/pkg/errors"
	"time"
)

var (
	CharacterFactory草主 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(4, elementals.Grass, weaponType.Sword,
			BaseCharacter(90, 10875, 216, 683, buff.AddAtkPercentage(24)),
			TalentsTemplateModifier(talent.NewTalentsTemplate(
				talent.BaseNormalAttack("异邦草翦", 1, 20),
				talent.BaseElementalSkill("草缘剑", 13, time.Second*8, 0),
				talent.BaseElementalBurst("偃草若化", 1, 80, time.Second*20, time.Second*12)).
				AddNormalAttacks(
					talent.LevelNormalAttack(1, []float64{44.5, 43.4, 53.0, 58.3, 70.8}, 63.9, 55.9, 60.8),
				).
				AddElementalSkills(
					talent.LevelElementalSkill(1, map[string]float64{"技能伤害": 230}),
					talent.LevelElementalSkill(2, map[string]float64{"技能伤害": 248}),
					talent.LevelElementalSkill(3, map[string]float64{"技能伤害": 265}),
					talent.LevelElementalSkill(4, map[string]float64{"技能伤害": 288}),
					talent.LevelElementalSkill(5, map[string]float64{"技能伤害": 305}),
					talent.LevelElementalSkill(6, map[string]float64{"技能伤害": 323}),
					talent.LevelElementalSkill(7, map[string]float64{"技能伤害": 346}),
					talent.LevelElementalSkill(8, map[string]float64{"技能伤害": 369}),
					talent.LevelElementalSkill(9, map[string]float64{"技能伤害": 392}),
					talent.LevelElementalSkill(10, map[string]float64{"技能伤害": 415}),
					talent.LevelElementalSkill(11, map[string]float64{"技能伤害": 438}),
					talent.LevelElementalSkill(12, map[string]float64{"技能伤害": 461}),
					talent.LevelElementalSkill(13, map[string]float64{"技能伤害": 490}),
				).
				AddElementalBursts(
					talent.LevelElementalBurst(1, map[string]float64{"草灯莲攻击伤害": 80.2, "激烈爆发伤害": 400.8}),
				).Check()),
		).Talents(normal, skill, burst)
	}
	CharacterFactory久岐忍 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(4, elementals.Electric, weaponType.Sword,
			BaseCharacter(90, 12289, 213, 751, buff.AddHpPercentage(24)),
			TalentsTemplateModifier(talent.NewTalentsTemplate(
				talent.BaseNormalAttack("忍流飞刃斩", 11, 20),
				talent.BaseElementalSkill("越祓雷草之轮", 13, time.Second*15, time.Second*12),
				talent.BaseElementalBurst("御咏鸣神刈山祭", 13, 60, time.Second*15, 0)).
				AddNormalAttacks(
					talent.LevelNormalAttack(1, []float64{48.8, 44.5, 59.3, 76.1}, 63.9, 55.6, 66.8),
					talent.LevelNormalAttack(2, []float64{52.7, 48.2, 64.2, 82.3}, 69.1, 60.2, 72.2),
					talent.LevelNormalAttack(3, []float64{56.7, 51.8, 69.0, 88.5}, 74.3, 64.7, 77.6),
					talent.LevelNormalAttack(4, []float64{62.4, 57.0, 75.9, 97.4}, 81.8, 71.2, 85.4),
					talent.LevelNormalAttack(5, []float64{66.3, 60.6, 80.7, 103.5}, 87.0, 75.7, 90.8),
					talent.LevelNormalAttack(6, []float64{70.9, 64.8, 86.3, 110.6}, 92.9, 80.9, 97.0),
					talent.LevelNormalAttack(7, []float64{77.1, 70.4, 93.8, 120.4}, 101.1, 88.0, 105.6),
					talent.LevelNormalAttack(8, []float64{83.3, 76.1, 101.4, 130.1}, 109.3, 95.1, 114.1),
					talent.LevelNormalAttack(9, []float64{89.6, 81.8, 109.0, 139.8}, 117.5, 102.2, 122.7),
					talent.LevelNormalAttack(10, []float64{96.4, 88.1, 117.3, 150.5}, 126.4, 110.0, 132.0),
					talent.LevelNormalAttack(11, []float64{103.2, 94.3, 125.6, 161.1}, 135.3, 117.7, 141.3),
				).
				AddElementalSkills(
					talent.LevelElementalSkillWithCure(1, map[string]float64{"技能伤害": 76, "越祓草轮伤害": 25.2}, 3.0, 289),
					talent.LevelElementalSkillWithCure(2, map[string]float64{"技能伤害": 81, "越祓草轮伤害": 27.1}, 3.2, 318),
					talent.LevelElementalSkillWithCure(3, map[string]float64{"技能伤害": 87, "越祓草轮伤害": 29.0}, 3.5, 349),
					talent.LevelElementalSkillWithCure(4, map[string]float64{"技能伤害": 95, "越祓草轮伤害": 31.6}, 3.8, 383),
					talent.LevelElementalSkillWithCure(5, map[string]float64{"技能伤害": 100, "越祓草轮伤害": 33.4}, 4.0, 419),
					talent.LevelElementalSkillWithCure(6, map[string]float64{"技能伤害": 106, "越祓草轮伤害": 35.3}, 4.2, 457),
					talent.LevelElementalSkillWithCure(7, map[string]float64{"技能伤害": 114, "越祓草轮伤害": 37.9}, 4.5, 498),
					talent.LevelElementalSkillWithCure(8, map[string]float64{"技能伤害": 121, "越祓草轮伤害": 40.4}, 4.8, 542),
					talent.LevelElementalSkillWithCure(9, map[string]float64{"技能伤害": 129, "越祓草轮伤害": 42.9}, 5.1, 587),
					talent.LevelElementalSkillWithCure(10, map[string]float64{"技能伤害": 136, "越祓草轮伤害": 45.4}, 5.4, 636),
					talent.LevelElementalSkillWithCure(11, map[string]float64{"技能伤害": 144, "越祓草轮伤害": 48.0}, 5.7, 686),
					talent.LevelElementalSkillWithCure(12, map[string]float64{"技能伤害": 151, "越祓草轮伤害": 50.5}, 6.0, 739),
					talent.LevelElementalSkillWithCure(13, map[string]float64{"技能伤害": 161, "越祓草轮伤害": 53.6}, 6.4, 795),
				).
				AddElementalBursts(
					talent.LevelElementalBurst(1, map[string]float64{"技能伤害": 3.6, "总伤害": 25.2, "半血总伤害": 43.3}),
					talent.LevelElementalBurst(2, map[string]float64{"技能伤害": 3.9, "总伤害": 27.1, "半血总伤害": 46.5}),
					talent.LevelElementalBurst(3, map[string]float64{"技能伤害": 4.1, "总伤害": 29.0, "半血总伤害": 49.8}),
					talent.LevelElementalBurst(4, map[string]float64{"技能伤害": 4.5, "总伤害": 31.5, "半血总伤害": 54.0}),
					talent.LevelElementalBurst(5, map[string]float64{"技能伤害": 4.8, "总伤害": 33.4, "半血总伤害": 57.3}),
					talent.LevelElementalBurst(6, map[string]float64{"技能伤害": 5.0, "总伤害": 35.3, "半血总伤害": 60.6}),
					talent.LevelElementalBurst(7, map[string]float64{"技能伤害": 5.4, "总伤害": 37.9, "半血总伤害": 64.9}),
					talent.LevelElementalBurst(8, map[string]float64{"技能伤害": 5.8, "总伤害": 40.4, "半血总伤害": 69.2}),
					talent.LevelElementalBurst(9, map[string]float64{"技能伤害": 6.1, "总伤害": 42.9, "半血总伤害": 73.5}),
					talent.LevelElementalBurst(10, map[string]float64{"技能伤害": 6.5, "总伤害": 45.4, "半血总伤害": 77.9}),
					talent.LevelElementalBurst(11, map[string]float64{"技能伤害": 6.8, "总伤害": 47.9, "半血总伤害": 82.1}),
					talent.LevelElementalBurst(12, map[string]float64{"技能伤害": 7.2, "总伤害": 50.5, "半血总伤害": 86.5}),
					talent.LevelElementalBurst(13, map[string]float64{"技能伤害": 7.7, "总伤害": 53.6, "半血总伤害": 91.9}),
				).Check()),
		).Talents(normal, skill, burst)
	}
	CharacterFactory迪卢克 = func(normal, skill, burst, constellation int) *Character {
		return NewCharacter(5, elementals.Fire, weaponType.Claymore,
			BaseCharacter(90, 12981, 335, 784, buff.AddCriticalRate(19.2)),
			TalentsTemplateModifier(talent.NewTalentsTemplate(
				talent.BaseNormalAttackByCyclic("淬炼之剑", 11, 40, time.Second*5),
				talent.BaseElementalSkill("逆焰之刃", 13, time.Second*10, 0),
				talent.BaseElementalBurst("黎明", 14, 40, time.Second*12, time.Second*8)).
				AddNormalAttacks(
					talent.LevelNormalAttackByCyclic(1, []float64{89.7, 87.6, 98.8, 134}, 89.5, 68.8, 125),
					talent.LevelNormalAttackByCyclic(2, []float64{97.0, 94.8, 107, 145}, 96.8, 74.4, 135),
					talent.LevelNormalAttackByCyclic(3, []float64{104, 102, 115, 156}, 104, 80, 145),
					talent.LevelNormalAttackByCyclic(4, []float64{115, 112, 126, 171}, 114, 88, 160),
					talent.LevelNormalAttackByCyclic(5, []float64{122, 119, 134, 182}, 122, 93.6, 170),
					talent.LevelNormalAttackByCyclic(6, []float64{130, 127, 144, 195}, 130, 100, 181),
					talent.LevelNormalAttackByCyclic(7, []float64{142, 139, 156, 212}, 142, 109, 197),
					talent.LevelNormalAttackByCyclic(8, []float64{153, 150, 169, 229}, 153, 118, 213),
					talent.LevelNormalAttackByCyclic(9, []float64{165, 161, 182, 246}, 164, 126, 229),
					talent.LevelNormalAttackByCyclic(10, []float64{177, 173, 195, 265}, 177, 136, 247),
					talent.LevelNormalAttackByCyclic(11, []float64{192, 187, 211, 286}, 189, 147, 266),
				).
				AddElementalSkills(
					talent.LevelElementalSkill(1, map[string]float64{"1段": 94.4, "2段": 97.6, "3段": 129}),
					talent.LevelElementalSkill(2, map[string]float64{"1段": 101, "2段": 105, "3段": 138}),
					talent.LevelElementalSkill(3, map[string]float64{"1段": 109, "2段": 112, "3段": 148}),
					talent.LevelElementalSkill(4, map[string]float64{"1段": 118, "2段": 122, "3段": 161}),
					talent.LevelElementalSkill(5, map[string]float64{"1段": 125, "2段": 129, "3段": 171}),
					talent.LevelElementalSkill(6, map[string]float64{"1段": 132, "2段": 137, "3段": 180}),
					talent.LevelElementalSkill(7, map[string]float64{"1段": 142, "2段": 146, "3段": 193}),
					talent.LevelElementalSkill(8, map[string]float64{"1段": 151, "2段": 156, "3段": 206}),
					talent.LevelElementalSkill(9, map[string]float64{"1段": 160, "2段": 166, "3段": 219}),
					talent.LevelElementalSkill(10, map[string]float64{"1段": 170, "2段": 176, "3段": 232}),
					talent.LevelElementalSkill(11, map[string]float64{"1段": 179, "2段": 185, "3段": 245}),
					talent.LevelElementalSkill(12, map[string]float64{"1段": 189, "2段": 195, "3段": 258}),
					talent.LevelElementalSkill(13, map[string]float64{"1段": 201, "2段": 207, "3段": 274}),
				).
				AddElementalBursts(
					talent.LevelElementalBurst(1, map[string]float64{"斩击": 204, "持续": 60.0, "爆裂": 204}),
					talent.LevelElementalBurst(2, map[string]float64{"斩击": 219, "持续": 64.5, "爆裂": 219}),
					talent.LevelElementalBurst(3, map[string]float64{"斩击": 235, "持续": 69.0, "爆裂": 235}),
					talent.LevelElementalBurst(4, map[string]float64{"斩击": 255, "持续": 75.0, "爆裂": 255}),
					talent.LevelElementalBurst(5, map[string]float64{"斩击": 270, "持续": 79.5, "爆裂": 270}),
					talent.LevelElementalBurst(6, map[string]float64{"斩击": 286, "持续": 84.0, "爆裂": 286}),
					talent.LevelElementalBurst(7, map[string]float64{"斩击": 306, "持续": 90.0, "爆裂": 306}),
					talent.LevelElementalBurst(8, map[string]float64{"斩击": 326, "持续": 96.0, "爆裂": 326}),
					talent.LevelElementalBurst(9, map[string]float64{"斩击": 347, "持续": 102.0, "爆裂": 347}),
					talent.LevelElementalBurst(10, map[string]float64{"斩击": 367, "持续": 108.0, "爆裂": 367}),
					talent.LevelElementalBurst(11, map[string]float64{"斩击": 388, "持续": 114.0, "爆裂": 388}),
					talent.LevelElementalBurst(12, map[string]float64{"斩击": 408, "持续": 120.0, "爆裂": 408}),
					talent.LevelElementalBurst(13, map[string]float64{"斩击": 434, "持续": 128.0, "爆裂": 434}),
					talent.LevelElementalBurst(14, map[string]float64{"斩击": 459, "持续": 135.0, "爆裂": 459}),
				).Check()),
		).Talents(normal, skill, burst)
	}
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

func TalentsTemplateModifier(talentsTemplate *talent.TalentsTemplate) CharacterModifier {
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

func (c *Character) Evaluate(replaceArtifacts map[position.Position]*Artifacts) map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	for _, artifact := range c.artifacts {
		for n, m := range artifact.Evaluate(replaceArtifacts[artifact.position]) {
			detects[n] = m
		}
	}
	for n, m := range c.weapon.Evaluate() {
		detects[n] = m
	}
	return detects
}
