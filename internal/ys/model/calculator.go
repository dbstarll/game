package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"go.uber.org/zap"
)

type Calculator struct {
	finalAttributes   *Attributes
	enemy             *Enemy
	action            *Action
	infusionElemental elemental.Elemental
	values            *Values
	init              map[string]float64
}

func NewCalculator(character *Character, enemy *Enemy, action *Action, infusionElemental elemental.Elemental) *Calculator {
	calculator := &Calculator{
		finalAttributes:   character.finalAttributes(),
		enemy:             enemy,
		action:            action,
		infusionElemental: infusionElemental,
		init: map[string]float64{
			"人物等级":  float64(character.level),
			"人物攻击力": character.base.Get(point.Atk).value,
			"武器攻击力": character.weapon.base.Get(point.Atk).value,
		},
	}
	calculator.calculate()
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
				c.set(p.String(), v.value/100)
			} else {
				c.set(p.String(), v.value)
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

func (c *Calculator) calculate() {
	c.prepare(true)

	zap.S().Debugf("Action: %s", c.action)
	zap.S().Debugf("Elemental: %s + %s = %s", c.action.elemental, c.infusionElemental, c.action.elemental.Infusion(c.infusionElemental))
	zap.S().Debugf("DamageBonusPoint: %s", c.action.elemental.Infusion(c.infusionElemental).DamageBonusPoint())

	基础伤害区, 增伤区, 防御区, 抗性区, 增幅区 := c.基础伤害区(), c.增伤区(), c.防御区(), c.抗性区(), c.增幅区()
	暴击收益, 暴伤倍率 := c.暴击区()
	总伤害 := 基础伤害区.multiply("总伤害", 增伤区, 防御区, 抗性区, 增幅区)
	总伤害平均 := 总伤害.multiply("总伤害(平均)", 暴击收益)
	总伤害最大 := 总伤害.multiply("总伤害(暴击)", 暴伤倍率)
	zap.S().Debugf("基础伤害区: %s", 基础伤害区.Algorithm())
	zap.S().Debugf("增伤区: %s", 增伤区.Algorithm())
	zap.S().Debugf("暴击区: %s, %s", 暴击收益.Algorithm(), 暴伤倍率.Algorithm())
	zap.S().Debugf("防御区: %s", 防御区.Algorithm())
	zap.S().Debugf("抗性区: %s", 抗性区.Algorithm())
	zap.S().Debugf("增幅区: %s", 增幅区.Algorithm())
	zap.S().Debugf("%s", 总伤害.Algorithm())
	zap.S().Debugf("%s", 总伤害平均.Algorithm())
	zap.S().Debugf("%s", 总伤害最大.Algorithm())
	zap.S().Debugf("总伤害: [%f, %f, %f]", 总伤害.value, 总伤害平均.value, 总伤害最大.value)

	//            // 抗性区
	//            const phyR = c.set("元素抗性",c.Get("怪物元素抗性") - c.Get("元素抗性减免"));
	//            if (phyR > 0.75) {
	//                c.set("元素抗性承伤",1 / (1 + 4 * phyR));
	//            } else if (phyR >= 0) {
	//                c.set("元素抗性承伤",1 - phyR);
	//            } else {
	//                c.set("元素抗性承伤",1 - phyR / 2);
	//            }
	//            const eleR = c.set("物理抗性",c.Get("怪物物理抗性") - c.Get("物理抗性减免"));
	//            if (eleR > 0.75) {
	//                c.set("物理抗性承伤",1 / (1 + 4 * eleR));
	//            } else if (eleR >= 0) {
	//                c.set("物理抗性承伤",1 - eleR);
	//            } else {
	//                c.set("物理抗性承伤",1 - eleR / 2);
	//            }
	//
	//            // 技能伤害
	//            c.set("普攻伤害",c.Get("基础伤害") * c.Get("普攻倍率") * c.Get("普通攻击增伤") * c.Get("物理抗性承伤"));
	//            c.set("普攻伤害(平均)",c.Get("普攻伤害") * c.Get("暴击收益"));
	//            c.set("普攻伤害(暴击)",c.Get("普攻伤害") * c.Get("暴伤倍率"));
	//            c.set("重击伤害",c.Get("基础伤害") * c.Get("重击倍率") * c.Get("重击增伤") * c.Get("物理抗性承伤"));
	//            c.set("重击伤害(平均)",c.Get("重击伤害") * c.Get("暴击收益"));
	//            c.set("重击伤害(暴击)",c.Get("重击伤害") * c.Get("暴伤倍率"));
	//            c.set("附魔普攻伤害",c.Get("基础伤害") * c.Get("普攻倍率") * c.Get("附魔普通攻击增伤") * c.Get("元素抗性承伤"));
	//            c.set("附魔普攻伤害(平均)",c.Get("附魔普攻伤害") * c.Get("暴击收益"));
	//            c.set("附魔普攻伤害(暴击)",c.Get("附魔普攻伤害") * c.Get("暴伤倍率"));
	//            c.set("附魔重击伤害",c.Get("基础伤害") * c.Get("重击倍率") * c.Get("附魔重击增伤") * c.Get("元素抗性承伤"));
	//            c.set("附魔重击伤害(平均)",c.Get("附魔重击伤害") * c.Get("暴击收益"));
	//            c.set("附魔重击伤害(暴击)",c.Get("附魔重击伤害") * c.Get("暴伤倍率"));
	//            c.set("元素战技伤害",c.Get("基础伤害") * c.Get("元素战技倍率") * c.Get("元素战技增伤") * c.Get("元素抗性承伤"));
	//            c.set("元素战技伤害(平均)",c.Get("元素战技伤害") * c.Get("暴击收益"));
	//            c.set("元素战技伤害(暴击)",c.Get("元素战技伤害") * c.Get("暴伤倍率"));
	//            c.set("元素爆发伤害",c.Get("基础伤害") * c.Get("元素爆发倍率") * c.Get("元素爆发增伤") * c.Get("元素抗性承伤"));
	//            c.set("元素爆发伤害(平均)",c.Get("元素爆发伤害") * c.Get("暴击收益"));
	//            c.set("元素爆发伤害(暴击)",c.Get("元素爆发伤害") * c.Get("暴伤倍率"));
	//
	//            // 增幅反应区
	//            c.set("增幅精通提升",(2.78 * c.Get("元素精通")) / (c.Get("元素精通") + 1400));
	//            c.set("增幅反应倍率",1 + c.Get("增幅精通提升") + c.Get("反应系数提高"));
	//
	//            // 剧变反应区
	//            c.set("剧变精通提升",16 * c.Get("元素精通") / (c.Get("元素精通") + 2000));
	//            c.set("剧变反应倍率",1 + c.Get("剧变精通提升") + c.Get("反应伤害提升"));
	//            c.set("等级系数",upheavals[c.Get("人物等级") - 1]);
	//            c.set("剧变基础伤害",c.Get("等级系数") * c.Get("元素抗性承伤") * c.Get("剧变反应倍率"));
	//
	//            // 反应伤害
	//            const element = $("select[column=元素类型]").val());
	//            const bonus = reactions.bonus[element]);
	//            if ("object" === typeof bonus) {
	//                Object.keys(bonus).forEach((key) => {
	//                    c.set("附魔普攻伤害(" + key + ")",c.Get("附魔普攻伤害") * c.Get("增幅反应倍率") * bonus[key].rate);
	//                    c.set("附魔重击伤害(" + key + ")",c.Get("附魔重击伤害") * c.Get("增幅反应倍率") * bonus[key].rate);
	//                    c.set("元素战技伤害(" + key + ")",c.Get("元素战技伤害") * c.Get("增幅反应倍率") * bonus[key].rate);
	//                    c.set("元素爆发伤害(" + key + ")",c.Get("元素爆发伤害") * c.Get("增幅反应倍率") * bonus[key].rate);
	//                    c.set("附魔普攻伤害(" + key.substr(0, 1) + "暴)",c.Get("附魔普攻伤害(" + key + ")") * c.Get("暴伤倍率"));
	//                    c.set("附魔重击伤害(" + key.substr(0, 1) + "暴)",c.Get("附魔重击伤害(" + key + ")") * c.Get("暴伤倍率"));
	//                    c.set("元素战技伤害(" + key.substr(0, 1) + "暴)",c.Get("元素战技伤害(" + key + ")") * c.Get("暴伤倍率"));
	//                    c.set("元素爆发伤害(" + key.substr(0, 1) + "暴)",c.Get("元素爆发伤害(" + key + ")") * c.Get("暴伤倍率"));
	//                });
	//            }
	//            const upheaval = reactions.upheaval[element];
	//            if ("object" === typeof upheaval) {
	//                Object.keys(upheaval).forEach((key) => {
	//                    c.set("剧变伤害(" + key + ")",c.Get("剧变基础伤害") * reactions.upheaval.rate[key]);
	//                });
	//            }
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
	prefix, elemental := c.action.mode.String(), c.action.elemental
	switch c.action.mode {
	case attackMode.NormalAttack, attackMode.ChargedAttack, attackMode.PlungeAttack:
		elemental = elemental.Infusion(c.infusionElemental)
		break
	}
	//TODO 对元素影响下的敌人伤害提高
	return c.add(prefix+"增伤", 1, elemental.DamageBonusPoint().String(), prefix+"伤害加成", "元素影响增伤", "伤害加成")
}

func (c *Calculator) 暴击区() (*Formula, *Formula) {
	return c.add("暴击收益", 1, c.multiply("暴击加成", "暴击率", "暴击伤害")), c.add("暴伤倍率", 1, "暴击伤害")
}

func (c *Calculator) 防御区() *Formula {
	怪物等级系数 := c.set("怪物等级", float64(c.enemy.level)).add("怪物等级系数", 100)
	怪物防御 := c.set("怪物防御%", c.enemy.base.Get(point.DefPercentage).value/100)
	人物等级系数 := c.add("人物等级系数", "人物等级", 100)
	减防系数 := c.add("怪物防御系数", 1, 怪物防御).reduce("减防系数", "防御减免")
	防御承伤基准 := c.reduce("穿防系数", 1, "无视防御").multiply("防御系数", 减防系数, 怪物等级系数).add("防御承伤基准", 人物等级系数)
	return 人物等级系数.divide("防御承伤", 防御承伤基准)
}

func (c *Calculator) 抗性区() *Formula {
	//TODO 待完善
	return c.set("抗性区", 1)
}

func (c *Calculator) 增幅区() *Formula {
	//TODO 待完善
	return c.set("增幅区", 1)
}

func (c *Calculator) String() string {
	return c.values.String()
}
