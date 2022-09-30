package elemental

// 元素类型
type Elemental int

const (
	Anemo   Elemental = iota // 风
	Pyro                     // 火
	Hydro                    // 水
	Cryo                     // 冰
	Electro                  // 雷
	Geo                      // 岩
	Dendro                   // 草
)

var Elements = []Elemental{
	Anemo,
	Pyro,
	Hydro,
	Cryo,
	Electro,
	Geo,
	Dendro,
}

func (e Elemental) String() string {
	switch e {
	case Anemo:
		return "风"
	case Pyro:
		return "火"
	case Hydro:
		return "水"
	case Cryo:
		return "冰"
	case Electro:
		return "雷"
	case Geo:
		return "岩"
	case Dendro:
		return "草"
	default:
		if e < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
