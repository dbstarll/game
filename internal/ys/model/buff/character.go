package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

func Character万叶扩散(mastery float64, dye elementals.Elemental) attr.AttributeModifier {
	return AddElementalDamageBonus(0.04*mastery, dye)
}
