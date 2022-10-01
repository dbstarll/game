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

func (a *Attribute) String() string {
	return fmt.Sprintf("%s[%v]", a.point, a.value)
}
