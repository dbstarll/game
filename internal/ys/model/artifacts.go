package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
)

var (
	ArtifactsFactory生之花 = func(star int, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.FlowerOfLife, BaseArtifacts(20, buff.AddHp(4780)), secondaryModifiers...)
	}
	ArtifactsFactory死之羽 = func(star int, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.PlumeOfDeath, BaseArtifacts(20, buff.AddAtk(311)), secondaryModifiers...)
	}
)

type Artifacts struct {
	star      int
	position  position.Position
	level     int
	primary   *attr.Attributes
	secondary *attr.Attributes
}

type ArtifactsModifier func(artifacts *Artifacts) func()

func BaseArtifacts(level int, primaryModifier attr.AttributeModifier) ArtifactsModifier {
	return func(artifacts *Artifacts) func() {
		oldLevel := artifacts.level
		artifacts.level = level
		callback := primaryModifier(artifacts.primary)
		return func() {
			callback()
			artifacts.level = oldLevel
		}
	}
}

func NewArtifacts(star int, position position.Position, baseModifier ArtifactsModifier, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
	a := &Artifacts{
		star:      star,
		position:  position,
		level:     1,
		primary:   attr.NewAttributes(),
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
