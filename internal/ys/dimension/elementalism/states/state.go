package states

// 元素状态
type State int

const (
	Frozen         State = iota // 冻结
	Superconduct                // 超导
	ElectroCharged              // 感电
	Crystallize                 // 结晶
	Burn                        // 燃烧
	Quicken                     // 原激化
	Bloom                       // 原绽放(模拟草原核)
)

var States = []State{
	Frozen,
	Superconduct,
	ElectroCharged,
	Crystallize,
	Burn,
	Quicken,
	Bloom,
}

func (s State) IsValid() bool {
	return s >= Frozen && s <= Bloom
}

func (s State) IsMiddle() bool {
	return s == Frozen || s == Quicken || s == Bloom
}

func (s State) String() string {
	switch s {
	case Frozen:
		return "冻结"
	case Superconduct:
		return "超导"
	case ElectroCharged:
		return "感电"
	case Crystallize:
		return "结晶"
	case Burn:
		return "燃烧"
	case Quicken:
		return "原激化"
	case Bloom:
		return "原绽放"
	default:
		return "未知"
	}
}

// TODO 状态的产生
//   草元素 ++ 水元素 = 草原核
//   草元素 <- 雷元素 = 原激化(时间由草元素量决定)
//   草元素 <- 火元素 = 燃烧(持续的火元素伤害，持续直到草元素消耗完，4次/秒)
//   水元素 ++ 冰元素 = 冻结(时间由冰元素量决定)
//   水元素 ++ 雷元素 = 感电(持续的雷元素伤害，反应伤害间隔为1秒)
//   冰元素 ++ 雷元素 = 超导(范围性冰元素伤害，持续8秒，降低目标40%物理抗性)
//   雷元素 ++ 火元素 = 超载(范围性火元素伤害，爆炸)
//   风元素 -> 其他元素 = 扩散(风元素伤害+被扩散元素伤害)
//   岩元素 -> 其他元素 = 结晶(拾取获得对应元素护盾)
//   所有草元素相关反应，都是吃后手触发者的精通

// TODO 状态反应
//   草原核 <- 雷元素 = 超绽放(草元素伤害)
//   草原核 <- 火元素 = 烈绽放(草元素伤害)
//   原激化 <- 雷元素 = 超激化(雷元素伤害，不影响原激化持续时间)
//   原激化 <- 草元素 = 蔓激化(草元素伤害，不影响原激化持续时间)
//   冻结 <- 强力攻击 = 碎冰(物理伤害)
