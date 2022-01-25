package position

import (
	"gopkg.in/yaml.v3"
)

//部位
type Position int

const (
	Unlimited Position = iota
	Weapon             // 武器
	Shield             // 副手
	Armor              // 盔甲
	Cloak              // 披风
	Shoes              // 鞋子
	Ring               // 饰品
	Head               // 头部
	Face               // 脸部
	Mouth              // 嘴部
	Back               // 背部
	Tail               // 尾部
	Ride               // 坐骑
	Fashion            // 时装
	God                // 神器
)

func (p Position) String() string {
	switch p {
	case Unlimited:
		return "不限"
	case Weapon:
		return "武器"
	case Shield:
		return "副手"
	case Armor:
		return "盔甲"
	case Cloak:
		return "披风"
	case Shoes:
		return "鞋子"
	case Ring:
		return "饰品"
	case Head:
		return "头部"
	case Face:
		return "脸部"
	case Mouth:
		return "嘴部"
	case Back:
		return "背部"
	case Tail:
		return "尾部"
	case Ride:
		return "坐骑"
	case Fashion:
		return "时装"
	case God:
		return "神器"
	default:
		return "未知"
	}
}

func (p *Position) UnmarshalYAML(value *yaml.Node) error {
	for i := Unlimited; i <= Fashion; i++ {
		if i.String() == value.Value {
			*p = i
			break
		}
	}
	return nil
}
