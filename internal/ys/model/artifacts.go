package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

var (
	ArtifactsFactory生之花 = func(star int, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.FlowerOfLife, BaseArtifacts(20, point.Hp, 4780), secondaryModifiers...)
	}
	ArtifactsFactory死之羽 = func(star int, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.PlumeOfDeath, BaseArtifacts(20, point.Atk, 311), secondaryModifiers...)
	}
)

type Artifacts struct {
	star      int
	position  position.Position
	level     int
	primary   *attr.Attribute
	secondary *attr.Attributes
}

type ArtifactsModifier func(artifacts *Artifacts) func()

func BaseArtifacts(level int, point point.Point, value float64) ArtifactsModifier {
	return func(artifacts *Artifacts) func() {
		oldLevel, oldPrimary := artifacts.level, artifacts.primary
		artifacts.level, artifacts.primary = level, attr.NewAttribute(point, value)
		return func() {
			artifacts.level, artifacts.primary = oldLevel, oldPrimary
		}
	}
}

func NewArtifacts(star int, position position.Position, baseModifier ArtifactsModifier, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
	a := &Artifacts{
		star:      star,
		position:  position,
		level:     1,
		secondary: attr.NewAttributes(),
	}
	baseModifier(a)
	attr.MergeAttributes(secondaryModifiers...)(a.secondary)
	return a
}

func (a *Artifacts) Accumulation() attr.AttributeModifier {
	return attr.MergeAttributes(a.primary.Accumulation(), a.secondary.Accumulation())
}

func (a *Artifacts) String() string {
	return fmt.Sprintf("%s{star:%d level:%d primary:%s secondary:%s}", a.position, a.star, a.level, a.primary, a.secondary)
}
