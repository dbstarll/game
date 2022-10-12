package attr

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
)

type Attributes struct {
	values map[point.Point]*Attribute
}

func NewAttributes(modifiers ...AttributeModifier) *Attributes {
	a := &Attributes{values: make(map[point.Point]*Attribute)}
	a.Apply(modifiers...)
	return a
}

func (a *Attributes) Add(attribute *Attribute) func() {
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
		a.Add(attribute.Reverse())
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

func (a *Attributes) Get(point point.Point) *Attribute {
	if value, exist := a.values[point]; exist {
		return value
	} else {
		return New(point, 0)
	}
}

func (a *Attributes) String() string {
	var values []*Attribute
	for _, point := range point.Points {
		if value, exist := a.values[point]; exist {
			values = append(values, value)
		}
	}
	return fmt.Sprintf("%s", values)
}
