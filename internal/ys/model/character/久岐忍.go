package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory久岐忍 = func(normal, skill, burst, constellation int) *Character {
	return New(4, elementals.Electric, weaponType.Sword,
		Base(90, 12289, 213, 751, buff.AddHpPercentage(24)),
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
