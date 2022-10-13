package detect

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"sort"
)

type Profit struct {
	Name  string
	Value float64
}

type Modifier struct {
	characterModifier attr.AttributeModifier
	enemyModifier     attr.AttributeModifier
}

type FinalDamage func(character *model.Character, enemy *enemy.Enemy, debug bool) float64

var (
	baseDetects = initBaseDetects(map[string]*Modifier{
		point.Hp.String():               NewCharacterModifier(buff.AddHp(209)),              // 生命值
		point.HpPercentage.String():     NewCharacterModifier(buff.AddHpPercentage(4.1)),    // 生命值%
		point.Atk.String():              NewCharacterModifier(buff.AddAtk(14)),              // 攻击力
		point.AtkPercentage.String():    NewCharacterModifier(buff.AddAtkPercentage(4.1)),   // 攻击力%
		point.Def.String():              NewCharacterModifier(buff.AddDef(16)),              // 防御力
		point.DefPercentage.String():    NewCharacterModifier(buff.AddDefPercentage(5.1)),   // 防御力%
		point.ElementalMastery.String(): NewCharacterModifier(buff.AddElementalMastery(16)), // 元素精通
		point.CriticalRate.String():     NewCharacterModifier(buff.AddCriticalRate(2.7)),    // 暴击率
		point.CriticalDamage.String():   NewCharacterModifier(buff.AddCriticalDamage(5.4)),  // 暴击伤害
		point.HealingBonus.String():     NewCharacterModifier(buff.AddHealingBonus(3.1)),    // 治疗加成
		//IncomingHealingBonus                   // 受治疗加成
		point.EnergyRecharge.String(): NewCharacterModifier(buff.AddEnergyRecharge(4.5)), // 元素充能效率
		//CDReduction                            // 冷却缩减
		//ShieldStrength                         // 护盾强效
		//DamageBonus                            // 伤害加成
		//IncomingDamageBonus                    // 受到的伤害加成
		//IgnoreDefence                          // 无视防御
		//DefenceReduction                       // 防御减免
		//NormalAttackDamageBonus                // 普通攻击伤害加成
		//ChargedAttackDamageBonus               // 重击伤害加成
		//PlungeAttackDamageBonus                // 下坠攻击伤害加成
		//ElementalSkillDamageBonus              // 元素战技伤害加成
		//ElementalBurstDamageBonus              // 元素爆发伤害加成
		//NormalAttackFactorBonus                // 普通攻击技能倍率加成
		//ChargedAttackFactorBonus               // 重击技能倍率加成
		//PlungeAttackFactorBonus                // 下坠攻击技能倍率加成
		//ElementalSkillFactorBonus              // 元素战技技能倍率加成
		//ElementalBurstFactorBonus              // 元素爆发技能倍率加成

	})
)

func NewCharacterModifier(characterModifier attr.AttributeModifier) *Modifier {
	return NewModifier(characterModifier, nil)
}

func NewEnemyModifier(enemyModifier attr.AttributeModifier) *Modifier {
	return NewModifier(nil, enemyModifier)
}

func NewModifier(characterModifier, enemyModifier attr.AttributeModifier) *Modifier {
	return &Modifier{
		characterModifier: characterModifier,
		enemyModifier:     enemyModifier,
	}
}

func initBaseDetects(detects map[string]*Modifier) map[string]*Modifier {
	// TODO
	//   元素抗性
	//   元素影响下增伤
	for _, ele := range elementals.Elements {
		detects[fmt.Sprintf("%s伤害加成", ele.Name())] = NewCharacterModifier(buff.AddElementalDamageBonus(4.1, ele))
	}
	return detects
}

func (m *Modifier) Apply(character *model.Character, enemy *enemy.Enemy) func() {
	var cancels []func()
	if m.characterModifier != nil {
		cancels = append(cancels, character.Apply(m.characterModifier))
	}
	if m.enemyModifier != nil {
		cancels = append(cancels, enemy.Apply(m.enemyModifier))
	}
	return func() {
		for _, cancel := range cancels {
			cancel()
		}
	}
}

func ProfitDetect(character *model.Character, enemy *enemy.Enemy, baseDetect bool, fn FinalDamage, customDetects map[string]*Modifier) []*Profit {
	base := fn(character, enemy, false)
	var profits []*Profit
	if baseDetect {
		for name, modifier := range baseDetects {
			cancel := modifier.Apply(character, enemy)
			value := fn(character, enemy, false)
			if value != base {
				profits = append(profits, &Profit{
					Name:  name,
					Value: 100 * (value - base) / base,
				})
			}
			cancel()
		}
	}
	for name, modifier := range customDetects {
		cancel := modifier.Apply(character, enemy)
		value := fn(character, enemy, false)
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
	return profits
}
