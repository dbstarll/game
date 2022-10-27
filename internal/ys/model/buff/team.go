package buff

import (
	"github.com/dbstarll/game/internal/ys/model/attr"
)

// 双火共鸣
func TeamFire() attr.AttributeModifier {
	return AddAtkPercentage(25)
}

// 双冰共鸣
func TeamIce() attr.AttributeModifier {
	return AddCriticalRate(15)
}
