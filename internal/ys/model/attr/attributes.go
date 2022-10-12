package attr

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
)

type Attributes struct {
	values                       map[point.Point]*Attribute
	elementalAttachedDamageBonus map[elemental.Elemental]float64
}

func NewAttributes(modifiers ...AttributeModifier) *Attributes {
	a := &Attributes{
		values:                       make(map[point.Point]*Attribute),
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

func (a *Attributes) addElementalAttachedDamageBonus(elemental elemental.Elemental, add float64) func() {
	if add == 0 {
		return NopCallBack
	}
	if oldValue, exist := a.elementalAttachedDamageBonus[elemental]; !exist {
		a.elementalAttachedDamageBonus[elemental] = add
	} else if newValue := oldValue + add; newValue == 0 {
		delete(a.elementalAttachedDamageBonus, elemental)
	} else {
		a.elementalAttachedDamageBonus[elemental] = newValue
	}
	return func() {
		a.addElementalAttachedDamageBonus(elemental, -add)
	}
}

func (a *Attributes) Accumulation() AttributeModifier {
	var modifiers []AttributeModifier
	for _, attr := range a.values {
		modifiers = append(modifiers, attr.Accumulation())
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
	values = append(values, fmt.Sprintf("元素影响下增伤: %v", a.elementalAttachedDamageBonus))
	return fmt.Sprintf("%s", values)
}
