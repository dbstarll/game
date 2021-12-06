package race

import "gopkg.in/yaml.v3"

//种族
type Race int

const (
	Unlimited Race = iota
	Animal         // 动物
	Human          // 人形
	Demon          // 恶魔
	Plant          // 植物
	Undead         // 不死
	Formless       // 无形
	Fish           // 鱼贝
	Angel          // 天使
	Insect         // 昆虫
	Dragon         // 龙
)

func (r Race) String() string {
	switch r {
	case Unlimited:
		return "不限"
	case Animal:
		return "动物"
	case Human:
		return "人形"
	case Demon:
		return "恶魔"
	case Plant:
		return "植物"
	case Undead:
		return "不死"
	case Formless:
		return "无形"
	case Fish:
		return "鱼贝"
	case Angel:
		return "天使"
	case Insect:
		return "昆虫"
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
