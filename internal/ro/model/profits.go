package model

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	Attack    int
	AttackPer float64

	Defence    int
	DefencePer float64
}

type Profits struct {
	physical Gains
	magical  Gains
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.Attack += incr.Attack
		g.AttackPer += incr.AttackPer

		g.Defence += incr.Defence
		g.DefencePer += incr.DefencePer
	}
}

func (p *Profits) Add(magic bool, incr *Gains) {
	if magic {
		p.magical.Add(incr)
	} else {
		p.physical.Add(incr)
	}
}

//装备攻击
func (p *Profits) Attack(magic bool) int {
	if magic {
		return p.magical.Attack
	} else {
		return p.physical.Attack
	}
}

//攻击%
func (p *Profits) AttackPer(magic bool) float64 {
	if magic {
		return p.magical.AttackPer
	} else {
		return p.physical.AttackPer
	}
}

//装备防御
func (p *Profits) Defence(magic bool) int {
	if magic {
		return p.magical.Defence
	} else {
		return p.physical.Defence
	}
}

//防御%
func (p *Profits) DefencePer(magic bool) float64 {
	if magic {
		return p.magical.DefencePer
	} else {
		return p.physical.DefencePer
	}
}
