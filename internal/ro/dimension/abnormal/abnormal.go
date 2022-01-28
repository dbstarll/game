package abnormal

//异常状态
type Abnormal int

const (
	Poisoning Abnormal = iota // 中毒
	Bleed                     // 流血
	Burn                      // 灼烧
	Vertigo                   // 眩晕
	Frozen                    // 冰冻
	Petrify                   // 石化
	Sleep                     // 睡眠
	Fear                      // 恐惧
	Fixed                     // 定身
	Silent                    // 沉默
	Cursed                    // 诅咒
	Dark                      // 黑暗
)

var Abnormals = []Abnormal{
	Poisoning,
	Bleed,
	Burn,
	Vertigo,
	Frozen,
	Petrify,
	Sleep,
	Fear,
	Fixed,
	Silent,
	Cursed,
	Dark,
}

func (n Abnormal) String() string {
	switch n {
	case Poisoning:
		return "无"
	case Bleed:
		return "流血"
	case Burn:
		return "灼烧"
	case Vertigo:
		return "眩晕"
	case Frozen:
		return "冰冻"
	case Petrify:
		return "石化"
	case Sleep:
		return "睡眠"
	case Fear:
		return "恐惧"
	case Fixed:
		return "定身"
	case Silent:
		return "沉默"
	case Cursed:
		return "诅咒"
	case Dark:
		return "黑暗"
	default:
		if n < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
