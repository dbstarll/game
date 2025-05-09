package model

//素质属性
type Quality struct {
	Str int `json:"str"`
	Agi int `json:"agi"`
	Vit int `json:"vit"`
	Int int `json:"int"`
	Dex int `json:"dex"`
	Luk int `json:"luk"`
}

//素质攻击
func (q *Quality) Attack(magic, remote bool) int {
	if magic {
		//魔法素质攻击 = 智力*2 + 取整(智力*智力/100)
		return q.Int*2 + q.Int*q.Int/100
	} else if remote {
		//远程素质物理攻击 = 灵巧*2 + 取整(灵巧*灵巧/100) + 取整(力量/5) + 取整(幸运/5)
		return q.Dex*2 + q.Dex*q.Dex/100 + q.Str/5 + q.Luk/5
	} else {
		//近战素质物理攻击 = 力量*2 + 取整(力量*力量/100) + 取整(灵巧/5) + 取整(幸运/5)
		return q.Str*2 + q.Str*q.Str/100 + q.Dex/5 + q.Luk/5
	}
}

//素质普攻攻击力
func (q *Quality) OrdinaryAttack(magic, remote bool) int {
	if magic {
		return q.Int * 3 // TODO 待确认
	} else if remote {
		return q.Dex * 3
	} else {
		return q.Str * 5
	}
}

//素质防御
func (q *Quality) Defence(magic bool) int {
	if magic {
		//素质魔法防御 = 智力
		return q.Int
	} else {
		//素质物理防御 = 体质
		return q.Vit
	}
}

func (q *Quality) Add(incr *Quality) {
	if incr != nil {
		q.Str += incr.Str
		q.Agi += incr.Agi
		q.Vit += incr.Vit
		q.Int += incr.Int
		q.Dex += incr.Dex
		q.Luk += incr.Luk
	}
}

func (q *Quality) Del(incr *Quality) {
	if incr != nil {
		q.Str -= incr.Str
		q.Agi -= incr.Agi
		q.Vit -= incr.Vit
		q.Int -= incr.Int
		q.Dex -= incr.Dex
		q.Luk -= incr.Luk
	}
}
