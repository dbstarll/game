package shape

import "gopkg.in/yaml.v3"

//体型
type Shape int

const (
	Unlimited Shape = iota
	Small           // 小型
	Medium          // 中型
	Large           // 大型
)

func (s Shape) String() string {
	switch s {
	case Unlimited:
		return "不限"
	case Small:
		return "小型"
	case Medium:
		return "中型"
	case Large:
		return "大型"
	default:
		return "未知"
	}
}

func (s *Shape) UnmarshalYAML(value *yaml.Node) error {
	for i := Unlimited; i <= Large; i++ {
		if i.String() == value.Value {
			*s = i
			break
		}
	}
	return nil
}
