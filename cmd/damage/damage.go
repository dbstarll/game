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
)

func main() {
	Template()
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

func Hunter() {
	if player, err := model.LoadPlayerFromYaml("小弓-游侠", true); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewMonster(types.Ordinary, race.Plant, nature.Neutral, shape.Medium)
		attack := player.AttackWithWeapon(weapon.Bow)
		generalAttack := player.GeneralAttack(monster, attack)
		fmt.Printf("%f, %f\n", generalAttack, player.FinalDamage(monster, attack))
		attack.WithCritical()
		fmt.Printf("%f, %f\n", generalAttack, player.FinalDamage(monster, attack))
		//4055 	6610
	}
}

func EarthBash() {
	if player, err := model.LoadPlayerFromYaml("猫爸-圣盾", false); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		monster := model.NewMonster(types.Ordinary, race.Human, nature.Wind, shape.Large)
		skillEarth, rate := player.SkillEarth(), player.SkillDamageRate(monster, false, nature.Earth)
		fmt.Printf("%f * %f = %f\n", skillEarth, rate, rate*skillEarth)
	}
}
