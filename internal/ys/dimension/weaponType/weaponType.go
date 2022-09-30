package weaponType

// 武器类型
type WeaponType int

const (
	Sword    WeaponType = iota // 单手剑
	Claymore                   // 双手剑
	Bow                        // 弓
	Polearm                    // 长柄武器
	Catalyst                   // 法器
)

var WeaponTypes = []WeaponType{
	Sword,
	Claymore,
	Bow,
	Polearm,
	Catalyst,
}

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "单手剑"
	case Claymore:
		return "双手剑"
	case Bow:
		return "弓"
	case Polearm:
		return "长柄武器"
	case Catalyst:
		return "法器"
	default:
		if w < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
