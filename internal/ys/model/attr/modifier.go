package attr

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/model/action"
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

type Character interface {
	GetAction(mode attackMode.AttackMode, name string) *action.Action
	BaseAttr(point point.Point) float64
	WeaponAttr(point point.Point) float64
	FinalAttributes() *Attributes
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

// 攻击模式技能倍率乘算加成
func AddAttackFactorMultiBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackFactorMultiBonus(r, add)
	}
}

// 攻击模式技能倍率加算加成
func AddAttackFactorAddBonus(r attackMode.AttackMode, add float64) AttributeModifier {
	return func(attributes *Attributes) func() {
		return attributes.addAttackFactorAddBonus(r, add)
	}
}

type Modifier struct {
	characterModifier AttributeModifier
	enemyModifier     AttributeModifier
	actionModifier    action.Modifier
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

func (m *Modifier) Apply(character Appliable, enemy Appliable, action *action.Action) func() {
	var cancels []func()
	if m.characterModifier != nil && character != nil {
		cancels = append(cancels, character.Apply(m.characterModifier))
	}
	if m.enemyModifier != nil && enemy != nil {
		cancels = append(cancels, enemy.Apply(m.enemyModifier))
	}
	if m.actionModifier != nil && action != nil {
		cancels = append(cancels, action.Apply(m.actionModifier))
	}
	return func() {
		for _, cancel := range cancels {
			cancel()
		}
	}
}

func (m *Modifier) Action(modifier action.Modifier) *Modifier {
	m.actionModifier = modifier
	return m
}

func (m *Modifier) EnemyModifier() AttributeModifier {
	return m.enemyModifier
}

func (m *Modifier) CharacterModifier() AttributeModifier {
	return m.characterModifier
}
