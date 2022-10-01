package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
)

type Attributes struct {
	values map[point.Point]*Attribute
}

func NewAttributes() *Attributes {
	return &Attributes{values: make(map[point.Point]*Attribute)}
}

func (a *Attributes) Modify(attribute *Attribute) func() {
	if attribute.IsZero() {
		return NopCallBack
	}
	point := attribute.GetPoint()
	if oldValue, exist := a.values[point]; !exist {
		a.values[point] = attribute.Clone()
	} else if newValue := oldValue.Add(attribute.GetValue()); newValue.IsZero() {
		delete(a.values, point)
	} else {
		a.values[point] = newValue
	}
	return func() {
		a.Modify(attribute.Reverse())
	}
}

func (a *Attributes) Apply(modifiers ...AttributeModifier) func() {
	return MergeAttributes(modifiers...)(a)
}

func (a *Attributes) Accumulation() AttributeModifier {
	var modifiers []AttributeModifier
	for _, attr := range a.values {
		modifiers = append(modifiers, attr.Accumulation())
	}
	return MergeAttributes(modifiers...)
}

func (a *Attributes) Clear(points ...point.Point) {
	for _, point := range points {
		delete(a.values, point)
	}
}

func (a *Attributes) Get(point point.Point) (float64, bool) {
	if value, exist := a.values[point]; exist {
		return value.value, true
	} else {
		return 0, false
	}
}
