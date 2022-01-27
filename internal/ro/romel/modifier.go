package romel

import (
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"github.com/dbstarll/game/internal/ro/model/general"
)

type BuffModifier func(val float64) model.CharacterModifier

var buffModifiers = &map[string]BuffModifier{
	// 素质属性
	"全能力": func(val float64) model.CharacterModifier {
		return buff.Quality(int(val))
	},
	"力量": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Str: int(val)})
	},
	"敏捷": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Agi: int(val)})
	},
	"体质": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Vit: int(val)})
	},
	"智力": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Int: int(val)})
	},
	"灵巧": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Dex: int(val)})
	},
	"灵巧分别": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Dex: int(val)})
	},
	"幸运": func(val float64) model.CharacterModifier {
		return model.AddQuality(&model.Quality{Luk: int(val)})
	},

	// 通用增益
	"暴击": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Critical: int(val)})
	},
	"暴击防护": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalResist: int(val)})
	},
	"普攻攻击力": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Ordinary: int(val)})
	},
	"普攻攻击": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Ordinary: int(val)})
	},
	"生命上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hp: int(val)})
	},
	"魔法上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Sp: int(val)})
	},
	"命中": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hit: int(val)})
	},
	"闪避": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Dodge: int(val)})
	},

	// 物理增益
	"物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"物理攻击分别": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Defence: val})
	},
	"精炼物攻": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Refine: val})
	},

	// 魔法增益
	"魔法攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Attack: val})
	},
	"魔法防御": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Defence: val})
	},
	"精炼魔攻": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Refine: val})
	},

	// 物理&魔法增益
	"物理、魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Defence: val}),
			model.AddGains(true, &model.Gains{Defence: val}),
		)
	},
}

var percentageBuffModifiers = &map[string]BuffModifier{}

func (b Buff) find(key string, val float64, percentage bool) (model.CharacterModifier, bool) {
	if percentage {
		if fn, exist := (*percentageBuffModifiers)[key]; exist {
			return fn(val), true
		} else {
			return nil, false
		}
	} else if fn, exist := (*buffModifiers)[key]; exist {
		return fn(val), true
	} else {
		return nil, false
	}
}
