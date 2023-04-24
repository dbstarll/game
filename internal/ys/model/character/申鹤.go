package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory申鹤 = func(normal, skill, burst, constellation int) *Character {
	return New(5, elementals.Ice, weaponType.Polearm,
		Base(90, 12993, 304, 830, buff.AddAtkPercentage(28.8)),
		TalentsTemplateModifier(talent.NewTalentsTemplate(
			talent.BaseNormalAttack("踏辰摄斗", 11, 25),
			talent.BaseElementalSkill("仰灵威召将役咒", 14, time.Second*10, time.Second*10),
			talent.BaseElementalBurst("神女遣灵真诀", 14, 80, time.Second*20, time.Second*12)).
			AddNormalAttacks(
				talent.LevelNormalAttack(1, []float64{43.3, 40.3, 53.3, 26.3 * 2, 65.6}, 63.9, 110.7),
				talent.LevelNormalAttack(2, []float64{46.8, 43.5, 57.7, 28.5 * 2, 71.0}, 69.1, 119.7),
				talent.LevelNormalAttack(3, []float64{50.3, 46.8, 62.0, 30.6 * 2, 76.3}, 74.3, 128.7),
				talent.LevelNormalAttack(4, []float64{55.3, 51.5, 68.2, 33.7 * 2, 83.9}, 81.8, 141.6),
				talent.LevelNormalAttack(5, []float64{58.9, 54.8, 72.5, 35.8 * 2, 89.3}, 87.0, 150.6),
				talent.LevelNormalAttack(6, []float64{62.9, 58.5, 77.5, 38.3 * 2, 95.4}, 92.9, 160.9),
				talent.LevelNormalAttack(7, []float64{68.4, 63.7, 84.3, 41.6 * 2, 103.8}, 101.1, 175.0),
				talent.LevelNormalAttack(8, []float64{73.9, 68.8, 91.1, 45.0 * 2, 112.2}, 109.3, 189.2),
				talent.LevelNormalAttack(9, []float64{79.5, 73.9, 98.0, 48.3 * 2, 120.6}, 117.5, 203.3),
				talent.LevelNormalAttack(10, []float64{85.5, 79.6, 105.4, 52.0 * 2, 129.7}, 126.4, 218.8),
				talent.LevelNormalAttack(11, []float64{91.6, 85.2, 112.8, 55.7 * 2, 138.9}, 135.3, 234.2),
			).
			AddElementalSkills(
				talent.LevelElementalSkill(1, map[string]float64{"点按技能伤害": 139.0, "长按技能伤害": 188.8, "伤害值提升": 45.7}),
				talent.LevelElementalSkill(2, map[string]float64{"点按技能伤害": 150.0, "长按技能伤害": 203.0, "伤害值提升": 49.1}),
				talent.LevelElementalSkill(3, map[string]float64{"点按技能伤害": 160.0, "长按技能伤害": 217.1, "伤害值提升": 52.5}),
				talent.LevelElementalSkill(4, map[string]float64{"点按技能伤害": 174.0, "长按技能伤害": 236.0, "伤害值提升": 57.1}),
				talent.LevelElementalSkill(5, map[string]float64{"点按技能伤害": 184.0, "长按技能伤害": 250.0, "伤害值提升": 60.5}),
				talent.LevelElementalSkill(6, map[string]float64{"点按技能伤害": 195.0, "长按技能伤害": 264.3, "伤害值提升": 63.9}),
				talent.LevelElementalSkill(7, map[string]float64{"点按技能伤害": 209.0, "长按技能伤害": 283.2, "伤害值提升": 68.5}),
				talent.LevelElementalSkill(8, map[string]float64{"点按技能伤害": 223.0, "长按技能伤害": 302.1, "伤害值提升": 73.0}),
				talent.LevelElementalSkill(9, map[string]float64{"点按技能伤害": 237.0, "长按技能伤害": 321.0, "伤害值提升": 77.6}),
				talent.LevelElementalSkill(10, map[string]float64{"点按技能伤害": 251.0, "长按技能伤害": 339.8, "伤害值提升": 82.2}),
				talent.LevelElementalSkill(11, map[string]float64{"点按技能伤害": 264.0, "长按技能伤害": 358.7, "伤害值提升": 86.8}),
				talent.LevelElementalSkill(12, map[string]float64{"点按技能伤害": 278.0, "长按技能伤害": 377.6, "伤害值提升": 91.3}),
				talent.LevelElementalSkill(13, map[string]float64{"点按技能伤害": 296.0, "长按技能伤害": 401.2, "伤害值提升": 97.0}),
				talent.LevelElementalSkill(14, map[string]float64{"点按技能伤害": 313.0, "长按技能伤害": 424.8, "伤害值提升": 102.7}),
			).
			AddElementalBursts(
				talent.LevelElementalBurst(1, map[string]float64{"技能伤害": 101, "抗性降低": 6, "持续伤害": 33.1}),
				talent.LevelElementalBurst(2, map[string]float64{"技能伤害": 108, "抗性降低": 7, "持续伤害": 35.6}),
				talent.LevelElementalBurst(3, map[string]float64{"技能伤害": 116, "抗性降低": 8, "持续伤害": 38.1}),
				talent.LevelElementalBurst(4, map[string]float64{"技能伤害": 126, "抗性降低": 9, "持续伤害": 41.4}),
				talent.LevelElementalBurst(5, map[string]float64{"技能伤害": 134, "抗性降低": 10, "持续伤害": 43.9}),
				talent.LevelElementalBurst(6, map[string]float64{"技能伤害": 141, "抗性降低": 11, "持续伤害": 46.4}),
				talent.LevelElementalBurst(7, map[string]float64{"技能伤害": 151, "抗性降低": 12, "持续伤害": 49.7}),
				talent.LevelElementalBurst(8, map[string]float64{"技能伤害": 161, "抗性降低": 13, "持续伤害": 53.0}),
				talent.LevelElementalBurst(9, map[string]float64{"技能伤害": 171, "抗性降低": 14, "持续伤害": 56.3}),
				talent.LevelElementalBurst(10, map[string]float64{"技能伤害": 181, "抗性降低": 15, "持续伤害": 59.6}),
				talent.LevelElementalBurst(11, map[string]float64{"技能伤害": 192, "抗性降低": 15, "持续伤害": 62.9}),
				talent.LevelElementalBurst(12, map[string]float64{"技能伤害": 202, "抗性降低": 15, "持续伤害": 66.2}),
				talent.LevelElementalBurst(13, map[string]float64{"技能伤害": 214, "抗性降低": 15, "持续伤害": 70.4}),
				talent.LevelElementalBurst(14, map[string]float64{"技能伤害": 227, "抗性降低": 15, "持续伤害": 74.5}),
			).Check()),
	).Talents(normal, skill, burst)
}
