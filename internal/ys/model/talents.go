package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
	"time"
)

type Talents struct {
	normalAttack   *NormalAttack
	elementalSkill *ElementalSkill
	elementalBurst *ElementalBurst
}

type NormalAttack struct {
	name      string
	lv        int
	hits      []float64
	charged   ChargedAttack
	plunge    PlungeAttack
	elemental elemental.Elemental
}

type ChargedAttack struct {
	cyclic   float64
	final    float64
	stamina  int
	duration time.Duration
}

type PlungeAttack struct {
	dmg  float64
	low  float64
	high float64
}

type ElementalSkill struct {
	name      string
	lv        int
	dmgs      map[string]float64
	cd        time.Duration // 冷却时间
	elemental elemental.Elemental
}

type ElementalBurst struct {
	name             string
	lv               int
	dmgs             map[string]float64
	cd               time.Duration // 冷却时间
	infusionDuration time.Duration // 附魔持续时间
	energyCost       int           // 元素能量
	elemental        elemental.Elemental
}

type TalentsTemplate struct {
	normalAttacks   []*NormalAttack
	elementalSkills []*ElementalSkill
	elementalBursts []*ElementalBurst
}

func NewTalentsTemplate(templateNormalAttack *NormalAttack, templateElementalSkill *ElementalSkill, templateElementalBurst *ElementalBurst) *TalentsTemplate {
	talentsTemplate := &TalentsTemplate{
		normalAttacks:   make([]*NormalAttack, templateNormalAttack.lv+1),
		elementalSkills: make([]*ElementalSkill, templateElementalSkill.lv+1),
		elementalBursts: make([]*ElementalBurst, templateElementalBurst.lv+1),
	}
	talentsTemplate.normalAttacks[0] = templateNormalAttack
	talentsTemplate.elementalSkills[0] = templateElementalSkill
	talentsTemplate.elementalBursts[0] = templateElementalBurst
	return talentsTemplate
}

func (t *TalentsTemplate) addNormalAttack(value *NormalAttack) *TalentsTemplate {
	if value != nil && value.lv > 0 {
		copy := *t.normalAttacks[0]
		copy.lv = value.lv
		copy.hits = value.hits
		copy.charged.cyclic = value.charged.cyclic
		copy.charged.final = value.charged.final
		copy.plunge = value.plunge
		t.normalAttacks[copy.lv] = &copy
	}
	return t
}

func (t *TalentsTemplate) addElementalSkill(value *ElementalSkill) *TalentsTemplate {
	if value != nil && value.lv > 0 {
		copy := *t.elementalSkills[0]
		copy.lv = value.lv
		copy.dmgs = value.dmgs
		t.elementalSkills[copy.lv] = &copy
	}
	return t
}

func (t *TalentsTemplate) addElementalBurst(value *ElementalBurst) *TalentsTemplate {
	if value != nil && value.lv > 0 {
		copy := *t.elementalBursts[0]
		copy.lv = value.lv
		copy.dmgs = value.dmgs
		t.elementalBursts[copy.lv] = &copy
	}
	return t
}

func (t *TalentsTemplate) addNormalAttacks(values ...*NormalAttack) *TalentsTemplate {
	for _, value := range values {
		t.addNormalAttack(value)
	}
	return t
}

func (t *TalentsTemplate) addElementalSkills(values ...*ElementalSkill) *TalentsTemplate {
	for _, value := range values {
		t.addElementalSkill(value)
	}
	return t
}

func (t *TalentsTemplate) addElementalBursts(values ...*ElementalBurst) *TalentsTemplate {
	for _, value := range values {
		t.addElementalBurst(value)
	}
	return t
}

func (t *TalentsTemplate) check() *TalentsTemplate {
	return t
}

func (t *Talents) DMGs() *Actions {
	actions := NewActions()
	actions.addAll(t.normalAttack.DMGs())
	actions.addAll(t.elementalSkill.DMGs())
	actions.addAll(t.elementalBurst.DMGs())
	return actions
}

func (a *NormalAttack) DMGs() *Actions {
	actions := NewActions()
	for idx, dmg := range a.hits {
		actions.add(NewAction(attackMode.NormalAttack, dmg, a.elemental, fmt.Sprintf("%s•%d段", a.name, idx+1)))
	}
	actions.addAll(a.charged.DMGs(a.name, a.elemental))
	actions.addAll(a.plunge.DMGs(a.name, a.elemental))
	return actions
}

func (a *ChargedAttack) DMGs(name string, elemental elemental.Elemental) *Actions {
	actions := NewActions()
	if a.cyclic > 0 {
		actions.add(NewAction(attackMode.ChargedAttack, a.cyclic, elemental, fmt.Sprintf("%s•重击持续", name)))
	}
	if a.final > 0 {
		actions.add(NewAction(attackMode.ChargedAttack, a.final, elemental, fmt.Sprintf("%s•重击终结", name)))
	}
	return actions
}

func (a *PlungeAttack) DMGs(name string, elemental elemental.Elemental) *Actions {
	actions := NewActions()
	if a.dmg > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.dmg, elemental, fmt.Sprintf("%s•下坠期间", name)))
	}
	if a.low > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.low, elemental, fmt.Sprintf("%s•低空坠地冲击", name)))
	}
	if a.high > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.high, elemental, fmt.Sprintf("%s•高空坠地冲击", name)))
	}
	return actions
}

func (a *ElementalSkill) DMGs() *Actions {
	actions := NewActions()
	for name, dmg := range a.dmgs {
		actions.add(NewAction(attackMode.ElementalSkill, dmg, a.elemental, fmt.Sprintf("%s•%s", a.name, name)))
	}
	return actions
}

func (a *ElementalBurst) DMGs() *Actions {
	actions := NewActions()
	for name, dmg := range a.dmgs {
		actions.add(NewAction(attackMode.ElementalBurst, dmg, a.elemental, fmt.Sprintf("%s•%s", a.name, name)))
	}
	return actions
}
