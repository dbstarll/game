package attackMode

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
)

type AttackMode int

const (
	NormalAttack   AttackMode = iota // 普通攻击
	ChargedAttack                    // 重击
	PlungeAttack                     // 下坠攻击
	ElementalSkill                   // 元素战技
	ElementalBurst                   // 元素爆发
)

var AttackModes = []AttackMode{
	NormalAttack,
	ChargedAttack,
	PlungeAttack,
	ElementalSkill,
	ElementalBurst,
}

func (m AttackMode) String() string {
	switch m {
	case NormalAttack:
		return "普通攻击"
	case ChargedAttack:
		return "重击"
	case PlungeAttack:
		return "下坠攻击"
	case ElementalSkill:
		return "元素战技"
	case ElementalBurst:
		return "元素爆发"
	default:
		if m < 0 {
			return "不限"
		} else {
			return "未知"
		}
	}
}

// 计算攻击模式对应的元素属性.
// @param weapon 所使用的武器类型
// @param characterElemental 角色的自身属性
// @return 攻击模式对应的元素属性
func (m AttackMode) Elemental(weapon weaponType.WeaponType, characterElemental elementals.Elemental) elementals.Elemental {
	switch m {
	case NormalAttack, PlungeAttack:
		if weapon == weaponType.Catalyst {
			return characterElemental
		} else {
			return elementals.Physical
		}
	case ChargedAttack:
		if weapon == weaponType.Bow || weapon == weaponType.Catalyst {
			return characterElemental
		} else {
			return elementals.Physical
		}
	default:
		return characterElemental
	}
}
