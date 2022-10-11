package main

import (
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/model"
	"go.uber.org/zap"
)

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
	迪卢克.Artifacts(魔女的炎之花)
	迪卢克.Artifacts(魔女常燃之羽)
	迪卢克.Artifacts(魔女破灭之时)
	迪卢克.Artifacts(魔女的心之火)
	迪卢克.Artifacts(渡火者的智慧)

	迪卢克.Apply(model.AddElementalDamageBonus(elemental.Fire, 15))

	enemy := model.NewEnemy(model.BaseEnemy(90, model.AddAllElementalResist(10)))
	enemy.Attach(elemental.Electric, 12)
	enemy.Attach(elemental.Water, 12)

	迪卢克.GetActions().Loop(func(index int, action *model.Action) bool {
		min, avg, max := 迪卢克.Calculate(enemy, action, elemental.Fire).Calculate()
		zap.S().Infof("总伤害: [%.0f, %.0f, %.0f] - %s", min.Value(), avg.Value(), max.Value(), action)
		return false
	})
}
