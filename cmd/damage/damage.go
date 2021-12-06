package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
	"github.com/dbstarll/game/internal/ro/model"
	"log"
	"sort"
)

type FinalDamage func(player *model.Player) float64

var (
	detect = map[string]model.CharacterModifier{
		"Str+1": model.AddQuality(&model.Quality{Str: 1}),
		"Agi+1": model.AddQuality(&model.Quality{Agi: 1}),
		"Vit+1": model.AddQuality(&model.Quality{Vit: 1}),
		"Int+1": model.AddQuality(&model.Quality{Int: 1}),
		"Dex+1": model.AddQuality(&model.Quality{Dex: 1}),
		"Luk+1": model.AddQuality(&model.Quality{Luk: 1}),

		"暴击+1":      model.AddGeneral(&model.General{Critical: 1}),
		"暴伤%+1":     model.AddGeneral(&model.General{CriticalDamage: 1}),
		"暴击防护+1":    model.AddGeneral(&model.General{CriticalResist: 1}),
		"爆伤减免%+1":   model.AddGeneral(&model.General{CriticalDamageResist: 1}),
		"普攻攻击力+20":  model.AddGeneral(&model.General{Ordinary: 20}),
		"普攻伤害加成%+1": model.AddGeneral(&model.General{OrdinaryDamage: 1}),
		"普攻伤害减免%+1": model.AddGeneral(&model.General{OrdinaryResist: 1}),
		"技能伤害加成%+1": model.AddGeneral(&model.General{Skill: 1}),
		"技能伤害减免%+1": model.AddGeneral(&model.General{SkillResist: 1}),
		"MVP增伤%+1":  model.AddGeneral(&model.General{MVP: 1}),
		"MVP减伤%+1":  model.AddGeneral(&model.General{MVPResist: 1}),

		"物理攻击+20":     model.AddGains(false, &model.Gains{Attack: 20}),
		"物理防御+20":     model.AddGains(false, &model.Gains{Defence: 20}),
		"物理穿刺%+1":     model.AddGains(false, &model.Gains{Spike: 1}),
		"物理攻击%+1":     model.AddGains(false, &model.Gains{AttackPer: 1}),
		"物理防御%+1":     model.AddGains(false, &model.Gains{DefencePer: 1}),
		"物伤加成%+1":     model.AddGains(false, &model.Gains{Damage: 1}),
		"忽视物防%+1":     model.AddGains(false, &model.Gains{Ignore: 1}),
		"物伤减免%+1":     model.AddGains(false, &model.Gains{Resist: 1}),
		"近战物理伤害减免%+1": model.AddGains(false, &model.Gains{NearResist: 1}),
		"远程物理伤害减免%+1": model.AddGains(false, &model.Gains{RemoteResist: 1}),
		"精炼物攻+20":     model.AddGains(false, &model.Gains{Refine: 20}),
		"精炼物免%+1":     model.AddGains(false, &model.Gains{RefineResist: 1}),

		"魔法攻击+20": model.AddGains(true, &model.Gains{Attack: 20}),
		"魔法防御+20": model.AddGains(true, &model.Gains{Defence: 20}),
		"魔法穿刺%+1": model.AddGains(true, &model.Gains{Spike: 1}),
		"魔法攻击%+1": model.AddGains(true, &model.Gains{AttackPer: 1}),
		"魔法防御%+1": model.AddGains(true, &model.Gains{DefencePer: 1}),
		"魔伤加成%+1": model.AddGains(true, &model.Gains{Damage: 1}),
		"忽视魔防%+1": model.AddGains(true, &model.Gains{Ignore: 1}),
		"魔伤减免%+1": model.AddGains(true, &model.Gains{Resist: 1}),
		"精炼魔攻+20": model.AddGains(true, &model.Gains{Refine: 20}),
		"精炼魔免%+1": model.AddGains(true, &model.Gains{RefineResist: 1}),

		"无属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Neutral: 1}),
		"地属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Earth: 1}),
		"风属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Wind: 1}),
		"水属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Water: 1}),
		"火属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Fire: 1}),
		"圣属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Holy: 1}),
		"暗属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Dark: 1}),
		"不死属性攻击%+1": model.AddNatureAttack(&map[nature.Nature]float64{nature.Undead: 1}),
		"念属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Ghost: 1}),
		"毒属性攻击%+1":  model.AddNatureAttack(&map[nature.Nature]float64{nature.Poison: 1}),

		"对无属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Neutral: 1}),
		"对地属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Earth: 1}),
		"对风属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Wind: 1}),
		"对水属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Water: 1}),
		"对火属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Fire: 1}),
		"对圣属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Holy: 1}),
		"对暗属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Dark: 1}),
		"对不死属性魔物增伤%+1": model.AddNatureDamage(&map[nature.Nature]float64{nature.Undead: 1}),
		"对念属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Ghost: 1}),
		"对毒属性魔物增伤%+1":  model.AddNatureDamage(&map[nature.Nature]float64{nature.Poison: 1}),

		"对无属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Neutral: 1}),
		"对地属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Earth: 1}),
		"对风属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Wind: 1}),
		"对水属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Water: 1}),
		"对火属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Fire: 1}),
		"对圣属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Holy: 1}),
		"对暗属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Dark: 1}),
		"对不死属性伤害减免%+1": model.AddNatureResist(&map[nature.Nature]float64{nature.Undead: 1}),
		"对念属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Ghost: 1}),
		"对毒属性伤害减免%+1":  model.AddNatureResist(&map[nature.Nature]float64{nature.Poison: 1}),

		"对动物增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Animal: 1}),
		"对人形增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Human: 1}),
		"对恶魔增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Demon: 1}),
		"对植物增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Plant: 1}),
		"对不死增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Undead: 1}),
		"对无形增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Formless: 1}),
		"对鱼贝增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Fish: 1}),
		"对天使增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Angel: 1}),
		"对昆虫增伤%+1": model.AddRaceDamage(&map[race.Race]float64{race.Insect: 1}),
		"对龙增伤%+1":  model.AddRaceDamage(&map[race.Race]float64{race.Dragon: 1}),

		"对动物减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Animal: 1}),
		"对人形减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Human: 1}),
		"对恶魔减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Demon: 1}),
		"对植物减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Plant: 1}),
		"对不死减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Undead: 1}),
		"对无形减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Formless: 1}),
		"对鱼贝减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Fish: 1}),
		"对天使减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Angel: 1}),
		"对昆虫减伤%+1": model.AddRaceResist(&map[race.Race]float64{race.Insect: 1}),
		"对龙减伤%+1":  model.AddRaceResist(&map[race.Race]float64{race.Dragon: 1}),

		"对小体型增伤%+1": model.AddShapeDamage(&map[shape.Shape]float64{shape.Small: 1}),
		"对中体型增伤%+1": model.AddShapeDamage(&map[shape.Shape]float64{shape.Medium: 1}),
		"对大体型增伤%+1": model.AddShapeDamage(&map[shape.Shape]float64{shape.Large: 1}),

		"对小体型减伤%+1": model.AddShapeResist(&map[shape.Shape]float64{shape.Small: 1}),
		"对中体型减伤%+1": model.AddShapeResist(&map[shape.Shape]float64{shape.Medium: 1}),
		"对大体型减伤%+1": model.AddShapeResist(&map[shape.Shape]float64{shape.Large: 1}),
	}
)

func main() {
	//Template()
	Hunter()
	//EarthBash()
}

func Template() {
	if player, err := model.LoadPlayerFromYaml("模版", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", player.Character)
	}
}

func Hunter() {
	if player, err := model.LoadPlayerFromYaml("小弓-游侠", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewMonster(types.Ordinary, race.Plant, nature.Neutral, shape.Medium)
		attack := player.AttackWithWeapon(weapon.Bow)

		model.Merge(model.AddQuality(&model.Quality{Str: 5, Agi: 5 - 30, Vit: 5, Int: 5 + 40, Dex: 5 + 40, Luk: 5}),
			model.AddGeneral(&model.General{Critical: 10 + 30, CriticalDamage: 100}),
			model.AddGains(false, &model.Gains{AttackPer: 10}))(player.Character)

		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f\n", player.FinalDamage(monster, attack))
		//22058
		ProfitDetect(player, func(player *model.Player) float64 {
			return player.FinalDamage(monster, attack)
		})
	}
}

func EarthBash() {
	if player, err := model.LoadPlayerFromYaml("猫爸-圣盾", false); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewMonster(types.MVP, race.Human, nature.Wind, shape.Large)
		skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)

		ProfitDetect(player, func(player *model.Player) float64 {
			return player.SkillEarth() * player.SkillDamageRate(monster, false, nature.Earth)
		})
	}
}

func ProfitDetect(player *model.Player, fn FinalDamage) error {
	base := fn(player)
	fmt.Printf("base: %f\n", base)
	var profits []*Profit
	for name, modifier := range detect {
		cancel := modifier(player.Character)
		value := fn(player)
		if value != base {
			profits = append(profits, &Profit{
				name:  name,
				value: value,
			})
		}
		cancel()
	}
	sort.Slice(profits, func(i, j int) bool {
		if profits[i].value < profits[j].value {
			return false
		} else if profits[i].value > profits[j].value {
			return true
		} else {
			return profits[i].name < profits[j].name
		}
	})
	for _, profit := range profits {
		fmt.Printf("增幅：%2.4f%% - %s\n", 100*(profit.value-base)/base, profit.name)
	}
	return nil
}

type Profit struct {
	name  string
	value float64
}
