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

// 基础攻击力 = 人物攻击力 + 武器攻击力
func (c *Character) baseAtk() *Calculator {
	characterBaseAtk, _ := c.base.Get(point.Atk)
	weaponBaseAtk, _ := c.weapon.base.Get(point.Atk)
	baseAtk := characterBaseAtk + weaponBaseAtk
	desc := fmt.Sprintf("基础攻击力[%v] = 人物攻击力[%v] + 武器攻击力[%v]", baseAtk, characterBaseAtk, weaponBaseAtk)
	return &Calculator{value: baseAtk, desc: desc}
}

// 额外攻击力 = 基础攻击力 * 攻击力% + 固定攻击力;
//func (c *Character) extraAtk() *Calculator {
//	baseAtk := c.baseAtk()
//	extraAtk := baseAtk.value //  + c.weapon.base.Atk
//	desc := fmt.Sprintf("基础攻击力[%d] = 人物攻击力[%d] + 武器攻击力[%d]", extraAtk, c.base.Atk, c.weapon.base.Atk)
//	return &Calculator{value: extraAtk, desc: desc}
//}

// 总攻击力 = 基础攻击力 + 额外攻击力;

func (c *Character) Calculate() *Calculator {
	// 基础区
	// 基础攻击力 = 人物攻击力 + 武器攻击力
	// 额外攻击力 = 基础攻击力 * 攻击力 % + 固定攻击力
	// 总攻击力 = 基础攻击力 + 额外攻击力

	basicAttributes, finalAttributes := c.basicAttributes(), c.finalAttributes()
	基础攻击力, _ := basicAttributes.Get(point.Atk)
	固定攻击力, _ := finalAttributes.Get(point.Atk)
	攻击力百分比, _ := finalAttributes.Get(point.AtkPercentage)
	额外攻击力 := 基础攻击力*攻击力百分比/100 + 固定攻击力
	总攻击力 := 基础攻击力 + 额外攻击力
	fmt.Printf("基础攻击力: %+v\n", 基础攻击力)
	fmt.Printf("固定攻击力: %+v\n", 固定攻击力)
	fmt.Printf("攻击力百分比: %+v\n", 攻击力百分比)
	fmt.Printf("额外攻击力: %+v\n", 额外攻击力)
	fmt.Printf("总攻击力: %+v\n", 总攻击力)
	fmt.Printf("basicAttributes: %+v\n", basicAttributes)
	fmt.Printf("finalAttributes: %+v\n", finalAttributes)

	//
	//            // 暴击区
	//            if (values['暴击率'] > 1) {
	//                values['暴击率'] = 1;
	//            }
	//            values['暴击收益'] = 1 + values['暴击率'] * values['暴击伤害'];
	//            values['暴伤倍率'] = 1 + values['暴击伤害'];
	//
	//            // 防御区
	//            values['穿防系数'] = 1 - values['无视防御'];
	//            values['减防系数'] = 1 - values['防御减免'] + values['怪物防御%'];
	//            values['防御系数'] = values['穿防系数'] * values['减防系数'] * (values['怪物等级'] + 100);
	//            values['防御承伤'] = (values['人物等级'] + 100) / (values['人物等级'] + 100 + values['防御系数']);
	//
	//            // 基础伤害
	//            values['基础伤害'] = values['总攻击力'] * values['防御承伤'];
	//            values['基础伤害(平均)'] = values['基础伤害'] * values['暴击收益'];
	//            values['基础伤害(暴击)'] = values['基础伤害'] * values['暴伤倍率'];
	//
	//            // 增伤区
	//            values['普攻伤害加成'] = 1 + values['物理伤害%'] + values['普攻增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//            values['重击伤害加成'] = 1 + values['物理伤害%'] + values['重击增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//            values['附魔普攻伤害加成'] = 1 + values['元素伤害%'] + values['普攻增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//            values['附魔重击伤害加成'] = 1 + values['元素伤害%'] + values['重击增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//            values['元素战技伤害加成'] = 1 + values['元素伤害%'] + values['元素战技增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//            values['元素爆发伤害加成'] = 1 + values['元素伤害%'] + values['元素爆发增伤'] + values['元素影响增伤'] + values['造成伤害提高'];
	//
	//            // 抗性区
	//            const phyR = values['元素抗性'] = values['怪物元素抗性'] - values['元素抗性减免'];
	//            if (phyR > 0.75) {
	//                values['元素抗性承伤'] = 1 / (1 + 4 * phyR);
	//            } else if (phyR >= 0) {
	//                values['元素抗性承伤'] = 1 - phyR;
	//            } else {
	//                values['元素抗性承伤'] = 1 - phyR / 2;
	//            }
	//            const eleR = values['物理抗性'] = values['怪物物理抗性'] - values['物理抗性减免'];
	//            if (eleR > 0.75) {
	//                values['物理抗性承伤'] = 1 / (1 + 4 * eleR);
	//            } else if (eleR >= 0) {
	//                values['物理抗性承伤'] = 1 - eleR;
	//            } else {
	//                values['物理抗性承伤'] = 1 - eleR / 2;
	//            }
	//
	//            // 技能伤害
	//            values['普攻伤害'] = values['基础伤害'] * values['普攻倍率'] * values['普攻伤害加成'] * values['物理抗性承伤'];
	//            values['普攻伤害(平均)'] = values['普攻伤害'] * values['暴击收益'];
	//            values['普攻伤害(暴击)'] = values['普攻伤害'] * values['暴伤倍率'];
	//            values['重击伤害'] = values['基础伤害'] * values['重击倍率'] * values['重击伤害加成'] * values['物理抗性承伤'];
	//            values['重击伤害(平均)'] = values['重击伤害'] * values['暴击收益'];
	//            values['重击伤害(暴击)'] = values['重击伤害'] * values['暴伤倍率'];
	//            values['附魔普攻伤害'] = values['基础伤害'] * values['普攻倍率'] * values['附魔普攻伤害加成'] * values['元素抗性承伤'];
	//            values['附魔普攻伤害(平均)'] = values['附魔普攻伤害'] * values['暴击收益'];
	//            values['附魔普攻伤害(暴击)'] = values['附魔普攻伤害'] * values['暴伤倍率'];
	//            values['附魔重击伤害'] = values['基础伤害'] * values['重击倍率'] * values['附魔重击伤害加成'] * values['元素抗性承伤'];
	//            values['附魔重击伤害(平均)'] = values['附魔重击伤害'] * values['暴击收益'];
	//            values['附魔重击伤害(暴击)'] = values['附魔重击伤害'] * values['暴伤倍率'];
	//            values['元素战技伤害'] = values['基础伤害'] * values['元素战技倍率'] * values['元素战技伤害加成'] * values['元素抗性承伤'];
	//            values['元素战技伤害(平均)'] = values['元素战技伤害'] * values['暴击收益'];
	//            values['元素战技伤害(暴击)'] = values['元素战技伤害'] * values['暴伤倍率'];
	//            values['元素爆发伤害'] = values['基础伤害'] * values['元素爆发倍率'] * values['元素爆发伤害加成'] * values['元素抗性承伤'];
	//            values['元素爆发伤害(平均)'] = values['元素爆发伤害'] * values['暴击收益'];
	//            values['元素爆发伤害(暴击)'] = values['元素爆发伤害'] * values['暴伤倍率'];
	//
	//            // 增幅反应区
	//            values['增幅精通提升'] = (2.78 * values['元素精通']) / (values['元素精通'] + 1400);
	//            values['增幅反应倍率'] = 1 + values['增幅精通提升'] + values['反应系数提高'];
	//
	//            // 剧变反应区
	//            values['剧变精通提升'] = 16 * values['元素精通'] / (values['元素精通'] + 2000);
	//            values['剧变反应倍率'] = 1 + values['剧变精通提升'] + values['反应伤害提升'];
	//            values['等级系数'] = upheavals[values['人物等级'] - 1];
	//            values['剧变基础伤害'] = values['等级系数'] * values['元素抗性承伤'] * values['剧变反应倍率'];
	//
	//            // 反应伤害
	//            const element = $('select[column=元素类型]').val();
	//            const bonus = reactions.bonus[element];
	//            if ('object' === typeof bonus) {
	//                Object.keys(bonus).forEach((key) => {
	//                    values['附魔普攻伤害(' + key + ')'] = values['附魔普攻伤害'] * values['增幅反应倍率'] * bonus[key].rate;
	//                    values['附魔重击伤害(' + key + ')'] = values['附魔重击伤害'] * values['增幅反应倍率'] * bonus[key].rate;
	//                    values['元素战技伤害(' + key + ')'] = values['元素战技伤害'] * values['增幅反应倍率'] * bonus[key].rate;
	//                    values['元素爆发伤害(' + key + ')'] = values['元素爆发伤害'] * values['增幅反应倍率'] * bonus[key].rate;
	//                    values['附魔普攻伤害(' + key.substr(0, 1) + '暴)'] = values['附魔普攻伤害(' + key + ')'] * values['暴伤倍率'];
	//                    values['附魔重击伤害(' + key.substr(0, 1) + '暴)'] = values['附魔重击伤害(' + key + ')'] * values['暴伤倍率'];
	//                    values['元素战技伤害(' + key.substr(0, 1) + '暴)'] = values['元素战技伤害(' + key + ')'] * values['暴伤倍率'];
	//                    values['元素爆发伤害(' + key.substr(0, 1) + '暴)'] = values['元素爆发伤害(' + key + ')'] * values['暴伤倍率'];
	//                });
	//            }
	//            const upheaval = reactions.upheaval[element];
	//            if ('object' === typeof upheaval) {
	//                Object.keys(upheaval).forEach((key) => {
	//                    values['剧变伤害(' + key + ')'] = values['剧变基础伤害'] * reactions.upheaval.rate[key];
	//                });
	//            }
	return nil
}

//<div id='calculator2'>
//<fieldset class='calculator-damage-bonus'>
//<legend>增伤区</legend>
//</fieldset>
//<fieldset class='calculator-resistance'>
//<legend>抗性区</legend>
//</fieldset>
//<fieldset class='calculator-damage-skill'>
//<legend>技能伤害</legend>
//</fieldset>
//</div>
//<div id='calculator3'>
//<fieldset class='calculator-reaction-bonus'>
//<legend>增幅反应区</legend>
//</fieldset>
//<fieldset class='calculator-reaction-upheaval'>
//<legend>剧变反应区</legend>
//</fieldset>
//<fieldset class='calculator-damage-reaction'>
//<legend>反应伤害</legend>
//</fieldset>
//</div>
