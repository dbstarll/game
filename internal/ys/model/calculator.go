package model

import (
	"fmt"
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
			"人物等级": float64(character.level),
			//"怪物等级":  float64(enemy.level),
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

func (c *Calculator) multiply(totalKey string, objs ...interface{}) *Formula {
	return c.values.Multiply(totalKey, objs...)
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
}

func (c *Calculator) calculate() {
	c.prepare(true)

	zap.S().Debugf("Action: %s", c.action)
	zap.S().Debugf("Elemental: %s + %s = %s", c.action.elemental, c.infusionElemental, c.action.elemental.Infusion(c.infusionElemental))
	zap.S().Debugf("DamageBonusPoint: %s", c.action.elemental.Infusion(c.infusionElemental).DamageBonusPoint())
	zap.S().Debugf("攻击区: %s", c.攻击区().Algorithm())

	// 基础伤害区 = 攻击区 * 倍率区
	//伤害 = (基础伤害区 + 激化区) * 增伤区 * 暴击区 * 增幅区 * 防御区 * 抗性区

	//c.set("怪物防御%", c.enemy.base.Get(point.DefPercentage).value/100)
	//c.set("基础攻击力", c.basicAttributes.Get(point.Atk).value)
	//
	//// 基础区
	//c.set("额外攻击力", c.Get("基础攻击力")*c.Get("攻击力%")+c.Get("攻击力"))
	//c.set("总攻击力", c.Get("基础攻击力")+c.Get("额外攻击力"))
	//
	//// 暴击区
	//if c.Get("暴击率") > 1.0 {
	//	c.set("暴击率", 1.0)
	//}
	//c.set("暴击收益", 1+c.Get("暴击率")*c.Get("暴击伤害"))
	//c.set("暴伤倍率", 1+c.Get("暴击伤害"))
	//
	//// 防御区
	//c.set("穿防系数", 1-c.Get("无视防御"))
	//c.set("减防系数", 1-c.Get("防御减免")+c.Get("怪物防御%"))
	//c.set("防御系数", c.Get("穿防系数")*c.Get("减防系数")*(c.Get("怪物等级")+100))
	//c.set("防御承伤", (c.Get("人物等级")+100)/(c.Get("人物等级")+100+c.Get("防御系数")))
	//
	//// 基础伤害
	//c.set("基础伤害", c.Get("总攻击力")*c.Get("防御承伤"))
	//c.set("基础伤害(平均)", c.Get("基础伤害")*c.Get("暴击收益"))
	//c.set("基础伤害(暴击)", c.Get("基础伤害")*c.Get("暴伤倍率"))
	//
	//// 增伤区
	//switch c.action.mode {
	//case attackMode.NormalAttack:
	//	c.set("普通攻击增伤", 1+c.Get(c.action.elemental.Infusion(c.infusionElemental).DamageBonusPoint().String())+c.Get("普通攻击伤害加成")+c.Get("元素影响增伤")+c.Get("伤害加成"))
	//	break
	//case attackMode.ChargedAttack:
	//	c.set("重击增伤", 1+c.Get(c.action.elemental.Infusion(c.infusionElemental).DamageBonusPoint().String())+c.Get("重击伤害加成")+c.Get("元素影响增伤")+c.Get("伤害加成"))
	//	break
	//case attackMode.ElementalSkill:
	//	c.set("元素战技增伤", 1+c.Get(c.action.elemental.DamageBonusPoint().String())+c.Get("元素战技伤害加成")+c.Get("元素影响增伤")+c.Get("伤害加成"))
	//	break
	//case attackMode.ElementalBurst:
	//	c.set("元素爆发增伤", 1+c.Get(c.action.elemental.DamageBonusPoint().String())+c.Get("元素爆发伤害加成")+c.Get("元素影响增伤")+c.Get("伤害加成"))
	//	break
	//}

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
	fmt.Printf("finalAttributes: %+v\n", c.finalAttributes)
}

func (c *Calculator) 攻击区() *Formula {
	基础攻击力 := c.add("基础攻击力", "人物攻击力", "武器攻击力")
	return 基础攻击力.add("总攻击力", 基础攻击力.multiply("百分比攻击力", "攻击力%").add("额外攻击力", "攻击力"))
}

func (c *Calculator) 倍率区() float64 {
	return 1
}

//func (c *Calculator) 基础伤害区() float64 {
//	return c.攻击区() * c.倍率区()
//}

func (c *Calculator) String() string {
	return c.values.String()
}
