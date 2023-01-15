package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

func Artifacts炽烈的炎之魔女2() attr.AttributeModifier {
	return AddElementalDamageBonus(15, elementals.Fire)
}

func Artifacts炽烈的炎之魔女4(layers int) attr.AttributeModifier {
	return attr.MergeAttributes(
		Artifacts炽烈的炎之魔女2(),
		AddReactionDamageBonus(40, reactions.Overload, reactions.Burn, reactions.Burgeon),
		AddReactionDamageBonus(15, reactions.Vaporize, reactions.Melt),
		AddElementalDamageBonus(15*[]float64{0, 0.5, 1.0, 1.5}[layers], elementals.Fire),
	)
}

func Artifacts冰风迷途的勇士2() attr.AttributeModifier {
	return AddElementalDamageBonus(15, elementals.Ice)
}

func Artifacts冰风迷途的勇士4(frozen bool) attr.AttributeModifier {
	criticalRate := 20.0
	if frozen {
		criticalRate += 20.0
	}
	return attr.MergeAttributes(
		Artifacts冰风迷途的勇士2(),
		AddCriticalRate(criticalRate),
	)
}

func Artifacts翠绿之影2() attr.AttributeModifier {
	return AddElementalDamageBonus(15, elementals.Wind)
}

func Artifacts翠绿之影4(dye elementals.Elemental) *attr.Modifier {
	return attr.NewModifier(
		attr.MergeAttributes(
			Artifacts翠绿之影2(),
			AddReactionDamageBonus(60, reactions.Swirl),
		),
		AddElementalResist(-40, dye),
	)
}
