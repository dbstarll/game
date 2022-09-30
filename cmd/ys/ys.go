package main

import (
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/element"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model"
	"log"
)

func main() {
	迪卢克 := model.NewCharacter(element.Fire, weaponType.BigSword,
		model.BaseCharacter(90, 12981, 335, 784, model.AddCritical(24.2)))
	无工之剑 := model.NewWeapon(weaponType.BigSword,
		model.BaseWeapon(90, 608, model.AddAttackPer(49.6)))
	log.Printf("%+v\n", 迪卢克)
	log.Printf("%+v\n", 无工之剑)
}
