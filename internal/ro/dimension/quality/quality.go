package quality

import "gopkg.in/yaml.v3"

//品质
type Quality int

const (
	Unlimited Quality = iota
	White             // 白色
	Green             // 绿色
	Blue              // 蓝色
	Purple            // 紫色
)

func (q Quality) String() string {
	switch q {
	case Unlimited:
		return "不限"
	case White:
		return "白色"
	case Green:
		return "绿色"
	case Blue:
		return "蓝色"
	case Purple:
		return "紫色"
	default:
		return "未知"
	}
}

func (q *Quality) UnmarshalYAML(value *yaml.Node) error {
	for i := Unlimited; i <= Purple; i++ {
		if i.String() == value.Value {
			*q = i
			break
		}
	}
	return nil
}
