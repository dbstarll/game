package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"time"
)

type Talents struct {
	normalAttack   *NormalAttack
	elementalSkill *ElementalSkill
	elementalBurst *ElementalBurst
}

type NormalAttack struct {
	name    string
	lv      int
	hits    []float64
	charged ChargedAttack
	plunge  PlungeAttack
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
	name string
	lv   int
	dmgs map[string]float64
	cd   time.Duration // 冷却时间
}

type ElementalBurst struct {
	name             string
	lv               int
	dmgs             map[string]float64
	cd               time.Duration // 冷却时间
	infusionDuration time.Duration // 附魔持续时间
	energyCost       int           // 元素能量
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
		actions.add(NewAction(attackMode.NormalAttack, dmg, fmt.Sprintf("%s•%d段伤害", a.name, idx+1)))
	}
	actions.addAll(a.charged.DMGs(a.name))
	actions.addAll(a.plunge.DMGs(a.name))
	return actions
}

func (a *ChargedAttack) DMGs(name string) *Actions {
	actions := NewActions()
	if a.cyclic > 0 {
		actions.add(NewAction(attackMode.ChargedAttack, a.cyclic, fmt.Sprintf("%s•重击持续伤害", name)))
	}
	if a.final > 0 {
		actions.add(NewAction(attackMode.ChargedAttack, a.final, fmt.Sprintf("%s•重击终结伤害", name)))
	}
	return actions
}

func (a *PlungeAttack) DMGs(name string) *Actions {
	actions := NewActions()
	if a.dmg > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.dmg, fmt.Sprintf("%s•下坠期间伤害", name)))
	}
	if a.low > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.low, fmt.Sprintf("%s•低空坠地冲击伤害", name)))
	}
	if a.high > 0 {
		actions.add(NewAction(attackMode.PlungeAttack, a.high, fmt.Sprintf("%s•高空坠地冲击伤害", name)))
	}
	return actions
}

func (s *ElementalSkill) DMGs() *Actions {
	actions := NewActions()
	for name, dmg := range s.dmgs {
		actions.add(NewAction(attackMode.ElementalSkill, dmg, fmt.Sprintf("%s•%s", s.name, name)))
	}
	return actions
}

func (b *ElementalBurst) DMGs() *Actions {
	actions := NewActions()
	for name, dmg := range b.dmgs {
		actions.add(NewAction(attackMode.ElementalBurst, dmg, fmt.Sprintf("%s•%s", b.name, name)))
	}
	return actions
}
