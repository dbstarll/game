package main

import (
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/model"
	"log"
)

func main() {
	迪卢克 := model.CharacterFactory迪卢克()
	无工之剑 := model.WeaponFactory无工之剑(1)
	if _, err := 迪卢克.Weapon(无工之剑); err != nil {
		log.Fatalf("%+v\n", err)
	}
	log.Printf("%+v\n", 迪卢克)
}
