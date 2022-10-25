package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
)

var (
	ArtifactsFactory生之花 = func(star int, name string, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.FlowerOfLife, name, BaseArtifacts(20, buff.AddHp(4780)), secondaryModifiers...)
	}
	ArtifactsFactory死之羽 = func(star int, name string, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
		return NewArtifacts(star, position.PlumeOfDeath, name, BaseArtifacts(20, buff.AddAtk(311)), secondaryModifiers...)
	}
)

type Artifacts struct {
	star      int
	position  position.Position
	name      string
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

func NewArtifacts(star int, position position.Position, name string, baseModifier ArtifactsModifier, secondaryModifiers ...attr.AttributeModifier) *Artifacts {
	a := &Artifacts{
		star:      star,
		position:  position,
		name:      name,
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

func (a *Artifacts) Evaluate() {
	primaryFactor, levelFactor, secondaryFactor := a.primaryFactor(), a.levelFactor(0), a.secondaryFactor()
	fmt.Printf("%s[%d], level: %d, (%v, %v, %v)\n", a.name, a.star, a.level, primaryFactor, levelFactor, secondaryFactor)
	for _, p := range point.Points {
		if r := p.Multiple(); r > 0 {
			fmt.Printf("\t%s: %v, (%v, %v, %v)\n", p, r, r*primaryFactor*levelFactor, r*primaryFactor, r*secondaryFactor)
		}
	}
	for _, e := range elementals.Elementals {
		if r := e.Multiple(); r > 0 {
			fmt.Printf("\t%s: %v, (%v, %v)\n", e.Name(), r, r*primaryFactor*levelFactor, r*primaryFactor)
		}
	}
}

func (a *Artifacts) primaryFactor() float64 {
	switch a.star {
	case 2:
		return 551.0 / point.Hp.Multiple()
	case 3:
		return 1893.0 / point.Hp.Multiple()
	case 4:
		return 3571.0 / point.Hp.Multiple()
	case 5:
		return 4780.0 / point.Hp.Multiple()
	default:
		return 0
	}
}

func (a *Artifacts) baseFactor() float64 {
	switch a.star {
	case 2:
		return 258.0 / 551.0
	case 3:
		return 430.0 / 1893.0
	case 4:
		return 645.0 / 3571.0
	case 5:
		return 717.0 / 4780.0
	default:
		return 0
	}
}

func (a *Artifacts) maxLevel() int {
	switch a.star {
	case 2:
		return 4
	case 3:
		return 12
	case 4:
		return 16
	case 5:
		return 20
	default:
		return 0
	}
}

func (a *Artifacts) levelFactor(level int) float64 {
	baseFactor, maxLevel := a.baseFactor(), a.maxLevel()
	return (baseFactor*float64(maxLevel-level) + float64(level)) / float64(maxLevel)
}

func (a *Artifacts) secondaryFactor() float64 {
	switch a.star {
	case 4:
		return 2.5 * 3571.0 / point.Hp.Multiple() / 20
	case 5:
		return 2.5 * 4780.0 / point.Hp.Multiple() / 20
	default:
		return 0
	}
}

func (a *Artifacts) String() string {
	return fmt.Sprintf("%s{star:%d level:%d primary:%s secondary:%s}", a.position, a.star, a.level, a.primary, a.secondary)
}
