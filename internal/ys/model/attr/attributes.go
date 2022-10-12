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

func (a *Attributes) Get(point point.Point) float64 {
	if value, exist := a.values[point]; exist && !value.isZero() {
		return value.value
	} else {
		return 0
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
