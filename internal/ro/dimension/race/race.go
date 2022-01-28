package race

import "gopkg.in/yaml.v3"

//种族
type Race int

const (
	Formless Race = iota // 无形
	Human                // 人形
	Plant                // 植物
	Animal               // 动物
	Insect               // 昆虫
	Fish                 // 鱼贝
	Angel                // 天使
	Demon                // 恶魔
	Undead               // 不死
	Dragon               // 龙
)

var Races = []Race{
	Formless,
	Human,
	Plant,
	Animal,
	Insect,
	Fish,
	Angel,
	Demon,
	Undead,
	Dragon,
}

func (r Race) String() string {
	switch r {
	case Formless:
		return "无形"
	case Human:
		return "人形"
	case Plant:
		return "植物"
	case Animal:
		return "动物"
	case Insect:
		return "昆虫"
	case Fish:
		return "鱼贝"
	case Angel:
		return "天使"
	case Demon:
		return "恶魔"
	case Undead:
		return "不死"
	case Dragon:
		return "龙"
	default:
		if r < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}

func (r Race) Name() string {
	switch r {
	case Formless:
		return "无形种族"
	case Human:
		return "人形种族"
	case Plant:
		return "植物种族"
	case Animal:
		return "动物种族"
	case Insect:
		return "昆虫种族"
	case Fish:
		return "鱼贝种族"
	case Angel:
		return "天使种族"
	case Demon:
		return "恶魔种族"
	case Undead:
		return "不死种族"
	case Dragon:
		return "龙族"
	default:
		if r < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}

func (r *Race) UnmarshalYAML(value *yaml.Node) error {
	for i := Formless; i <= Dragon; i++ {
		if i.String() == value.Value {
			*r = i
			break
		}
	}
	return nil
}
