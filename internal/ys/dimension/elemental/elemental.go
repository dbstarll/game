package elemental

// 元素类型
type Elemental int

const (
	Pyro    Elemental = iota // 火
	Hydro                    // 水
	Dendro                   // 草
	Electro                  // 雷
	Anemo                    // 风
	Cryo                     // 冰
	Geo                      // 岩
)

var Elementals = []Elemental{
	Pyro,
	Hydro,
	Dendro,
	Electro,
	Anemo,
	Cryo,
	Geo,
}

func (e Elemental) String() string {
	switch e {
	case Pyro:
		return "火"
	case Hydro:
		return "水"
	case Dendro:
		return "草"
	case Electro:
		return "雷"
	case Anemo:
		return "风"
	case Cryo:
		return "冰"
	case Geo:
		return "岩"
	default:
		if e < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
