package weaponType

// 武器类型
type WeaponType int

const (
	Sword    WeaponType = iota //单手剑
	BigSword                   //双手剑
	Bow                        //弓
	Spear                      //长柄武器
	Magic                      //法器
)

var WeaponTypes = []WeaponType{
	Sword,
	BigSword,
	Bow,
	Spear,
	Magic,
}

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "单手剑"
	case BigSword:
		return "双手剑"
	case Bow:
		return "弓"
	case Spear:
		return "长柄武器"
	case Magic:
		return "法器"
	default:
		if w < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}
