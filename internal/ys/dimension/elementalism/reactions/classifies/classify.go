package classifies

// 元素反应类型
type Classify int

const (
	Amplify   Classify = iota // 增幅
	Upheaval                  // 剧变
	Crystal                   // 结晶
	Intensify                 // 激化
)

var Classifies = []Classify{
	Amplify,
	Upheaval,
	Crystal,
	Intensify,
}

func (c Classify) String() string {
	switch c {
	case Amplify:
		return "增幅"
	case Upheaval:
		return "剧变"
	case Crystal:
		return "结晶"
	case Intensify:
		return "激化"
	default:
		if c < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
