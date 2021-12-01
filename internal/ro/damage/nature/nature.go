package nature

//属性
type Nature int

const (
	Unlimited Nature = iota
	Neutral          // 无
	Earth            // 地
	Wind             // 风
	Water            // 水
	Fire             // 火
	Holy             // 圣
	Dark             // 暗
	Ghost            // 念
	Undead           // 不死
	Poison           // 毒
)

func (n Nature) String() string {
	switch n {
	case Unlimited:
		return "不限"
	case Neutral:
		return "无"
	case Earth:
		return "地"
	case Wind:
		return "风"
	case Water:
		return "水"
	case Fire:
		return "火"
	case Holy:
		return "圣"
	case Dark:
		return "暗"
	case Ghost:
		return "念"
	case Undead:
		return "不死"
	case Poison:
		return "毒"
	default:
		return "未知"
	}
}
