package dimension

//素质属性
type Quality struct {
	str int
	agi int
	vit int
	int int
	dex int
	luk int
}

//素质攻击
func (q *Quality) Attack(magic bool, remote bool) int {
	if magic {
		//魔法素质攻击 = 智力*2 + 取整(智力*智力/100)
		return q.int*2 + q.int*q.int/100
	} else if remote {
		//远程素质物理攻击 = 灵巧*2 + 取整(灵巧*灵巧/100) + 取整(力量/5) + 取整(幸运/5)
		return q.dex*2 + q.dex*q.dex/100 + q.str/5 + q.luk/5
	} else {
		//近战素质物理攻击 = 力量*2 + 取整(力量*力量/100) + 取整(灵巧/5) + 取整(幸运/5)
		return q.str*2 + q.str*q.str/100 + q.dex/5 + q.luk/5
	}
}

//素质防御
func (q *Quality) Defence(magic bool) int {
	if magic {
		//素质魔法防御 = 智力
		return q.int
	} else {
		//素质物理防御 = 体质
		return q.vit
	}
}
