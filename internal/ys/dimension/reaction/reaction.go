package reaction

type Reaction int

const (
	Vaporize       Reaction = iota // 蒸发
	Melt                           // 融化
	Overload                       // 超载
	Superconduct                   // 超导
	ElectroCharged                 // 感电
	Shattered                      // 碎冰
	Swirl                          // 扩散
	Crystallize                    // 结晶
	Frozen                         // 冻结
	Burn                           // 燃烧
	Bloom                          // 绽放
	Hyperbloom                     // 超绽放
	Burgeon                        // 烈绽放
	Catalyze                       // 激化
	Quicken                        // 原激化
	Aggravate                      // 超激化
	Spread                         // 蔓激化
)

var (
	Reactions = []Reaction{
		Vaporize,
		Melt,
		Overload,
		Superconduct,
		ElectroCharged,
		Shattered,
		Swirl,
		Crystallize,
		Frozen,
		Burn,
		Bloom,
		Hyperbloom,
		Burgeon,
		Catalyze,
		Quicken,
		Aggravate,
		Spread,
	}
)

func (r Reaction) String() string {
	switch r {
	case Vaporize:
		return "蒸发"
	case Melt:
		return "融化"
	case Overload:
		return "超载"
	case Superconduct:
		return "超导"
	case ElectroCharged:
		return "感电"
	case Shattered:
		return "碎冰"
	case Swirl:
		return "扩散"
	case Crystallize:
		return "结晶"
	case Frozen:
		return "冻结"
	case Burn:
		return "燃烧"
	case Bloom:
		return "绽放"
	case Hyperbloom:
		return "超绽放"
	case Burgeon:
		return "烈绽放"
	case Catalyze:
		return "激化"
	case Quicken:
		return "原激化"
	case Aggravate:
		return "超激化"
	case Spread:
		return "蔓激化"
	default:
		if r < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}

func (r Reaction) Classify() Classify {
	switch r {
	case Vaporize, Melt:
		return Amplify
	case Crystallize:
		return Crystal
	case Overload, Superconduct, ElectroCharged, Shattered, Swirl, Frozen, Burn, Bloom, Hyperbloom, Burgeon:
		return Upheaval
	case Catalyze, Quicken, Aggravate, Spread:
		return Intensify
	default:
		return -1
	}
}

type Factor struct {
	reaction Reaction
	factor   float64
}

func NewFactor(reaction Reaction, factor float64) *Factor {
	return &Factor{reaction: reaction, factor: factor}
}

func (f *Factor) GetReaction() Reaction {
	return f.reaction
}

func (f *Factor) GetFactor() float64 {
	return f.factor
}
