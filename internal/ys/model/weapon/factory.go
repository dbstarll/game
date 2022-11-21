package weapon

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"time"
)

var (
	Factory无工之剑 = func(refine int) *Weapon {
		return New(5, weaponType.Claymore, "无工之剑", Base(90, refine, 608, buff.AddAtkPercentage(49.6)),
			buff.AddShieldStrength(float64(15+refine*5)),
			buff.Superposition(5, time.Second*8, time.Millisecond*300, buff.AddAtkPercentage(2*float64(3+refine))),
		)
	}
	Factory黑岩斩刀 = func(refine int) *Weapon {
		return New(4, weaponType.Claymore, "黑岩斩刀", Base(90, refine, 510, buff.AddCriticalDamage(55.1)),
			buff.Superposition(3, time.Second*30, 0, buff.AddAtkPercentage(float64(9+3*refine))),
		)
	}
	Factory螭骨剑 = func(refine int) *Weapon {
		return New(5, weaponType.Claymore, "螭骨剑", Base(90, refine, 509, buff.AddCriticalRate(27.6)),
			buff.Superposition(5, 0, time.Second*4, buff.AddDamageBonus(5.0+float64(refine))),
			buff.Superposition(5, 0, time.Second*4, buff.AddIncomingDamageBonus([]float64{3.0, 2.7, 2.4, 2.2, 2.0}[refine-1])),
		)
	}
	Factory原木刀 = func(refine int) *Weapon {
		return New(4, weaponType.Sword, "原木刀", Base(90, refine, 565, buff.AddEnergyRecharge(30.6)))
	}
	Factory雾切之回光 = func(refine, layers int, elemental elementals.Elemental) *Weapon {
		return New(5, weaponType.Sword, "雾切之回光", Base(90, refine, 674, buff.AddCriticalDamage(44.1)),
			buff.AddAllElementalDamageBonus(float64(9+refine*3)),
			buff.AddElementalDamageBonus(float64([]int{0, 6, 12, 21}[layers]+refine*[]int{0, 2, 4, 7}[layers]), elemental), // 巴印层数
		)
	}
	Factory祭礼残章 = func(refine int) *Weapon {
		return New(4, weaponType.Catalyst, "祭礼残章", Base(90, refine, 454, buff.AddElementalMastery(221)))
	}
)
