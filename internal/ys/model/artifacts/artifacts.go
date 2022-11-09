package artifacts

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/pkg/errors"
	"math"
	"reflect"
)

var (
	Factory生之花 = func(star int, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.FlowerOfLife, entry.Hp, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory死之羽 = func(star int, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.PlumeOfDeath, entry.Atk, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory时之沙 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.SandsOfEon, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory空之杯 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.GobletOfEonothem, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory理之冠 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.CircletOfLogos, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	starHpRect = [][]float64{
		{0, 1, 1, 0, 0, 0, 0, 0, 0},
		{4, 551, 258, 0, 0, 0, 0, 0, 0},
		{12, 1893, 430, 0, 0, 0, 0, 0, 0},
		{16, 3571, 645, 167, 191, 215, 239, 2, 3},
		{20, 4780, 717, 209.13, 239, 268.88, 298.75, 3, 4},
	}
)

type FloatEntries map[entry.Entry]float64
type IntEntries map[entry.Entry]int

type EntriesLooper interface {
	LoopEntries(looper func(entry entry.Entry, value interface{}) error) error
}

func (e FloatEntries) LoopEntries(looper func(entry entry.Entry, value interface{}) error) error {
	for entry, value := range e {
		if err := looper(entry, value); err != nil {
			return err
		}
	}
	return nil
}

func (e IntEntries) LoopEntries(looper func(entry entry.Entry, value interface{}) error) error {
	for entry, value := range e {
		if err := looper(entry, value); err != nil {
			return err
		}
	}
	return nil
}

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

type Modifier func(artifacts *Artifacts) (func(), error)

func base(level int) Modifier {
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

func secondary(secondaryEntries EntriesLooper) Modifier {
	return func(artifacts *Artifacts) (f func(), err error) {
		total, secondaryFactors, secondaryModifiers := 0, artifacts.secondaryFactors(), make([]attr.AttributeModifier, 0)
		if err := secondaryEntries.LoopEntries(func(entry entry.Entry, value interface{}) error {
			if artifacts.primaryEntry.entry == entry {
				return errors.Errorf("圣遗物主副词条[%s]不能相同", entry)
			} else if rate, fn := entry.Multiple(); rate == 0 || fn == nil {
				return errors.Errorf("圣遗物副词条[%s]增幅未定义", entry)
			} else if ratio, exist := entry.Secondary(); !exist {
				return errors.Errorf("圣遗物不支持副词条[%s]", entry)
			} else if intValue, ok := value.(int); ok {
				matchRect, matchFactor := make([]int, intValue), secondaryFactors[3]*float64(intValue)/ratio
				for ; intValue > 0; intValue-- {
					matchRect[intValue-1] = 3
				}
				secondaryEntry := &SecondaryEntry{
					rect:   matchRect,
					value:  matchFactor * rate,
					load:   fn(matchFactor * rate),
					unload: fn(-matchFactor * rate),
				}
				total, secondaryModifiers = total+len(matchRect), append(secondaryModifiers, secondaryEntry.load)
				artifacts.secondaryEntries[entry] = secondaryEntry
				return nil
			} else if floatValue, ok := value.(float64); !ok {
				return errors.Errorf("不支持的圣遗物副词条[%s]增幅类型: %s", entry, reflect.TypeOf(value))
			} else if matchRect, matchFactor, err := detectSecondaryEntry(entry, floatValue/rate, ratio, secondaryFactors); err != nil {
				return err
			} else {
				secondaryEntry := &SecondaryEntry{
					rect:   matchRect,
					value:  matchFactor * rate,
					load:   fn(matchFactor * rate),
					unload: fn(-matchFactor * rate),
				}
				total, secondaryModifiers = total+len(matchRect), append(secondaryModifiers, secondaryEntry.load)
				artifacts.secondaryEntries[entry] = secondaryEntry
				return nil
			}
		}); err != nil {
			return nil, err
		} else if min, max := artifacts.secondaryEntryNumber(); total > max || total < min {
			return nil, errors.Errorf("圣遗物副词条目数量[%d]超过限制[%d, %d]", total, min, max)
		}
		callback := attr.MergeAttributes(secondaryModifiers...)(artifacts.secondary)
		return func() {
			callback()
			secondaryEntries.LoopEntries(func(entry entry.Entry, value interface{}) error {
				delete(artifacts.secondaryEntries, entry)
				return nil
			})
		}, nil
	}
}

func detectSecondaryEntry(entry entry.Entry, value, ratio float64, secondaryFactors []float64) ([]int, float64, error) {
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

func New(star int, position position.Position, primaryEntry entry.Entry, baseModifier, secondaryModifier Modifier) (*Artifacts, error) {
	if star < 3 || star > 5 {
		return nil, errors.Errorf("不支持的星级: %d", star)
	} else if !primaryEntry.Primary(position) {
		return nil, errors.Errorf("圣遗物[%s]不支持主词条[%s]", position, primaryEntry)
	}
	a := &Artifacts{
		star:             star,
		position:         position,
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

func (a *Artifacts) Evaluate(replaceArtifacts ...*Artifacts) map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	detects[a.position.String()] = attr.NewCharacterModifier(a.Accumulation(true))
	detects[fmt.Sprintf("%s - [主]%s: %.2f", a.position, a.primaryEntry.entry, a.primaryEntry.value)] = attr.NewCharacterModifier(a.primaryEntry.unload)
	for ent, secondaryEntry := range a.secondaryEntries {
		detects[fmt.Sprintf("%s - [副]%s%v: %.2f", a.position, ent, secondaryEntry.rect, secondaryEntry.value)] = attr.NewCharacterModifier(secondaryEntry.unload)
	}
	for _, replace := range replaceArtifacts {
		if replace.position == a.position {
			modifier := attr.NewCharacterModifier(a.Accumulation(true), replace.Accumulation(false))
			detects[fmt.Sprintf("%s - [替换]%s", a.position, replace)] = modifier
			detects[fmt.Sprintf("[替换]%s = %s", a.position, replace)] = modifier
		}
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

// 副词条数量限制
func (a *Artifacts) secondaryEntryNumber() (int, int) {
	min, max, levelCount := int(starHpRect[a.star-1][7]), int(starHpRect[a.star-1][8]), a.maxLevel()/4
	return min + levelCount, max + levelCount
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
