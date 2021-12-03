package dimension

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	Attack  int
	Defence int
}

type Profits struct {
	physical Gains
	magical  Gains
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.Attack += incr.Attack
		g.Defence += incr.Defence
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

//装备防御
func (p *Profits) Defence(magic bool) int {
	if magic {
		return p.magical.Defence
	} else {
		return p.physical.Defence
	}
}
