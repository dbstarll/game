package _type

//类型
type Type int

const (
	Unlimited Type = iota
	Ordinary       // 普通魔物
	Star           // 星怪
	Special        // 特殊
	MINI           // MINI Boss
	MVP            // MVP Boss
	DEAD           // 亡者 Boss
)

func (t Type) String() string {
	switch t {
	case Unlimited:
		return "不限"
	case Ordinary:
		return "普通魔物"
	case Star:
		return "星怪"
	case Special:
		return "特殊"
	case MINI:
		return "MINI Boss"
	case MVP:
		return "MVP Boss"
	case DEAD:
		return "亡者 Boss"
	default:
		return "未知"
	}
}

func (t Type) IsBoss() bool {
	switch t {
	case MINI, MVP, DEAD:
		return true
	default:
		return false
	}
}
