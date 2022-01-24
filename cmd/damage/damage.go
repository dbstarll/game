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

func ff(b, d float64) float64 {
	return 51061*(4000+3027*d)/(4000+3027*d*10)*b + 1071
}

func Template() {
	if player, err := model.LoadPlayerFromYaml("模版", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", player.Character)
	}
}

func Shooter() {
	if player, err := model.LoadPlayerFromYaml("璐璐-暴君", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else if monster, err := model.LoadMonsterFromYaml("鳄鱼人"); err != nil {
		//} else if monster, err := model.LoadMonsterFromYaml("木桩"); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		attack := player.AttackWithWeapon(weapon.Rifle).WithNature(nature.Wind)

		player.Apply(
			buff.Manor(),
			//buff.HuntingGround(),
			//model.AddGains(false, &model.Gains{Spike: 30, Ignore: 18}), //大君之怒

			//双面硬币
			//model.AddQuality(&model.Quality{Luk: 30, Agi: 42}),
			//model.AddGains(true, &model.Gains{AttackPer: 32.2}),
			//model.AddGeneral(&general.General{MoveSpeed: 20}),

			buff.CardAdmiral(),
			//model.AddGains(false, &model.Gains{Ignore: 0.8, Refine: 4}),
		)

		monster.Apply(
		//model.AddGains(false, &model.Gains{Resist: 30, DefencePer: 40}),
		//model.AddRaceResist(&map[race.Race]float64{race.Human: 15}),
		//model.AddGeneral(&general.General{CriticalDamageResist: 22}),
		)
		//0.8/4/607208/506597/703985/628183
		//1.6/8/607246/509071/704022/631252
		//2.4/12/607283/511567/704060/634355
		//3.2/16/607321/514090/704098/637494
		//4.0/20/607359/516636/704135/640668

		//魔术子弹：魔法攻击+31%，魔法攻击50%的有视物理防御的攻击
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		//attack.WithSkill(1.74 * 5)
		//fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		buff.ProfitDetect(player, func(player *model.Player) float64 {
			return player.FinalDamage(monster, attack)
		}, map[string]model.CharacterModifier{
			"战役斗篷": model.Merge(
				model.AddQuality(&model.Quality{Str: 5, Dex: 5}),
				model.AddGeneral(&general.General{Critical: 5, Ordinary: 240, OrdinaryDamage: 3}),
				model.AddGains(false, &model.Gains{Defence: 100}),
			),
			"伯爵斗篷": model.Merge(
				model.AddGains(false, &model.Gains{Defence: 31, Ignore: 15}),
			),
			"勇士肩甲": model.Merge(
				model.AddQuality(&model.Quality{Str: 10}),
				model.AddGains(false, &model.Gains{Defence: 100, Ignore: 18}),
			),
			"王室骑士披风": model.Merge(
				model.AddQuality(&model.Quality{Luk: 20}),
				model.AddGeneral(&general.General{Critical: 5, CriticalDamage: 15 + 7.5}),
				model.AddGains(false, &model.Gains{Defence: 100}),
			),
			"远航者战靴": model.Merge(
				model.AddQuality(&model.Quality{Dex: 12}),
				model.AddGeneral(&general.General{MoveSpeed: 12}),
				model.AddGains(false, &model.Gains{Defence: 120, RemoteDamage: 4}),
			),
			"平衡之理靴子": model.Merge(
				model.AddQuality(&model.Quality{Str: 6, Int: 6}),
				model.AddGeneral(&general.General{MoveSpeed: 12}),
				model.AddGains(false, &model.Gains{Defence: 120, AttackPer: 6}),
				model.AddGains(true, &model.Gains{AttackPer: 6}),
			),
			"统治者战靴": model.Merge(
				model.AddQuality(&model.Quality{Str: 12}),
				model.AddGeneral(&general.General{MoveSpeed: 12}),
				model.AddGains(false, &model.Gains{Defence: 120, NearDamage: 4}),
			),
			"斩龙者战靴": model.Merge(
				model.AddQuality(&model.Quality{Str: 6, Dex: 6}),
				model.AddGeneral(&general.General{MoveSpeed: 12}),
				model.AddGains(false, &model.Gains{Defence: 120, AttackPer: 8}),
			),
			"轻灵之鞋": model.Merge(
				model.AddQuality(&model.Quality{Agi: 12}),
				model.AddGeneral(&general.General{MoveSpeed: 12, CriticalDamage: 10}),
				model.AddGains(false, &model.Gains{Defence: 120}),
			),
			"虚无之晶": model.Merge(
				model.AddQuality(&model.Quality{Str: 8, Dex: 8}),
				model.AddGains(false, &model.Gains{Attack: 224, Ignore: 6}),
			),
			"远洋银币": model.Merge(
				model.AddQuality(&model.Quality{Str: 12}),
				model.AddGeneral(&general.General{Ordinary: 120}),
				model.AddGains(false, &model.Gains{Attack: 224}),
			),
			"镶金竖琴": model.Merge(
				model.AddQuality(&model.Quality{Dex: 12}),
				model.AddGains(false, &model.Gains{Attack: 224, RemoteDamage: 6}),
			),
			"乌金之坠": model.Merge(
				model.AddQuality(&model.Quality{Str: 12}),
				model.AddGains(false, &model.Gains{Attack: 224, NearDamage: 6}),
			),
			"灼炎之精": model.Merge(
				model.AddQuality(&model.Quality{Str: 8, Int: 8}),
				model.AddGains(true, &model.Gains{Attack: 200}),
				model.AddNatureAttack(&map[nature.Nature]float64{nature.Fire: 8}),
			),
			"热爱胸针": model.Merge(
				model.AddQuality(&model.Quality{Luk: 16}),
				model.AddGeneral(&general.General{CriticalDamage: 8}),
				model.AddGains(false, &model.Gains{Attack: 224}),
			),
			"黄金耳环": model.Merge(
				model.AddGeneral(&general.General{Critical: 10 + 5}),
				model.AddGains(false, &model.Gains{Attack: 224 + 60 + 50}),
			),
		})

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
	} else if monster1, err := model.LoadMonsterFromYaml("鳄鱼人"); err != nil {
		log.Fatalf("%+v\n", err)
	} else if monster2, err := model.LoadMonsterFromYaml("奈吉鸟"); err != nil {
		log.Fatalf("%+v\n", err)
	} else if monster3, err := model.LoadMonsterFromYaml("月夜蝙蝠"); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		player.Apply(
			buff.Manor(),
			model.AddGains(false, &model.Gains{Defence: 6}),
		)

		skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster1, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
		skillEarth, rate = player.SkillEarth(), player.SkillDamageRate(monster2, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
		skillEarth, rate = player.SkillEarth(), player.SkillDamageRate(monster3, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)

		//8021913/9222201/8021913
		//15968862/19529367
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
