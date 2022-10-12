package attackMode

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elemental"
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

func (m AttackMode) Elemental(weapon weaponType.WeaponType, elemental elemental.Elemental) elemental.Elemental {
	switch m {
	case NormalAttack, PlungeAttack:
		if weapon == weaponType.Catalyst {
			return elemental
		} else {
			return -1
		}
	case ChargedAttack:
		if weapon == weaponType.Bow || weapon == weaponType.Catalyst {
			return elemental
		} else {
			return -1
		}
	default:
		return elemental
	}
}

func (m AttackMode) DamageBonusPoint() point.Point {
	switch m {
	case NormalAttack:
		return point.NormalAttackDamageBonus
	case ChargedAttack:
		return point.ChargedAttackDamageBonus
	case PlungeAttack:
		return point.PlungeAttackDamageBonus
	case ElementalSkill:
		return point.ElementalSkillDamageBonus
	case ElementalBurst:
		return point.ElementalBurstDamageBonus
	default:
		return -1
	}
}

func (m AttackMode) FactorBonusPoint() point.Point {
	switch m {
	case NormalAttack:
		return point.NormalAttackFactorBonus
	case ChargedAttack:
		return point.ChargedAttackFactorBonus
	case PlungeAttack:
		return point.PlungeAttackFactorBonus
	case ElementalSkill:
		return point.ElementalSkillFactorBonus
	case ElementalBurst:
		return point.ElementalBurstFactorBonus
	default:
		return -1
	}
}
