package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var Factory雷电将军 = func(normal, skill, burst, constellation int) *Character {
	return New(5, elementals.Electric, weaponType.Polearm,
		Base(90, 12907, 337, 789, buff.AddEnergyRecharge(32)),
		TalentsTemplateModifier(talent.NewTalentsTemplate(
			talent.BaseNormalAttack("普通攻击·源流", 11, 20),
			talent.BaseElementalSkill("神变·恶曜开眼", 13, time.Second*10, time.Second*25),
			talent.BaseElementalBurst("奥义·梦想真说", 14, 90, time.Second*18, time.Second*7)).
			AddNormalAttacks(
				talent.LevelNormalAttack(1, []float64{39.6, 39.7, 49.9, 29.0 * 2, 65.4}, 63.9, 99.6),
				talent.LevelNormalAttack(2, []float64{42.9, 43.0, 53.9, 31.3 * 2, 70.8}, 69.1, 107.7),
				talent.LevelNormalAttack(3, []float64{46.1, 46.2, 58.0, 33.7 * 2, 76.1}, 74.3, 115.8),
				talent.LevelNormalAttack(4, []float64{50.7, 50.8, 63.8, 37.1 * 2, 83.7}, 81.8, 127.4),
				talent.LevelNormalAttack(5, []float64{53.9, 54.1, 67.9, 39.4 * 2, 89.0}, 87.0, 135.5),
				talent.LevelNormalAttack(6, []float64{57.6, 57.8, 72.5, 42.1 * 2, 95.1}, 92.9, 144.9),
				talent.LevelNormalAttack(7, []float64{62.7, 62.8, 78.9, 45.8 * 2, 103.5}, 101.1, 157.5),
				talent.LevelNormalAttack(8, []float64{67.8, 67.9, 85.3, 49.5 * 2, 111.9}, 109.3, 170.2),
				talent.LevelNormalAttack(9, []float64{72.9, 73.0, 91.6, 53.2 * 2, 120.2}, 117.5, 183.0),
				talent.LevelNormalAttack(10, []float64{78.4, 78.5, 98.6, 57.3 * 2, 129.4}, 126.4, 196.9),
				talent.LevelNormalAttack(11, []float64{84.7, 84.9, 106.6, 61.9 * 2, 139.8}, 135.3, 212.8),
			).
			AddElementalSkills(
				talent.LevelElementalSkill(1, map[string]float64{"技能": 117.2, "协同攻击": 42.0, "元素爆发伤害提高": 0.22}),
				talent.LevelElementalSkill(2, map[string]float64{"技能": 126.0, "协同攻击": 45.2, "元素爆发伤害提高": 0.23}),
				talent.LevelElementalSkill(3, map[string]float64{"技能": 134.8, "协同攻击": 48.3, "元素爆发伤害提高": 0.24}),
				talent.LevelElementalSkill(4, map[string]float64{"技能": 146.5, "协同攻击": 52.5, "元素爆发伤害提高": 0.25}),
				talent.LevelElementalSkill(5, map[string]float64{"技能": 155.3, "协同攻击": 55.7, "元素爆发伤害提高": 0.26}),
				talent.LevelElementalSkill(6, map[string]float64{"技能": 164.1, "协同攻击": 58.8, "元素爆发伤害提高": 0.27}),
				talent.LevelElementalSkill(7, map[string]float64{"技能": 175.8, "协同攻击": 63.0, "元素爆发伤害提高": 0.28}),
				talent.LevelElementalSkill(8, map[string]float64{"技能": 187.5, "协同攻击": 67.2, "元素爆发伤害提高": 0.29}),
				talent.LevelElementalSkill(9, map[string]float64{"技能": 199.2, "协同攻击": 71.4, "元素爆发伤害提高": 0.30}),
				talent.LevelElementalSkill(10, map[string]float64{"技能": 211.0, "协同攻击": 75.6, "元素爆发伤害提高": 0.30}),
				talent.LevelElementalSkill(11, map[string]float64{"技能": 222.7, "协同攻击": 79.8, "元素爆发伤害提高": 0.30}),
				talent.LevelElementalSkill(12, map[string]float64{"技能": 234.4, "协同攻击": 84.0, "元素爆发伤害提高": 0.30}),
				talent.LevelElementalSkill(13, map[string]float64{"技能": 249.1, "协同攻击": 89.3, "元素爆发伤害提高": 0.30}),
			).
			AddElementalBursts(
				talent.LevelElementalBurst(1, map[string]float64{"梦想一刀": 401, "梦想一刀愿力加成": 3.89, "梦想一心愿力加成": 0.73, "积攒愿力层数": 0.15, "1段": 44.7, "2段": 44.0, "3段": 53.8, "4段": 30.9 + 31.0, "5段": 73.9, "重击": 61.6 + 74.4, "下坠": 63.9, "梦想一心能量恢复": 1.6}),
				talent.LevelElementalBurst(2, map[string]float64{"梦想一刀": 431, "梦想一刀愿力加成": 4.18, "梦想一心愿力加成": 0.78, "积攒愿力层数": 0.16, "1段": 47.8, "2段": 47.0, "3段": 57.5, "4段": 33.0 + 33.1, "5段": 79.0, "重击": 65.8 + 79.4, "下坠": 69.1, "梦想一心能量恢复": 1.7}),
				talent.LevelElementalBurst(3, map[string]float64{"梦想一刀": 461, "梦想一刀愿力加成": 4.47, "梦想一心愿力加成": 0.84, "积攒愿力层数": 0.16, "1段": 50.8, "2段": 50.0, "3段": 61.2, "4段": 35.1 + 35.2, "5段": 84.0, "重击": 70.0 + 84.5, "下坠": 74.3, "梦想一心能量恢复": 1.8}),
				talent.LevelElementalBurst(4, map[string]float64{"梦想一刀": 501, "梦想一刀愿力加成": 4.86, "梦想一心愿力加成": 0.91, "积攒愿力层数": 0.17, "1段": 54.9, "2段": 53.9, "3段": 66.1, "4段": 37.9 + 38.0, "5段": 90.7, "重击": 75.6 + 91.3, "下坠": 81.8, "梦想一心能量恢复": 1.9}),
				talent.LevelElementalBurst(5, map[string]float64{"梦想一刀": 531, "梦想一刀愿力加成": 5.15, "梦想一心愿力加成": 0.96, "积攒愿力层数": 0.17, "1段": 58.0, "2段": 56.9, "3段": 69.7, "4段": 40.0 + 40.1, "5段": 95.8, "重击": 79.8 + 96.3, "下坠": 87.0, "梦想一心能量恢复": 2.0}),
				talent.LevelElementalBurst(6, map[string]float64{"梦想一刀": 561, "梦想一刀愿力加成": 5.44, "梦想一心愿力加成": 1.02, "积攒愿力层数": 0.18, "1段": 61.5, "2段": 60.4, "3段": 74.0, "4段": 42.5 + 42.6, "5段": 101.7, "重击": 84.7 + 102.2, "下坠": 92.9, "梦想一心能量恢复": 2.1}),
				talent.LevelElementalBurst(7, map[string]float64{"梦想一刀": 601, "梦想一刀愿力加成": 5.83, "梦想一心愿力加成": 1.09, "积攒愿力层数": 0.18, "1段": 66.1, "2段": 64.9, "3段": 79.5, "4段": 45.6 + 45.8, "5段": 109.2, "重击": 91.0 + 109.9, "下坠": 101.1, "梦想一心能量恢复": 2.2}),
				talent.LevelElementalBurst(8, map[string]float64{"梦想一刀": 641, "梦想一刀愿力加成": 6.22, "梦想一心愿力加成": 1.16, "积攒愿力层数": 0.19, "1段": 70.7, "2段": 69.4, "3段": 85.0, "4段": 48.8 + 48.9, "5段": 115.8, "重击": 97.3 + 117.5, "下坠": 109.3, "梦想一心能量恢复": 2.3}),
				talent.LevelElementalBurst(9, map[string]float64{"梦想一刀": 681, "梦想一刀愿力加成": 6.61, "梦想一心愿力加成": 1.23, "积攒愿力层数": 0.19, "1段": 75.2, "2段": 73.9, "3段": 90.5, "4段": 51.9 + 52.1, "5段": 124.4, "重击": 103.6 + 125.1, "下坠": 117.5, "梦想一心能量恢复": 2.4}),
				talent.LevelElementalBurst(10, map[string]float64{"梦想一刀": 721, "梦想一刀愿力加成": 7.00, "梦想一心愿力加成": 1.31, "积攒愿力层数": 0.20, "1段": 79.8, "2段": 78.4, "3段": 96.0, "4段": 55.1 + 55.3, "5段": 131.9, "重击": 109.9 + 132.7, "下坠": 126.4, "梦想一心能量恢复": 2.5}),
				talent.LevelElementalBurst(11, map[string]float64{"梦想一刀": 762, "梦想一刀愿力加成": 7.39, "梦想一心愿力加成": 1.38, "积攒愿力层数": 0.20, "1段": 84.4, "2段": 82.9, "3段": 101.5, "4段": 58.3 + 58.4, "5段": 139.5, "重击": 116.2 + 140.3, "下坠": 135.3, "梦想一心能量恢复": 2.5}),
				talent.LevelElementalBurst(12, map[string]float64{"梦想一刀": 802, "梦想一刀愿力加成": 7.78, "梦想一心愿力加成": 1.45, "积攒愿力层数": 0.20, "1段": 89.0, "2段": 87.4, "3段": 107.0, "4段": 61.4 + 61.6, "5段": 147.1, "重击": 122.5 + 147.9, "下坠": 144.2, "梦想一心能量恢复": 2.5}),
				talent.LevelElementalBurst(13, map[string]float64{"梦想一刀": 852, "梦想一刀愿力加成": 8.26, "梦想一心愿力加成": 1.54, "积攒愿力层数": 0.20, "1段": 93.5, "2段": 91.9, "3段": 112.5, "4段": 64.6 + 64.8, "5段": 154.6, "重击": 128.8 + 155.5, "下坠": 153.1, "梦想一心能量恢复": 2.5}),
				talent.LevelElementalBurst(14, map[string]float64{"梦想一刀": 902, "梦想一刀愿力加成": 8.75, "梦想一心愿力加成": 1.63, "积攒愿力层数": 0.20, "1段": 98.1, "2段": 96.4, "3段": 118.0, "4段": 67.7 + 67.9, "5段": 162.2, "重击": 135.1 + 163.1, "下坠": 162.1, "梦想一心能量恢复": 2.5}),
			).Check()),
	).Talents(normal, skill, burst)
}
