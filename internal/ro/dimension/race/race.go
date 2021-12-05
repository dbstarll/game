package race

import "gopkg.in/yaml.v3"

//种族
type Race int

const (
	Unlimited Race = iota
	Formless       // 无形
	Human          // 人形
	Plant          // 植物
	Animal         // 动物
	Insect         // 昆虫
	Fish           // 鱼贝
	Angel          // 天使
	Demon          // 恶魔
	Undead         // 不死
	Dragon         // 龙
)

func (r Race) String() string {
	switch r {
	case Unlimited:
		return "不限"
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
		return "未知"
	}
}

func (r *Race) UnmarshalYAML(value *yaml.Node) error {
	for i := Unlimited; i <= Dragon; i++ {
		if i.String() == value.Value {
			*r = i
			break
		}
	}
	return nil
}
