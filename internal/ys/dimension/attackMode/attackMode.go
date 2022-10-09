package attackMode

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
