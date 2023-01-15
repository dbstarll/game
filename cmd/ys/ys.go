package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/states"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/artifacts"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/calculator"
	"github.com/dbstarll/game/internal/ys/model/character"
	"github.com/dbstarll/game/internal/ys/model/detect"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"github.com/dbstarll/game/internal/ys/model/weapon"
	"log"
	"math"
	"strings"
	"time"
)

func main() {
	//迪卢克()
	//神里绫华()
	//超绽放队()
	//绽放队()
	//纳西妲()
	雷神一刀队()
}

var (
	damage detect.FinalDamage = func(character *character.Character, enemy *enemy.Enemy, action *action.Action, debug bool, finalModifiers ...attr.AttributeModifier) float64 {
		_, avg, _ := calculator.New(character, enemy, action, finalModifiers...).Calculate(debug)
		return avg.Value()
	}
	Weapon = func(character *character.Character, weapon *weapon.Weapon) {
		if _, err := character.Weapon(weapon); err != nil {
			log.Fatalf("%+v", err)
		}
	}
	Artifacts = func(artifacts *artifacts.Artifacts, err error) *artifacts.Artifacts {
		if err != nil {
			log.Fatalf("%+v", err)
		}
		return artifacts
	}
)

func 雷神一刀队() {
	雷电将军 := character.Factory雷电将军(9, 9, 9, 0)

	Weapon(雷电将军, weapon.Factory渔获(5))

	雷电将军.Artifacts(Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.AtkPercentage, 14}, {entry.CriticalRate, 7}, {entry.Atk, 54}, {entry.Def, 16}})))
	雷电将军.Artifacts(Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 7.4}, {entry.Def, 32}, {entry.AtkPercentage, 14.6}, {entry.EnergyRecharge, 6.5}})))
	雷电将军.Artifacts(Artifacts(artifacts.Factory时之沙(5, entry.AtkPercentage, &artifacts.FloatEntries{{entry.Atk, 18}, {entry.CriticalRate, 8.9}, {entry.CriticalDamage, 21}, {entry.Def, 23}})))
	雷电将军.Artifacts(Artifacts(artifacts.Factory空之杯(5, entry.ElectricDamageBonus, &artifacts.FloatEntries{{entry.Def, 35}, {entry.CriticalDamage, 22.5}, {entry.HpPercentage, 5.3}, {entry.CriticalRate, 6.6}})))
	雷电将军.Artifacts(Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.AtkPercentage, 15.2}, {entry.CriticalRate, 6.6}, {entry.EnergyRecharge, 11.7}, {entry.Hp, 269}})))

	actions := 雷电将军.GetActions()
	梦想一刀愿力加成 := actions.Get(attackMode.ElementalBurst, "梦想一刀愿力加成")
	元素爆发伤害提高 := actions.Get(attackMode.ElementalSkill, "元素爆发伤害提高")

	雷电将军.Apply(
		buff.Character雷电将军恶曜开眼(90, 元素爆发伤害提高.DMG()),                                 // 元素战技加成
		buff.AddAttackFactorAddBonus(60*梦想一刀愿力加成.DMG(), attackMode.ElementalBurst), // 梦想一刀愿力加成
		//buff.Character万叶扩散(1000, elementals.Electric),
		//buff.AddAtkPercentage(48), // 讨龙英杰谭
		//buff.AddAtkPercentage(20), // 岩四件套 or 宗室
		//buff.AddDamageBonus(60),   // 莫娜星异
	)

	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Apply(buff.AddDefPercentage(-30))
	//挨揍的.Attach(elementals.Ice, 12)
	//挨揍的.AttachState(states.Quicken, 12)
	//buff.Artifacts翠绿之影4(elementals.Electric).Apply(nil, 挨揍的, nil)

	replaceArtifacts := []*artifacts.Artifacts{
		Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.EnergyRecharge, 4.5}, {entry.CriticalRate, 10.5}, {entry.CriticalDamage, 19.4}, {entry.Def, 39}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.AtkPercentage, 15.7}, {entry.EnergyRecharge, 4.5}, {entry.Def, 32}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalDamage, 12.4}, {entry.CriticalRate, 6.6}, {entry.HpPercentage, 15.7}, {entry.ElementalMastery, 16}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 11.7}, {entry.Def, 23}, {entry.DefPercentage, 13.1}, {entry.CriticalDamage, 7.8}})),
		Artifacts(artifacts.Factory时之沙(5, entry.AtkPercentage, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.CriticalDamage, 6.2}, {entry.EnergyRecharge, 20.1}, {entry.Def, 21}})),
		Artifacts(artifacts.Factory空之杯(5, entry.IceDamageBonus, &artifacts.FloatEntries{{entry.Atk, 29}, {entry.CriticalDamage, 14}, {entry.AtkPercentage, 9.9}, {entry.Hp, 807}})),
		Artifacts(artifacts.Factory空之杯(5, entry.IceDamageBonus, &artifacts.FloatEntries{{entry.CriticalRate, 2.7}, {entry.EnergyRecharge, 5.2}, {entry.ElementalMastery, 63}, {entry.AtkPercentage, 16.3}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.Def, 37}, {entry.AtkPercentage, 9.3}, {entry.EnergyRecharge, 11.7}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.Def, 44}, {entry.EnergyRecharge, 11.7}, {entry.HpPercentage, 15.7}, {entry.Atk, 18}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.EnergyRecharge, 10.4}, {entry.DefPercentage, 13.9}, {entry.CriticalRate, 3.5}, {entry.AtkPercentage, 15.7}})),
	}
	攻击 := actions.Get(attackMode.ElementalBurst, "梦想一刀")
	profitDetect(雷电将军, 挨揍的, 攻击, damage, CustomDetects(elementals.Electric), replaceArtifacts, buff.Artifacts绝缘之旗印4(), buff.Character雷电将军殊胜之御体())
}

func 神里绫华() {
	神里绫华 := character.Factory神里绫华(9, 9, 10, 0)

	Weapon(神里绫华, weapon.Factory雾切之回光(1, 3, elementals.Ice))

	神里绫华.Artifacts(Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.AtkPercentage, 14}, {entry.CriticalRate, 7}, {entry.Atk, 54}, {entry.Def, 16}})))
	神里绫华.Artifacts(Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 7.4}, {entry.Def, 32}, {entry.AtkPercentage, 14.6}, {entry.EnergyRecharge, 6.5}})))
	神里绫华.Artifacts(Artifacts(artifacts.Factory时之沙(5, entry.AtkPercentage, &artifacts.FloatEntries{{entry.Atk, 18}, {entry.CriticalRate, 8.9}, {entry.CriticalDamage, 21}, {entry.Def, 23}})))
	神里绫华.Artifacts(Artifacts(artifacts.Factory空之杯(5, entry.IceDamageBonus, &artifacts.FloatEntries{{entry.Def, 35}, {entry.CriticalDamage, 22.5}, {entry.HpPercentage, 5.3}, {entry.CriticalRate, 6.6}})))
	神里绫华.Artifacts(Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.AtkPercentage, 15.2}, {entry.CriticalRate, 6.6}, {entry.EnergyRecharge, 11.7}, {entry.Hp, 269}})))

	神里绫华.Apply(
		buff.AddAttackDamageBonus(30, attackMode.NormalAttack, attackMode.ChargedAttack), // 绫华固有天赋5
		buff.AddElementalDamageBonus(18, elementals.Ice),                                 // 绫华固有天赋6
		buff.Artifacts冰风迷途的勇士4(true),
		buff.TeamIce(),
		buff.Character万叶扩散(1000, elementals.Ice),
		buff.AddAtkPercentage(48), // 讨龙英杰谭
		buff.AddAtkPercentage(20), // 岩四件套 or 宗室
		buff.AddDamageBonus(60),   // 莫娜星异
	)

	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Apply(buff.AddDefPercentage(-30))
	//挨揍的.Attach(elementals.Ice, 12)
	//挨揍的.AttachState(states.Frozen, 12)
	buff.Artifacts翠绿之影4(elementals.Ice).Apply(nil, 挨揍的, nil)

	replaceArtifacts := []*artifacts.Artifacts{
		Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.EnergyRecharge, 4.5}, {entry.CriticalRate, 10.5}, {entry.CriticalDamage, 19.4}, {entry.Def, 39}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.AtkPercentage, 15.7}, {entry.EnergyRecharge, 4.5}, {entry.Def, 32}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalDamage, 12.4}, {entry.CriticalRate, 6.6}, {entry.HpPercentage, 15.7}, {entry.ElementalMastery, 16}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 11.7}, {entry.Def, 23}, {entry.DefPercentage, 13.1}, {entry.CriticalDamage, 7.8}})),
		Artifacts(artifacts.Factory时之沙(5, entry.AtkPercentage, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.CriticalDamage, 6.2}, {entry.EnergyRecharge, 20.1}, {entry.Def, 21}})),
		Artifacts(artifacts.Factory空之杯(5, entry.IceDamageBonus, &artifacts.FloatEntries{{entry.Atk, 29}, {entry.CriticalDamage, 14}, {entry.AtkPercentage, 9.9}, {entry.Hp, 807}})),
		Artifacts(artifacts.Factory空之杯(5, entry.IceDamageBonus, &artifacts.FloatEntries{{entry.CriticalRate, 2.7}, {entry.EnergyRecharge, 5.2}, {entry.ElementalMastery, 63}, {entry.AtkPercentage, 16.3}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.CriticalRate, 6.6}, {entry.Def, 37}, {entry.AtkPercentage, 9.3}, {entry.EnergyRecharge, 11.7}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.Def, 44}, {entry.EnergyRecharge, 11.7}, {entry.HpPercentage, 15.7}, {entry.Atk, 18}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalDamage, &artifacts.FloatEntries{{entry.EnergyRecharge, 10.4}, {entry.DefPercentage, 13.9}, {entry.CriticalRate, 3.5}, {entry.AtkPercentage, 15.7}})),
	}

	攻击 := 神里绫华.GetActions().Get(attackMode.ElementalBurst, "切割")
	攻击.Apply(action.Infusion(elementals.Ice))
	profitDetect(神里绫华, 挨揍的, 攻击, damage, CustomDetects(elementals.Ice), replaceArtifacts)
}

func 绽放队() {
	草主 := character.Factory草主(1, 1, 1, 6)
	Weapon(草主, weapon.Factory原木刀(1))
	草主.Apply(buff.AddElementalMastery(1000))
	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Attach(elemental.Electric, 12)
	挨揍的.Attach(elementals.Water, 12)
	攻击 := 草主.GetActions().Get(attackMode.ElementalSkill, "技能伤害")
	profitDetect(草主, 挨揍的, 攻击, damage, CustomDetects(elementals.Grass), nil)
}

func 超绽放队() {
	久岐忍 := character.Factory久岐忍(1, 6, 1, 0)
	Weapon(久岐忍, weapon.Factory原木刀(1))
	久岐忍.Apply(buff.AddElementalMastery(860))
	挨揍的 := enemy.New(enemy.Base(90))
	挨揍的.Apply(buff.AddElementalResist(-30, elementals.Grass))
	挨揍的.Attach(elementals.Grass, 12)
	挨揍的.AttachState(states.Bloom, 12)
	挨揍的.AttachState(states.Quicken, 12)
	攻击 := 久岐忍.GetActions().Get(attackMode.ElementalSkill, "")
	profitDetect(久岐忍, 挨揍的, 攻击, damage, CustomDetects(elementals.Electric), nil)
}

func 迪卢克() {
	迪卢克 := character.Factory迪卢克(10, 9, 9, 0)

	Weapon(迪卢克, weapon.Factory无工之剑(1))

	迪卢克.Artifacts(Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.Atk, 51}, {entry.AtkPercentage, 12.8}, {entry.CriticalRate, 3.1}, {entry.DefPercentage, 6.6}})))
	迪卢克.Artifacts(Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 7.8}, {entry.Hp, 239}, {entry.CriticalDamage, 14}, {entry.ElementalMastery, 54}})))
	迪卢克.Artifacts(Artifacts(artifacts.Factory时之沙(5, entry.AtkPercentage, &artifacts.FloatEntries{{entry.CriticalDamage, 11.7}, {entry.ElementalMastery, 61}, {entry.EnergyRecharge, 15.5}, {entry.CriticalRate, 3.1}})))
	迪卢克.Artifacts(Artifacts(artifacts.Factory空之杯(5, entry.FireDamageBonus, &artifacts.FloatEntries{{entry.EnergyRecharge, 14.2}, {entry.ElementalMastery, 16}, {entry.CriticalDamage, 12.4}, {entry.Hp, 448}})))
	迪卢克.Artifacts(Artifacts(artifacts.Factory理之冠(5, entry.CriticalRate, &artifacts.FloatEntries{{entry.AtkPercentage, 9.9}, {entry.Atk, 18}, {entry.ElementalMastery, 35}, {entry.CriticalDamage, 20.2}})))

	迪卢克.Apply(
		buff.AddElementalDamageBonus(20, elementals.Fire), // 卢姥爷大招
		buff.Artifacts炽烈的炎之魔女4(3),
		buff.TeamFire(),
	)

	挨揍的 := enemy.New(enemy.Base(90))
	挨揍的.Attach(elementals.Water, 12)

	replaceArtifacts := []*artifacts.Artifacts{
		Artifacts(artifacts.Factory生之花(5, &artifacts.IntEntries{{entry.CriticalDamage, 6}, {entry.ElementalMastery, 1}, {entry.CriticalRate, 1}, {entry.AtkPercentage, 1}})),
		Artifacts(artifacts.Factory死之羽(5, &artifacts.IntEntries{{entry.CriticalDamage, 5}, {entry.ElementalMastery, 2}, {entry.CriticalRate, 1}, {entry.AtkPercentage, 1}})),
		Artifacts(artifacts.Factory时之沙(5, entry.ElementalMastery, &artifacts.IntEntries{{entry.CriticalDamage, 5}, {entry.Atk, 1}, {entry.CriticalRate, 2}, {entry.AtkPercentage, 1}})),
		Artifacts(artifacts.Factory空之杯(5, entry.FireDamageBonus, &artifacts.IntEntries{{entry.CriticalDamage, 5}, {entry.ElementalMastery, 1}, {entry.CriticalRate, 2}, {entry.AtkPercentage, 1}})),
		Artifacts(artifacts.Factory理之冠(5, entry.CriticalRate, &artifacts.IntEntries{{entry.CriticalDamage, 6}, {entry.ElementalMastery, 1}, {entry.Atk, 1}, {entry.AtkPercentage, 1}})),
		Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.CriticalDamage, 26.4}, {entry.AtkPercentage, 4.1}, {entry.CriticalRate, 3.9}, {entry.Def, 42}})),
		Artifacts(artifacts.Factory空之杯(5, entry.FireDamageBonus, &artifacts.FloatEntries{{entry.Hp, 986}, {entry.HpPercentage, 9.3}, {entry.CriticalRate, 3.9}, {entry.Def, 35}})),
	}

	攻击 := 迪卢克.GetActions().Get(attackMode.ElementalSkill, "1段")
	//攻击.Apply(action.Infusion(elementals.Fire))
	profitDetect(迪卢克, 挨揍的, 攻击, damage, CustomDetects(elementals.Fire), replaceArtifacts)
}

func 纳西妲() {
	纳西妲 := character.Factory纳西妲(1, 9, 9, 0)
	Weapon(纳西妲, weapon.Factory祭礼残章(5))
	纳西妲.Artifacts(Artifacts(artifacts.Factory生之花(5, &artifacts.FloatEntries{{entry.AtkPercentage, 14}, {entry.CriticalRate, 7}, {entry.Atk, 54}, {entry.Def, 16}})))
	纳西妲.Artifacts(Artifacts(artifacts.Factory死之羽(5, &artifacts.FloatEntries{{entry.CriticalRate, 11.7}, {entry.Def, 23}, {entry.DefPercentage, 13.1}, {entry.CriticalDamage, 7.8}})))
	纳西妲.Artifacts(Artifacts(artifacts.Factory时之沙(5, entry.ElementalMastery, &artifacts.IntEntries{{entry.CriticalRate, 1}, {entry.CriticalDamage, 6}, {entry.AtkPercentage, 1}, {entry.Atk, 1}})))
	纳西妲.Artifacts(Artifacts(artifacts.Factory空之杯(5, entry.ElementalMastery, &artifacts.IntEntries{{entry.CriticalRate, 3}, {entry.CriticalDamage, 3}, {entry.AtkPercentage, 1}, {entry.Atk, 1}})))
	纳西妲.Artifacts(Artifacts(artifacts.Factory理之冠(5, entry.ElementalMastery, &artifacts.IntEntries{{entry.CriticalRate, 3}, {entry.CriticalDamage, 3}, {entry.AtkPercentage, 1}, {entry.Atk, 1}})))

	actions := 纳西妲.GetActions()
	actionQ := actions.Get(attackMode.ElementalBurst, "伤害1")
	纳西妲.Apply(
		buff.Character纳西妲净善摄受明论(1000),                                       // 草神固有天赋4
		buff.AddAttackDamageBonus(actionQ.DMG(), attackMode.ElementalSkill), // Q增伤
	)

	挨揍的 := enemy.New(enemy.Base(90))
	//挨揍的.Apply(buff.AddElementalResist(-30, elementals.Grass))
	挨揍的.Attach(elementals.Water, 12)
	//挨揍的.AttachState(states.Quicken, 12)
	攻击 := actions.Get(attackMode.ElementalSkill, "灭净三业")
	profitDetect(纳西妲, 挨揍的, 攻击, damage, CustomDetects(elementals.Grass), nil, buff.Character纳西妲慧明缘觉智论())
}

func CustomDetects(dye elementals.Elemental) map[string]*attr.Modifier {
	return map[string]*attr.Modifier{
		"钟离+岩四件套":     attr.NewModifier(attr.MergeAttributes(buff.AddAtkPercentage(20), buff.Superposition(5, time.Second*20, 0, buff.AddShieldStrength(5))), buff.AddAllElementalResist(-20)),
		"钟离":          attr.NewModifier(buff.Superposition(5, time.Second*20, 0, buff.AddShieldStrength(5)), buff.AddAllElementalResist(-20)),
		"万叶":          attr.NewCharacterModifier(buff.Character万叶扩散(1000, dye)),
		"风四件套":        attr.NewEnemyModifier(buff.Artifacts翠绿之影4(dye).EnemyModifier()),
		"万叶+风四件套":     attr.NewModifier(buff.Character万叶扩散(1000, dye), buff.Artifacts翠绿之影4(dye).EnemyModifier()),
		"班尼特+宗室四件套":   attr.NewCharacterModifier(buff.AddAtkPercentage(20), buff.AddAtk(int(math.Round(1.19*(191+565))))),
		"班尼特6命+宗室四件套": attr.NewCharacterModifier(buff.AddAtkPercentage(20), buff.AddAtk(int(math.Round(1.19*(191+565)))), buff.AddElementalDamageBonus(15, elementals.Fire)).Action(action.Infusion(elementals.Fire)),
		"讨龙英杰谭":       attr.NewCharacterModifier(buff.AddAtkPercentage(48)),
		"砂糖":          attr.NewCharacterModifier(buff.AddElementalMastery(50 + 200)),
		"砂糖+风四件套":     attr.NewModifier(buff.AddElementalMastery(50+200), buff.Artifacts翠绿之影4(dye).EnemyModifier()),
		"砂糖6命":        attr.NewCharacterModifier(buff.AddElementalMastery(50+200), buff.AddElementalDamageBonus(20, dye)),
		"莫娜":          attr.NewCharacterModifier(buff.AddDamageBonus(60)),
		"莫娜+讨龙":       attr.NewCharacterModifier(buff.AddDamageBonus(60), buff.AddAtkPercentage(48)),
		"岩四件套":        attr.NewCharacterModifier(buff.AddAtkPercentage(20)),
		"岩主Q":         attr.NewCharacterModifier(buff.AddCriticalRate(15)),
		"深林四件套":       attr.NewEnemyModifier(buff.AddElementalResist(-30, elementals.Grass)),
		"减防30":        attr.NewEnemyModifier(buff.AddDefPercentage(-30)),
		"如雷四件套":       attr.NewCharacterModifier(buff.AddReactionDamageBonus(40, reactions.Overload, reactions.ElectroCharged, reactions.Superconduct, reactions.Hyperbloom)),
		"九条":          attr.NewCharacterModifier(buff.AddAtk(600)),
		"九条6命":        attr.NewCharacterModifier(buff.AddAtk(600), buff.Character九条裟罗六命(dye)),
		"雷神2命":        attr.NewCharacterModifier(buff.AddIgnoreDefence(60)),
	}
}

func profitDetect(character *character.Character, enemy *enemy.Enemy, action *action.Action, fn detect.FinalDamage,
	customDetects map[string]*attr.Modifier, replaceArtifacts []*artifacts.Artifacts, finalModifiers ...attr.AttributeModifier) {
	fmt.Printf("action: %s\n", action)
	fmt.Printf("base: %f\n", fn(character, enemy, action, true, finalModifiers...))
	profits := detect.ProfitDetect(character, enemy, action, true, fn, nil, finalModifiers...)
	fmt.Printf("素质增益:\n")
	for _, p := range profits {
		fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
	}
	if len(customDetects) > 0 {
		profits = detect.ProfitDetect(character, enemy, action, false, fn, customDetects, finalModifiers...)
		fmt.Printf("队友增益:\n")
		for _, p := range profits {
			fmt.Printf("\t增幅：%2.4f%% - %s\n", p.Value, p.Name)
		}
	}
	if evaluateDetects := character.Evaluate(replaceArtifacts...); len(evaluateDetects) > 0 {
		profits = detect.ProfitDetect(character, enemy, action, false, fn, evaluateDetects, finalModifiers...)
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
