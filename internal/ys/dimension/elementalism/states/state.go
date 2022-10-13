package states

// 元素状态
type State int

const (
	Frozen  State = iota // 冻结
	Burn                 // 燃烧
	Quicken              // 原激化
	Bloom                // 绽放(模拟草原核)
)

var States = []State{
	Frozen,
	Burn,
	Quicken,
	Bloom,
}

func (s State) String() string {
	switch s {
	case Frozen:
		return "冻结"
	case Burn:
		return "燃烧"
	case Quicken:
		return "原激化"
	case Bloom:
		return "草原核"
	default:
		return "未知"
	}
}
