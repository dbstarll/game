package size

//体型
type Size int

const (
	Unlimited Size = iota
	Small
	Medium
	Large
)

func (s Size) String() string {
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
