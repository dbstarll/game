package romel

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"github.com/dbstarll/game/internal/ro/model/general"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
)

type Buff string

func (b Buff) Contains(o Buff) bool {
	return strings.Index(string(b), string(o)) >= 0
}

func (b Buff) Effect() ([]model.CharacterModifier, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var modifiers []model.CharacterModifier
	for _, line := range strings.Split(string(b), "\n") {
		for _, item := range strings.Split(line, "；") {
			if ms, err := b.resolveItem(item); err != nil {
				return nil, err
			} else if len(ms) > 0 {
				modifiers = append(modifiers, ms...)
			}
		}
	}
	return modifiers, nil
}

func (b Buff) resolveItem(item string) ([]model.CharacterModifier, error) {
	if match, modifiers, err := b.resolveCombineItem(item); err != nil {
		return nil, err
	} else if match {
		return modifiers, nil
	} else if match, modifiers, err := b.resolveRefineItem(item); err != nil {
		return nil, err
	} else if match {
		return modifiers, nil
	} else {
		return b.resolveEffects(item)
	}
}

func (b Buff) resolveCombineItem(item string) (bool, []model.CharacterModifier, error) {
	//TODO 忽略装备组合增益
	return strings.Index(item, "】+【") > 0 || strings.Index(item, "）+【") > 0, nil, nil
}

func (b Buff) resolveRefineItem(item string) (bool, []model.CharacterModifier, error) {
	cap := 0
	if strings.HasPrefix(item, "精炼+") {
		cap = 6
	} else if strings.HasPrefix(item, "当精炼+") {
		cap = 9
	} else if strings.HasPrefix(item, "武器精炼+") {
		cap = 12
	} else if strings.HasPrefix(item, "当盔甲精炼+") {
		cap = 15
	} else if strings.HasPrefix(item, "当武器精炼+") {
		cap = 15
	} else {
		return false, nil, nil
	}
	if idx := strings.Index(item, "时"); idx < 0 {
		return false, nil, nil
	} else {
		refineStr, effectStr := item[cap:idx], strings.TrimPrefix(item[idx+3:], "，")
		if strings.Index(refineStr, "、") >= 0 {
			modifiers, err := b.resolveRefineItemSplit(refineStr, effectStr, "、")
			return err == nil, modifiers, err
		} else if strings.Index(refineStr, "，") >= 0 {
			modifiers, err := b.resolveRefineItemSplit(refineStr, effectStr, "，")
			return err == nil, modifiers, err
		} else if strings.Index(refineStr, ",") >= 0 {
			modifiers, err := b.resolveRefineItemSplit(refineStr, effectStr, ",")
			return err == nil, modifiers, err
		} else if idx := strings.Index(refineStr, "时且"); idx >= 0 {
			modifiers, err := b.resolveRefineItemCond(refineStr[:idx], effectStr, refineStr[idx+6:])
			return err == nil, modifiers, err
		} else if idx := strings.Index(refineStr, "且"); idx >= 0 {
			modifiers, err := b.resolveRefineItemCond(refineStr[:idx], effectStr, refineStr[idx+3:])
			return err == nil, modifiers, err
		} else if idx := strings.Index(refineStr, "以后"); idx >= 0 {
			modifiers, err := b.resolveRefineItemCond(refineStr[:idx], effectStr, refineStr[idx+6:])
			return err == nil, modifiers, err
		} else if strings.HasSuffix(refineStr, "及以上") {
			refineStr = strings.TrimSuffix(refineStr, "及以上")
		} else if strings.HasSuffix(refineStr, "以上") {
			refineStr = strings.TrimSuffix(refineStr, "以上")
		}
		if refine, err := strconv.Atoi(refineStr); err != nil {
			return false, nil, errors.Wrapf(err, "item: %s, refineStr: %s", item, refineStr)
		} else {
			modifiers, err := b.resolveRefine(effectStr, refine)
			return err == nil, modifiers, err
		}
	}
}

func (b Buff) resolveRefineItemSplit(refineStr, effectStr, split string) ([]model.CharacterModifier, error) {
	var refines []int
	for _, refineStrSplit := range strings.Split(refineStr, split) {
		if refine, err := strconv.Atoi(refineStrSplit); err != nil {
			return nil, errors.WithStack(err)
		} else {
			refines = append(refines, refine)
		}
	}
	if idx := strings.Index(effectStr, "+"); idx >= 0 {
		return b.resolveRefinesSplit(effectStr[:idx], effectStr[idx:], split, refines...)
	} else if idx := strings.Index(effectStr, "-"); idx >= 0 {
		return b.resolveRefinesSplit(effectStr[:idx], effectStr[idx:], split, refines...)
	} else if idx := strings.Index(effectStr, "上升"); idx >= 0 {
		return b.resolveRefinesSplit(effectStr[:idx+6], effectStr[idx+6:], split, refines...)
	} else if idx := strings.Index(effectStr, "减低"); idx >= 0 {
		return b.resolveRefinesSplit(effectStr[:idx+6], effectStr[idx+6:], split, refines...)
	} else if idx := strings.Index(effectStr, "提升"); idx >= 0 {
		return b.resolveRefinesSplit(effectStr[:idx+6], effectStr[idx+6:], split, refines...)
	} else {
		return nil, errors.Errorf("resolveRefines: %d%s --> %s\n", refines, split, effectStr)
	}
}

func (b Buff) resolveRefinesSplit(base, effectStr, split string, refines ...int) ([]model.CharacterModifier, error) {
	if effects := strings.Split(effectStr, split); len(effects) != len(refines) {
		return nil, errors.Errorf("count mismatch: [%d]%d --> [%d]%s", len(refines), refines, len(effects), effects)
	} else {
		var modifiers []model.CharacterModifier
		for idx, refine := range refines {
			if sub, err := b.resolveRefine(base+effects[idx], refine); err != nil {
				return nil, err
			} else if len(sub) > 0 {
				modifiers = append(modifiers, sub...)
			}
		}
		return modifiers, nil
	}
}

func (b Buff) resolveRefineItemCond(refineStr, effectStr, cond string) ([]model.CharacterModifier, error) {
	if refine, err := strconv.Atoi(refineStr); err != nil {
		return nil, errors.WithStack(err)
	} else if modifiers, err := b.resolveRefine(effectStr, refine); err != nil {
		return nil, err
	} else if len(modifiers) == 0 {
		return nil, nil
	} else {
		//TODO 解析条件并限制
		log.Printf("resolveRefineItemCond: %s\n", cond)
		return modifiers, nil
	}
}

func (b Buff) resolveRefine(effectStr string, refine int) ([]model.CharacterModifier, error) {
	if modifiers, err := b.resolveEffects(effectStr); err != nil {
		return nil, err
	} else if len(modifiers) == 0 {
		return nil, nil
	} else {
		// TODO 限制精炼等级
		//fmt.Printf("resolveRefine: %d --> %s\n", refine, effectStr)
		return modifiers, nil
	}
}

func (b Buff) resolveEffects(effectStr string) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	pos, runeArray, perRefine := 0, []rune(effectStr), false
	for idx, char := range runeArray {
		if (char == '、' || char == '，') && b.isEndOfDigit(runeArray[idx-1]) {
			if effect := string(runeArray[pos:idx]); strings.HasSuffix(effect, "每精炼+1") {
				perRefine = true
				break
			} else if sub, err := b.resolveEffect(effect); err != nil {
				return nil, err
			} else if sub != nil {
				modifiers = append(modifiers, sub)
			}
			pos = idx + 1
		}
	}
	if effect := string(runeArray[pos:]); perRefine {
		//TODO 每精炼+1
		log.Printf("resolveEffects: %s", effect)
	} else if sub, err := b.resolveEffect(effect); err != nil {
		return nil, err
	} else if sub != nil {
		modifiers = append(modifiers, sub)
	}
	return modifiers, nil
}

func (b Buff) isEndOfDigit(s rune) bool {
	return s == '%' || s == '）' || (s >= '0' && s <= '9')
}

func (b Buff) resolveEffect(effectStr string) (model.CharacterModifier, error) {
	runeArray := []rune(effectStr)
	for idx, char := range runeArray {
		if char == '+' || char == '-' {
			if key, val := string(runeArray[:idx]), string(runeArray[idx+1:]); strings.HasSuffix(val, "%") {
				val = strings.TrimSuffix(val, "%")
				if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
					if char == '-' {
						floatVal = -floatVal
					}
					switch key {
					case "暴击伤害":
						return model.AddGeneral(&general.General{CriticalDamage: floatVal}), nil
					case "受到暴击伤害":
						return model.AddGeneral(&general.General{CriticalDamageResist: -floatVal}), nil
					case "普攻伤害":
						return model.AddGeneral(&general.General{OrdinaryDamage: floatVal}), nil
					case "对MVP、Mini魔物伤害", "对MVP、Mini魔物的伤害", "对MVP、Mini魔物增伤", "对MVP、Mini魔物魔物增伤", "对MVP、Mini魔物造成的伤害":
						return model.AddGeneral(&general.General{MVP: floatVal}), nil
					case "受到MVP、Mini魔物造成的伤害":
						return model.AddGeneral(&general.General{MVPResist: -floatVal}), nil
					case "技能伤害":
						return model.AddGeneral(&general.General{Skill: floatVal}), nil
					case "技能伤害减免":
						return model.AddGeneral(&general.General{SkillResist: floatVal}), nil
					case "装备攻速":
						return model.AddGeneral(&general.General{AttackSpeed: floatVal}), nil
					case "移动速度":
						return model.AddGeneral(&general.General{MoveSpeed: floatVal}), nil
					case "闪避":
						return model.AddGeneral(&general.General{DodgePer: floatVal}), nil
					case "生命上限":
						return model.AddGeneral(&general.General{HpPer: floatVal}), nil
					case "魔法上限":
						return model.AddGeneral(&general.General{SpPer: floatVal}), nil
					case "治疗加成":
						return model.AddGeneral(&general.General{Cure: floatVal}), nil
					case "受治疗加成":
						return model.AddGeneral(&general.General{Cured: floatVal}), nil
					case "可变吟唱时间":
						return model.AddGeneral(&general.General{SingElasticity: floatVal}), nil
					case "固定吟唱时间":
						return model.AddGeneral(&general.General{SingFixed: floatVal}), nil
					case "技能冷却":
						return model.AddGeneral(&general.General{SkillCooling: floatVal}), nil
					case "所有技能SP消耗", "使用技能Sp消耗量":
						return model.AddGeneral(&general.General{SpCost: floatVal}), nil
					case "法术普攻暴击概率":
						return model.AddGeneral(&general.General{MagicOrdinaryCriticalRate: floatVal}), nil
					case "法术普攻暴击伤害":
						return model.AddGeneral(&general.General{MagicOrdinaryCriticalDamage: floatVal}), nil

					// 物理、魔法兼得的增益
					case "物理、魔法攻击":
						return model.Merge(model.AddGains(false, &model.Gains{AttackPer: floatVal}), model.AddGains(true, &model.Gains{AttackPer: floatVal})), nil
					case "最终伤害":
						return model.Merge(model.AddGains(false, &model.Gains{Damage: floatVal}), model.AddGains(true, &model.Gains{Damage: floatVal})), nil
					case "物伤、魔伤减免":
						return model.Merge(model.AddGains(false, &model.Gains{Resist: floatVal}), model.AddGains(true, &model.Gains{Resist: floatVal})), nil

					// 物理增益
					case "物理穿刺":
						return model.AddGains(false, &model.Gains{Spike: floatVal}), nil
					case "物理攻击", "装备者的物理防御和魔法防御不再降低，物理攻击":
						return model.AddGains(false, &model.Gains{AttackPer: floatVal}), nil
					case "远程物理攻击":
						//TODO 限制远程武器
						return model.AddGains(false, &model.Gains{AttackPer: floatVal}), nil
					case "近战物理攻击":
						//TODO 限制近战物理
						return model.AddGains(false, &model.Gains{AttackPer: floatVal}), nil
					case "物理防御":
						return model.AddGains(false, &model.Gains{DefencePer: floatVal}), nil
					case "物理伤害":
						return model.AddGains(false, &model.Gains{Damage: floatVal}), nil
					case "远程物理伤害":
						return model.AddGains(false, &model.Gains{RemoteDamage: floatVal}), nil
					case "近战物理伤害":
						return model.AddGains(false, &model.Gains{NearDamage: floatVal}), nil
					case "忽视物理防御":
						return model.AddGains(false, &model.Gains{Ignore: floatVal}), nil
					case "物伤减免":
						return model.AddGains(false, &model.Gains{Resist: floatVal}), nil
					case "受到远距离物理伤害减免":
						return model.AddGains(false, &model.Gains{RemoteResist: floatVal}), nil

					// 魔法增益
					case "魔法穿刺":
						return model.AddGains(true, &model.Gains{Spike: floatVal}), nil
					case "魔法攻击":
						return model.AddGains(true, &model.Gains{AttackPer: floatVal}), nil
					case "魔法防御", "且额外获得魔法防御":
						return model.AddGains(true, &model.Gains{DefencePer: floatVal}), nil
					case "魔法伤害":
						return model.AddGains(true, &model.Gains{Damage: floatVal}), nil
					case "忽视魔法防御":
						return model.AddGains(true, &model.Gains{Ignore: floatVal}), nil
					case "魔伤减免":
						return model.AddGains(true, &model.Gains{Resist: floatVal}), nil

					// 属性攻击%
					case "风、地、水、火属性攻击":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Wind: floatVal, nature.Earth: floatVal, nature.Water: floatVal, nature.Fire: floatVal}), nil
					case "火属性攻击":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Fire: floatVal}), nil
					case "水属性攻击":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Water: floatVal}), nil
					case "圣属性攻击":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Holy: floatVal}), nil
					case "无属性攻击":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Neutral: floatVal}), nil

					// 属性增伤%
					case "对水属性魔物伤害":
						return model.AddNatureDamage(&map[nature.Nature]float64{nature.Water: floatVal}), nil
					case "对火属性魔物伤害":
						return model.AddNatureDamage(&map[nature.Nature]float64{nature.Fire: floatVal}), nil

					// 种族增伤%
					case "全种族伤害":
						return buff.RaceDamage(floatVal), nil
					case "人形种族伤害", "人形种族加伤", "人形种族加伤额外":
						return model.AddRaceDamage(&map[race.Race]float64{race.Human: floatVal}), nil

					// 种族减伤%
					case "全种族减伤":
						return buff.AddRaceResist(floatVal), nil
					case "人形种族减伤":
						return model.AddRaceResist(&map[race.Race]float64{race.Human: floatVal}), nil

					// 属性减伤%
					case "受到风、地、水、火、无属性伤害":
						return model.AddNatureResist(&map[nature.Nature]float64{nature.Wind: -floatVal, nature.Earth: -floatVal, nature.Water: -floatVal, nature.Fire: -floatVal, nature.Neutral: -floatVal}), nil
					case "受到无、圣、暗、念、毒属性伤害":
						return model.AddNatureResist(&map[nature.Nature]float64{nature.Neutral: -floatVal, nature.Holy: -floatVal, nature.Dark: -floatVal, nature.Ghost: -floatVal, nature.Poison: -floatVal}), nil
					case "受到无属性伤害":
						return model.AddNatureAttack(&map[nature.Nature]float64{nature.Neutral: -floatVal}), nil

					// 体型增伤%
					case "对小、中、大型魔物伤害", "对小、中、大型魔物的伤害":
						return model.AddShapeDamage(&map[shape.Shape]float64{shape.Large: floatVal, shape.Medium: floatVal, shape.Small: floatVal}), nil
					case "对大型魔物伤害", "对大型魔物的伤害":
						return model.AddShapeDamage(&map[shape.Shape]float64{shape.Large: floatVal}), nil
					case "对中型魔物伤害", "对中型魔物的伤害":
						return model.AddShapeDamage(&map[shape.Shape]float64{shape.Medium: floatVal}), nil
					case "对小型魔物伤害", "对小型魔物的伤害":
						return model.AddShapeDamage(&map[shape.Shape]float64{shape.Small: floatVal}), nil

					// 体型减伤%
					case "受到大、中、小型魔物伤害":
						return model.AddShapeResist(&map[shape.Shape]float64{shape.Large: -floatVal, shape.Medium: -floatVal, shape.Small: -floatVal}), nil
					case "受到小型魔物伤害":
						return model.AddShapeResist(&map[shape.Shape]float64{shape.Small: -floatVal}), nil
					case "受到中型魔物伤害":
						return model.AddShapeResist(&map[shape.Shape]float64{shape.Medium: -floatVal}), nil
					case "SP恢复", "“艾米斯可鲁”物理攻击", "对中毒的目标造成伤害额外", "反伤率", "成功概率", "陷阱类技能的伤害":
						//忽略以上
						return nil, nil
					default:
						//过滤掉技能
						if strings.Index(key, "【") < 0 {
							log.Printf("resolveEffect: %s %s %f || %s", key, string(char), floatVal, effectStr)
						}
					}
				} else {
					zap.S().Warnf("resolveEffect: %s %s %s || %s", key, string(char), val, effectStr)
				}
			} else {
				if intVal, err := strconv.Atoi(val); err == nil {
					if char == '-' {
						intVal = -intVal
					}
					switch key {
					case "全能力":
						return buff.Quality(intVal), nil
					case "力量":
						return model.AddQuality(&model.Quality{Str: intVal}), nil
					case "敏捷":
						return model.AddQuality(&model.Quality{Agi: intVal}), nil
					case "体质":
						return model.AddQuality(&model.Quality{Vit: intVal}), nil
					case "智力":
						return model.AddQuality(&model.Quality{Int: intVal}), nil
					case "灵巧":
						return model.AddQuality(&model.Quality{Dex: intVal}), nil
					case "幸运":
						return model.AddQuality(&model.Quality{Luk: intVal}), nil
					case "暴击":
						return model.AddGeneral(&general.General{Critical: intVal}), nil
					case "普攻攻击力", "普攻攻击":
						return model.AddGeneral(&general.General{Ordinary: intVal}), nil
					case "暴击防护":
						return model.AddGeneral(&general.General{CriticalResist: intVal}), nil
					case "命中":
						return model.AddGeneral(&general.General{Hit: intVal}), nil
					case "闪避":
						return model.AddGeneral(&general.General{Dodge: intVal}), nil
					case "生命上限":
						return model.AddGeneral(&general.General{Hp: intVal}), nil
					case "魔法上限":
						return model.AddGeneral(&general.General{Sp: intVal}), nil
					case "物理攻击", "物理攻击分别":
						return model.AddGains(false, &model.Gains{Attack: float64(intVal)}), nil
					case "物理防御":
						return model.AddGains(false, &model.Gains{Defence: float64(intVal)}), nil
					case "精炼物攻":
						return model.AddGains(false, &model.Gains{Refine: float64(intVal)}), nil
					case "魔法攻击":
						return model.AddGains(true, &model.Gains{Attack: float64(intVal)}), nil
					case "魔法防御":
						return model.AddGains(true, &model.Gains{Defence: float64(intVal)}), nil
					case "精炼魔攻":
						return model.AddGains(true, &model.Gains{Refine: float64(intVal)}), nil
					case "物理、魔法防御":
						return model.Merge(model.AddGains(false, &model.Gains{Defence: float64(intVal)}), model.AddGains(true, &model.Gains{Defence: float64(intVal)})), nil
					case "生命自然恢复", "SP恢复", "Sp恢复", "魔法恢复", "生命恢复", "Hp恢复":
						//忽略
						return nil, nil
					default:
						//过滤掉技能
						if strings.Index(key, "【") < 0 {
							//fmt.Printf("\tresolveEffect: %s %s %d || %s\n", key, string(char), intVal, effectStr)
							log.Printf("resolveEffect: %s %s %d || %s", key, string(char), intVal, effectStr)
						}
					}
				} else if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
					if char == '-' {
						floatVal = -floatVal
					}
					switch key {
					case "物理攻击":
						return model.AddGains(false, &model.Gains{Attack: floatVal}), nil
					case "物理防御":
						return model.AddGains(false, &model.Gains{Defence: floatVal}), nil
					case "魔法攻击":
						return model.AddGains(true, &model.Gains{Attack: floatVal}), nil
					case "魔法防御":
						return model.AddGains(true, &model.Gains{Defence: floatVal}), nil
					default:
						//过滤掉技能
						if strings.Index(key, "【") < 0 {
							fmt.Printf("\tresolveEffect: %s %s %f || %s\n", key, string(char), floatVal, effectStr)
							log.Printf("resolveEffect: %s %s %f || %s", key, string(char), floatVal, effectStr)
						}
					}
				} else {
					zap.S().Warnf("resolveEffect: %s %s %s || %s", key, string(char), val, effectStr)
				}
			}
			break
		}
	}
	return nil, nil
}
