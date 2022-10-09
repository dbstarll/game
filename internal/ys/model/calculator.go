package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
)

type Calculator struct {
	value float64
	desc  string
}

func (c *Calculator) Val() float64 {
	return c.value
}

func (c *Calculator) Desc() string {
	return c.desc
}

func (c *Calculator) String() string {
	return c.desc
}

// 额外攻击力 = 基础攻击力 * 攻击力% + 固定攻击力;
//func (c *Character) extraAtk() *Calculator {
//	baseAtk := c.baseAtk()
//	extraAtk := baseAtk.value //  + c.weapon.base.Atk
//	desc := fmt.Sprintf("基础攻击力[%d] = 人物攻击力[%d] + 武器攻击力[%d]", extraAtk, c.base.Atk, c.weapon.base.Atk)
//	return &Calculator{value: extraAtk, desc: desc}
//}

// 总攻击力 = 基础攻击力 + 额外攻击力;

func (c *Character) Calculate(enemy *Enemy) *Calculator {
	values := NewValues()
	basicAttributes, finalAttributes := c.basicAttributes(), c.finalAttributes()

	// 源数据
	values.Set("人物攻击力", c.base.Get(point.Atk).value)
	values.Set("武器攻击力", c.weapon.base.Get(point.Atk).value)
	for _, p := range point.Points {
		if v := finalAttributes.Get(p); !v.IsZero() {
			if p.IsPercentage() {
				values.Set(p.String(), v.value/100)
			} else {
				values.Set(p.String(), v.value)
			}
		}
	}
	values.Set("人物等级", float64(c.level))
	values.Set("怪物等级", float64(enemy.level))
	values.Set("怪物防御%", enemy.base.Get(point.DefPercentage).value/100)

	// 基础区
	values.Set("基础攻击力", values.Get("人物攻击力")+values.Get("武器攻击力"))
	values.Set("额外攻击力", values.Get("基础攻击力")*values.Get("攻击力%")+values.Get("攻击力"))
	values.Set("总攻击力", values.Get("基础攻击力")+values.Get("额外攻击力"))

	// 暴击区
	if values.Get("暴击率") > 1.0 {
		values.Set("暴击率", 1.0)
	}
	values.Set("暴击收益", 1+values.Get("暴击率")*values.Get("暴击伤害"))
	values.Set("暴伤倍率", 1+values.Get("暴击伤害"))

	// 防御区
	values.Set("穿防系数", 1-values.Get("无视防御"))
	values.Set("减防系数", 1-values.Get("防御减免")+values.Get("怪物防御%"))
	values.Set("防御系数", values.Get("穿防系数")*values.Get("减防系数")*(values.Get("怪物等级")+100))
	values.Set("防御承伤", (values.Get("人物等级")+100)/(values.Get("人物等级")+100+values.Get("防御系数")))

	// 基础伤害
	values.Set("基础伤害", values.Get("总攻击力")*values.Get("防御承伤"))
	values.Set("基础伤害(平均)", values.Get("基础伤害")*values.Get("暴击收益"))
	values.Set("基础伤害(暴击)", values.Get("基础伤害")*values.Get("暴伤倍率"))

	//            // 增伤区
	//            values.Set("普攻伤害加成",1 + values.Get("物理伤害%") + values.Get("普攻增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//            values.Set("重击伤害加成",1 + values.Get("物理伤害%") + values.Get("重击增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//            values.Set("附魔普攻伤害加成",1 + values.Get("元素伤害%") + values.Get("普攻增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//            values.Set("附魔重击伤害加成",1 + values.Get("元素伤害%") + values.Get("重击增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//            values.Set("元素战技伤害加成",1 + values.Get("元素伤害%") + values.Get("元素战技增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//            values.Set("元素爆发伤害加成",1 + values.Get("元素伤害%") + values.Get("元素爆发增伤") + values.Get("元素影响增伤") + values.Get("造成伤害提高"));
	//
	//            // 抗性区
	//            const phyR = values.Set("元素抗性",values.Get("怪物元素抗性") - values.Get("元素抗性减免"));
	//            if (phyR > 0.75) {
	//                values.Set("元素抗性承伤",1 / (1 + 4 * phyR));
	//            } else if (phyR >= 0) {
	//                values.Set("元素抗性承伤",1 - phyR);
	//            } else {
	//                values.Set("元素抗性承伤",1 - phyR / 2);
	//            }
	//            const eleR = values.Set("物理抗性",values.Get("怪物物理抗性") - values.Get("物理抗性减免"));
	//            if (eleR > 0.75) {
	//                values.Set("物理抗性承伤",1 / (1 + 4 * eleR));
	//            } else if (eleR >= 0) {
	//                values.Set("物理抗性承伤",1 - eleR);
	//            } else {
	//                values.Set("物理抗性承伤",1 - eleR / 2);
	//            }
	//
	//            // 技能伤害
	//            values.Set("普攻伤害",values.Get("基础伤害") * values.Get("普攻倍率") * values.Get("普攻伤害加成") * values.Get("物理抗性承伤"));
	//            values.Set("普攻伤害(平均)",values.Get("普攻伤害") * values.Get("暴击收益"));
	//            values.Set("普攻伤害(暴击)",values.Get("普攻伤害") * values.Get("暴伤倍率"));
	//            values.Set("重击伤害",values.Get("基础伤害") * values.Get("重击倍率") * values.Get("重击伤害加成") * values.Get("物理抗性承伤"));
	//            values.Set("重击伤害(平均)",values.Get("重击伤害") * values.Get("暴击收益"));
	//            values.Set("重击伤害(暴击)",values.Get("重击伤害") * values.Get("暴伤倍率"));
	//            values.Set("附魔普攻伤害",values.Get("基础伤害") * values.Get("普攻倍率") * values.Get("附魔普攻伤害加成") * values.Get("元素抗性承伤"));
	//            values.Set("附魔普攻伤害(平均)",values.Get("附魔普攻伤害") * values.Get("暴击收益"));
	//            values.Set("附魔普攻伤害(暴击)",values.Get("附魔普攻伤害") * values.Get("暴伤倍率"));
	//            values.Set("附魔重击伤害",values.Get("基础伤害") * values.Get("重击倍率") * values.Get("附魔重击伤害加成") * values.Get("元素抗性承伤"));
	//            values.Set("附魔重击伤害(平均)",values.Get("附魔重击伤害") * values.Get("暴击收益"));
	//            values.Set("附魔重击伤害(暴击)",values.Get("附魔重击伤害") * values.Get("暴伤倍率"));
	//            values.Set("元素战技伤害",values.Get("基础伤害") * values.Get("元素战技倍率") * values.Get("元素战技伤害加成") * values.Get("元素抗性承伤"));
	//            values.Set("元素战技伤害(平均)",values.Get("元素战技伤害") * values.Get("暴击收益"));
	//            values.Set("元素战技伤害(暴击)",values.Get("元素战技伤害") * values.Get("暴伤倍率"));
	//            values.Set("元素爆发伤害",values.Get("基础伤害") * values.Get("元素爆发倍率") * values.Get("元素爆发伤害加成") * values.Get("元素抗性承伤"));
	//            values.Set("元素爆发伤害(平均)",values.Get("元素爆发伤害") * values.Get("暴击收益"));
	//            values.Set("元素爆发伤害(暴击)",values.Get("元素爆发伤害") * values.Get("暴伤倍率"));
	//
	//            // 增幅反应区
	//            values.Set("增幅精通提升",(2.78 * values.Get("元素精通")) / (values.Get("元素精通") + 1400));
	//            values.Set("增幅反应倍率",1 + values.Get("增幅精通提升") + values.Get("反应系数提高"));
	//
	//            // 剧变反应区
	//            values.Set("剧变精通提升",16 * values.Get("元素精通") / (values.Get("元素精通") + 2000));
	//            values.Set("剧变反应倍率",1 + values.Get("剧变精通提升") + values.Get("反应伤害提升"));
	//            values.Set("等级系数",upheavals[values.Get("人物等级") - 1]);
	//            values.Set("剧变基础伤害",values.Get("等级系数") * values.Get("元素抗性承伤") * values.Get("剧变反应倍率"));
	//
	//            // 反应伤害
	//            const element = $("select[column=元素类型]").val());
	//            const bonus = reactions.bonus[element]);
	//            if ("object" === typeof bonus) {
	//                Object.keys(bonus).forEach((key) => {
	//                    values.Set("附魔普攻伤害(" + key + ")",values.Get("附魔普攻伤害") * values.Get("增幅反应倍率") * bonus[key].rate);
	//                    values.Set("附魔重击伤害(" + key + ")",values.Get("附魔重击伤害") * values.Get("增幅反应倍率") * bonus[key].rate);
	//                    values.Set("元素战技伤害(" + key + ")",values.Get("元素战技伤害") * values.Get("增幅反应倍率") * bonus[key].rate);
	//                    values.Set("元素爆发伤害(" + key + ")",values.Get("元素爆发伤害") * values.Get("增幅反应倍率") * bonus[key].rate);
	//                    values.Set("附魔普攻伤害(" + key.substr(0, 1) + "暴)",values.Get("附魔普攻伤害(" + key + ")") * values.Get("暴伤倍率"));
	//                    values.Set("附魔重击伤害(" + key.substr(0, 1) + "暴)",values.Get("附魔重击伤害(" + key + ")") * values.Get("暴伤倍率"));
	//                    values.Set("元素战技伤害(" + key.substr(0, 1) + "暴)",values.Get("元素战技伤害(" + key + ")") * values.Get("暴伤倍率"));
	//                    values.Set("元素爆发伤害(" + key.substr(0, 1) + "暴)",values.Get("元素爆发伤害(" + key + ")") * values.Get("暴伤倍率"));
	//                });
	//            }
	//            const upheaval = reactions.upheaval[element];
	//            if ("object" === typeof upheaval) {
	//                Object.keys(upheaval).forEach((key) => {
	//                    values.Set("剧变伤害(" + key + ")",values.Get("剧变基础伤害") * reactions.upheaval.rate[key]);
	//                });
	//            }
	fmt.Printf("basicAttributes: %+v\n", basicAttributes)
	fmt.Printf("finalAttributes: %+v\n", finalAttributes)
	fmt.Printf("values: %+v\n", values)
	fmt.Printf("%+v\n", c.talents.normalAttack)
	fmt.Printf("%+v\n", c.talents.elementalSkill)
	fmt.Printf("%+v\n", c.talents.elementalBurst)
	fmt.Printf("%+v\n", c.talents.DMGs())
	return nil
}

//<div id="calculator2">
//<fieldset class="calculator-damage-bonus">
//<legend>增伤区</legend>
//</fieldset>
//<fieldset class="calculator-resistance">
//<legend>抗性区</legend>
//</fieldset>
//<fieldset class="calculator-damage-skill">
//<legend>技能伤害</legend>
//</fieldset>
//</div>
//<div id="calculator3">
//<fieldset class="calculator-reaction-bonus">
//<legend>增幅反应区</legend>
//</fieldset>
//<fieldset class="calculator-reaction-upheaval">
//<legend>剧变反应区</legend>
//</fieldset>
//<fieldset class="calculator-damage-reaction">
//<legend>反应伤害</legend>
//</fieldset>
//</div>
