package damage

//体型
type Somatotype int

const (
	UnlimitedSized Somatotype = iota
	SmallSized
	MediumSized
	LargeSized
)

func (s Somatotype) String() string {
	switch s {
	case UnlimitedSized:
		return "不限"
	case SmallSized:
		return "小型"
	case MediumSized:
		return "中型"
	case LargeSized:
		return "大型"
	default:
		return "未知"
	}
}
