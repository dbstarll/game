package damage

//体型
type Somatotype int

const (
	Small Somatotype = iota
	Medium
	Large
)

func (s Somatotype) String() string {
	switch s {
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
