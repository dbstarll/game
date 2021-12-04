package weapon

import "github.com/dbstarll/game/internal/ro/dimension/shape"

//武器类型
type Weapon int

const (
	Unlimited Weapon = iota
	空手
	短剑
	长剑
	长矛
	拳刃
	拳套
	钝器
	斧子
	法杖
	书
	弓
	乐器
	鞭子
	风魔
	手枪
	来复枪
)

var restraints = [][]float64{
	//{Small,Medium,Large}
	{1.0, 1.0, 1.0},   //空手
	{1.0, 0.75, 0.5},  //短剑
	{0.75, 1.0, 0.75}, //长剑
	{0.75, 0.75, 1.0}, //长矛
	{0.75, 1.0, 0.75}, //拳刃
	{1.0, 0.75, 0.5},  //拳套
	{0.75, 1.0, 1.0},  //钝器
	{0.5, 0.75, 1.0},  //斧子
	{1.0, 1.0, 1.0},   //法杖
	{1.0, 1.0, 0.5},   //书
	{1.0, 1.0, 0.75},  //弓
	{0.75, 1.0, 0.75}, //乐器
	{0.75, 1.0, 0.75}, //鞭子
	{1.0, 1.0, 1.0},   //风魔
	{1.0, 1.0, 1.0},   //手枪
	{1.0, 1.0, 1.0},   //来复枪
}

func (w Weapon) String() string {
	switch w {
	case Unlimited:
		return "不限"
	default:
		return "未知"
	}
}

func (w Weapon) Restraint(s shape.Shape) float64 {
	if w > Unlimited && w <= 来复枪 {
		if s > shape.Unlimited && s <= shape.Large {
			return restraints[w-1][s-1]
		}
	}
	return 1
}
