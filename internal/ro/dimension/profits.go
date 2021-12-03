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
