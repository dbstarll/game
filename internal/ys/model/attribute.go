package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
)

type Attribute struct {
	point point.Point
	value float64
}

func NewAttribute(point point.Point, value float64) *Attribute {
	return &Attribute{
		point: point,
		value: value,
	}
}

func (a *Attribute) Accumulation() AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.Modify(a)
	}
}

func (a *Attribute) GetPoint() point.Point {
	return a.point
}

func (a *Attribute) GetValue() float64 {
	return a.value
}

func (a *Attribute) IsZero() bool {
	return a == nil || a.value == 0.0
}

func (a *Attribute) Clone() *Attribute {
	return NewAttribute(a.point, a.value)
}

func (a *Attribute) Reverse() *Attribute {
	return NewAttribute(a.point, -a.value)
}

func (a *Attribute) Add(value float64) *Attribute {
	return NewAttribute(a.point, a.value+value)
}

func (a *Attribute) String() string {
	return fmt.Sprintf("%s[%v]", a.point, a.value)
}
