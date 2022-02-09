package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"github.com/dbstarll/game/internal/ro/model/general"
	"github.com/dbstarll/game/internal/ro/romel"
	"go.uber.org/zap"
	"log"
)

func main() {
	//Template()
	//Shooter()
	Hunter()
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
	if player, err := model.LoadPlayerFromYaml("猫爸-暴君", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else if monster, err := model.LoadMonsterFromYaml("鳄鱼人"); err != nil {
		//} else if monster, err := model.LoadMonsterFromYaml("木桩"); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		attack := player.AttackWithWeapon(weapon.Rifle).WithNature(nature.Wind)

		player.Apply(
			buff.Manor(),
			//buff.HuntingGround(),
			model.AddGains(false, &model.Gains{Spike: 30, Ignore: 18}), //大君之怒

			//双面硬币
			//model.AddQuality(&model.Quality{Luk: 30, Agi: 42}),
			//model.AddGains(true, &model.Gains{AttackPer: 32.2}),
			//model.AddGeneral(&general.General{MoveSpeed: 20}),

			buff.CardAdmiral(),
			//model.AddGains(false, &model.Gains{Ignore: 0.8, Refine: 4}),
		)

		monster.Apply(
			model.AddGains(false, &model.Gains{Resist: 30, DefencePer: 100}),
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

		buff.ProfitDetect(player, true, func(player *model.Player) float64 {
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
	if player, err := model.LoadPlayerFromYaml("璐璐-群星猎手", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else if monster, err := model.LoadMonsterFromYaml("鳄鱼人"); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		attack := player.AttackWithWeapon(weapon.Bow).WithNature(nature.Wind).WithCritical()

		player.Apply(
			buff.Manor(),
			buff.DexA(),
			model.Merge( //狙杀瞄准
				buff.Quality(5),
				model.AddGeneral(&general.General{CriticalDamage: 7.8}),
				model.AddGains(false, &model.Gains{AttackPer: 10}),
			),
			model.Merge( //蓄势待发
				model.AddQuality(&model.Quality{Int: 40, Dex: 40}),
				model.AddGeneral(&general.General{CriticalDamage: 100, OrdinaryDamage: 12.2}),
			),
			model.AddGains(false, &model.Gains{Spike: 20}),   //无限星辰
			model.AddGains(false, &model.Gains{Attack: 240}), //勿忘初心.鱼排
		)
		monster.Apply(
			model.AddGains(false, &model.Gains{Resist: 30, DefencePer: 30}),
		)

		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		profitDetect(player, func(player *model.Player) float64 {
			return player.FinalDamage(monster, attack)
		})
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
		//buff.Manor(),
		//model.AddGains(false, &model.Gains{Defence: 6}),
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

		monster2.Apply(
			model.AddGains(false, &model.Gains{Resist: 30, DefencePer: 100}),
			//model.AddRaceResist(&map[race.Race]float64{race.Human: 15}),
			//model.AddGeneral(&general.General{CriticalDamageResist: 22}),
		)

		profitDetect(player, func(player *model.Player) float64 {
			return player.SkillEarth() * player.SkillDamageRate(monster2, false, nature.Earth)
		})
	}
}

func profitDetect(player *model.Player, fn buff.FinalDamage) {
	for _, pos := range []position.Position{position.Weapon, position.Shield, position.Armor, position.Cloak, position.Shoes, position.Ring} {
		fmt.Printf("装备: %s\n", pos)
		for idx, p := range buff.ProfitDetect(player, false, fn, Equips(pos, player.Job())) {
			if idx < 10 {
				fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
			}
		}
	}
	for _, pos := range []position.Position{position.Head, position.Face, position.Mouth, position.Back, position.Tail} {
		fmt.Printf("头饰: %s\n", pos)
		for idx, p := range buff.ProfitDetect(player, false, fn, Hats(pos)) {
			if idx < 10 {
				fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
			}
		}
	}
	for _, pos := range []position.Position{position.Weapon, position.Shield, position.Armor, position.Cloak, position.Shoes, position.Ring, position.Head} {
		fmt.Printf("卡片: %s\n", pos)
		for idx, p := range buff.ProfitDetect(player, false, fn, Cards(pos)) {
			if idx < 10 {
				fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
			}
		}
	}
	fmt.Printf("素质:\n")
	for _, p := range buff.ProfitDetect(player, true, fn, nil) {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
}

func Equips(pos position.Position, _job job.Job) map[string]model.CharacterModifier {
	modifiers := make(map[string]model.CharacterModifier)
	if _, err := romel.Equips.Filter(func(equip *romel.Equip) error {
		var ms []model.CharacterModifier
		if m := equip.Effect.Effect(); len(m) > 0 {
			ms = append(ms, m...)
		}
		if m := equip.Buff.Effect(); len(m) > 0 {
			ms = append(ms, m...)
		}
		if equip.RandomBuff.Empty() {
			if len(ms) > 0 {
				modifiers[equip.Name] = model.Merge(ms...)
			}
		} else if m := equip.RandomBuff.Effect(); len(m) == equip.RandomBuff.Size() {
			for idx, item := range equip.RandomBuff.Items() {
				modifiers[equip.Name+":"+item] = model.Merge(model.Merge(ms...), m[idx])
			}
		}
		return nil
	}, func(filter *romel.Equip) {
		filter.Position = pos
		filter.Job = &[]job.Job{_job}
	}); err != nil {
		zap.S().Errorf("%+v", err)
	}
	return modifiers
}

func Cards(pos position.Position) map[string]model.CharacterModifier {
	modifiers := make(map[string]model.CharacterModifier)
	if _, err := romel.Cards.Filter(func(card *romel.Card) error {
		if m := card.Buff.Effect(); len(m) > 0 {
			modifiers[card.Name] = model.Merge(m...)
		}
		return nil
	}, func(filter *romel.Card) {
		filter.Position = pos
	}); err != nil {
		zap.S().Errorf("%+v", err)
	}
	return modifiers
}

func Hats(pos position.Position) map[string]model.CharacterModifier {
	modifiers := make(map[string]model.CharacterModifier)
	if _, err := romel.Hats.Filter(func(card *romel.Hat) error {
		if m := card.Buff.Effect(); len(m) > 0 {
			modifiers[card.Name] = model.Merge(m...)
		}
		return nil
	}, func(filter *romel.Hat) {
		filter.Position = pos
	}); err != nil {
		zap.S().Errorf("%+v", err)
	}
	return modifiers
}
