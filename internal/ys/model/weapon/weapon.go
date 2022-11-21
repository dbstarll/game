package weapon

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"strings"
)

type Weapon struct {
	star        int
	weaponType  weaponType.WeaponType
	name        string
	level       int
	refineLevel int
	base        *attr.Attributes
	entry       *attr.Attributes
	refine      *attr.Attributes
}

type Modifier func(weapon *Weapon) func()

func Base(level, refineLevel, baseAtk int, entryModifier attr.AttributeModifier) Modifier {
	return func(weapon *Weapon) func() {
		oldLevel, oldRefineLevel := weapon.level, weapon.refineLevel
		weapon.level, weapon.refineLevel = level, refineLevel
		callback := buff.AddAtk(baseAtk)(weapon.base)
		callback2 := entryModifier(weapon.entry)
		return func() {
			callback2()
			callback()
			weapon.level, weapon.refineLevel = oldLevel, oldRefineLevel
		}
	}
}

func New(star int, weaponType weaponType.WeaponType, name string, baseModifier Modifier, refineModifiers ...attr.AttributeModifier) *Weapon {
	w := &Weapon{
		star:        star,
		weaponType:  weaponType,
		name:        name,
		level:       1,
		refineLevel: 1,
		base:        attr.NewAttributes(),
		entry:       attr.NewAttributes(),
		refine:      attr.NewAttributes(),
	}
	baseModifier(w)
	attr.MergeAttributes(refineModifiers...)(w.refine)
	return w
}

func (w *Weapon) Get(point point.Point) float64 {
	return w.base.Get(point)
}

func (w *Weapon) Type() weaponType.WeaponType {
	return w.weaponType
}

func (w *Weapon) AccumulationBase() attr.AttributeModifier {
	if w == nil {
		return attr.NopAttributeModifier
	} else {
		return attr.MergeAttributes(w.base.Accumulation(false), w.entry.Accumulation(false))
	}
}

func (w *Weapon) AccumulationRefine() attr.AttributeModifier {
	if w == nil {
		return attr.NopAttributeModifier
	} else {
		return w.refine.Accumulation(false)
	}
}

func (w *Weapon) Evaluate() map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	detects[fmt.Sprintf("%s - 白值", w.name)] = attr.NewCharacterModifier(w.base.Accumulation(true))
	detects[fmt.Sprintf("%s - 主词条", w.name)] = attr.NewCharacterModifier(w.entry.Accumulation(true))
	detects[fmt.Sprintf("%s - 精炼", w.name)] = attr.NewCharacterModifier(w.refine.Accumulation(true))
	detects[w.name] = attr.NewCharacterModifier(w.base.Accumulation(true), w.entry.Accumulation(true), w.refine.Accumulation(true))
	return detects
}

func (w *Weapon) String() string {
	if w == nil {
		return "未装备"
	} else {
		return fmt.Sprintf("%s{%s, %s[%d][%d], base: %s, entry: %s, refine: %s}", w.name, w.weaponType, strings.Repeat("⭐", w.star), w.level, w.refineLevel, w.base, w.entry, w.refine)
	}
}
