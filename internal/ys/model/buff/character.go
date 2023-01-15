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
