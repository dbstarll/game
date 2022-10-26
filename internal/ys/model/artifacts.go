package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/pkg/errors"
)

var (
	ArtifactsFactory生之花 = func(star int, name string, level int, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
		return NewArtifacts(star, position.FlowerOfLife, entry.Hp, name, BaseArtifacts(level), secondaryModifiers...)
	}
	ArtifactsFactory死之羽 = func(star int, name string, level int, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
		return NewArtifacts(star, position.PlumeOfDeath, entry.Atk, name, BaseArtifacts(level), secondaryModifiers...)
	}
	ArtifactsFactory时之沙 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
		return NewArtifacts(star, position.SandsOfEon, primaryEntry, name, BaseArtifacts(level), secondaryModifiers...)
	}
	ArtifactsFactory空之杯 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
		return NewArtifacts(star, position.GobletOfEonothem, primaryEntry, name, BaseArtifacts(level), secondaryModifiers...)
	}
	ArtifactsFactory理之冠 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
		return NewArtifacts(star, position.CircletOfLogos, primaryEntry, name, BaseArtifacts(level), secondaryModifiers...)
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
	star         int
	position     position.Position
	name         string
	level        int
	primaryEntry entry.Entry
	primary      *attr.Attributes
	secondary    *attr.Attributes
}

type ArtifactsModifier func(artifacts *Artifacts) (func(), error)

func BaseArtifacts(level int) ArtifactsModifier {
	return func(artifacts *Artifacts) (func(), error) {
		if level < 0 {
			return nil, errors.Errorf("圣遗物等级[%d]不能为负数", level)
		} else if level > artifacts.maxLevel() {
			return nil, errors.Errorf("圣遗物等级[%d]超过上限[%d]", level, artifacts.maxLevel())
		} else if rate, fn := artifacts.primaryEntry.Multiple(); rate == 0 || fn == nil {
			return nil, errors.Errorf("圣遗物主词条[%s]增幅未定义", artifacts.primaryEntry)
		} else {
			oldLevel := artifacts.level
			artifacts.level = level
			primaryFactor, levelFactor := artifacts.primaryFactor(), artifacts.levelFactor(level)
			callback := fn(rate * primaryFactor * levelFactor)(artifacts.primary)
			return func() {
				callback()
				artifacts.level = oldLevel
			}, nil
		}
	}
}

func NewArtifacts(star int, position position.Position, primaryEntry entry.Entry, name string, baseModifier ArtifactsModifier, secondaryModifiers ...attr.AttributeModifier) (*Artifacts, error) {
	if star < 3 || star > 5 {
		return nil, errors.Errorf("不支持的星级: %d", star)
	} else if !primaryEntry.Primary(position) {
		return nil, errors.Errorf("圣遗物[%s]不支持主词条[%s]", position, primaryEntry)
	}
	a := &Artifacts{
		star:         star,
		position:     position,
		name:         name,
		level:        0,
		primaryEntry: primaryEntry,
		primary:      attr.NewAttributes(),
		secondary:    attr.NewAttributes(),
	}
	if _, err := baseModifier(a); err != nil {
		return nil, err
	}
	attr.MergeAttributes(secondaryModifiers...)(a.secondary)
	return a, nil
}

func (a *Artifacts) Accumulation() attr.AttributeModifier {
	return attr.MergeAttributes(a.primary.Accumulation(), a.secondary.Accumulation())
}

func (a *Artifacts) Evaluate() {
	primaryFactor, levelFactor := a.primaryFactor(), a.levelFactor(0)
	secondaryFactors := a.secondaryFactors()
	fmt.Printf("%s[%d], level: %d, (%.2f, %.2f, %.2f)\n", a.name, a.star, a.level, primaryFactor, levelFactor, secondaryFactors)
	for _, p := range entry.Entries {
		if r, _ := p.Multiple(); r > 0 {
			sr, max := r, r*primaryFactor
			if p == entry.Hp || p == entry.Atk {
				sr /= 2
			}
			fmt.Printf("\t%s: 主词条(%.2f, %.2f), 副词条(%.2f, %.2f, %.2f, %.2f)\n", p, max*levelFactor, max, sr*secondaryFactors[0], sr*secondaryFactors[1], sr*secondaryFactors[2], sr*secondaryFactors[3])
		}
	}
}

func (a *Artifacts) primaryFactor() float64 {
	base, _ := entry.Hp.Multiple()
	return starHpRect[a.star-1][1] / base
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
	return fmt.Sprintf("%s: %s{star:%d level:%d primary:%s secondary:%s}", a.name, a.position, a.star, a.level, a.primary, a.secondary)
}
