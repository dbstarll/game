package romel

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/abnormal"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"github.com/dbstarll/game/internal/ro/model/general"
)

type BuffModifier func(val float64) model.CharacterModifier

var inited = false

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
	"使暴击翻倍": func(val float64) model.CharacterModifier {
		return func(character *model.Character) func() {
			return model.AddGeneral(&general.General{Critical: character.Profits.General.Critical})(character)
		}
	},
	"暴击防护": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalResist: int(val)})
	},
	"普攻攻击力": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Ordinary: int(val)})
	},
	"普通攻击力": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Ordinary: int(val)})
	},
	"普攻攻击": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Ordinary: int(val)})
	},
	"生命上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hp: int(val)})
	},
	"MaxHp": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hp: int(val)})
	},
	"魔法上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Sp: int(val)})
	},
	"MaxSp": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Sp: int(val)})
	},
	"命中": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hit: int(val)})
	},
	"闪避": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Dodge: int(val)})
	},
	"可变吟唱时间": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SingElasticity: val})
	},
	"Zeny": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Zeny: int(val)})
	},
	"包包格子": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Bag: int(val)})
	},

	// 物理增益
	"物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"增加自身物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"自身物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"获得物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Attack: val})
	},
	"使自身与所有生命体的物理攻击": func(val float64) model.CharacterModifier {
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
	"增加自身魔法攻击": func(val float64) model.CharacterModifier {
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
	"物理防御和魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Defence: val}),
			model.AddGains(true, &model.Gains{Defence: val}),
		)
	},
	"物理攻击和魔法攻击": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Attack: val}),
			model.AddGains(true, &model.Gains{Attack: val}),
		)
	},
	"物理攻击力与魔法攻击力": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Attack: val}),
			model.AddGains(true, &model.Gains{Attack: val}),
		)
	},
	"避免陷入晕眩的状态": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{abnormal.Vertigo: 100})
	},
}

var percentageBuffModifiers = &map[string]BuffModifier{
	"暴击": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Critical: int(val)})
	},
	"暴击伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamage: val})
	},
	"暴击伤害提": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamage: val})
	},
	"暴击伤害增加": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamage: val})
	},
	"暴伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamage: val})
	},
	"暴击伤害分别额外增加": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamage: val})
	},
	"受到暴击伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamageResist: -val})
	},
	"暴伤减免": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{CriticalDamageResist: val})
	},
	"普攻伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{OrdinaryDamage: val})
	},
	"普通攻击伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{OrdinaryDamage: val})
	},
	"剑士系普通攻击伤害": func(val float64) model.CharacterModifier {
		//TODO 限制剑士系
		return model.AddGeneral(&general.General{OrdinaryDamage: val})
	},
	"普攻伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{OrdinaryResist: val})
	},
	"对MVP、Mini魔物造成的伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对MVP、Mini魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对MVP、Mini魔物魔物增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对MVP、Mini魔物增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对MVP、Mini魔物的伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对MVP/Mini增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val})
	},
	"对普通魔物（不包含MVP、Mini魔物）增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{NoMVP: val})
	},
	"对魔物（包含MVP、Mini魔物）增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVP: val, NoMVP: val})
	},
	"受到MVP、Mini魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVPResist: -val})
	},
	"受MVP、Mini魔物造成的伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVPResist: -val})
	},
	"受到MVP、Mini魔物造成的伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MVPResist: -val})
	},
	"受普通魔物（不包含MVP、Mini魔物）造成的伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{NoMVPResist: -val})
	},
	"PVP/GVG中对玩家伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Player: val})
	},
	"对玩家增伤": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Player: val})
	},
	"PVP/GVG中受玩家伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{PlayerResist: -val})
	},
	"技能伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Skill: val})
	},
	"技能攻击伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Skill: val})
	},
	"技能伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SkillResist: val})
	},
	"装备攻速": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{AttackSpeed: val})
	},
	"攻击速度": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{AttackSpeed: val})
	},
	"攻速": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{AttackSpeed: val})
	},
	"且自身移动速度": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MoveSpeed: val})
	},
	"移动速度": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MoveSpeed: val})
	},
	"自身移动速度": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MoveSpeed: val})
	},
	"移速": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MoveSpeed: val})
	},
	"命中": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Hit: int(val)})
	},
	"闪避": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{DodgePer: val})
	},
	"生命上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{HpPer: val})
	},
	"魔法上限": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SpPer: val})
	},
	"治疗加成": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Cure: val})
	},
	"受治疗加成": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{Cured: val})
	},
	"可变吟唱时间": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SingPerElasticity: val})
	},
	"可变吟唱": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SingPerElasticity: val})
	},
	"固定吟唱时间": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SingPerFixed: val})
	},
	"技能冷却": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SkillCooling: val})
	},
	"技能延迟": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{SkillDelay: val})
	},
	"法术普攻暴击概率": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MagicOrdinaryCriticalRate: val})
	},
	"法术普攻暴击伤害": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{MagicOrdinaryCriticalDamage: val})
	},
	"击杀魔物Base经验": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{BaseExp: val})
	},
	"击杀魔物Job经验": func(val float64) model.CharacterModifier {
		return model.AddGeneral(&general.General{JobExp: val})
	},

	// 物理、魔法兼得的增益
	"物理和魔法穿刺": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Spike: val}),
			model.AddGains(true, &model.Gains{Spike: val}),
		)
	},
	"物理穿刺和魔法穿刺": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Spike: val}),
			model.AddGains(true, &model.Gains{Spike: val}),
		)
	},
	"物理和魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{DefencePer: val}),
			model.AddGains(true, &model.Gains{DefencePer: val}),
		)
	},
	"物理、魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{DefencePer: val}),
			model.AddGains(true, &model.Gains{DefencePer: val}),
		)
	},
	"物理防御、魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{DefencePer: val}),
			model.AddGains(true, &model.Gains{DefencePer: val}),
		)
	},
	"物理、魔法攻击": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{AttackPer: val}),
			model.AddGains(true, &model.Gains{AttackPer: val}),
		)
	},
	"物理、魔法伤害": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Damage: val}),
			model.AddGains(true, &model.Gains{Damage: val}),
		)
	},
	"物理伤害、魔法伤害": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Damage: val}),
			model.AddGains(true, &model.Gains{Damage: val}),
		)
	},
	"在隐匿的状态下，自身物理、魔法伤害": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Damage: val}),
			model.AddGains(true, &model.Gains{Damage: val}),
		)
	},
	"最终伤害": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Damage: val}),
			model.AddGains(true, &model.Gains{Damage: val}),
		)
	},
	"忽视物理、魔法防御": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Ignore: val}),
			model.AddGains(true, &model.Gains{Ignore: val}),
		)
	},
	"物伤、魔伤减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"物理、魔法伤害减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"冰冻状态下，物伤、魔伤减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"物理和魔法伤害减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"物伤减免与魔伤减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"物伤减免和魔伤减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"物理伤害减免和魔法伤害减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"魔伤减免、物伤减免": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: val}),
			model.AddGains(true, &model.Gains{Resist: val}),
		)
	},
	"受到的伤害": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddGains(false, &model.Gains{Resist: -val}),
			model.AddGains(true, &model.Gains{Resist: -val}),
		)
	},

	// 物理增益
	"物理穿刺": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Spike: val})
	},
	"物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"物理攻击力": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"装备者的物理防御和魔法防御不再降低，物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"使自身与所有生命体的物理攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"远程物理攻击": func(val float64) model.CharacterModifier {
		//TODO 限制远程武器
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"远距离物理攻击力": func(val float64) model.CharacterModifier {
		//TODO 限制远程武器
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"近战物理攻击": func(val float64) model.CharacterModifier {
		//TODO 限制近战物理
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"近战物理攻击力": func(val float64) model.CharacterModifier {
		//TODO 限制近战物理
		return model.AddGains(false, &model.Gains{AttackPer: val})
	},
	"物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{DefencePer: val})
	},
	"展开绝对领域：物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{DefencePer: val})
	},
	"异常状态下，物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{DefencePer: val})
	},
	"则物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{DefencePer: val})
	},
	"物理伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Damage: val})
	},
	"远程物理伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{RemoteDamage: val})
	},
	"近战物理伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{NearDamage: val})
	},
	"忽视物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Ignore: val})
	},
	"洞察：忽视物理防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Ignore: val})
	},
	"忽视物防": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Ignore: val})
	},
	"物伤减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Resist: val})
	},
	"但物伤减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Resist: val})
	},
	"物理伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Resist: val})
	},
	"精炼物伤减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{RefineResist: val})
	},
	"物伤防御": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Resist: val})
	},
	"受到远距离物理伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{RemoteResist: val})
	},
	"受到近战物理伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{NearResist: val})
	},
	"受到物理伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(false, &model.Gains{Resist: -val})
	},

	// 魔法增益
	"魔法穿刺": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Spike: val})
	},
	"魔法攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{AttackPer: val})
	},
	"获得魔法攻击": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{AttackPer: val})
	},
	"魔法防御": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{DefencePer: val})
	},
	"获得魔法防御": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{DefencePer: val})
	},
	"且额外获得魔法防御": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{DefencePer: val})
	},
	"魔法伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"自身魔法伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"获得魔法伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"魔法范围型技能伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"提升贤者系范围型技能伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"法师系范围型技能伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"魔语者系范围型技能伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Damage: val})
	},
	"忽视魔法防御": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Ignore: val})
	},
	"魔伤减免": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Resist: val})
	},
	"魔法伤害减免": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Resist: val})
	},
	"魔法减免": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Resist: val})
	},
	"受到魔法伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Resist: -val})
	},
	"受到的所有魔法伤害": func(val float64) model.CharacterModifier {
		return model.AddGains(true, &model.Gains{Resist: -val})
	},

	// 属性攻击%
	"风、地、水、火属性攻击": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Wind:  val,
			nature.Earth: val,
			nature.Water: val,
			nature.Fire:  val,
		})
	},
	"风、地、水、火、无属性攻击": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Wind:    val,
			nature.Earth:   val,
			nature.Water:   val,
			nature.Fire:    val,
			nature.Neutral: val,
		})
	},
	"火属性、水属性、地属性、风属性攻击": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Fire:  val,
			nature.Water: val,
			nature.Earth: val,
			nature.Wind:  val,
		})
	},
	"风、火、水、地属性攻击": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Wind:  val,
			nature.Fire:  val,
			nature.Water: val,
			nature.Earth: val,
		})
	},
	"圣属性、暗属性、无属性攻击": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Holy:    val,
			nature.Dark:    val,
			nature.Neutral: val,
		})
	},
	"水属性和地属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureAttack(&map[nature.Nature]float64{
			nature.Water: val,
			nature.Earth: val,
		})
	},
	"所有属性攻击": func(val float64) model.CharacterModifier {
		return buff.AddNatureAttack(val)
	},

	// 属性增伤%
	"对水、地属性魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Water: val,
			nature.Earth: val,
		})
	},
	"对圣、暗属性魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Holy: val,
			nature.Dark: val,
		})
	},
	"对圣、念属性魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Holy:  val,
			nature.Ghost: val,
		})
	},
	"对暗、不死属性魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Dark:   val,
			nature.Undead: val,
		})
	},
	"对风、地属性魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureDamage(&map[nature.Nature]float64{
			nature.Wind:  val,
			nature.Earth: val,
		})
	},

	// 属性减伤%
	"受到风、地、水、火、无属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Wind:    -val,
			nature.Earth:   -val,
			nature.Water:   -val,
			nature.Fire:    -val,
			nature.Neutral: -val,
		})
	},
	"受无、风、地、水、火属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Neutral: -val,
			nature.Wind:    -val,
			nature.Earth:   -val,
			nature.Water:   -val,
			nature.Fire:    -val,
		})
	},
	"受到火、水、地、风属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Fire:  -val,
			nature.Water: -val,
			nature.Earth: -val,
			nature.Wind:  -val,
		})
	},
	"受到风、地、水、火、圣属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Wind:  -val,
			nature.Earth: -val,
			nature.Water: -val,
			nature.Fire:  -val,
			nature.Holy:  -val,
		})
	},
	"受到风、地、水、火属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Wind:  -val,
			nature.Earth: -val,
			nature.Water: -val,
			nature.Fire:  -val,
		})
	},
	"受火、水、地、风属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Fire:  -val,
			nature.Water: -val,
			nature.Earth: -val,
			nature.Wind:  -val,
		})
	},
	"火、水、地、风属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Fire:  -val,
			nature.Water: -val,
			nature.Earth: -val,
			nature.Wind:  -val,
		})
	},
	"受到无、圣、暗、念、毒属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Neutral: -val,
			nature.Holy:    -val,
			nature.Dark:    -val,
			nature.Ghost:   -val,
			nature.Poison:  -val,
		})
	},
	"受到无、毒、圣、暗、念属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Neutral: -val,
			nature.Poison:  -val,
			nature.Holy:    -val,
			nature.Dark:    -val,
			nature.Ghost:   -val,
		})
	},
	"受到毒属性、圣属性、暗属性、念属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Poison: -val,
			nature.Holy:   -val,
			nature.Dark:   -val,
			nature.Ghost:  -val,
		})
	},
	"受到毒、圣、暗、念、无属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Poison:  -val,
			nature.Holy:    -val,
			nature.Dark:    -val,
			nature.Ghost:   -val,
			nature.Neutral: -val,
		})
	},
	"受毒、圣、暗、念、无属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Poison:  -val,
			nature.Holy:    -val,
			nature.Dark:    -val,
			nature.Ghost:   -val,
			nature.Neutral: -val,
		})
	},
	"受到暗属性、念属性、不死属性、毒属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Dark:   -val,
			nature.Ghost:  -val,
			nature.Undead: -val,
			nature.Poison: -val,
		})
	},
	"受毒、圣、暗、念、无、火、水、地、风属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Poison:  -val,
			nature.Holy:    -val,
			nature.Dark:    -val,
			nature.Ghost:   -val,
			nature.Neutral: -val,
			nature.Fire:    -val,
			nature.Water:   -val,
			nature.Earth:   -val,
			nature.Wind:    -val,
		})
	},
	"受到地、火属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Earth: -val,
			nature.Fire:  -val,
		})
	},
	"受到其他所有属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Earth:  -val,
			nature.Wind:   -val,
			nature.Water:  -val,
			nature.Fire:   -val,
			nature.Holy:   -val,
			nature.Dark:   -val,
			nature.Ghost:  -val,
			nature.Undead: -val,
			nature.Poison: -val,
		})
	},
	"使自身5米范围内的队友（包括自身）受到地、火属性伤害": func(val float64) model.CharacterModifier {
		return model.AddNatureResist(&map[nature.Nature]float64{
			nature.Earth: -val,
			nature.Fire:  -val,
		})
	},

	// 种族增伤%
	"对全种族伤害": func(val float64) model.CharacterModifier {
		return buff.RaceDamage(val)
	},
	"全种族伤害": func(val float64) model.CharacterModifier {
		return buff.RaceDamage(val)
	},
	"对所有种族魔物伤害": func(val float64) model.CharacterModifier {
		return buff.RaceDamage(val)
	},
	"人形种族伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Human: val})
	},
	"人形种族加伤额外": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Human: val})
	},
	"对人形种族加伤": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Human: val})
	},
	"对恶魔系魔物增伤": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Demon: val})
	},
	"对植物系伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Plant: val})
	},
	"龙形种族加伤": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Dragon: val})
	},
	"对龙族种族增伤": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{race.Dragon: val})
	},
	"对动物、恶魔种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Animal: val,
			race.Demon:  val,
		})
	},
	"对昆虫、鱼贝种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Insect: val,
			race.Fish:   val,
		})
	},
	"对天使、恶魔种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Angel: val,
			race.Demon: val,
		})
	},
	"对不死、恶魔种族伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Undead: val,
			race.Demon:  val,
		})
	},
	"对动物、植物种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Animal: val,
			race.Plant:  val,
		})
	},
	"对植物、动物种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Plant:  val,
			race.Animal: val,
		})
	},
	"对植物、动物种族加伤": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Plant:  val,
			race.Animal: val,
		})
	},
	"对动物、昆虫种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Animal: val,
			race.Insect: val,
		})
	},
	"对非人形种族伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceDamage(&map[race.Race]float64{
			race.Formless: val,
			race.Plant:    val,
			race.Animal:   val,
			race.Insect:   val,
			race.Fish:     val,
			race.Angel:    val,
			race.Demon:    val,
			race.Undead:   val,
			race.Dragon:   val,
		})
	},

	// 种族减伤%
	"全种族减伤": func(val float64) model.CharacterModifier {
		return buff.AddRaceResist(val)
	},
	"对恶魔系魔物减伤": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{race.Demon: val})
	},
	"恶魔种族伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{race.Demon: -val})
	},
	"受恶魔种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{race.Demon: -val})
	},
	"恶魔种族减伤和不死种族减伤": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{
			race.Demon:  val,
			race.Undead: val,
		})
	},
	"其他种族减伤": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{
			race.Formless: val,
			race.Human:    val,
			race.Plant:    val,
			race.Animal:   val,
			race.Insect:   val,
			race.Fish:     val,
			race.Angel:    val,
			race.Dragon:   val,
		})
	},
	"受到植物、动物种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{
			race.Plant:  -val,
			race.Animal: -val,
		})
	},
	"受植物、动物种族魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddRaceResist(&map[race.Race]float64{
			race.Plant:  -val,
			race.Animal: -val,
		})
	},
	"人形种族加伤、减伤": func(val float64) model.CharacterModifier {
		return model.Merge(
			model.AddRaceDamage(&map[race.Race]float64{race.Human: val}),
			model.AddRaceResist(&map[race.Race]float64{race.Human: val}),
		)
	},

	//体型增伤%
	"对小、中、大型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeDamage(&map[shape.Shape]float64{
			shape.Large:  val,
			shape.Medium: val,
			shape.Small:  val,
		})
	},
	"对中体型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeDamage(&map[shape.Shape]float64{shape.Medium: val})
	},
	"对小体型增伤": func(val float64) model.CharacterModifier {
		return model.AddShapeDamage(&map[shape.Shape]float64{shape.Small: val})
	},
	"对中体型增伤": func(val float64) model.CharacterModifier {
		return model.AddShapeDamage(&map[shape.Shape]float64{shape.Medium: val})
	},
	"对大体型增伤": func(val float64) model.CharacterModifier {
		return model.AddShapeDamage(&map[shape.Shape]float64{shape.Large: val})
	},

	//体型减伤%
	"受到大、中、小型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{
			shape.Large:  -val,
			shape.Medium: -val,
			shape.Small:  -val,
		})
	},
	"受到中、小型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{
			shape.Medium: -val,
			shape.Small:  -val,
		})
	},
	"受到大、中型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{
			shape.Large:  -val,
			shape.Medium: -val,
		})
	},
	"受中体型魔物伤害": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{shape.Medium: -val})
	},
	"对小体型减伤": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{shape.Small: val})
	},
	"对中体型减伤": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{shape.Medium: val})
	},
	"对大体型减伤": func(val float64) model.CharacterModifier {
		return model.AddShapeResist(&map[shape.Shape]float64{shape.Large: val})
	},

	//异常状态抵抗%
	"定身、恐惧抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Fixed: val,
			abnormal.Fear:  val,
		})
	},
	"眩晕、冰冻抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Vertigo: val,
			abnormal.Frozen:  val,
		})
	},
	"眩晕，冰冻抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Vertigo: val,
			abnormal.Frozen:  val,
		})
	},
	"定身、冰冻抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Fixed:  val,
			abnormal.Frozen: val,
		})
	},
	"中毒抵抗和冰冻抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Poisoning: val,
			abnormal.Frozen:    val,
		})
	},
	"恐惧抵抗和黑暗抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Fear: val,
			abnormal.Dark: val,
		})
	},
	"恐惧抵抗和冰冻抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Fear:   val,
			abnormal.Frozen: val,
		})
	},
	"晕眩抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{abnormal.Vertigo: val})
	},
	"晕眩抵抗和定身抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Vertigo: val,
			abnormal.Fixed:   val,
		})
	},
	"晕眩抵抗和恐惧抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Vertigo: val,
			abnormal.Fear:    val,
		})
	},
	"晕眩抵抗、恐惧抵抗和石化抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Vertigo: val,
			abnormal.Fear:    val,
			abnormal.Petrify: val,
		})
	},
	"异常状态抵抗": func(val float64) model.CharacterModifier {
		return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{
			abnormal.Poisoning: val,
			abnormal.Bleed:     val,
			abnormal.Burn:      val,
			abnormal.Vertigo:   val,
			abnormal.Frozen:    val,
			abnormal.Petrify:   val,
			abnormal.Sleep:     val,
			abnormal.Fear:      val,
			abnormal.Fixed:     val,
			abnormal.Silent:    val,
			abnormal.Cursed:    val,
			abnormal.Dark:      val,
		})
	},
}

func init() {
	initModifier()
}

func initModifier() {
	//177
	if !inited {
		inited = true
		for _, item := range nature.Natures {
			_nature := item
			//属性攻击%
			(*percentageBuffModifiers)[fmt.Sprintf("%s属性攻击", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureAttack(&map[nature.Nature]float64{_nature: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("自身%s属性攻击", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureAttack(&map[nature.Nature]float64{_nature: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s属性攻击", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureAttack(&map[nature.Nature]float64{_nature: val})
			}
			//属性增伤%
			(*percentageBuffModifiers)[fmt.Sprintf("对%s属性魔物伤害", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureDamage(&map[nature.Nature]float64{_nature: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s属性魔物增伤", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureDamage(&map[nature.Nature]float64{_nature: val})
			}
			//属性减伤%
			(*percentageBuffModifiers)[fmt.Sprintf("受到%s属性伤害", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureResist(&map[nature.Nature]float64{_nature: -val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("受%s属性伤害", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureResist(&map[nature.Nature]float64{_nature: -val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s属性伤害减免", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureResist(&map[nature.Nature]float64{_nature: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s属性减伤", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureResist(&map[nature.Nature]float64{_nature: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s属性伤害减免", _nature)] = func(val float64) model.CharacterModifier {
				return model.AddNatureResist(&map[nature.Nature]float64{_nature: val})
			}
		}
		for _, item := range race.Races {
			_race := item
			//种族增伤%
			(*percentageBuffModifiers)[fmt.Sprintf("%s加伤", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s加伤", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s增伤", _race)] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s魔物伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s增伤", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s增伤", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceDamage(&map[race.Race]float64{_race: val})
			}

			//种族减伤%
			(*percentageBuffModifiers)[fmt.Sprintf("%s减伤", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s减伤", _race)] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s伤害减免", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("受到%s魔物伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: -val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("受%s伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: -val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("受到%s伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: -val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("受%s种族魔物伤害", _race.Name())] = func(val float64) model.CharacterModifier {
				return model.AddRaceResist(&map[race.Race]float64{_race: -val})
			}
		}
		for _, item := range shape.Shapes {
			_shape := item
			//体型增伤%
			(*percentageBuffModifiers)[fmt.Sprintf("对%s魔物伤害", _shape)] = func(val float64) model.CharacterModifier {
				return model.AddShapeDamage(&map[shape.Shape]float64{_shape: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s魔物的伤害", _shape)] = func(val float64) model.CharacterModifier {
				return model.AddShapeDamage(&map[shape.Shape]float64{_shape: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s体型增伤", _shape)] = func(val float64) model.CharacterModifier {
				return model.AddShapeDamage(&map[shape.Shape]float64{_shape: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("对%s体型魔物增伤", _shape)] = func(val float64) model.CharacterModifier {
				return model.AddShapeDamage(&map[shape.Shape]float64{_shape: val})
			}
			//体型减伤%
			(*percentageBuffModifiers)[fmt.Sprintf("受到%s魔物伤害", _shape)] = func(val float64) model.CharacterModifier {
				return model.AddShapeResist(&map[shape.Shape]float64{_shape: -val})
			}
		}
		for _, item := range abnormal.Abnormals {
			_abnormal := item
			//异常状态抵抗%
			(*percentageBuffModifiers)[fmt.Sprintf("%s抵抗", _abnormal)] = func(val float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: val})
			}
			(*percentageBuffModifiers)[fmt.Sprintf("%s抗性", _abnormal)] = func(val float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: val})
			}
			(*buffModifiers)[fmt.Sprintf("避免陷入%s状态", _abnormal)] = func(float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: 100})
			}
			(*buffModifiers)[fmt.Sprintf("避免陷入%s的状态", _abnormal)] = func(float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: 100})
			}
			(*buffModifiers)[fmt.Sprintf("%s免疫", _abnormal)] = func(float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: 100})
			}
			(*buffModifiers)[fmt.Sprintf("不会陷入%s效果", _abnormal)] = func(float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: 100})
			}
			(*buffModifiers)[fmt.Sprintf("避免进入%s状态", _abnormal)] = func(float64) model.CharacterModifier {
				return model.AddAbnormalResist(&map[abnormal.Abnormal]float64{_abnormal: 100})
			}
		}
	}
}

func (b *Buff) find(key string, val float64, percentage bool) (model.CharacterModifier, bool) {
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
