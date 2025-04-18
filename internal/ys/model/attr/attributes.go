package attr

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"
	"reflect"
)

type Attributes struct {
	values                       map[point.Point]*Attribute
	elementalDamageBonus         map[elementals.Elemental]float64  // 元素伤害加成
	elementalResist              map[elementals.Elemental]float64  // 元素抗性
	elementalAttachedDamageBonus map[elementals.Elemental]float64  // 元素影响下增伤
	reactionDamageBonus          map[reactions.Reaction]float64    // 元素反应系数提高/元素反应伤害提升
	attackModeDamageBonus        map[attackMode.AttackMode]float64 // 攻击模式伤害加成
	attackModeFactorMultiBonus   map[attackMode.AttackMode]float64 // 攻击模式技能倍率乘算加成
	attackModeFactorAddBonus     map[attackMode.AttackMode]float64 // 攻击模式技能倍率加算加成
}

func NewAttributes(modifiers ...AttributeModifier) *Attributes {
	a := &Attributes{
		values:                       make(map[point.Point]*Attribute),
		elementalDamageBonus:         make(map[elementals.Elemental]float64),
		elementalResist:              make(map[elementals.Elemental]float64),
		elementalAttachedDamageBonus: make(map[elementals.Elemental]float64),
		reactionDamageBonus:          make(map[reactions.Reaction]float64),
		attackModeDamageBonus:        make(map[attackMode.AttackMode]float64),
		attackModeFactorMultiBonus:   make(map[attackMode.AttackMode]float64),
		attackModeFactorAddBonus:     make(map[attackMode.AttackMode]float64),
	}
	MergeAttributes(modifiers...)(a)
	return a
}

func (a *Attributes) add(attribute *Attribute) func() {
	if attribute.isZero() {
		return NopCallBack
	}
	point := attribute.point
	if oldValue, exist := a.values[point]; !exist {
		a.values[point] = attribute.clone()
	} else if newValue := oldValue.add(attribute.value); newValue.isZero() {
		delete(a.values, point)
	} else {
		a.values[point] = newValue
	}
	return func() {
		a.add(attribute.reverse())
	}
}

func (a *Attributes) addElementalDamageBonus(elemental elementals.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalDamageBonus, a.addElementalDamageBonus)
}

func (a *Attributes) addElementalResist(elemental elementals.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalResist, a.addElementalResist)
}

func (a *Attributes) addElementalAttachedDamageBonus(elemental elementals.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalAttachedDamageBonus, a.addElementalAttachedDamageBonus)
}

func addElementalMap(e elementals.Elemental, v float64, values map[elementals.Elemental]float64, cancel func(elementals.Elemental, float64) func()) func() {
	if v == 0 {
		return NopCallBack
	}
	if oldValue, exist := values[e]; !exist {
		values[e] = v
	} else if newValue := oldValue + v; newValue == 0 {
		delete(values, e)
	} else {
		values[e] = newValue
	}
	return func() {
		cancel(e, -v)
	}
}

func (a *Attributes) addReactionDamageBonus(r reactions.Reaction, v float64) func() {
	if v == 0 {
		return NopCallBack
	}
	if oldValue, exist := a.reactionDamageBonus[r]; !exist {
		a.reactionDamageBonus[r] = v
	} else if newValue := oldValue + v; newValue == 0 {
		delete(a.reactionDamageBonus, r)
	} else {
		a.reactionDamageBonus[r] = newValue
	}
	return func() {
		a.addReactionDamageBonus(r, -v)
	}
}

func (a *Attributes) addAttackDamageBonus(mode attackMode.AttackMode, add float64) func() {
	return addAttackModeMap(mode, add, a.attackModeDamageBonus, a.addAttackDamageBonus)
}

func (a *Attributes) addAttackFactorMultiBonus(mode attackMode.AttackMode, add float64) func() {
	return addAttackModeMap(mode, add, a.attackModeFactorMultiBonus, a.addAttackFactorMultiBonus)
}

func (a *Attributes) addAttackFactorAddBonus(mode attackMode.AttackMode, add float64) func() {
	return addAttackModeMap(mode, add, a.attackModeFactorAddBonus, a.addAttackFactorAddBonus)
}

func addAttackModeMap(e attackMode.AttackMode, v float64, values map[attackMode.AttackMode]float64, cancel func(attackMode.AttackMode, float64) func()) func() {
	if v == 0 {
		return NopCallBack
	}
	if oldValue, exist := values[e]; !exist {
		values[e] = v
	} else if newValue := oldValue + v; newValue == 0 {
		delete(values, e)
	} else {
		values[e] = newValue
	}
	return func() {
		cancel(e, -v)
	}
}

func (a *Attributes) Accumulation(unload bool) AttributeModifier {
	modifiers, sign := make([]AttributeModifier, 0), 1.0
	if unload {
		sign = -1.0
	}
	for _, attr := range a.values {
		if unload {
			modifiers = append(modifiers, attr.reverse().Accumulation())
		} else {
			modifiers = append(modifiers, attr.Accumulation())
		}
	}
	for ele, val := range a.elementalDamageBonus {
		modifiers = append(modifiers, AddElementalDamageBonus(ele, val*sign))
	}
	for ele, val := range a.elementalResist {
		modifiers = append(modifiers, AddElementalResist(ele, val*sign))
	}
	for ele, val := range a.elementalAttachedDamageBonus {
		modifiers = append(modifiers, AddElementalAttachedDamageBonus(ele, val*sign))
	}
	for r, val := range a.reactionDamageBonus {
		modifiers = append(modifiers, AddReactionDamageBonus(r, val*sign))
	}
	for r, val := range a.attackModeDamageBonus {
		modifiers = append(modifiers, AddAttackDamageBonus(r, val*sign))
	}
	for r, val := range a.attackModeFactorMultiBonus {
		modifiers = append(modifiers, AddAttackFactorMultiBonus(r, val*sign))
	}
	for r, val := range a.attackModeFactorAddBonus {
		modifiers = append(modifiers, AddAttackFactorAddBonus(r, val*sign))
	}
	return MergeAttributes(modifiers...)
}

func (a *Attributes) Clear(points ...point.Point) {
	for _, point := range points {
		delete(a.values, point)
	}
}

func (a *Attributes) Get(point point.Point) float64 {
	if value, exist := a.values[point]; exist && !value.isZero() {
		return value.value
	} else {
		return 0
	}
}

func (a *Attributes) GetElementalDamageBonus(elemental elementals.Elemental) float64 {
	if value, exist := a.elementalDamageBonus[elemental]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetElementalResist(elemental elementals.Elemental) float64 {
	if value, exist := a.elementalResist[elemental]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetElementalAttachedDamageBonus(elemental elementals.Elemental) float64 {
	if value, exist := a.elementalAttachedDamageBonus[elemental]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetReactionDamageBonus(reaction reactions.Reaction) float64 {
	if value, exist := a.reactionDamageBonus[reaction]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetAttackDamageBonus(mode attackMode.AttackMode) float64 {
	if value, exist := a.attackModeDamageBonus[mode]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetAttackFactorMultiBonus(mode attackMode.AttackMode) float64 {
	if value, exist := a.attackModeFactorMultiBonus[mode]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetAttackFactorAddBonus(mode attackMode.AttackMode) float64 {
	if value, exist := a.attackModeFactorAddBonus[mode]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) String() string {
	var values []string
	for _, point := range point.Points {
		if value, exist := a.values[point]; exist {
			values = append(values, value.String())
		}
	}
	values = a.append(values, "元素伤害加成", a.elementalDamageBonus)
	values = a.append(values, "元素抗性", a.elementalResist)
	values = a.append(values, "元素影响下增伤", a.elementalAttachedDamageBonus)
	values = a.append(values, "元素反应系数提高", a.reactionDamageBonus)
	values = a.append(values, "伤害加成", a.attackModeDamageBonus)
	values = a.append(values, "技能倍率乘算加成", a.attackModeFactorMultiBonus)
	values = a.append(values, "技能倍率加算加成", a.attackModeFactorAddBonus)
	return fmt.Sprintf("%s", values)
}

func (a *Attributes) append(values []string, title string, field interface{}) []string {
	if reflect.ValueOf(field).Len() == 0 {
		return values
	} else if items, ok := field.(map[elementals.Elemental]float64); ok {
		var sub []string
		for ele, val := range items {
			sub = append(sub, fmt.Sprintf("%s%s[%.1f%%]", ele, title, val))
		}
		return append(values, sub...)
	} else if items, ok := field.(map[attackMode.AttackMode]float64); ok {
		var sub []string
		for mode, val := range items {
			sub = append(sub, fmt.Sprintf("%s%s[%.1f%%]", mode, title, val))
		}
		return append(values, sub...)
	} else if items, ok := field.(map[reactions.Reaction]float64); ok {
		var sub []string
		for reaction, val := range items {
			switch reaction.Classify() {
			case classifies.Amplify, classifies.Intensify:
				sub = append(sub, fmt.Sprintf("%s反应系数提高[%.1f%%]", reaction, val))
				break
			case classifies.Upheaval:
				sub = append(sub, fmt.Sprintf("%s反应伤害提升[%.1f%%]", reaction, val))
				break
			default:
				//TODO
				// Crystal                   // 结晶
				// Intensify                 // 激化
			}
		}
		return append(values, sub...)
	} else {
		return append(values, fmt.Sprintf("%s: %v", title, field))
	}
}
