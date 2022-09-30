package main

import (
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/model"
	"log"
)

func main() {
	迪卢克 := model.CharacterFactory迪卢克()
	if _, err := 迪卢克.Weapon(model.Weapon无工之剑); err != nil {
		log.Fatalf("%+v\n", err)
	}
	if _, err := 迪卢克.Weapon(model.Weapon原木刀); err != nil {
		log.Printf("%+v\n", err)
	}
	log.Printf("%+v\n", 迪卢克)
}
