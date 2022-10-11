package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
)

type Enemy struct {
	level           int
	base            *Attributes
	attachedAmounts map[elemental.Elemental]float64 // 附着的元素量
}

type EnemyModifier func(enemy *Enemy) func()

func BaseEnemy(level int, modifiers ...AttributeModifier) EnemyModifier {
	return func(enemy *Enemy) func() {
		oldLevel := enemy.level
		enemy.level = level
		callback := MergeAttributes(modifiers...)(enemy.base)
		return func() {
			callback()
			enemy.level = oldLevel
		}
	}
}

func NewEnemy(modifiers ...EnemyModifier) *Enemy {
	enemy := &Enemy{level: 1, base: NewAttributes(), attachedAmounts: make(map[elemental.Elemental]float64)}
	for _, modifier := range modifiers {
		modifier(enemy)
	}
	return enemy
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
