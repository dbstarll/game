package buff

import (
	"github.com/dbstarll/game/internal/ys/dimension/attribute/point"
	"github.com/dbstarll/game/internal/ys/model/attr"
)

func Weapon赤沙之杖(refine, layers int) attr.AttributeModifier {
	primary, second := 39+13*refine, (21+7*refine)*layers
	return func(attributes *attr.Attributes) func() {
		return AddAtk(int(attributes.Get(point.ElementalMastery) * float64(primary+second) / 100.0))(attributes)
	}
}
