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
	"github.com/dbstarll/game/internal/ro/model/general"
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
	} else if monster, err := model.LoadMonsterFromYaml("鳄鱼人"); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		attack := player.AttackWithWeapon(weapon.Rifle).WithNature(nature.Wind)

		player.Apply(
			buff.Manor(),
			model.AddGains(false, &model.Gains{Attack: 9, Spike: 27 - 6 + 1.2}), //大君之怒
			buff.CardAdmiral(),
		)

		monster.Apply(
		//model.AddGains(false, &model.Gains{Resist: 30}),
		//model.AddRaceResist(&map[race.Race]float64{race.Human: 10}),
		//model.AddGeneral(&general.General{CriticalDamageResist: 0, OrdinaryResist: 10}),
		)

		//0.4/1409981/1786425/2195794/
		//0.8/1414322/1791883/2201252/4340/5458/0.7951
		//1.2/1418662/1797341/2206710/4340/5458/0.7951

		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		//buff.ProfitDetect(player, func(player *model.Player) float64 {
		//	return player.FinalDamage(monster, attack)
		//}, nil)

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

		model.AddGeneral(&general.General{Critical: 30, CriticalDamage: 100})(player.Character)
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
		monster := model.NewCharacter(types.Ordinary, race.Animal, nature.Water, shape.Medium)

		player.Apply(
			buff.Manor(),
			model.AddGains(false, &model.Gains{Defence: 6}),
		)

		monster.Apply(
			model.AddGains(false, &model.Gains{Resist: 30}),
			model.AddRaceResist(&map[race.Race]float64{race.Human: 10}),
		)

		skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)

		//39.7/6577060/9306344
		//34.1/6241312/8933291/8039962/
		//34.5/6265294/8959937/8063944/23982/26646
		//34.9/6289276/8986584/8087926/23982/26647

		//buff.ProfitDetect(player, func(player *model.Player) float64 {
		//	return player.SkillEarth() * player.SkillDamageRate(monster, false, nature.Earth)
		//}, map[string]model.CharacterModifier{
		//	"物理防御%+15": model.AddGains(false, &model.Gains{DefencePer: 15}),
		//	"物理防御+240": model.AddGains(false, &model.Gains{Defence: 240}),
		//	"Vit+12":   model.AddQuality(&model.Quality{Vit: 12}),
		//	"体质料理B":    buff.VitB(),
		//	"力量料理B":    buff.StrB(),
		//})
	}
}
