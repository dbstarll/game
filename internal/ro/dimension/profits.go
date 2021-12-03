package dimension

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	attack  int
	defence int
}

type Profits struct {
	physical Gains
	magical  Gains
}

//装备攻击
func (p *Profits) Attack(magic bool) int {
	if magic {
		return p.magical.attack
	} else {
		return p.physical.attack
	}
}

//装备防御
func (p *Profits) Defence(magic bool) int {
	if magic {
		return p.magical.defence
	} else {
		return p.physical.defence
	}
}
