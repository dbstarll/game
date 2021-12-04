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
		model.AddQuality(&model.Quality{Str: 103, Agi: 131, Vit: 402, Int: 196, Dex: 78, Luk: 42}),
		model.AddGains(false, &model.Gains{AttackPer: 31, Spike: 37 - 30, Damage: 70.5, DefencePer: 112.5}),
		model.AddGains(true, &model.Gains{AttackPer: 31, Spike: 5, DefencePer: 60}),
		model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Wind: 21, nature.Earth: 20, nature.Fire: 21, nature.Water: 1, nature.Neutral: 2, nature.Holy: 28,
		}),
		model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Wind: 8, nature.Earth: 6, nature.Water: 10, nature.Fire: 7, nature.Holy: 6, nature.Dark: 2, nature.Poison: 1,
		}),
		model.AddRaceDamage(&map[race.Race]float64{
			race.Animal: 22, race.Human: 26.7, race.Demon: 15, race.Plant: 3, race.Undead: 22,
			race.Formless: 4, race.Fish: 9, race.Angel: 3, race.Insect: 16, race.Dragon: 5,
		}),
		model.AddShapeDamage(&map[shape.Shape]float64{
			shape.Small:  65,
			shape.Medium: 6,
			shape.Large:  55,
		}), model.AddDamage(&model.Damage{
			MVP: 28,
		}), model.DetectDefenceByPanel(7766, 3308))

	monster := model.NewMonster(types.MVP, race.Demon, nature.Dark, shape.Large)
	skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)*0.7
	fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
	//5631426
}
