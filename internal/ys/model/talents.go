package model

import (
	"time"
)

type Talents struct {
	normalAttack   *NormalAttack
	elementalSkill *ElementalSkill
	elementalBurst *ElementalBurst
}

type NormalAttack struct {
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
	DMG  float64
	low  float64
	high float64
}

type ElementalSkill struct {
	lv   int
	DMGs map[string]float64
	cd   time.Duration // 冷却时间
}

type ElementalBurst struct {
	lv               int
	DMGs             map[string]float64
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
		copy.DMGs = value.DMGs
		t.elementalSkills[copy.lv] = &copy
	}
	return t
}

func (t *TalentsTemplate) addElementalBurst(value *ElementalBurst) *TalentsTemplate {
	if value != nil && value.lv > 0 {
		copy := *t.elementalBursts[0]
		copy.lv = value.lv
		copy.DMGs = value.DMGs
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
