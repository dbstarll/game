package character

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/talent"
	"time"
)

var (
	Factory迪卢克 = func(normal, skill, burst, constellation int) *Character {
		return New(5, elementals.Fire, weaponType.Claymore,
			Base(90, 12981, 335, 784, buff.AddCriticalRate(19.2)),
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
