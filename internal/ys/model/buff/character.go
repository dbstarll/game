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
