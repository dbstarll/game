package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

type Calculator struct {
	finalAttributes *attr.Attributes
	enemy           *Enemy
	action          *Action
	elemental       elemental.Elemental
	values          *Values
	init            map[string]float64
}

func NewCalculator(character *Character, enemy *Enemy, action *Action, infusionElemental elemental.Elemental) *Calculator {
	calculator := &Calculator{
		finalAttributes: character.finalAttributes(),
		enemy:           enemy,
		action:          action,
		elemental:       action.elemental,
		init: map[string]float64{
			"人物等级":  float64(character.level),
			"人物攻击力": character.base.Get(point.Atk).GetValue(),
			"武器攻击力": character.weapon.base.Get(point.Atk).GetValue(),
		},
	}
	switch action.mode {
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
		if v := c.finalAttributes.Get(p); putZero || !v.IsZero() {
			if p.IsPercentage() {
				c.set(p.String(), v.GetValue()/100)
			} else {
				c.set(p.String(), v.GetValue())
			}
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

func (c *Calculator) Calculate() (*Formula, *Formula, *Formula) {
	c.prepare(true)

	基础伤害区, 增伤区, 防御区, 抗性区, 增幅区 := c.基础伤害区(), c.增伤区(), c.防御区(), c.抗性区(), c.增幅区()
	暴击收益, 暴伤倍率 := c.暴击区()
	增幅总伤害 := 基础伤害区.multiply("增幅总伤害", 增伤区, 防御区, 抗性区, 增幅区)
	增幅总伤害平均 := 增幅总伤害.multiply("增幅总伤害(平均)", 暴击收益)
	增幅总伤害最大 := 增幅总伤害.multiply("增幅总伤害(暴击)", 暴伤倍率)

	剧变区 := c.剧变区(抗性区)
	总伤害 := 增幅总伤害.add("总伤害", 剧变区...)
	总伤害平均 := 增幅总伤害平均.add("总伤害(平均)", 剧变区...)
	总伤害最大 := 增幅总伤害最大.add("总伤害(暴击)", 剧变区...)
	//zap.S().Debugf("基础伤害区: %s", 基础伤害区.Algorithm())
	//zap.S().Debugf("增伤区: %s", 增伤区.Algorithm())
	//zap.S().Debugf("暴击区: %s, %s", 暴击收益.Algorithm(), 暴伤倍率.Algorithm())
	//zap.S().Debugf("防御区: %s", 防御区.Algorithm())
	//zap.S().Debugf("抗性区: %s", 抗性区.Algorithm())
	//zap.S().Debugf("增幅区: %s", 增幅区.Algorithm())
	//zap.S().Debugf("剧变区: %s", 剧变区)
	return 总伤害, 总伤害平均, 总伤害最大
}

func (c *Calculator) 攻击区() *Formula {
	基础攻击力 := c.add("基础攻击力", "人物攻击力", "武器攻击力")
	return 基础攻击力.add("总攻击力", 基础攻击力.multiply("百分比攻击力", "攻击力%").add("额外攻击力", "攻击力"))
}

func (c *Calculator) 倍率区() *Formula {
	prefix := c.action.mode.String()
	//c.set(prefix+"技能倍率加成", 0.5)
	return c.set(prefix+"技能倍率", c.action.dmg/100).multiply(prefix+"伤害倍率", c.add(prefix+"技能倍率增伤", 1, prefix+"技能倍率加成"))
}

func (c *Calculator) 基础倍率区() *Formula {
	return c.攻击区().multiply("基础倍率", c.倍率区())
}

func (c *Calculator) 激化区() *Formula {
	//TODO 待完善
	return c.set("激化加成值", 0)
}

func (c *Calculator) 基础伤害区() *Formula {
	return c.基础倍率区().add("基础伤害", c.激化区())
}

func (c *Calculator) 增伤区() *Formula {
	prefix := c.action.mode.String()
	//TODO 对元素影响下的敌人伤害提高
	return c.add(prefix+"增伤", 1, c.elemental.DamageBonusPoint().String(), prefix+"伤害加成", "元素影响增伤", "伤害加成")
}

func (c *Calculator) 暴击区() (*Formula, *Formula) {
	return c.add("暴击收益", 1, c.multiply("暴击加成", "暴击率", "暴击伤害")), c.add("暴伤倍率", 1, "暴击伤害")
}

func (c *Calculator) 防御区() *Formula {
	怪物等级系数 := c.set("怪物等级", float64(c.enemy.level)).add("怪物等级系数", 100)
	怪物防御 := c.set("怪物防御%", c.enemy.base.Get(point.DefPercentage).GetValue()/100)
	人物等级系数 := c.add("人物等级系数", "人物等级", 100)
	减防系数 := c.add("怪物防御系数", 1, 怪物防御).reduce("减防系数", "防御减免")
	防御承伤基准 := c.reduce("穿防系数", 1, "无视防御").multiply("防御系数", 减防系数, 怪物等级系数).add("防御承伤基准", 人物等级系数)
	return 人物等级系数.divide("防御承伤", 防御承伤基准)
}

func (c *Calculator) 抗性区() *Formula {
	prefix := c.elemental.ResistPoint().String()
	抗性 := c.set("怪物"+prefix, c.enemy.base.Get(c.elemental.ResistPoint()).GetValue()/100)
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
	for _, factor := range c.enemy.DetectReaction(c.elemental, reaction.Amplify) {
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
	for _, factor := range c.enemy.DetectReaction(c.elemental, reaction.Upheaval) {
		reactionName := factor.GetReaction().String()
		剧变反应倍率 := c.add(reactionName+"反应倍率", 1, 剧变精通提升, reactionName+"反应伤害提升")
		damages = append(damages, c.set(reactionName+"反应基础倍率", factor.GetFactor()).multiply(reactionName+"反应伤害", 剧变等级系数, 抗性承伤, 剧变反应倍率))
	}
	return damages
}

func (c *Calculator) String() string {
	return fmt.Sprintf("%s", c.values)
}
