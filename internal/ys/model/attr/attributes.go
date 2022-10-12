package attr

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
)

type Attributes struct {
	values                       map[point.Point]*Attribute
	elementalDamageBonus         map[elemental.Elemental]float64 // 元素伤害加成
	elementalResist              map[elemental.Elemental]float64 // 元素抗性
	elementalAttachedDamageBonus map[elemental.Elemental]float64 // 元素影响下增伤
}

func NewAttributes(modifiers ...AttributeModifier) *Attributes {
	a := &Attributes{
		values:                       make(map[point.Point]*Attribute),
		elementalDamageBonus:         make(map[elemental.Elemental]float64),
		elementalResist:              make(map[elemental.Elemental]float64),
		elementalAttachedDamageBonus: make(map[elemental.Elemental]float64),
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

func (a *Attributes) addElementalDamageBonus(elemental elemental.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalDamageBonus, a.addElementalDamageBonus)
}

func (a *Attributes) addElementalResist(elemental elemental.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalResist, a.addElementalResist)
}

func (a *Attributes) addElementalAttachedDamageBonus(elemental elemental.Elemental, add float64) func() {
	return addElementalMap(elemental, add, a.elementalAttachedDamageBonus, a.addElementalAttachedDamageBonus)
}

func addElementalMap(e elemental.Elemental, v float64, values map[elemental.Elemental]float64, cancel func(elemental.Elemental, float64) func()) func() {
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

func (a *Attributes) Accumulation() AttributeModifier {
	var modifiers []AttributeModifier
	for _, attr := range a.values {
		modifiers = append(modifiers, attr.Accumulation())
	}
	for ele, val := range a.elementalDamageBonus {
		modifiers = append(modifiers, AddElementalDamageBonus(ele, val))
	}
	for ele, val := range a.elementalResist {
		modifiers = append(modifiers, AddElementalResist(ele, val))
	}
	for ele, val := range a.elementalAttachedDamageBonus {
		modifiers = append(modifiers, AddElementalAttachedDamageBonus(ele, val))
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

func (a *Attributes) GetElementalDamageBonus(elemental elemental.Elemental) float64 {
	if value, exist := a.elementalDamageBonus[elemental]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetElementalResist(elemental elemental.Elemental) float64 {
	if value, exist := a.elementalResist[elemental]; exist {
		return value
	} else {
		return 0
	}
}

func (a *Attributes) GetElementalAttachedDamageBonus(elemental elemental.Elemental) float64 {
	if value, exist := a.elementalAttachedDamageBonus[elemental]; exist {
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
	values = append(values, fmt.Sprintf("元素伤害加成: %v", a.elementalDamageBonus))
	values = append(values, fmt.Sprintf("元素抗性: %v", a.elementalResist))
	values = append(values, fmt.Sprintf("元素影响下增伤: %v", a.elementalAttachedDamageBonus))
	return fmt.Sprintf("%s", values)
}
