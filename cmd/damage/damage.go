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
func main() {
	Hunter()
}

func Hunter() {
	player := model.NewPlayer(job.Hunter3,
		model.AddQuality(&model.Quality{Str: 15, Agi: 159, Vit: 34, Int: 42, Dex: 264, Luk: 73}),
		model.AddGains(false, &model.Gains{AttackPer: 20, DefencePer: 5, Refine: 58}),
		model.AddGains(true, &model.Gains{AttackPer: 19, DefencePer: 5}),
		model.AddNatureAttack(&map[nature.Nature]float64{nature.Fire: 1, nature.Water: 1, nature.Holy: 1}),
		model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Wind: 6, nature.Earth: 7, nature.Water: 8, nature.Fire: 7, nature.Holy: 6, nature.Ghost: 2, nature.Poison: 1,
		}),
		model.AddRaceDamage(&map[race.Race]float64{
			race.Animal: 4, race.Human: 1, race.Demon: 2, race.Undead: 6, race.Fish: 5, race.Insect: 13, race.Dragon: 2,
		}), model.DetectDefenceByPanel(494, 396),
		model.DetectAttackByPanel(true, 3295, 1117))

	monster := model.NewMonster(types.Ordinary, race.Plant, nature.Neutral, shape.Medium)
	generalAttack := player.GeneralAttack(monster, false, true, nature.Neutral)
	fmt.Printf("%f\n", generalAttack)
}

func EarthBash() {
	player := model.NewPlayer(job.Crusader4,
		model.AddQuality(&model.Quality{Str: 103, Agi: 131, Vit: 402, Int: 196, Dex: 78, Luk: 42}),
		model.AddGains(false, &model.Gains{AttackPer: 31, Spike: 37, Damage: 70.5, DefencePer: 112.5}),
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

	monster := model.NewMonster(types.MVP, race.Demon, nature.Dark, shape.Large,
		model.AddGains(false, &model.Gains{Resist: 30}),
		model.AddRaceResist(&map[race.Race]float64{race.Human: 30}))
	skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
	fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
}
