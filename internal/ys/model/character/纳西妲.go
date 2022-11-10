package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory纳西妲 = func(normal, skill, burst, constellation int) *Character {
	return New(5, elementals.Grass, weaponType.Catalyst,
		Base(90, 10360, 299, 630, buff.AddElementalMastery(115)),
		TalentsTemplateModifier(talent.NewTalentsTemplate(
			talent.BaseNormalAttack("行相", 11, 50),
			talent.BaseElementalSkillWithInterval("所闻遍计", 13, time.Second*5, time.Second*25, time.Millisecond*2500),
			talent.BaseElementalBurst("心景幻成", 13, 50, time.Millisecond*13500, time.Second*15)).
			AddNormalAttacks(
				talent.LevelNormalAttack(1, []float64{40.3, 37.0, 45.9, 58.4}, 56.8, 132.0),
				talent.LevelNormalAttack(2, []float64{43.3, 39.7, 49.3, 62.8}, 61.5, 141.9),
				talent.LevelNormalAttack(3, []float64{46.4, 42.5, 52.8, 67.2}, 66.1, 151.8),
				talent.LevelNormalAttack(4, []float64{50.4, 46.2, 57.3, 73.0}, 72.7, 165.0),
				talent.LevelNormalAttack(5, []float64{53.4, 49.0, 60.8, 77.4}, 77.3, 174.9),
				talent.LevelNormalAttack(6, []float64{56.4, 51.8, 64.2, 81.8}, 82.6, 184.8),
				talent.LevelNormalAttack(7, []float64{60.5, 55.5, 68.8, 87.6}, 89.9, 198.0),
				talent.LevelNormalAttack(8, []float64{64.5, 59.2, 73.4, 93.5}, 97.1, 211.2),
				talent.LevelNormalAttack(9, []float64{68.5, 62.9, 78.0, 99.3}, 104.4, 224.4),
				talent.LevelNormalAttack(10, []float64{72.5, 66.6, 82.6, 105.1}, 112.3, 237.6),
				talent.LevelNormalAttack(11, []float64{76.6, 70.3, 87.2, 111.0}, 120.3, 250.8),
			).
			AddElementalSkills(
				talent.LevelElementalSkill(1, map[string]float64{"点按": 98.4, "长按": 130.4, "灭净三业": 103.2}),
				talent.LevelElementalSkill(2, map[string]float64{"点按": 105.8, "长按": 140.2, "灭净三业": 110.9}),
				talent.LevelElementalSkill(3, map[string]float64{"点按": 113.2, "长按": 150.0, "灭净三业": 118.7}),
				talent.LevelElementalSkill(4, map[string]float64{"点按": 123.0, "长按": 163.0, "灭净三业": 129.0}),
				talent.LevelElementalSkill(5, map[string]float64{"点按": 130.4, "长按": 172.8, "灭净三业": 136.7}),
				talent.LevelElementalSkill(6, map[string]float64{"点按": 137.8, "长按": 182.6, "灭净三业": 144.5}),
				talent.LevelElementalSkill(7, map[string]float64{"点按": 147.6, "长按": 195.6, "灭净三业": 154.8}),
				talent.LevelElementalSkill(8, map[string]float64{"点按": 157.4, "长按": 208.6, "灭净三业": 165.1}),
				talent.LevelElementalSkill(9, map[string]float64{"点按": 167.3, "长按": 221.7, "灭净三业": 175.4}),
				talent.LevelElementalSkill(10, map[string]float64{"点按": 177.1, "长按": 234.7, "灭净三业": 185.8}),
				talent.LevelElementalSkill(11, map[string]float64{"点按": 187.0, "长按": 247.8, "灭净三业": 196.1}),
				talent.LevelElementalSkill(12, map[string]float64{"点按": 196.8, "长按": 260.8, "灭净三业": 206.4}),
				talent.LevelElementalSkill(13, map[string]float64{"点按": 209.1, "长按": 277.1, "灭净三业": 219.3}),
			).
			AddElementalBursts(
				talent.LevelElementalBurst(1, map[string]float64{"伤害1": 14.9, "伤害2": 22.3, "间隔1": 0.25, "间隔2": 0.37, "持续1": 3.34, "持续2": 5.02}),
				talent.LevelElementalBurst(2, map[string]float64{"伤害1": 16.0, "伤害2": 24.0, "间隔1": 0.27, "间隔2": 0.40, "持续1": 3.59, "持续2": 5.39}),
				talent.LevelElementalBurst(3, map[string]float64{"伤害1": 17.1, "伤害2": 25.7, "间隔1": 0.29, "间隔2": 0.43, "持续1": 3.85, "持续2": 5.77}),
				talent.LevelElementalBurst(4, map[string]float64{"伤害1": 18.6, "伤害2": 27.9, "间隔1": 0.31, "间隔2": 0.47, "持续1": 4.18, "持续2": 6.27}),
				talent.LevelElementalBurst(5, map[string]float64{"伤害1": 19.7, "伤害2": 29.6, "间隔1": 0.33, "间隔2": 0.49, "持续1": 4.43, "持续2": 6.65}),
				talent.LevelElementalBurst(6, map[string]float64{"伤害1": 20.8, "伤害2": 31.3, "间隔1": 0.35, "间隔2": 0.52, "持续1": 4.68, "持续2": 7.02}),
				talent.LevelElementalBurst(7, map[string]float64{"伤害1": 22.3, "伤害2": 33.5, "间隔1": 0.37, "间隔2": 0.56, "持续1": 5.02, "持续2": 7.52}),
				talent.LevelElementalBurst(8, map[string]float64{"伤害1": 23.8, "伤害2": 35.7, "间隔1": 0.40, "间隔2": 0.60, "持续1": 5.35, "持续2": 8.03}),
				talent.LevelElementalBurst(9, map[string]float64{"伤害1": 25.3, "伤害2": 37.9, "间隔1": 0.42, "间隔2": 0.63, "持续1": 5.68, "持续2": 8.53}),
				talent.LevelElementalBurst(10, map[string]float64{"伤害1": 26.8, "伤害2": 40.2, "间隔1": 0.45, "间隔2": 0.67, "持续1": 6.02, "持续2": 9.03}),
				talent.LevelElementalBurst(11, map[string]float64{"伤害1": 28.27, "伤害2": 42.41, "间隔1": 0.47, "间隔2": 0.70, "持续1": 6.35, "持续2": 9.53}),
				talent.LevelElementalBurst(12, map[string]float64{"伤害1": 29.8, "伤害2": 44.6, "间隔1": 0.50, "间隔2": 0.74, "持续1": 6.69, "持续2": 10.03}),
				talent.LevelElementalBurst(13, map[string]float64{"伤害1": 31.6, "伤害2": 47.4, "间隔1": 0.53, "间隔2": 0.79, "持续1": 7.11, "持续2": 10.66}),
			).Check()),
	).Talents(normal, skill, burst)
}
