package weapon

import (
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

//武器类型
type Weapon int

const (
	Unlimited Weapon = iota
	Empty            //空手
	Dagger           //短剑
	Sword            //长剑
	Spear            //长矛
	Katar            //拳刃
	Glove            //拳套
	Blunt            //钝器
	Axe              //斧子
	Wand             //法杖
	Book             //书
	Bow              //弓
	Musical          //乐器
	Whip             //鞭子
	Dart             //风魔
	Pistol           //手枪
	Rifle            //来复枪
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
	if w > Unlimited && w <= Rifle {
		if s > shape.Unlimited && s <= shape.Large {
			return restraints[w-1][s-1]
		}
	}
	return 1
}

func (w Weapon) IsMagic(j job.Job) bool {
	switch w {
	case Book, Wand:
		return true
	case Blunt:
		return j >= job.Acolyte && j <= job.Priest4
	default:
		return false
	}
}

func (w Weapon) IsRemote(j job.Job) bool {
	switch w {
	case Book, Wand, Bow, Musical, Whip, Dart, Pistol, Rifle:
		return true
	default:
		return false
	}
}
