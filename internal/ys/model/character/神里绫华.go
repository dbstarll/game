package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory神里绫华 = func(normal, skill, burst, constellation int) *Character {
	return New(5, elementals.Ice, weaponType.Sword,
		Base(90, 12858, 342, 784, buff.AddCriticalDamage(38.4)),
		TalentsTemplateModifier(talent.NewTalentsTemplate(
			talent.BaseNormalAttack("神里流·倾", 11, 20),
			talent.BaseElementalSkill("神里流·冰华", 13, time.Second*10, 0),
			talent.BaseElementalBurst("神里流·霜灭", 14, 80, time.Second*20, time.Second*5)).
			AddNormalAttacks(
				talent.LevelNormalAttack(1, []float64{45.7, 48.7, 62.6, 22.7 * 3, 78.2}, 64, 55.1, 55.1, 55.1),
				talent.LevelNormalAttack(2, []float64{49.5, 52.7, 67.7, 24.5 * 3, 84.6}, 69, 59.6, 59.6, 59.6),
				talent.LevelNormalAttack(3, []float64{53.2, 56.6, 72.8, 26.2 * 3, 90.9}, 74, 64.1, 64.1, 64.1),
				talent.LevelNormalAttack(4, []float64{58.5, 62.3, 80.1, 29.0 * 3, 100.0}, 82, 70.5, 70.5, 70.5),
				talent.LevelNormalAttack(5, []float64{62.2, 66.2, 85.2, 30.8 * 3, 106.4}, 87, 75.0, 75.0, 75.0),
				talent.LevelNormalAttack(6, []float64{66.5, 70.8, 91.0, 32.9 * 3, 113.6}, 93, 80.1, 80.1, 80.1),
				talent.LevelNormalAttack(7, []float64{72.3, 77.0, 99.0, 35.8 * 3, 123.6}, 101, 87.2, 87.2, 87.2),
				talent.LevelNormalAttack(8, []float64{78.2, 83.2, 107.0, 38.7 * 3, 133.6}, 109, 94.2, 94.2, 94.2),
				talent.LevelNormalAttack(9, []float64{84.0, 89.4, 115.1, 41.6 * 3, 143.6}, 117, 101.3, 101.3, 101.3),
				talent.LevelNormalAttack(10, []float64{90.4, 96.2, 123.8, 44.8 * 3, 154.6}, 126, 109.0, 109.0, 109.0),
				talent.LevelNormalAttack(11, []float64{96.8, 103.0, 132.5, 47.9 * 3, 165.5}, 135, 116.7, 116.7, 116.7),
			).
			AddElementalSkills(
				talent.LevelElementalSkill(1, map[string]float64{"技能": 239}),
				talent.LevelElementalSkill(2, map[string]float64{"技能": 257}),
				talent.LevelElementalSkill(3, map[string]float64{"技能": 275}),
				talent.LevelElementalSkill(4, map[string]float64{"技能": 299}),
				talent.LevelElementalSkill(5, map[string]float64{"技能": 317}),
				talent.LevelElementalSkill(6, map[string]float64{"技能": 335}),
				talent.LevelElementalSkill(7, map[string]float64{"技能": 359}),
				talent.LevelElementalSkill(8, map[string]float64{"技能": 383}),
				talent.LevelElementalSkill(9, map[string]float64{"技能": 407}),
				talent.LevelElementalSkill(10, map[string]float64{"技能": 431}),
				talent.LevelElementalSkill(11, map[string]float64{"技能": 454}),
				talent.LevelElementalSkill(12, map[string]float64{"技能": 478}),
				talent.LevelElementalSkill(13, map[string]float64{"技能": 508}),
			).
			AddElementalBursts(
				talent.LevelElementalBurst(1, map[string]float64{"切割": 112, "绽放": 168}),
				talent.LevelElementalBurst(2, map[string]float64{"切割": 121, "绽放": 181}),
				talent.LevelElementalBurst(3, map[string]float64{"切割": 129, "绽放": 194}),
				talent.LevelElementalBurst(4, map[string]float64{"切割": 140, "绽放": 211}),
				talent.LevelElementalBurst(5, map[string]float64{"切割": 149, "绽放": 223}),
				talent.LevelElementalBurst(6, map[string]float64{"切割": 157, "绽放": 236}),
				talent.LevelElementalBurst(7, map[string]float64{"切割": 168, "绽放": 253}),
				talent.LevelElementalBurst(8, map[string]float64{"切割": 180, "绽放": 270}),
				talent.LevelElementalBurst(9, map[string]float64{"切割": 191, "绽放": 286}),
				talent.LevelElementalBurst(10, map[string]float64{"切割": 202, "绽放": 303}),
				talent.LevelElementalBurst(11, map[string]float64{"切割": 213, "绽放": 320}),
				talent.LevelElementalBurst(12, map[string]float64{"切割": 225, "绽放": 337}),
				talent.LevelElementalBurst(13, map[string]float64{"切割": 239, "绽放": 358}),
				talent.LevelElementalBurst(14, map[string]float64{"切割": 253, "绽放": 379}),
			).Check()),
	).Talents(normal, skill, burst)
}
