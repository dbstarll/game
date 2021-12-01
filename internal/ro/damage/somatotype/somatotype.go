package somatotype

//体型
type Somatotype int

const (
	Unlimited Somatotype = iota
	Small
	Medium
	Large
)

func (s Somatotype) String() string {
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
