package position

// 圣遗物位置
type Position int

const (
	FlowerOfLife     Position = iota // 生之花
	PlumeOfDeath                     // 死之羽
	SandsOfEon                       // 时之沙
	GobletOfEonothem                 // 空之杯
	CircletOfLogos                   // 理之冠
)

var Positions = []Position{
	FlowerOfLife,
	PlumeOfDeath,
	SandsOfEon,
	GobletOfEonothem,
	CircletOfLogos,
}

func (p Position) String() string {
	switch p {
	case FlowerOfLife:
		return "生之花"
	case PlumeOfDeath:
		return "死之羽"
	case SandsOfEon:
		return "时之沙"
	case GobletOfEonothem:
		return "空之杯"
	case CircletOfLogos:
		return "理之冠"
	default:
		if p < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
