package nature

import "gopkg.in/yaml.v3"

//属性
type Nature int

const (
	Neutral Nature = iota // 无
	Earth                 // 地
	Wind                  // 风
	Water                 // 水
	Fire                  // 火
	Holy                  // 圣
	Dark                  // 暗
	Ghost                 // 念
	Undead                // 不死
	Poison                // 毒
)

var Natures = []Nature{
	Neutral,
	Earth,
	Wind,
	Water,
	Fire,
	Holy,
	Dark,
	Ghost,
	Undead,
	Poison,
}

var restraints = [][]float64{
	//{无,地,风,水,火,圣,暗,念,不死,毒}
	{1, 1, 1, 1, 1, 1, 1, 0.25, 1, 1},                    //无
	{1, 0.25, 2, 1, 0.5, 0.75, 1, 1, 1, 0.75},            //地
	{1, 0.5, 0.25, 2, 1, 0.75, 1, 1, 1, 0.75},            //风
	{1, 1, 0.5, 0.25, 2, 0.75, 1, 1, 1.5, 0.75},          //水
	{1, 2, 1, 0.5, 0.25, 0.75, 1, 1, 2, 0.75},            //火
	{1, 1, 1, 1, 1, 0.25, 2, 1, 2, 1.25},                 //圣
	{1, 1, 1, 1, 1, 2, 0.25, 1, 0.25, 0.25},              //暗
	{0.25, 1, 1, 1, 1, 0.75, 0.75, 2, 1.75, 1},           //念
	{1, 0.5, 0.5, 0.5, 0.5, 1.75, 0.25, 1, 0.25, 0.25},   //不死
	{1, 1.25, 1.25, 1, 1.25, 0.5, 0.25, 0.5, 0.25, 0.25}, //毒
}

func (n Nature) String() string {
	switch n {
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
	case Dark:
		return "暗"
	case Ghost:
		return "念"
	case Undead:
		return "不死"
	case Poison:
		return "毒"
	default:
		if n < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}

func (n Nature) Restraint(defence Nature) float64 {
	if n >= Neutral && n <= Poison {
		if defence >= Neutral && defence <= Poison {
			return restraints[n][defence]
		}
	}
	return 1
}

func (n *Nature) UnmarshalYAML(value *yaml.Node) error {
	for i := Neutral; i <= Poison; i++ {
		if i.String() == value.Value {
			*n = i
			break
		}
	}
	return nil
}
