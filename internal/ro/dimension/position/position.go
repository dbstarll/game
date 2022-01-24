package position

import "gopkg.in/yaml.v3"

//部位
type Position int

const (
	Unlimited Position = iota
	Neutral            // 无
	Earth              // 地
	Wind               // 风
	Water              // 水
	Fire               // 火
	Holy               // 圣
	Head               // 头部
	Face               // 脸部
	Mouth              // 嘴部
	Back               // 背部
	Tail               // 尾部
	Ride               // 坐骑
	Fashion            // 时装
)

func (p Position) String() string {
	switch p {
	case Unlimited:
		return "不限"
	case Neutral:
		return "无"
	case Earth:
		return "地"
	case Wind:
		return "风"
	case Water:
		return "水"
	case Fire:
		return "火"
	case Holy:
		return "圣"
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
