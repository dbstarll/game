package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"log"
)

func main() {
	//Template()
	Shooter()
	//Hunter()
	//EarthBash()
}

func Template() {
	if player, err := model.LoadPlayerFromYaml("模版", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", player.Character)
	}
}
func Shooter() {
	if player, err := model.LoadPlayerFromYaml("猫爸-暴君", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewCharacter(types.Ordinary, race.Animal, nature.Water, shape.Medium)
		attack := player.AttackWithWeapon(weapon.Rifle).WithNature(nature.Wind)

		buff.Manor()(player.Character)
		buff.DexA()(player.Character)
		model.Merge(model.AddGains(false, &model.Gains{Resist: 30}))(monster)

		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		buff.ProfitDetect(player, func(player *model.Player) float64 {
			return player.FinalDamage(monster, attack)
		}, nil)

		//武器体型修正
		//技能伤害加成%
		//MVP增伤%
		//元素加伤
		//状态加伤
		//		体型减伤%
		//		属性减伤%
		//		种族减伤%
		//		物伤减免%
	}
}

func Hunter() {
	if player, err := model.LoadPlayerFromYaml("晴天有时下猪", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewCharacter(types.Ordinary, race.Animal, nature.Water, shape.Medium)
		attack := player.AttackWithWeapon(weapon.Bow).WithNature(nature.Wind)

		model.AddGeneral(&model.General{Critical: 30, CriticalDamage: 100})(player.Character)
		buff.Manor()(player.Character)
		buff.DexA()(player.Character)
		model.Merge(model.AddGains(false, &model.Gains{Resist: 30}))(monster)
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		buff.ProfitDetect(player, func(player *model.Player) float64 {
			return player.FinalDamage(monster, attack)
		}, nil)
	}
}

func EarthBash() {
	if player, err := model.LoadPlayerFromYaml("猫爸-圣盾", false); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewCharacter(types.MVP, race.Human, nature.Wind, shape.Large)
		skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)

		buff.ProfitDetect(player, func(player *model.Player) float64 {
			return player.SkillEarth() * player.SkillDamageRate(monster, false, nature.Earth)
		}, map[string]model.CharacterModifier{
			"物理防御%+15": model.AddGains(false, &model.Gains{DefencePer: 15}),
			"物理防御+240": model.AddGains(false, &model.Gains{Defence: 240}),
			"Vit+12":   model.AddQuality(&model.Quality{Vit: 12}),
			"体质料理B":    buff.VitB(),
			"力量料理B":    buff.StrB(),
		})
	}
}
