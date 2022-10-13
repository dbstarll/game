package enemy

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
)

type Enemy struct {
	level           int
	base            *attr.Attributes
	attachedAmounts map[elementals.Elemental]float64 // 附着的元素量
}

type Modifier func(enemy *Enemy) func()

func Base(level int, modifiers ...attr.AttributeModifier) Modifier {
	return func(enemy *Enemy) func() {
		oldLevel := enemy.level
		enemy.level = level
		callback := attr.MergeAttributes(modifiers...)(enemy.base)
		return func() {
			callback()
			enemy.level = oldLevel
		}
	}
}

func New(modifiers ...Modifier) *Enemy {
	enemy := &Enemy{
		level:           1,
		base:            attr.NewAttributes(buff.AddAllElementalResist(10)),
		attachedAmounts: make(map[elementals.Elemental]float64),
	}
	for _, modifier := range modifiers {
		modifier(enemy)
	}
	return enemy
}

func (e *Enemy) Apply(modifiers ...attr.AttributeModifier) func() {
	return attr.MergeAttributes(modifiers...)(e.base)
}

func (e *Enemy) Level() int {
	return e.level
}

func (e *Enemy) Get(point point.Point) float64 {
	return e.base.Get(point)
}

func (e *Enemy) GetElementalResist(elemental elementals.Elemental) float64 {
	return e.base.GetElementalResist(elemental)
}

// TODO 附着元素量
func (e *Enemy) Attach(attached elementals.Elemental, amount float64) {
	e.attachedAmounts[attached] = amount
}

func (e *Enemy) DetectReaction(trigger elementals.Elemental, classify classifies.Classify) []*reactions.Factor {
	factors := make([]*reactions.Factor, 0)
	for attached, amount := range e.attachedAmounts {
		if amount > 0 {
			if factor := trigger.Reaction(attached); factor != nil && factor.Match(classify) {
				factors = append(factors, factor)
			}
		}
	}
	return factors
}

func (e *Enemy) Attached() []elementals.Elemental {
	elements := make([]elementals.Elemental, 0)
	for attached, amount := range e.attachedAmounts {
		if amount > 0 {
			elements = append(elements, attached)
		}
	}
	return elements
}
