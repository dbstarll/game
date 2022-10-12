package enemy

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
)

type Enemy struct {
	level           int
	base            *attr.Attributes
	attachedAmounts map[elemental.Elemental]float64 // 附着的元素量
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
		attachedAmounts: make(map[elemental.Elemental]float64),
	}
	for _, modifier := range modifiers {
		modifier(enemy)
	}
	return enemy
}

func (e *Enemy) Level() int {
	return e.level
}

func (e *Enemy) Get(point point.Point) *attr.Attribute {
	return e.base.Get(point)
}

func (e *Enemy) Attach(attached elemental.Elemental, amount float64) {
	e.attachedAmounts[attached] = amount
}

func (e *Enemy) DetectReaction(trigger elemental.Elemental, classify reaction.Classify) []*reaction.Factor {
	factors := make([]*reaction.Factor, 0)
	for attached, amount := range e.attachedAmounts {
		if amount > 0 {
			if factor := trigger.Reaction(attached); factor != nil && factor.Match(classify) {
				factors = append(factors, factor)
			}
		}
	}
	return factors
}
