package model

type Enemy struct {
	level int
	base  *Attributes
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
	enemy := &Enemy{level: 1, base: NewAttributes()}
	for _, modifier := range modifiers {
		modifier(enemy)
	}
	return enemy
}
