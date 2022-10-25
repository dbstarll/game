package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/states"
	"github.com/dbstarll/game/internal/ys/model"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/detect"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"github.com/dbstarll/game/internal/ys/model/weapon"
	"math"
	"time"
)

func main() {
	迪卢克1()
	//超绽放队()
	//绽放队()
}

func 绽放队() {
	草主 := model.CharacterFactory草主(1, 1, 1, 6)
	草主.Weapon(weapon.Factory原木刀(1))
	草主.Apply(buff.AddElementalMastery(1000))
	草主.GetActions().Loop(func(index int, action *action.Action) bool {
		fmt.Println(action)
		return false
	})
	action := 草主.GetActions().Get(attackMode.ElementalSkill, "技能伤害")
	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Attach(elemental.Electric, 12)
	挨揍的.Attach(elementals.Water, 12)
	profitDetect(草主, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, -1).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Grass))
}

func 超绽放队() {
	久岐忍 := model.CharacterFactory久岐忍(1, 6, 1, 0)
	久岐忍.Weapon(weapon.Factory原木刀(1))
	久岐忍.Apply(buff.AddElementalMastery(860))
	久岐忍.GetActions().Loop(func(index int, action *action.Action) bool {
		fmt.Println(action)
		return false
	})
	action := 久岐忍.GetActions().Get(attackMode.ElementalSkill, "")
	挨揍的 := enemy.New(enemy.Base(90))
	挨揍的.Apply(buff.AddElementalResist(-30, elementals.Grass))
	挨揍的.Attach(elementals.Grass, 12)
	挨揍的.AttachState(states.Bloom, 12)
	挨揍的.AttachState(states.Quicken, 12)
	profitDetect(久岐忍, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, -1).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Electric))
}

func 迪卢克1() {
	迪卢克 := model.CharacterFactory迪卢克(10, 9, 9, 0)
	魔女的炎之花 := model.ArtifactsFactory生之花(5, "魔女的炎之花", buff.AddAtk(51), buff.AddAtkPercentage(12.8),
		buff.AddCriticalRate(3.1), buff.AddDefPercentage(6.6))
	魔女常燃之羽 := model.ArtifactsFactory死之羽(4, "魔女常燃之羽", buff.AddCriticalRate(7.8), buff.AddHp(239),
		buff.AddCriticalDamage(14), buff.AddElementalMastery(54))
	魔女破灭之时 := model.NewArtifacts(3, position.SandsOfEon, "魔女破灭之时", model.BaseArtifacts(20, buff.AddAtkPercentage(46.6)),
		buff.AddCriticalDamage(11.7), buff.AddElementalMastery(61), buff.AddEnergyRecharge(15.5), buff.AddCriticalRate(3.1))
	魔女的心之火 := model.NewArtifacts(2, position.GobletOfEonothem, "魔女的心之火", model.BaseArtifacts(20, buff.AddElementalDamageBonus(46.6, elementals.Fire)),
		buff.AddHp(986), buff.AddHpPercentage(9.3), buff.AddCriticalRate(3.9), buff.AddDef(35))
	渡火者的智慧 := model.NewArtifacts(5, position.CircletOfLogos, "渡火者的智慧", model.BaseArtifacts(20, buff.AddCriticalDamage(62.2)),
		buff.AddAtkPercentage(15.2), buff.AddCriticalRate(6.6), buff.AddEnergyRecharge(11.7), buff.AddHp(269))

	迪卢克.Weapon(weapon.Factory螭骨剑(3))
	//迪卢克.Weapon(weapon.Factory无工之剑(1))
	迪卢克.Artifacts(魔女的炎之花)
	迪卢克.Artifacts(魔女常燃之羽)
	迪卢克.Artifacts(魔女破灭之时)
	迪卢克.Artifacts(魔女的心之火)
	迪卢克.Artifacts(渡火者的智慧)

	迪卢克.Apply(
		buff.AddElementalDamageBonus(20, elementals.Fire),   // 卢姥爷大招
		buff.AddElementalDamageBonus(37.5, elementals.Fire), // 魔女套
		buff.AddReactionDamageBonus(40, reactions.Overload, reactions.Burn, reactions.Burgeon),
		buff.AddReactionDamageBonus(15, reactions.Vaporize, reactions.Melt),
		buff.AddAtkPercentage(25), // 双火共鸣
	)

	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Attach(elemental.Electric, 12)
	挨揍的.Attach(elementals.Water, 12)

	迪卢克.GetActions().Loop(func(index int, action *action.Action) bool {
		fmt.Println(action)
		return false
	})

	action := 迪卢克.GetActions().Get(attackMode.ElementalSkill, "1段")
	profitDetect(迪卢克, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, -1).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Fire))
}

func CustomDetects(dye elementals.Elemental) map[string]*detect.Modifier {
	return map[string]*detect.Modifier{
		"玉璋护盾":      detect.NewModifier(buff.Superposition(5, time.Second*20, 0, buff.AddShieldStrength(5)), buff.AddAllElementalResist(-20)),
		"万叶扩散":      detect.NewCharacterModifier(buff.AddElementalDamageBonus(0.04*1000, dye)),
		"风四件套":      detect.NewEnemyModifier(buff.AddElementalResist(-40, dye)),
		"万叶扩散+风四件套": detect.NewModifier(buff.AddElementalDamageBonus(0.04*1000, dye), buff.AddElementalResist(-40, dye)),
		"班尼特":       detect.NewCharacterModifier(buff.AddAtk(int(math.Round(1.19 * (191 + 565))))),
		"班尼特6命":     detect.NewCharacterModifier(attr.MergeAttributes(buff.AddAtk(int(math.Round(1.19*(191+565)))), buff.AddElementalDamageBonus(15, dye))),
		"讨龙英杰谭":     detect.NewCharacterModifier(buff.AddAtkPercentage(48)),
		"砂糖":        detect.NewCharacterModifier(buff.AddElementalMastery(50 + 200)),
		"砂糖6命":      detect.NewCharacterModifier(attr.MergeAttributes(buff.AddElementalMastery(50+200), buff.AddElementalDamageBonus(20, dye))),
		"莫娜星异":      detect.NewCharacterModifier(buff.AddDamageBonus(60)),
		"深林四件套":     detect.NewEnemyModifier(buff.AddElementalResist(-30, elementals.Grass)),
		"如雷四件套":     detect.NewCharacterModifier(buff.AddReactionDamageBonus(40, reactions.Overload, reactions.ElectroCharged, reactions.Superconduct, reactions.Hyperbloom)),
	}
}

func profitDetect(character *model.Character, enemy *enemy.Enemy, fn detect.FinalDamage, customDetects map[string]*detect.Modifier) {
	fmt.Printf("base: %f\n", fn(character, enemy, true))
	profits := detect.ProfitDetect(character, enemy, true, fn, nil)
	fmt.Printf("素质增益:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
	profits = detect.ProfitDetect(character, enemy, false, fn, customDetects)
	fmt.Printf("角色增益:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
}
