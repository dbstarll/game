package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory草主 = func(normal, skill, burst, constellation int) *Character {
	return New(4, elementals.Grass, weaponType.Sword,
		Base(90, 10875, 216, 683, buff.AddAtkPercentage(24)),
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
