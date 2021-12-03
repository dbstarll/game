package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/model"
)

//最终伤害 = 基础伤害 * 元素加伤 * 状态加伤 *(1+真实伤害)
//
//*面板物理攻击 = 物理攻击 * (1+物理攻击%)
//
//魔法攻击 = 素质魔法攻击 + 装备魔法攻击
//*面板魔法攻击 = 魔法攻击 * (1+魔法攻击%)
func main() {
	player := model.NewPlayer(job.Crusader4,
		model.AddQuality(&model.Quality{
			Vit: 402,
		}), model.AddGains(false, &model.Gains{
			Spike:      37,
			Damage:     70.5,
			DefencePer: 112.5,
		}), model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Earth: 20,
		}), model.AddRaceDamage(&map[race.Race]float64{
			race.Human: 26.7,
		}), model.AddShapeDamage(&map[shape.Shape]float64{
			shape.Small:  65,
			shape.Medium: 6,
			shape.Large:  55,
		}), model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Wind: 8,
		}), model.AddDamage(&model.Damage{
			Skill: 2,
			MVP:   28,
		}), model.DetectDefenceByPanel(7752, 3000))

	monster := model.NewMonster(types.MINI, nature.Wind, race.Human, shape.Small)
	skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
	fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
}
