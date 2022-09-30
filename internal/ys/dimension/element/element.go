package element

// 元素类型
type Element int

const (
	Wind    Element = iota // 风
	Fire                   // 火
	Water                  // 水
	Ice                    // 冰
	Thunder                // 雷
	Earth                  // 岩
	Grass                  // 草
)

var Elements = []Element{
	Wind,
	Fire,
	Water,
	Ice,
	Thunder,
	Earth,
	Grass,
}

func (e Element) String() string {
	switch e {
	case Wind:
		return "风"
	case Fire:
		return "火"
	case Water:
		return "水"
	case Ice:
		return "冰"
	case Thunder:
		return "雷"
	case Earth:
		return "岩"
	case Grass:
		return "草"
	default:
		if e < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
