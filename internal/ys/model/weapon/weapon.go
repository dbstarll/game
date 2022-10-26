package weapon

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"time"
)

var (
	Factory无工之剑 = func(refine int) *Weapon {
		return New(5, weaponType.Claymore, "无工之剑", Base(90, 608, buff.AddAtkPercentage(49.6)),
			buff.AddShieldStrength(float64(15+refine*5)),
			buff.Superposition(5, time.Second*8, time.Millisecond*300, buff.AddAtkPercentage(float64(3+refine))),
		)
	}
	Factory螭骨剑 = func(refine int) *Weapon {
		return New(5, weaponType.Claymore, "螭骨剑", Base(90, 509, buff.AddCriticalRate(27.6)),
			buff.Superposition(5, 0, time.Second*4, buff.AddDamageBonus(5.0+float64(refine))),
			buff.Superposition(5, 0, time.Second*4, buff.AddIncomingDamageBonus([]float64{3.0, 2.7, 2.4, 2.2, 2.0}[refine-1])),
		)
	}
	Factory原木刀 = func(refine int) *Weapon {
		return New(4, weaponType.Sword, "原木刀", Base(90, 565, buff.AddEnergyRecharge(30.6)))
	}
)

type Weapon struct {
	star            int
	weaponType      weaponType.WeaponType
	name            string
	level           int
	base            *attr.Attributes
	refineModifiers []attr.AttributeModifier
}

type Modifier func(weapon *Weapon) func()

func Base(level, baseAtk int, baseModifier attr.AttributeModifier) Modifier {
	return func(weapon *Weapon) func() {
		oldLevel := weapon.level
		weapon.level = level
		callback := attr.MergeAttributes(buff.AddAtk(baseAtk), baseModifier)(weapon.base)
		return func() {
			callback()
			weapon.level = oldLevel
		}
	}
}

func New(star int, weaponType weaponType.WeaponType, name string, baseModifier Modifier, refineModifiers ...attr.AttributeModifier) *Weapon {
	w := &Weapon{
		star:            star,
		weaponType:      weaponType,
		name:            name,
		level:           1,
		base:            attr.NewAttributes(),
		refineModifiers: refineModifiers,
	}
	baseModifier(w)
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
		return w.base.Accumulation(false)
	}
}

func (w *Weapon) AccumulationRefine() attr.AttributeModifier {
	if w == nil {
		return attr.NopAttributeModifier
	} else {
		return attr.MergeAttributes(w.refineModifiers...)
	}
}

func (w *Weapon) Evaluate() map[string]*attr.Modifier {
	detects := make(map[string]*attr.Modifier)
	weaponRefine := attr.NewAttributes()
	w.AccumulationRefine()(weaponRefine)
	detects[fmt.Sprintf("%s - 基础", w.name)] = attr.NewCharacterModifier(w.base.Accumulation(true))
	detects[fmt.Sprintf("%s - 精炼", w.name)] = attr.NewCharacterModifier(weaponRefine.Accumulation(true))
	detects[w.name] = attr.NewCharacterModifier(attr.MergeAttributes(w.base.Accumulation(true), weaponRefine.Accumulation(true)))
	return detects
}
