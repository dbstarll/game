package size

//体型
type Size int

const (
	Unlimited Size = iota
	Small          // 小型
	Medium         // 中型
	Large          // 大型
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
