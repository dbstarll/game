package talent

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
		copy.charged.hits = value.charged.hits
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
		copy.cure = value.cure
		copy.curePercentage = value.curePercentage
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

func (t *TalentsTemplate) AddNormalAttacks(values ...*NormalAttack) *TalentsTemplate {
	for _, value := range values {
		t.addNormalAttack(value)
	}
	return t
}

func (t *TalentsTemplate) AddElementalSkills(values ...*ElementalSkill) *TalentsTemplate {
	for _, value := range values {
		t.addElementalSkill(value)
	}
	return t
}

func (t *TalentsTemplate) AddElementalBursts(values ...*ElementalBurst) *TalentsTemplate {
	for _, value := range values {
		t.addElementalBurst(value)
	}
	return t
}

func (t *TalentsTemplate) Talents(normal, skill, burst int) *Talents {
	return &Talents{
		normalAttack:   t.normalAttacks[normal],
		elementalSkill: t.elementalSkills[skill],
		elementalBurst: t.elementalBursts[burst],
	}
}

func (t *TalentsTemplate) Check() *TalentsTemplate {
	return t
}
