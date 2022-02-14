package weapon

import (
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

//武器类型
type Weapon int

const (
	Unlimited Weapon = iota
	Spear            //长矛
	Sword            //长剑
	Wand             //法杖
	Katar            //拳刃
	Bow              //弓
	Blunt            //钝器
	Axe              //斧子
	Book             //书
	Dagger           //短剑
	Musical          //乐器
	Whip             //鞭子
	_
	Glove  //拳套
	Dart   //风魔
	Pistol //手枪
	Rifle  //来复枪
)

var Weapons = []Weapon{
	Spear,
	Sword,
	Wand,
	Katar,
	Bow,
	Blunt,
	Axe,
	Book,
	Dagger,
	Musical,
	Whip,
	Glove,
	Dart,
	Pistol,
	Rifle,
}

var restraints = map[Weapon][]float64{
	//{Small,Medium,Large}
	Dagger:  {1.0, 0.75, 0.5},  //短剑
	Sword:   {0.75, 1.0, 0.75}, //长剑
	Spear:   {0.75, 0.75, 1.0}, //长矛
	Katar:   {0.75, 1.0, 0.75}, //拳刃
	Glove:   {1.0, 0.75, 0.5},  //拳套
	Blunt:   {0.75, 1.0, 1.0},  //钝器
	Axe:     {0.5, 0.75, 1.0},  //斧子
	Book:    {1.0, 1.0, 0.5},   //书
	Bow:     {1.0, 1.0, 0.75},  //弓
	Musical: {0.75, 1.0, 0.75}, //乐器
	Whip:    {0.75, 1.0, 0.75}, //鞭子
}

func (w Weapon) String() string {
	switch w {
	case Unlimited:
		return "不限"
	case Spear:
		return "长矛"
	case Sword:
		return "长剑"
	case Wand:
		return "法杖"
	case Katar:
		return "拳刃"
	case Bow:
		return "弓"
	case Blunt:
		return "钝器"
	case Axe:
		return "斧子"
	case Book:
		return "书"
	case Dagger:
		return "短剑"
	case Musical:
		return "乐器"
	case Whip:
		return "鞭子"
	case Glove:
		return "拳套"
	case Dart:
		return "风魔"
	case Pistol:
		return "手枪"
	case Rifle:
		return "来复枪"
	default:
		return "未知"
	}
}

func (w Weapon) Restraint(s shape.Shape) float64 {
	if values, exist := restraints[w]; exist {
		if s >= shape.Small && s <= shape.Large {
			return values[s-1]
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
