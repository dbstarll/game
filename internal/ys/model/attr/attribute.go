package attr

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
)

type Attribute struct {
	point point.Point
	value float64
}

func New(point point.Point, value float64) *Attribute {
	return &Attribute{
		point: point,
		value: value,
	}
}

func (a *Attribute) Accumulation() AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.add(a)
	}
}

func (a *Attribute) isZero() bool {
	return a == nil || a.value == 0.0
}

func (a *Attribute) clone() *Attribute {
	return New(a.point, a.value)
}

func (a *Attribute) reverse() *Attribute {
	return New(a.point, -a.value)
}

func (a *Attribute) add(value float64) *Attribute {
	return New(a.point, a.value+value)
}

func (a Attribute) String() string {
	if a.point.IsPercentage() {
		return fmt.Sprintf("%s[%.1f%%]", a.point, a.value)
	} else {
		return fmt.Sprintf("%s[%.0f]", a.point, a.value)
	}
}
