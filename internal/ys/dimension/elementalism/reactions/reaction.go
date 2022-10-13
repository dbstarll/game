package reactions

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"
)

// 元素反应
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
		return "未知"
	}
}

func (r Reaction) Classify() classifies.Classify {
	switch r {
	case Vaporize, Melt:
		return classifies.Amplify
	case Crystallize:
		return classifies.Crystal
	case Overload, Superconduct, ElectroCharged, Shattered, Swirl, Frozen, Burn, Bloom, Hyperbloom, Burgeon:
		return classifies.Upheaval
	case Catalyze, Quicken, Aggravate, Spread:
		return classifies.Intensify
	default:
		return -1
	}
}
