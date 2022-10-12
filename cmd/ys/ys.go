package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/model"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/detect"
)

func main() {
	迪卢克 := model.CharacterFactory迪卢克(10, 9, 9, 0)
	魔女的炎之花 := model.ArtifactsFactory生之花(5, buff.AddAtk(51), buff.AddAtkPercentage(12.8),
		buff.AddCriticalRate(3.1), buff.AddDefPercentage(6.6))
	魔女常燃之羽 := model.ArtifactsFactory死之羽(5, buff.AddCriticalRate(7.8), buff.AddHp(239),
		buff.AddCriticalDamage(14), buff.AddElementalMastery(54))
	魔女破灭之时 := model.NewArtifacts(5, position.SandsOfEon, model.BaseArtifacts(20, point.AtkPercentage, 46.6),
		buff.AddCriticalDamage(11.7), buff.AddElementalMastery(61), buff.AddEnergyRecharge(15.5), buff.AddCriticalRate(3.1))
	魔女的心之火 := model.NewArtifacts(5, position.GobletOfEonothem, model.BaseArtifacts(20, point.PyroDamageBonus, 46.6),
		buff.AddHp(986), buff.AddHpPercentage(9.3), buff.AddCriticalRate(3.9), buff.AddDef(35))
	渡火者的智慧 := model.NewArtifacts(5, position.CircletOfLogos, model.BaseArtifacts(20, point.CriticalDamage, 62.2),
		buff.AddAtkPercentage(15.2), buff.AddCriticalRate(6.6), buff.AddEnergyRecharge(11.7), buff.AddHp(269))

	迪卢克.Weapon(model.WeaponFactory螭骨剑(3))
	//迪卢克.Weapon(model.WeaponFactory无工之剑(1))
	迪卢克.Artifacts(魔女的炎之花)
	迪卢克.Artifacts(魔女常燃之羽)
	迪卢克.Artifacts(魔女破灭之时)
	迪卢克.Artifacts(魔女的心之火)
	迪卢克.Artifacts(渡火者的智慧)

	迪卢克.Apply(buff.AddElementalDamageBonus(elemental.Fire, 15))

	enemy := model.NewEnemy(model.BaseEnemy(90, buff.AddAllElementalResist(10)))
	//enemy.Attach(elemental.Electric, 12)
	enemy.Attach(elemental.Water, 12)

	action := 迪卢克.GetActions().GetAction(attackMode.ElementalSkill, "逆焰之刃•1段")
	profitDetect(迪卢克, func(player *model.Character) float64 {
		_, avg, _ := 迪卢克.Calculate(enemy, action, -1).Calculate()
		return avg.Value()
	}, map[string]attr.AttributeModifier{
		//"玉璋护盾": nil,
	})
}

func profitDetect(character *model.Character, fn detect.FinalDamage, customDetects map[string]attr.AttributeModifier) {
	fmt.Printf("base: %f\n", fn(character))
	profits := detect.ProfitDetect(character, true, fn, nil)
	fmt.Printf("素质:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
	profits = detect.ProfitDetect(character, false, fn, customDetects)
	fmt.Printf("custom:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
}
