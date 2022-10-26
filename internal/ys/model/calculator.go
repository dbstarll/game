package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"go.uber.org/zap"
)

type Calculator struct {
	finalAttributes *attr.Attributes
	enemy           *enemy.Enemy
	action          *action.Action
	elemental       elementals.Elemental
	values          *Values
	init            map[string]float64
}

func NewCalculator(character *Character, enemy *enemy.Enemy, action *action.Action, infusionElemental elementals.Elemental) *Calculator {
	calculator := &Calculator{
		finalAttributes: character.finalAttributes(),
		enemy:           enemy,
		action:          action,
		elemental:       action.Elemental(),
		init: map[string]float64{
			"人物等级":  float64(character.level),
			"人物攻击力": character.base.Get(point.Atk),
			"武器攻击力": character.weapon.Get(point.Atk),
		},
	}
	switch action.Mode() {
	case attackMode.NormalAttack, attackMode.ChargedAttack, attackMode.PlungeAttack:
		calculator.elemental = calculator.elemental.Infusion(infusionElemental)
		break
	}
	return calculator
}

func (c *Calculator) Get(key string) float64 {
	if v, exist := c.values.Get(key); exist && v != nil {
		return v.value
	} else {
		return 0
	}
}

func (c *Calculator) set(key string, value float64) *Formula {
	return c.values.Set(key, value)
}

func (c *Calculator) add(totalKey string, objs ...interface{}) *Formula {
	return c.values.Add(totalKey, objs...)
}

func (c *Calculator) reduce(totalKey string, objs ...interface{}) *Formula {
	return c.values.Reduce(totalKey, objs...)
}

func (c *Calculator) multiply(totalKey string, objs ...interface{}) *Formula {
	return c.values.Multiply(totalKey, objs...)
}

func (c *Calculator) divide(totalKey string, objs ...interface{}) *Formula {
	return c.values.Divide(totalKey, objs...)
}

func (c *Calculator) prepare(putZero bool) {
	c.values = NewValues()
	for _, p := range point.Points {
		if v := c.finalAttributes.Get(p); putZero || v != 0 {
			if p.IsPercentage() {
				c.set(p.String(), v/100)
			} else {
				c.set(p.String(), v)
			}
		}
	}
	for _, ele := range elementals.Elementals {
		if v := c.finalAttributes.GetElementalDamageBonus(ele); putZero || v != 0 {
			c.set(fmt.Sprintf("%s伤害加成", ele.Name()), v/100)
		}
		if v := c.finalAttributes.GetElementalResist(ele); putZero || v != 0 {
			c.set(fmt.Sprintf("%s抗性", ele.Name()), v/100)
		}
		if ele != elementals.Physical {
			// 没有物理影响下增伤
			if v := c.finalAttributes.GetElementalAttachedDamageBonus(ele); putZero || v != 0 {
				c.set(fmt.Sprintf("%s影响下增伤", ele.Name()), v/100)
			}
		}
	}
	for _, ra := range reactions.Reactions {
		if v := c.finalAttributes.GetReactionDamageBonus(ra); putZero || v != 0 {
			switch ra.Classify() {
			case classifies.Amplify, classifies.Intensify:
				c.set(fmt.Sprintf("%s反应系数提高", ra), v/100)
				break
			case classifies.Upheaval:
				c.set(fmt.Sprintf("%s反应伤害提升", ra), v/100)
				break
				//TODO
				// Crystal                   // 结晶
				// Intensify                 // 激化
			}
		}
	}
	for _, mode := range attackMode.AttackModes {
		if v := c.finalAttributes.GetAttackDamageBonus(mode); putZero || v != 0 {
			c.set(fmt.Sprintf("%s伤害加成", mode), v/100)
		}
		if v := c.finalAttributes.GetAttackFactorBonus(mode); putZero || v != 0 {
			c.set(fmt.Sprintf("%s技能倍率加成", mode), v/100)
		}
	}

	for key, value := range c.init {
		if putZero || value != 0 {
			c.set(key, value)
		}
	}
	if c.Get("暴击率") > 1.0 {
		c.set("暴击率", 1.0)
	}
}

func (c *Calculator) Calculate(debug bool) (*Formula, *Formula, *Formula) {
	c.prepare(true)

	基础伤害区, 增伤区, 防御区, 抗性区, 增幅区 := c.基础伤害区(), c.增伤区(), c.防御区(), c.抗性区(c.elemental), c.增幅区()
	暴击收益, 暴伤倍率 := c.暴击区()
	增幅总伤害 := 基础伤害区.multiply("增幅总伤害", 增伤区, 防御区, 抗性区, 增幅区)
	增幅总伤害平均 := 增幅总伤害.multiply("增幅总伤害(平均)", 暴击收益)
	增幅总伤害最大 := 增幅总伤害.multiply("增幅总伤害(暴击)", 暴伤倍率)

	剧变区 := c.剧变区(抗性区)
	总伤害 := 增幅总伤害.add("总伤害", 剧变区...)
	总伤害平均 := 增幅总伤害平均.add("总伤害(平均)", 剧变区...)
	总伤害最大 := 增幅总伤害最大.add("总伤害(暴击)", 剧变区...)
	if debug {
		zap.S().Debugf("Action: %s", c.action)
		zap.S().Debugf("Attributes: %s", c.finalAttributes)
		zap.S().Debugf("基础伤害区: %s", 基础伤害区.Algorithm())
		zap.S().Debugf("增伤区: %s", 增伤区.Algorithm())
		zap.S().Debugf("暴击区: %s, %s", 暴击收益.Algorithm(), 暴伤倍率.Algorithm())
		zap.S().Debugf("防御区: %s", 防御区.Algorithm())
		zap.S().Debugf("抗性区: %s", 抗性区.Algorithm())
		zap.S().Debugf("增幅区: %s", 增幅区.Algorithm())
		zap.S().Debugf("剧变区: %s", 剧变区)
		zap.S().Debugf("%s", 增幅总伤害.Algorithm())
		zap.S().Debugf("%s", 增幅总伤害平均.Algorithm())
		zap.S().Debugf("%s", 增幅总伤害最大.Algorithm())
		zap.S().Debugf("%s", 总伤害.Algorithm())
		zap.S().Debugf("%s", 总伤害平均.Algorithm())
		zap.S().Debugf("%s", 总伤害最大.Algorithm())
	}
	return 总伤害, 总伤害平均, 总伤害最大
}

func (c *Calculator) 攻击区() *Formula {
	基础攻击力 := c.add("基础攻击力", "人物攻击力", "武器攻击力")
	return 基础攻击力.add("总攻击力", 基础攻击力.multiply("百分比攻击力", "攻击力%").add("额外攻击力", "攻击力"))
}

func (c *Calculator) 倍率区() *Formula {
	prefix := c.action.Mode().String()
	return c.set(prefix+"技能倍率", c.action.DMG()/100).multiply(prefix+"伤害倍率", c.add(prefix+"技能倍率增伤", 1, prefix+"技能倍率加成"))
}

func (c *Calculator) 基础倍率区() *Formula {
	return c.攻击区().multiply("基础倍率", c.倍率区())
}

func (c *Calculator) 激化区() *Formula {
	for _, react := range c.enemy.DetectStateReaction(c.elemental, classifies.Intensify) {
		激化等级系数 := c.set("激化等级系数", 1446.85)
		激化精通提升 := c.multiply("激化精通系数1", 5, "元素精通").divide("激化精通提升", c.add("激化精通系数2", 1200, "元素精通"))
		reactionName := react.Reaction.String()
		激化反应倍率 := c.add(reactionName+"反应倍率", 1, 激化精通提升, reactionName+"反应系数提高")
		return c.set(reactionName+"反应基础倍率", react.Factor).multiply(reactionName+"反应伤害", 激化等级系数, 激化反应倍率)
	}
	return c.set("无激化加成", 0)
}

func (c *Calculator) 基础伤害区() *Formula {
	return c.基础倍率区().add("基础伤害", c.激化区())
}

func (c *Calculator) 增伤区() *Formula {
	prefix := c.action.Mode().String()
	objs := []interface{}{1, fmt.Sprintf("%s伤害加成", c.elemental.Name()), prefix + "伤害加成", "伤害加成"}
	for _, element := range c.enemy.Attached() {
		objs = append(objs, fmt.Sprintf("%s影响下增伤", element.Name()))
	}
	return c.add(prefix+"增伤", objs...)
}

func (c *Calculator) 暴击区() (*Formula, *Formula) {
	return c.add("暴击收益", 1, c.multiply("暴击加成", "暴击率", "暴击伤害")), c.add("暴伤倍率", 1, "暴击伤害")
}

func (c *Calculator) 防御区() *Formula {
	怪物等级系数 := c.set("怪物等级", float64(c.enemy.Level())).add("怪物等级系数", 100)
	怪物防御 := c.set("怪物防御%", c.enemy.Get(point.DefPercentage)/100)
	人物等级系数 := c.add("人物等级系数", "人物等级", 100)
	减防系数 := c.add("怪物防御系数", 1, 怪物防御).reduce("减防系数", "防御减免")
	防御承伤基准 := c.reduce("穿防系数", 1, "无视防御").multiply("防御系数", 减防系数, 怪物等级系数).add("防御承伤基准", 人物等级系数)
	return 人物等级系数.divide("防御承伤", 防御承伤基准)
}

func (c *Calculator) 抗性区(elemental elementals.Elemental) *Formula {
	prefix := fmt.Sprintf("%s抗性", elemental.Name())
	抗性 := c.set("怪物"+prefix, c.enemy.GetElementalResist(elemental)/100)
	if 抗性.value > 0.75 {
		return c.divide(prefix+"承伤", 1, 抗性.multiply("怪物"+prefix+"系数1", 4).add("怪物"+prefix+"系数2", 1))
	} else if 抗性.value >= 0 {
		return c.reduce(prefix+"承伤", 1, 抗性)
	} else {
		return c.reduce(prefix+"承伤", 1, 抗性.divide("怪物"+prefix+"系数3", 2))
	}
}

func (c *Calculator) 增幅区() *Formula {
	增幅精通提升 := c.multiply("增幅精通系数1", 25.0/9, "元素精通").divide("增幅精通提升", c.add("增幅精通系数2", 1400, "元素精通"))
	for _, factor := range c.enemy.DetectReaction(c.elemental, classifies.Amplify) {
		reactionName := factor.GetReaction().String()
		增幅反应倍率 := c.add("增幅反应倍率", 1, 增幅精通提升, reactionName+"反应系数提高")
		return c.set(reactionName+"反应基础倍率", factor.GetFactor()).multiply(reactionName+"反应总倍率", 增幅反应倍率)
	}
	return c.set("无增幅反应", 1)
}

func (c *Calculator) 剧变区(抗性承伤 *Formula) []interface{} {
	damages := make([]interface{}, 0)
	剧变等级系数 := c.set("剧变等级系数", 1446.85)
	剧变精通提升 := c.multiply("剧变精通系数1", 16, "元素精通").divide("剧变精通提升", c.add("剧变精通系数2", 2000, "元素精通"))
	for _, factor := range c.enemy.DetectReaction(c.elemental, classifies.Upheaval) {
		reactionName := factor.GetReaction().String()
		剧变反应倍率 := c.add(reactionName+"反应倍率", 1, 剧变精通提升, reactionName+"反应伤害提升")
		damages = append(damages, c.set(reactionName+"反应基础倍率", factor.GetFactor()).multiply(reactionName+"反应伤害", 剧变等级系数, 抗性承伤, 剧变反应倍率))
	}
	for _, factor := range c.enemy.DetectStateReaction(c.elemental, classifies.Upheaval) {
		reactionName := factor.Reaction.String()
		剧变反应倍率 := c.add(reactionName+"反应倍率", 1, 剧变精通提升, reactionName+"反应伤害提升")
		if factor.Elemental.IsValid() && factor.Elemental != c.elemental {
			状态抗性承伤 := c.抗性区(factor.Elemental)
			damages = append(damages, c.set(reactionName+"反应基础倍率", factor.Factor).multiply(reactionName+"反应伤害", 剧变等级系数, 状态抗性承伤, 剧变反应倍率))
		} else {
			damages = append(damages, c.set(reactionName+"反应基础倍率", factor.Factor).multiply(reactionName+"反应伤害", 剧变等级系数, 抗性承伤, 剧变反应倍率))
		}
	}
	return damages
}

func (c *Calculator) String() string {
	return fmt.Sprintf("%s", c.values)
}
