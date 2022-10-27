package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
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
	"log"
	"math"
	"strings"
	"time"
)

func main() {
	//if err := 迪卢克2(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	if err := 神里绫华(); err != nil {
		log.Fatalf("%+v", err)
	}
	//超绽放队()
	//绽放队()
}

func 神里绫华() error {
	神里绫华 := model.CharacterFactory神里绫华(9, 9, 9, 0)

	if _, err := 神里绫华.Weapon(weapon.Factory雾切之回光(1, 2, elementals.Ice)); err != nil {
		return err
	} else if 生之花, err := model.ArtifactsFactory生之花(5, "历经风雪的思念", 20, model.FloatEntries{
		entry.Def: 42, entry.AtkPercentage: 15.2, entry.HpPercentage: 9.3, entry.CriticalDamage: 12.4}); err != nil {
		return err
	} else if 死之羽, err := model.ArtifactsFactory死之羽(5, "摧冰而行的执望", 20, model.FloatEntries{
		entry.DefPercentage: 27, entry.EnergyRecharge: 6.5, entry.CriticalDamage: 12.4, entry.AtkPercentage: 5.3}); err != nil {
		return err
	} else if 时之沙, err := model.ArtifactsFactory时之沙(5, "冰雪故园的终期", entry.AtkPercentage, 20, model.FloatEntries{
		entry.Atk: 18, entry.CriticalRate: 8.9, entry.CriticalDamage: 21, entry.Def: 23}); err != nil {
		return err
	} else if 空之杯, err := model.ArtifactsFactory空之杯(5, "度火者的醒悟", entry.IceDamageBonus, 20, model.FloatEntries{
		entry.CriticalRate: 2.7, entry.EnergyRecharge: 5.2, entry.ElementalMastery: 63, entry.AtkPercentage: 16.3}); err != nil {
		return err
	} else if 理之冠, err := model.ArtifactsFactory理之冠(5, "破冰踏雪的回音", entry.CriticalDamage, 20, model.FloatEntries{
		entry.Def: 44, entry.EnergyRecharge: 11.7, entry.HpPercentage: 15.7, entry.Atk: 18}); err != nil {
		return err
	} else {
		神里绫华.Artifacts(生之花)
		神里绫华.Artifacts(死之羽)
		神里绫华.Artifacts(时之沙)
		神里绫华.Artifacts(空之杯)
		神里绫华.Artifacts(理之冠)
	}

	神里绫华.Apply(
		buff.AddAttackDamageBonus(30, attackMode.NormalAttack, attackMode.ChargedAttack), // 绫华固有天赋5
		buff.AddElementalDamageBonus(18, elementals.Ice),                                 // 绫华固有天赋6
		buff.AddElementalDamageBonus(15, elementals.Ice),                                 // 冰二件套
		buff.AddCriticalRate(40),                                                         // 冰四件套
		buff.AddCriticalRate(15),                                                         // 双冰共鸣
	)

	挨揍的 := enemy.New(enemy.Base(90))
	挨揍的.Attach(elementals.Ice, 12)
	//挨揍的.AttachState(states.Frozen, 12)

	replaceArtifacts := make(map[position.Position]*model.Artifacts)
	if 生之花, err := model.ArtifactsFactory生之花(5, "历经风雪的思念", 20, model.IntEntries{
		entry.CriticalRate: 6, entry.AtkPercentage: 1, entry.CriticalDamage: 1, entry.Atk: 1}); err != nil {
		return err
	} else if 死之羽, err := model.ArtifactsFactory死之羽(5, "摧冰而行的执望", 20, model.IntEntries{
		entry.CriticalRate: 6, entry.AtkPercentage: 1, entry.CriticalDamage: 1, entry.EnergyRecharge: 1}); err != nil {
		return err
	} else if 时之沙, err := model.ArtifactsFactory时之沙(5, "冰雪故园的终期", entry.AtkPercentage, 20, model.IntEntries{
		entry.CriticalRate: 6, entry.CriticalDamage: 1, entry.Atk: 1, entry.EnergyRecharge: 1}); err != nil {
		return err
	} else if 空之杯, err := model.ArtifactsFactory空之杯(5, "度火者的醒悟", entry.IceDamageBonus, 20, model.IntEntries{
		entry.CriticalRate: 6, entry.AtkPercentage: 1, entry.CriticalDamage: 1, entry.Atk: 1}); err != nil {
		return err
	} else if 理之冠, err := model.ArtifactsFactory理之冠(5, "破冰踏雪的回音", entry.CriticalRate, 20, model.IntEntries{
		entry.CriticalDamage: 4, entry.AtkPercentage: 3, entry.Atk: 1, entry.EnergyRecharge: 1}); err != nil {
		return err
	} else {
		replaceArtifacts[生之花.Position()] = 生之花
		replaceArtifacts[死之羽.Position()] = 死之羽
		replaceArtifacts[时之沙.Position()] = 时之沙
		replaceArtifacts[空之杯.Position()] = 空之杯
		replaceArtifacts[理之冠.Position()] = 理之冠
	}

	action := 神里绫华.GetActions().Get(attackMode.ChargedAttack, "")
	profitDetect(神里绫华, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, elementals.Ice).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Ice), replaceArtifacts)
	return nil
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
	}, CustomDetects(elementals.Grass), nil)
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
	}, CustomDetects(elementals.Electric), nil)
}

func 迪卢克1() error {
	迪卢克 := model.CharacterFactory迪卢克(10, 9, 9, 0)

	if _, err := 迪卢克.Weapon(weapon.Factory螭骨剑(3)); err != nil {
		return err
		//} else if _, err := 迪卢克.Weapon(weapon.Factory无工之剑(1)); err != nil {
		//	return err
	} else if 生之花, err := model.ArtifactsFactory生之花(5, "魔女的炎之花", 20, model.FloatEntries{
		entry.Atk: 51, entry.AtkPercentage: 12.8, entry.CriticalRate: 3.1, entry.DefPercentage: 6.6}); err != nil {
		return err
	} else if 死之羽, err := model.ArtifactsFactory死之羽(5, "魔女常燃之羽", 20, model.FloatEntries{
		entry.CriticalRate: 7.8, entry.Hp: 239, entry.CriticalDamage: 14, entry.ElementalMastery: 54}); err != nil {
		return err
	} else if 时之沙, err := model.ArtifactsFactory时之沙(5, "魔女破灭之时", entry.AtkPercentage, 20, model.FloatEntries{
		entry.CriticalDamage: 11.7, entry.ElementalMastery: 61, entry.EnergyRecharge: 15.5, entry.CriticalRate: 3.1}); err != nil {
		return err
	} else if 空之杯, err := model.ArtifactsFactory空之杯(5, "魔女的心之火", entry.FireDamageBonus, 20, model.FloatEntries{
		entry.Hp: 986, entry.HpPercentage: 9.3, entry.CriticalRate: 3.9, entry.Def: 35}); err != nil {
		return err
	} else if 理之冠, err := model.ArtifactsFactory理之冠(5, "渡火者的智慧", entry.CriticalDamage, 20, model.FloatEntries{
		entry.AtkPercentage: 15.2, entry.CriticalRate: 6.6, entry.EnergyRecharge: 11.7, entry.Hp: 269}); err != nil {
		return err
	} else {
		迪卢克.Artifacts(生之花)
		迪卢克.Artifacts(死之羽)
		迪卢克.Artifacts(时之沙)
		迪卢克.Artifacts(空之杯)
		迪卢克.Artifacts(理之冠)
	}

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

	action := 迪卢克.GetActions().Get(attackMode.ElementalSkill, "1段")
	profitDetect(迪卢克, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, -1).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Fire), nil)
	return nil
}

func 迪卢克2() error {
	迪卢克 := model.CharacterFactory迪卢克(10, 9, 9, 0)

	if _, err := 迪卢克.Weapon(weapon.Factory无工之剑(1)); err != nil {
		return err
	} else if 生之花, err := model.ArtifactsFactory生之花(5, "明威之镡", 20, model.FloatEntries{
		entry.CriticalDamage: 26.4, entry.AtkPercentage: 4.1, entry.CriticalRate: 3.9, entry.Def: 42}); err != nil {
		return err
	} else if 生之花2, err := model.ArtifactsFactory生之花(5, "魔女的炎之花", 20, model.FloatEntries{
		entry.Atk: 51, entry.AtkPercentage: 12.8, entry.CriticalRate: 3.1, entry.DefPercentage: 6.6}); err != nil {
		return err
	} else if 死之羽, err := model.ArtifactsFactory死之羽(5, "魔女常燃之羽", 20, model.FloatEntries{
		entry.CriticalRate: 7.8, entry.Hp: 239, entry.CriticalDamage: 14, entry.ElementalMastery: 54}); err != nil {
		return err
	} else if 时之沙, err := model.ArtifactsFactory时之沙(5, "魔女破灭之时", entry.AtkPercentage, 20, model.FloatEntries{
		entry.CriticalDamage: 11.7, entry.ElementalMastery: 61, entry.EnergyRecharge: 15.5, entry.CriticalRate: 3.1}); err != nil {
		return err
	} else if 空之杯, err := model.ArtifactsFactory空之杯(5, "魔女的心之火", entry.FireDamageBonus, 20, model.FloatEntries{
		entry.Hp: 986, entry.HpPercentage: 9.3, entry.CriticalRate: 3.9, entry.Def: 35}); err != nil {
		return err
	} else if 空之杯2, err := model.ArtifactsFactory空之杯(5, "超越之盏", entry.FireDamageBonus, 20, model.FloatEntries{
		entry.EnergyRecharge: 14.2, entry.ElementalMastery: 16, entry.CriticalDamage: 12.4, entry.Hp: 448}); err != nil {
		return err
	} else if 理之冠, err := model.ArtifactsFactory理之冠(5, "焦灼的魔女帽", entry.CriticalRate, 20, model.FloatEntries{
		entry.AtkPercentage: 9.9, entry.Atk: 18, entry.ElementalMastery: 35, entry.CriticalDamage: 20.2}); err != nil {
		return err
	} else {
		迪卢克.Artifacts(生之花2)
		迪卢克.Artifacts(生之花)
		迪卢克.Artifacts(死之羽)
		迪卢克.Artifacts(时之沙)
		迪卢克.Artifacts(空之杯2)
		迪卢克.Artifacts(空之杯)
		迪卢克.Artifacts(理之冠)
	}

	迪卢克.Apply(
		buff.AddElementalDamageBonus(20, elementals.Fire),   // 卢姥爷大招
		buff.AddElementalDamageBonus(37.5, elementals.Fire), // 魔女套
		buff.AddReactionDamageBonus(40, reactions.Overload, reactions.Burn, reactions.Burgeon),
		buff.AddReactionDamageBonus(15, reactions.Vaporize, reactions.Melt),
		buff.AddAtkPercentage(25), // 双火共鸣
	)

	挨揍的 := enemy.New(enemy.Base(90))
	挨揍的.Attach(elementals.Water, 12)

	replaceArtifacts := make(map[position.Position]*model.Artifacts)
	if 生之花, err := model.ArtifactsFactory生之花(5, "明威之镡", 20, model.IntEntries{
		entry.CriticalDamage: 6, entry.ElementalMastery: 1, entry.CriticalRate: 1, entry.AtkPercentage: 1}); err != nil {
		return err
	} else if 死之羽, err := model.ArtifactsFactory死之羽(5, "魔女常燃之羽", 20, model.IntEntries{
		entry.CriticalDamage: 5, entry.ElementalMastery: 2, entry.CriticalRate: 1, entry.AtkPercentage: 1}); err != nil {
		return err
	} else if 时之沙, err := model.ArtifactsFactory时之沙(5, "魔女破灭之时", entry.ElementalMastery, 20, model.IntEntries{
		entry.CriticalDamage: 5, entry.Atk: 1, entry.CriticalRate: 2, entry.AtkPercentage: 1}); err != nil {
		return err
	} else if 空之杯, err := model.ArtifactsFactory空之杯(5, "魔女的心之火", entry.FireDamageBonus, 20, model.IntEntries{
		entry.CriticalDamage: 5, entry.ElementalMastery: 1, entry.CriticalRate: 2, entry.AtkPercentage: 1}); err != nil {
		return err
	} else if 理之冠, err := model.ArtifactsFactory理之冠(5, "焦灼的魔女帽", entry.CriticalRate, 20, model.IntEntries{
		entry.CriticalDamage: 6, entry.ElementalMastery: 1, entry.Atk: 1, entry.AtkPercentage: 13}); err != nil {
		return err
	} else {
		replaceArtifacts[生之花.Position()] = 生之花
		replaceArtifacts[死之羽.Position()] = 死之羽
		replaceArtifacts[时之沙.Position()] = 时之沙
		replaceArtifacts[空之杯.Position()] = 空之杯
		replaceArtifacts[理之冠.Position()] = 理之冠
	}

	action := 迪卢克.GetActions().Get(attackMode.ElementalSkill, "1段")
	profitDetect(迪卢克, 挨揍的, func(player *model.Character, enemy *enemy.Enemy, debug bool) float64 {
		_, avg, _ := player.Calculate(enemy, action, -1).Calculate(debug)
		return avg.Value()
	}, CustomDetects(elementals.Fire), replaceArtifacts)
	return nil
}

func CustomDetects(dye elementals.Elemental) map[string]*attr.Modifier {
	return map[string]*attr.Modifier{
		"玉璋护盾":    attr.NewModifier(buff.Superposition(5, time.Second*20, 0, buff.AddShieldStrength(5)), buff.AddAllElementalResist(-20)),
		"万叶":      attr.NewCharacterModifier(buff.AddElementalDamageBonus(0.04*1000, dye)),
		"风四件套":    attr.NewEnemyModifier(buff.AddElementalResist(-40, dye)),
		"万叶+风四件套": attr.NewModifier(buff.AddElementalDamageBonus(0.04*1000, dye), buff.AddElementalResist(-40, dye)),
		"班尼特":     attr.NewCharacterModifier(buff.AddAtk(int(math.Round(1.19 * (191 + 565))))),
		"班尼特6命":   attr.NewCharacterModifier(buff.AddAtk(int(math.Round(1.19*(191+565)))), buff.AddElementalDamageBonus(15, elementals.Fire)),
		"讨龙英杰谭":   attr.NewCharacterModifier(buff.AddAtkPercentage(48)),
		"砂糖":      attr.NewCharacterModifier(buff.AddElementalMastery(50 + 200)),
		"砂糖+风四件套": attr.NewModifier(buff.AddElementalMastery(50+200), buff.AddElementalResist(-40, dye)),
		"砂糖6命":    attr.NewCharacterModifier(buff.AddElementalMastery(50+200), buff.AddElementalDamageBonus(20, dye)),
		"莫娜星异":    attr.NewCharacterModifier(buff.AddDamageBonus(60)),
		"深林四件套":   attr.NewEnemyModifier(buff.AddElementalResist(-30, elementals.Grass)),
		"如雷四件套":   attr.NewCharacterModifier(buff.AddReactionDamageBonus(40, reactions.Overload, reactions.ElectroCharged, reactions.Superconduct, reactions.Hyperbloom)),
	}
}

func profitDetect(character *model.Character, enemy *enemy.Enemy, fn detect.FinalDamage,
	customDetects map[string]*attr.Modifier, replaceArtifacts map[position.Position]*model.Artifacts) {
	fmt.Printf("base: %f\n", fn(character, enemy, true))
	profits := detect.ProfitDetect(character, enemy, true, fn, nil)
	fmt.Printf("素质增益:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
	if len(customDetects) > 0 {
		profits = detect.ProfitDetect(character, enemy, false, fn, customDetects)
		fmt.Printf("队友增益:\n")
		for _, p := range profits {
			fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
		}
	}
	if evaluateDetects := character.Evaluate(replaceArtifacts); len(evaluateDetects) > 0 {
		profits = detect.ProfitDetect(character, enemy, false, fn, evaluateDetects)
		fmt.Printf("角色增益:\n")
		for _, p := range profits {
			if strings.Index(p.Name, "-") < 0 {
				fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
				for _, s := range profits {
					if strings.HasPrefix(s.Name, p.Name) && s.Name != p.Name {
						fmt.Printf("\t\t增幅：%2.4f%% - %s\n", s.Value, s.Name)
					}
				}
			}
		}
	}
}
