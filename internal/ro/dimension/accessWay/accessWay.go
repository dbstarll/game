package accessWay

import "gopkg.in/yaml.v3"

//宠物获取途径
type AccessWay int

const (
	Rank     AccessWay = iota // 组队竞技赛段位奖励
	Catch                     // 通过捕捉魔物获得
	Shop                      // 通过宠物材料商店获得
	Activity                  // 活动期间购买获得
	Compound                  // 通过宠物融合获得
	Pvp                       // 组队竞技赛结算奖励
	Checkin                   // 通过签到获得
	Tower                     // 达纳托斯之塔
	None                      // 暂无获取途径
)

func (n AccessWay) String() string {
	switch n {
	case Rank:
		return "组队竞技赛段位奖励"
	case Catch:
		return "通过捕捉魔物获得"
	case Shop:
		return "通过宠物材料商店获得"
	case Activity:
		return "活动期间购买获得"
	case Compound:
		return "通过宠物融合获得"
	case Pvp:
		return "组队竞技赛结算奖励"
	case Checkin:
		return "通过签到获得"
	case Tower:
		return "达纳托斯之塔"
	case None:
		return "暂无获取途径"
	default:
		return "未知"
	}
}

func (n *AccessWay) UnmarshalYAML(value *yaml.Node) error {
	for i := Rank; i <= None; i++ {
		if i.String() == value.Value {
			*n = i
			break
		}
	}
	return nil
}
