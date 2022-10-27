package attr

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
)

var (
	NopCallBack          = func() {}
	NopAttributeModifier = func(_ *Attributes) func() {
		return NopCallBack
	}
)

type AttributeModifier func(attributes *Attributes) func()

type Appliable interface {
	Apply(modifiers ...AttributeModifier) func()
}

func MergeAttributes(modifiers ...AttributeModifier) AttributeModifier {
	switch len(modifiers) {
	case 0:
		return NopAttributeModifier
	case 1:
		return modifiers[0]
	default:
		return func(attributes *Attributes) func() {
			size := len(modifiers)
			cancelList := make([]func(), size)
			for idx, modifier := range modifiers {
				cancelList[size-idx-1] = modifier(attributes)
			}
			return func() {
				for _, cancel := range cancelList {
					cancel()
				}
			}
		}
	}
}

// 单个元素伤害加成
func AddElementalDamageBonus(e elementals.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalDamageBonus(e, add)
	}
}

// 单个元素抗性
func AddElementalResist(e elementals.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalResist(e, add)
	}
}

// 单个元素影响下增伤
func AddElementalAttachedDamageBonus(e elementals.Elemental, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addElementalAttachedDamageBonus(e, add)
	}
}

// 单个元素反应系数提高/元素反应伤害提升
func AddReactionDamageBonus(r reactions.Reaction, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addReactionDamageBonus(r, add)
	}
}

// 单个攻击模式伤害加成
func AddAttackDamageBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackDamageBonus(r, add)
	}
}

// 攻击模式技能倍率加成
func AddAttackFactorBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackFactorBonus(r, add)
	}
}

type Modifier struct {
	characterModifier AttributeModifier
	enemyModifier     AttributeModifier
}

func NewCharacterModifier(characterModifiers ...AttributeModifier) *Modifier {
	switch len(characterModifiers) {
	case 0:
		return NewModifier(nil, nil)
	case 1:
		return NewModifier(characterModifiers[0], nil)
	default:
		return NewModifier(MergeAttributes(characterModifiers...), nil)
	}
}

func NewEnemyModifier(enemyModifiers ...AttributeModifier) *Modifier {
	switch len(enemyModifiers) {
	case 0:
		return NewModifier(nil, nil)
	case 1:
		return NewModifier(nil, enemyModifiers[0])
	default:
		return NewModifier(nil, MergeAttributes(enemyModifiers...))
	}
}

func NewModifier(characterModifier, enemyModifier AttributeModifier) *Modifier {
	return &Modifier{
		characterModifier: characterModifier,
		enemyModifier:     enemyModifier,
	}
}

func (m *Modifier) Apply(character Appliable, enemy Appliable) func() {
	var cancels []func()
	if m.characterModifier != nil {
		cancels = append(cancels, character.Apply(m.characterModifier))
	}
	if m.enemyModifier != nil {
		cancels = append(cancels, enemy.Apply(m.enemyModifier))
	}
	return func() {
		for _, cancel := range cancels {
			cancel()
		}
	}
}
