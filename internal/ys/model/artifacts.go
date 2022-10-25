package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
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
	starHpRect = [][]float64{
		{0, 1, 1, 0, 0, 0, 0},
		{4, 551, 258, 0, 0, 0, 0},
		{12, 1893, 430, 0, 0, 0, 0},
		{16, 3571, 645, 167, 191, 215, 239},
		{20, 4780, 717, 209.13, 239, 268.88, 298.75},
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
	primaryFactor, levelFactor := a.primaryFactor(), a.levelFactor(0)
	secondaryFactors := a.secondaryFactors()
	fmt.Printf("%s[%d], level: %d, (%.2f, %.2f, %.2f)\n", a.name, a.star, a.level, primaryFactor, levelFactor, secondaryFactors)
	for _, p := range point.Points {
		if r := p.Multiple(); r > 0 {
			sr, max := r, r*primaryFactor
			if p == point.Hp || p == point.Atk {
				sr /= 2
			}
			fmt.Printf("\t%s: 主词条(%.2f, %.2f), 副词条(%.2f, %.2f, %.2f, %.2f)\n", p, max*levelFactor, max, sr*secondaryFactors[0], sr*secondaryFactors[1], sr*secondaryFactors[2], sr*secondaryFactors[3])
		}
	}
}

func (a *Artifacts) primaryFactor() float64 {
	return starHpRect[a.star-1][1] / point.Hp.Multiple()
}

func (a *Artifacts) baseFactor() float64 {
	return starHpRect[a.star-1][2] / starHpRect[a.star-1][1]
}

func (a *Artifacts) maxLevel() int {
	return int(starHpRect[a.star-1][0])
}

func (a *Artifacts) levelFactor(level int) float64 {
	baseFactor, maxLevel := a.baseFactor(), a.maxLevel()
	return (baseFactor*float64(maxLevel-level) + float64(level)) / float64(maxLevel)
}

func (a *Artifacts) secondaryFactors() []float64 {
	rect := starHpRect[a.star-1]
	base := a.primaryFactor() * 2 / rect[1]
	return []float64{
		base * rect[3],
		base * rect[4],
		base * rect[5],
		base * rect[6],
	}
}

func (a *Artifacts) String() string {
	return fmt.Sprintf("%s{star:%d level:%d primary:%s secondary:%s}", a.position, a.star, a.level, a.primary, a.secondary)
}
