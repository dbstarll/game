package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var (
	Factory胡桃 = func(normal, skill, burst, constellation int) *Character {
		return New(5, elementals.Fire, weaponType.Polearm,
			Base(90, 15552, 107, 876, buff.AddCriticalDamage(38.4)),
			TalentsTemplateModifier(talent.NewTalentsTemplate(
				talent.BaseNormalAttack("往生秘传枪法", 11, 25),
				talent.BaseElementalSkill("蝶引来生", 14, time.Second*16, time.Second*9),
				talent.BaseElementalBurst("安神秘法", 14, 60, time.Second*15, 0)).
				AddNormalAttacks(
					talent.LevelNormalAttack(1, []float64{46.9, 48.3, 61.1, 65.6, 33.3 + 35.2, 86.0}, 65.4, 136.0),
					talent.LevelNormalAttack(2, []float64{50.1, 51.5, 65.2, 70.1, 35.5 + 37.6, 91.8}, 69.9, 145.2),
					talent.LevelNormalAttack(3, []float64{53.3, 54.8, 69.4, 74.6, 37.8 + 40.0, 97.7}, 74.3, 154.5),
					talent.LevelNormalAttack(4, []float64{57.5, 59.2, 74.9, 80.6, 40.8 + 43.2, 105.5}, 80.3, 166.9),
					talent.LevelNormalAttack(5, []float64{60.7, 62.5, 79.1, 85.0, 43.1 + 45.6, 111.4}, 84.7, 176.1),
					talent.LevelNormalAttack(6, []float64{65.5, 66.3, 83.9, 90.3, 45.8 + 48.4, 118.2}, 90.0, 186.9),
					talent.LevelNormalAttack(7, []float64{69.3, 71.3, 90.2, 97.0, 49.2 + 52.0, 127.0}, 96.6, 200.9),
					talent.LevelNormalAttack(8, []float64{74.1, 76.2, 96.4, 103.7, 52.6 + 55.6, 135.8}, 103.3, 214.8),
					talent.LevelNormalAttack(9, []float64{78.9, 81.2, 102.7, 110.4, 56.0 + 59.2, 114.6}, 110.0, 228.7),
					talent.LevelNormalAttack(10, []float64{83.6, 86.1, 108.9, 117.1, 59.4 + 62.8, 153.4}, 116.7, 242.6),
					talent.LevelNormalAttack(11, []float64{88.4, 91.0, 115.2, 123.8, 62.8 + 66.4, 162.1}, 123.4, 256.5),
				).
				AddElementalSkills(
					talent.LevelElementalSkill(1, map[string]float64{"攻击力提高": 3.84, "血梅香": 64}),
					talent.LevelElementalSkill(2, map[string]float64{"攻击力提高": 4.07, "血梅香": 69}),
					talent.LevelElementalSkill(3, map[string]float64{"攻击力提高": 4.30, "血梅香": 74}),
					talent.LevelElementalSkill(4, map[string]float64{"攻击力提高": 4.60, "血梅香": 80}),
					talent.LevelElementalSkill(5, map[string]float64{"攻击力提高": 4.83, "血梅香": 85}),
					talent.LevelElementalSkill(6, map[string]float64{"攻击力提高": 5.06, "血梅香": 90}),
					talent.LevelElementalSkill(7, map[string]float64{"攻击力提高": 5.36, "血梅香": 96}),
					talent.LevelElementalSkill(8, map[string]float64{"攻击力提高": 5.66, "血梅香": 102}),
					talent.LevelElementalSkill(9, map[string]float64{"攻击力提高": 5.96, "血梅香": 109}),
					talent.LevelElementalSkill(10, map[string]float64{"攻击力提高": 6.26, "血梅香": 115}),
					talent.LevelElementalSkill(11, map[string]float64{"攻击力提高": 6.56, "血梅香": 122}),
					talent.LevelElementalSkill(12, map[string]float64{"攻击力提高": 6.85, "血梅香": 128}),
					talent.LevelElementalSkill(13, map[string]float64{"攻击力提高": 7.15, "血梅香": 136}),
					talent.LevelElementalSkill(14, map[string]float64{"攻击力提高": 7.45, "血梅香": 144}),
				).
				AddElementalBursts(
					talent.LevelElementalBurst(1, map[string]float64{"伤害": 303, "低血量伤害": 379, "治疗量": 6.26, "低血量治疗量": 8.35}),
					talent.LevelElementalBurst(2, map[string]float64{"伤害": 321, "低血量伤害": 402, "治疗量": 6.64, "低血量治疗量": 8.85}),
					talent.LevelElementalBurst(3, map[string]float64{"伤害": 340, "低血量伤害": 424, "治疗量": 7.01, "低血量治疗量": 9.35}),
					talent.LevelElementalBurst(4, map[string]float64{"伤害": 363, "低血量伤害": 454, "治疗量": 7.50, "低血量治疗量": 10.0}),
					talent.LevelElementalBurst(5, map[string]float64{"伤害": 381, "低血量伤害": 477, "治疗量": 7.88, "低血量治疗量": 10.5}),
					talent.LevelElementalBurst(6, map[string]float64{"伤害": 400, "低血量伤害": 499, "治疗量": 8.25, "低血量治疗量": 11.0}),
					talent.LevelElementalBurst(7, map[string]float64{"伤害": 423, "低血量伤害": 529, "治疗量": 8.74, "低血量治疗量": 11.65}),
					talent.LevelElementalBurst(8, map[string]float64{"伤害": 447, "低血量伤害": 558, "治疗量": 9.23, "低血量治疗量": 12.3}),
					talent.LevelElementalBurst(9, map[string]float64{"伤害": 470, "低血量伤害": 588, "治疗量": 9.71, "低血量治疗量": 12.95}),
					talent.LevelElementalBurst(10, map[string]float64{"伤害": 494, "低血量伤害": 617, "治疗量": 10.2, "低血量治疗量": 13.6}),
					talent.LevelElementalBurst(11, map[string]float64{"伤害": 518, "低血量伤害": 647, "治疗量": 10.69, "低血量治疗量": 14.25}),
					talent.LevelElementalBurst(12, map[string]float64{"伤害": 541, "低血量伤害": 676, "治疗量": 11.18, "低血量治疗量": 14.9}),
					talent.LevelElementalBurst(13, map[string]float64{"伤害": 565, "低血量伤害": 706, "治疗量": 11.66, "低血量治疗量": 15.55}),
					talent.LevelElementalBurst(14, map[string]float64{"伤害": 588, "低血量伤害": 735, "治疗量": 12.15, "低血量治疗量": 16.2}),
				).Check()),
		).Talents(normal, skill, burst)
	}
)
