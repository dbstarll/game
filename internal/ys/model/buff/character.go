package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"math"
)

func Character万叶扩散(mastery float64, dye elementals.Elemental) attr.AttributeModifier {
	return AddElementalDamageBonus(0.04*mastery, dye)
}

func Character纳西妲净善摄受明论(mastery float64) attr.AttributeModifier {
	return AddElementalMastery(math.Min(mastery, 1000) * 0.25)
}

func Character纳西妲慧明缘觉智论() attr.AttributeModifier {
	return func(attributes *attr.Attributes) func() {
		if mastery := attributes.Get(point.ElementalMastery); mastery > 200 {
			base := math.Min(mastery-200, 800)
			return attr.MergeAttributes(
				AddAttackDamageBonus(base*0.1, attackMode.ElementalSkill),
				AddCriticalRate(base*0.03),
			)(attributes)
		} else {
			return attr.NopCallBack
		}
	}
}

func Character雷电将军殊胜之御体() attr.AttributeModifier {
	return func(attributes *attr.Attributes) func() {
		if recharge := attributes.Get(point.EnergyRecharge); recharge > 100 {
			return AddElementalDamageBonus((recharge-100)*0.4, elementals.Electric)(attributes)
		} else {
			return attr.NopCallBack
		}
	}
}

func Character雷电将军恶曜开眼(energyCost int, add float64) attr.AttributeModifier {
	return AddAttackDamageBonus(float64(energyCost)*add, attackMode.ElementalBurst)
}

func Character九条裟罗六命(dye elementals.Elemental) attr.AttributeModifier {
	if dye == elementals.Electric {
		return AddCriticalDamage(60)
	} else {
		return attr.NopAttributeModifier
	}
}

func Character胡桃彼岸蝶舞(baseHp, baseAtk, atkPreMaxHp float64) attr.AttributeModifier {
	return func(attributes *attr.Attributes) func() {
		finalHp := baseHp*(1.0+attributes.Get(point.HpPercentage)/100) + attributes.Get(point.Hp)
		return AddAtk(int(math.Round(math.Min(finalHp*atkPreMaxHp/100, baseAtk*4))))(attributes)
	}
}

func Character神里绫华天罪国罪镇词() attr.AttributeModifier {
	return AddAttackDamageBonus(30, attackMode.NormalAttack, attackMode.ChargedAttack)
}

func Character神里绫华寒天宣命祝词() attr.AttributeModifier {
	return AddElementalDamageBonus(18, elementals.Ice)
}

func Character申鹤大洞弥罗尊法() attr.AttributeModifier {
	return AddElementalDamageBonus(15, elementals.Ice)
}

func Character申鹤缚灵通真法印(longPress bool) attr.AttributeModifier {
	if longPress {
		return AddAttackDamageBonus(15, attackMode.NormalAttack, attackMode.ChargedAttack, attackMode.PlungeAttack)
	} else {
		return AddAttackDamageBonus(15, attackMode.ElementalSkill, attackMode.ElementalBurst)
	}
}

func Character申鹤Q(character attr.Character) *attr.Modifier {
	if character != nil {
		if 抗性降低 := character.GetAction(attackMode.ElementalBurst, "抗性降低"); 抗性降低 != nil {
			return attr.NewModifier(Character申鹤大洞弥罗尊法(), AddElementalResist(-抗性降低.DMG(), elementals.Ice, elementals.Physical))
		}
	}
	return attr.NewCharacterModifier(Character申鹤大洞弥罗尊法())
}

func Character申鹤E(character attr.Character, longPress bool, attackElement elementals.Elemental) attr.AttributeModifier {
	if character != nil && attackElement == elementals.Ice {
		if 伤害值提升 := character.GetAction(attackMode.ElementalSkill, "伤害值提升"); 伤害值提升 != nil {
			baseAtk, final := character.BaseAttr(point.Atk)+character.WeaponAttr(point.Atk), character.FinalAttributes()
			addAtk, addAtkPercentage := final.Get(point.Atk), final.Get(point.AtkPercentage)
			finalAtk := baseAtk*(1+addAtkPercentage/100) + addAtk
			return attr.MergeAttributes(AddAtk(int(finalAtk*伤害值提升.DMG()/100)), Character申鹤缚灵通真法印(longPress))
		}
	}
	return Character申鹤缚灵通真法印(longPress)
}
