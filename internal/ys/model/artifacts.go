package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/pkg/errors"
	"math"
)

var (
	ArtifactsFactory生之花 = func(star int, name string, level int, secondaryEntries map[entry.Entry]float64) (*Artifacts, error) {
		return newArtifacts(star, position.FlowerOfLife, entry.Hp, name, baseArtifacts(level), secondaryArtifacts(secondaryEntries))
	}
	ArtifactsFactory死之羽 = func(star int, name string, level int, secondaryEntries map[entry.Entry]float64) (*Artifacts, error) {
		return newArtifacts(star, position.PlumeOfDeath, entry.Atk, name, baseArtifacts(level), secondaryArtifacts(secondaryEntries))
	}
	ArtifactsFactory时之沙 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryEntries map[entry.Entry]float64) (*Artifacts, error) {
		return newArtifacts(star, position.SandsOfEon, primaryEntry, name, baseArtifacts(level), secondaryArtifacts(secondaryEntries))
	}
	ArtifactsFactory空之杯 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryEntries map[entry.Entry]float64) (*Artifacts, error) {
		return newArtifacts(star, position.GobletOfEonothem, primaryEntry, name, baseArtifacts(level), secondaryArtifacts(secondaryEntries))
	}
	ArtifactsFactory理之冠 = func(star int, name string, primaryEntry entry.Entry, level int, secondaryEntries map[entry.Entry]float64) (*Artifacts, error) {
		return newArtifacts(star, position.CircletOfLogos, primaryEntry, name, baseArtifacts(level), secondaryArtifacts(secondaryEntries))
	}
	starHpRect = [][]float64{
		{0, 1, 1, 0, 0, 0, 0},
		{4, 551, 258, 0, 0, 0, 0},
		{12, 1893, 430, 0, 0, 0, 0},
		{16, 3571, 645, 167, 191, 215, 239},
		{20, 4780, 717, 209.13, 239, 268.88, 298.75},
	}
)

type PrimaryEntry struct {
	entry  entry.Entry
	value  float64
	load   attr.AttributeModifier
	unload attr.AttributeModifier
}

type SecondaryEntry struct {
	rect   []int
	value  float64
	load   attr.AttributeModifier
	unload attr.AttributeModifier
}

type Artifacts struct {
	star             int
	position         position.Position
	name             string
	level            int
	primaryEntry     *PrimaryEntry
	secondaryEntries map[entry.Entry]*SecondaryEntry
	primary          *attr.Attributes
	secondary        *attr.Attributes
}

type ArtifactsModifier func(artifacts *Artifacts) (func(), error)

func baseArtifacts(level int) ArtifactsModifier {
	return func(artifacts *Artifacts) (func(), error) {
		if level < 0 {
			return nil, errors.Errorf("圣遗物等级[%d]不能为负数", level)
		} else if level > artifacts.maxLevel() {
			return nil, errors.Errorf("圣遗物等级[%d]超过上限[%d]", level, artifacts.maxLevel())
		} else if rate, fn := artifacts.primaryEntry.entry.Multiple(); rate == 0 || fn == nil {
			return nil, errors.Errorf("圣遗物主词条[%s]增幅未定义", artifacts.primaryEntry.entry)
		} else {
			oldLevel := artifacts.level
			artifacts.level = level
			primaryFactor, levelFactor := artifacts.primaryFactor(), artifacts.levelFactor(level)
			value := rate * primaryFactor * levelFactor
			artifacts.primaryEntry.value = value
			artifacts.primaryEntry.load = fn(value)
			artifacts.primaryEntry.unload = fn(-value)
			callback := artifacts.primaryEntry.load(artifacts.primary)
			return func() {
				callback()
				artifacts.level = oldLevel
			}, nil
		}
	}
}

func secondaryArtifacts(secondaryEntries map[entry.Entry]float64) ArtifactsModifier {
	return func(artifacts *Artifacts) (f func(), err error) {
		secondaryFactors, secondaryModifiers := artifacts.secondaryFactors(), make([]attr.AttributeModifier, 0)
		for ent, value := range secondaryEntries {
			if rate, fn := ent.Multiple(); rate == 0 || fn == nil {
				return nil, errors.Errorf("圣遗物副词条[%s]增幅未定义", ent)
			} else if matchRect, matchFactor, err := detectSecondaryEntry(ent, value/rate, secondaryFactors); err != nil {
				return nil, err
			} else {
				secondaryEntry := &SecondaryEntry{
					rect:   matchRect,
					value:  matchFactor * rate,
					load:   fn(matchFactor * rate),
					unload: fn(-matchFactor * rate),
				}
				secondaryModifiers = append(secondaryModifiers, secondaryEntry.load)
				artifacts.secondaryEntries[ent] = secondaryEntry
			}
		}
		callback := attr.MergeAttributes(secondaryModifiers...)(artifacts.secondary)
		return func() {
			callback()
			for ent, _ := range secondaryEntries {
				delete(artifacts.secondaryEntries, ent)
			}
		}, nil
	}
}

func detectSecondaryEntry(entry entry.Entry, value float64, secondaryFactors []float64) ([]int, float64, error) {
	if ratio, exist := entry.Secondary(); !exist {
		return nil, 0, errors.Errorf("圣遗物不支持副词条[%s]", entry)
	} else {
		factor, matchRet, matchFactor := value*ratio, make([]int, 0), 0.0
		var ret [][]int
		for ite := int(math.Floor((factor + 0.1) / secondaryFactors[0])); ite >= 0; ite-- {
			if len(ret) == 0 {
				for idx, secondaryFactor := range secondaryFactors {
					testRet, testFactor := []int{idx}, secondaryFactor
					if math.Abs(factor-testFactor) < math.Abs(factor-matchFactor) {
						matchRet, matchFactor = testRet, testFactor
					}
					ret = append(ret, testRet)
				}
			} else {
				var newRet [][]int
				for _, sub := range ret {
					subFactor, minIndex := 0.0, len(secondaryFactors)-1
					for _, i := range sub {
						subFactor += secondaryFactors[i]
						if i < minIndex {
							minIndex = i
						}
					}
					for idx, secondaryFactor := range secondaryFactors {
						if idx >= minIndex {
							testRet, testFactor := append(sub, idx), subFactor+secondaryFactor
							if math.Abs(factor-testFactor) < math.Abs(factor-matchFactor) {
								matchRet, matchFactor = testRet, testFactor
							}
							newRet = append(newRet, testRet)
						}
					}
				}
				ret = newRet
			}
		}
		if len(matchRet) == 0 {
			return nil, 0, errors.Errorf("无法解析圣遗物副词条[%s]增幅: %v, %v", entry, value, secondaryFactors)
		} else {
			return matchRet, matchFactor / ratio, nil
		}
	}
}

func newArtifacts(star int, position position.Position, primaryEntry entry.Entry, name string, baseModifier, secondaryModifier ArtifactsModifier) (*Artifacts, error) {
	if star < 3 || star > 5 {
		return nil, errors.Errorf("不支持的星级: %d", star)
	} else if !primaryEntry.Primary(position) {
		return nil, errors.Errorf("圣遗物[%s]不支持主词条[%s]", position, primaryEntry)
	}
	a := &Artifacts{
		star:             star,
		position:         position,
		name:             name,
		level:            0,
		primaryEntry:     &PrimaryEntry{entry: primaryEntry},
		secondaryEntries: make(map[entry.Entry]*SecondaryEntry),
		primary:          attr.NewAttributes(),
		secondary:        attr.NewAttributes(),
	}
	if _, err := baseModifier(a); err != nil {
		return nil, err
	} else if _, err := secondaryModifier(a); err != nil {
		return nil, err
	} else {
		return a, nil
	}
}

func (a *Artifacts) Position() position.Position {
	return a.position
}

func (a *Artifacts) Accumulation(unload bool) attr.AttributeModifier {
	return attr.MergeAttributes(a.primary.Accumulation(unload), a.secondary.Accumulation(unload))
}

func (a *Artifacts) Evaluate(replace *Artifacts) map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	detects[a.name] = attr.NewCharacterModifier(a.Accumulation(true))
	if replace != nil {
		detects[fmt.Sprintf("%s - [替换]", a.name)] = attr.NewCharacterModifier(attr.MergeAttributes(a.Accumulation(true), replace.Accumulation(false)))
	}
	detects[fmt.Sprintf("%s - [主]%s: %.2f", a.name, a.primaryEntry.entry, a.primaryEntry.value)] = attr.NewCharacterModifier(a.primaryEntry.unload)
	for ent, secondaryEntry := range a.secondaryEntries {
		detects[fmt.Sprintf("%s - [副]%s%v: %.2f", a.name, ent, secondaryEntry.rect, secondaryEntry.value)] = attr.NewCharacterModifier(secondaryEntry.unload)
	}
	return detects
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
