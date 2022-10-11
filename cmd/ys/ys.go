package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/model"
	"sort"
)

type FinalDamage func(player *model.Character) float64

type Profit struct {
	Name  string
	Value float64
}

func main() {
	迪卢克 := model.CharacterFactory迪卢克(10, 9, 9, 0)
	魔女的炎之花 := model.ArtifactsFactory生之花(5, model.AddAtk(51), model.AddAtkPercentage(12.8),
		model.AddCriticalRate(3.1), model.AddDefPercentage(6.6))
	魔女常燃之羽 := model.ArtifactsFactory死之羽(5, model.AddCriticalRate(7.8), model.AddHp(239),
		model.AddCriticalDamage(14), model.AddElementalMastery(54))
	魔女破灭之时 := model.NewArtifacts(5, position.SandsOfEon, model.BaseArtifacts(20, point.AtkPercentage, 46.6),
		model.AddCriticalDamage(11.7), model.AddElementalMastery(61), model.AddEnergyRecharge(15.5), model.AddCriticalRate(3.1))
	魔女的心之火 := model.NewArtifacts(5, position.GobletOfEonothem, model.BaseArtifacts(20, point.PyroDamageBonus, 46.6),
		model.AddHp(986), model.AddHpPercentage(9.3), model.AddCriticalRate(3.9), model.AddDef(35))
	渡火者的智慧 := model.NewArtifacts(5, position.CircletOfLogos, model.BaseArtifacts(20, point.CriticalDamage, 62.2),
		model.AddAtkPercentage(15.2), model.AddCriticalRate(6.6), model.AddEnergyRecharge(11.7), model.AddHp(269))

	迪卢克.Weapon(model.WeaponFactory螭骨剑(3))
	//迪卢克.Weapon(model.WeaponFactory无工之剑(1))
	迪卢克.Artifacts(魔女的炎之花)
	迪卢克.Artifacts(魔女常燃之羽)
	迪卢克.Artifacts(魔女破灭之时)
	迪卢克.Artifacts(魔女的心之火)
	迪卢克.Artifacts(渡火者的智慧)

	迪卢克.Apply(model.AddElementalDamageBonus(elemental.Fire, 15))

	enemy := model.NewEnemy(model.BaseEnemy(90, model.AddAllElementalResist(10)))
	//enemy.Attach(elemental.Electric, 12)
	enemy.Attach(elemental.Water, 12)

	action := 迪卢克.GetActions().GetAction(attackMode.ElementalSkill, "逆焰之刃•1段")
	profitDetect(迪卢克, func(player *model.Character) float64 {
		_, avg, _ := 迪卢克.Calculate(enemy, action, -1).Calculate()
		return avg.Value()
	}, map[string]model.AttributeModifier{
		point.CriticalRate.String():     model.AddCriticalRate(2.7),
		point.HpPercentage.String():     model.AddHpPercentage(4.1),
		point.AtkPercentage.String():    model.AddAtkPercentage(4.1),
		point.EnergyRecharge.String():   model.AddEnergyRecharge(4.5),
		point.DefPercentage.String():    model.AddDefPercentage(5.1),
		point.CriticalDamage.String():   model.AddCriticalDamage(5.4),
		point.Atk.String():              model.AddAtk(14),
		point.Def.String():              model.AddDef(16),
		point.ElementalMastery.String(): model.AddElementalMastery(16),
		point.Hp.String():               model.AddHp(209),
	})
}

func profitDetect(character *model.Character, fn FinalDamage, customDetects map[string]model.AttributeModifier) {
	fmt.Printf("base: %f\n", fn(character))
	base := fn(character)
	var profits []*Profit
	for name, modifier := range customDetects {
		cancel := character.Apply(modifier)
		value := fn(character)
		if value != base {
			profits = append(profits, &Profit{
				Name:  name,
				Value: 100 * (value - base) / base,
			})
		}
		cancel()
	}
	sort.Slice(profits, func(i, j int) bool {
		if profits[i].Value < profits[j].Value {
			return false
		} else if profits[i].Value > profits[j].Value {
			return true
		} else {
			return profits[i].Name < profits[j].Name
		}
	})
	fmt.Printf("素质:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
}
