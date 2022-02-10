package romel

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Buff struct {
	buff      string
	modifiers []model.CharacterModifier
}

var (
	BuffTotal       = 0
	BuffUnknown     = 0
	BuffError       = 0
	BuffIgnore      = 0
	BuffDetected    = 0
	Buffs           = make(map[string]int)
	perQualityRates = map[string]func(character *model.Character, num int) int{
		"力量": func(character *model.Character, num int) int {
			return character.Quality.Str / num
		},
		"敏捷": func(character *model.Character, num int) int {
			return character.Quality.Agi / num
		},
		"智力": func(character *model.Character, num int) int {
			return character.Quality.Int / num
		},
		"智力可": func(character *model.Character, num int) int {
			return character.Quality.Int / num
		},
		"灵巧": func(character *model.Character, num int) int {
			return character.Quality.Dex / num
		},
		"体质": func(character *model.Character, num int) int {
			return character.Quality.Vit / num
		},
		"幸运": func(character *model.Character, num int) int {
			return character.Quality.Luk / num
		},
		"暴击": func(character *model.Character, num int) int {
			return character.Profits.General.Critical / num
		},
		"闪避": func(character *model.Character, num int) int {
			return character.Profits.General.Dodge / num
		},
		"生命上限": func(character *model.Character, num int) int {
			//TODO 计算生命上限
			return character.Profits.General.Hp / num
		},
		"百分之一魔伤减免": func(character *model.Character, num int) int {
			return int(character.Profits.Gains(true).Resist) / num
		},
		"1%忽视魔法防御": func(character *model.Character, num int) int {
			return int(character.Profits.Gains(true).Ignore) / num
		},
		"1%忽视物理防御": func(character *model.Character, num int) int {
			return int(character.Profits.Gains(false).Ignore) / num
		},
		"物理攻击": func(character *model.Character, num int) int {
			return int(character.Profits.Gains(false).Attack) / num
		},
		"魔法防御": func(character *model.Character, num int) int {
			return int(character.Profits.Gains(true).Defence) / num
		},
		"1%装备攻速将": func(character *model.Character, num int) int {
			return int(character.Profits.General.AttackSpeed) / num
		},
	}
)

func (b *Buff) UnmarshalJSON(data []byte) error {
	if buff, err := strconv.Unquote(string(data)); err != nil {
		return errors.WithStack(err)
	} else if len(buff) == 0 {
		return nil
	} else if modifiers, err := b.parseBuff(buff); err != nil {
		return err
	} else {
		b.buff = buff
		b.modifiers = modifiers
		return nil
	}
}

func (b *Buff) Empty() bool {
	return b == nil || len(b.buff) == 0
}

func (b *Buff) Contains(o *Buff) bool {
	if o.Empty() {
		return true
	} else if b.Empty() {
		return false
	} else {
		return strings.Index(b.buff, o.buff) >= 0
	}
}

func (b *Buff) Effect() []model.CharacterModifier {
	if b.Empty() {
		return nil
	} else {
		return b.modifiers
	}
}

func (b *Buff) Size() int {
	return len(b.Items())
}

func (b *Buff) Items() []string {
	return strings.Split(b.buff, "\n")
}

func (b *Buff) parseBuff(effect string) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	for _, line := range strings.Split(effect, "\n") {
		for _, item := range strings.Split(line, "；") {
			if ms, err := b.parseItem(item); err != nil {
				return nil, errors.Wrapf(err, "%s", item)
			} else if len(ms) > 0 {
				modifiers = append(modifiers, ms...)
			}
		}
	}
	return modifiers, nil
}

func (b *Buff) parseItem(item string) ([]model.CharacterModifier, error) {
	if b.parseIgnoreItem(item) {
		return nil, nil
	} else if match, modifiers, err := b.parsePerRefineItem(item); match || err != nil {
		return modifiers, err
	} else if match, modifiers, err := b.parsePerIntensifyItem(item); match || err != nil {
		return modifiers, err
	} else if match, modifiers, err := b.parseRefineItem(item); match || err != nil {
		return modifiers, err
	} else if match, modifiers, err := b.parsePetItem(item); match || err != nil {
		return modifiers, err
	} else if match, modifiers, err := b.parseContinuedItem(item); match || err != nil {
		return modifiers, err
	} else if b.parseSkillItem(item) {
		return nil, nil
	} else if match, modifiers, err := b.parseConditionItem(item); match || err != nil {
		return modifiers, err
	} else if match, modifiers, err := b.parsePerQualityItem(item); match || err != nil {
		return modifiers, err
	} else {
		return b.parseEffects(item, 1)
	}
}

func (b *Buff) parseIgnoreItem(item string) bool {
	if strings.Index(item, "竞技场模式") >= 0 {
		return true
	} else if strings.Index(item, "达纳托斯之塔") >= 0 {
		return true
	} else if strings.Index(item, "野外地图") >= 0 {
		return true
	} else if b.parseCombineItem(item) {
		return true
	} else {
		return false
	}
}

func (b *Buff) parseCombineItem(item string) bool {
	// 忽略装备组合增益
	return strings.Index(item, "】+【") > 0 || strings.Index(item, "）+【") > 0
}

func (b *Buff) parsePerRefineItem(item string) (bool, []model.CharacterModifier, error) {
	condition, effect, rate := "", "", 1
	if idx := strings.Index(item, "每精炼+1时"); idx >= 0 {
		condition, effect = strings.TrimSuffix(item[:idx], "，"), strings.TrimPrefix(item[idx+14:], "，")
	} else if idx := strings.Index(item, "再每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(item[:idx], "，"), strings.TrimPrefix(item[idx+14:], "，")
	} else if idx := strings.Index(item, "每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(item[:idx], "，"), strings.TrimPrefix(item[idx+11:], "，")
	} else if idx := strings.Index(item, "精炼每+1时"); idx >= 0 {
		condition, effect = strings.TrimSuffix(item[:idx], "，"), strings.TrimPrefix(item[idx+14:], "，")
	} else if idx := strings.Index(item, "精炼每+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(item[:idx], "，"), strings.TrimPrefix(item[idx+11:], "，")
	} else {
		return false, nil, nil
	}
	var modifiers []model.CharacterModifier
	switch condition {
	case "", "头饰", "副手", "盔甲", "装备这张卡片的武器", "使用演奏类技能时":
		rate = 15
	case "精炼+5开始", "精炼+5以后":
		rate = 10
	case "精炼+8开始":
		rate = 8
	case "精炼+1至+6", "精炼+10开始":
		rate = 6
	case "精炼+10以后", "精炼+10以后使用斧头类武器时":
		rate = 5
	default:
		if strings.HasSuffix(condition, "，精炼+1至+6") {
			rate = 6
			condition = strings.TrimSuffix(condition, "，精炼+1至+6")
		} else if strings.HasPrefix(condition, "精炼+10时，") {
			rate = 5
		} else {
			rate = 15
		}
		if match, ms, err := b.parseRefineItem(condition); err != nil {
			return false, nil, err
		} else if match {
			modifiers = append(modifiers, ms...)
		} else if ms, err := b.parseItem(condition); err != nil {
			return false, nil, err
		} else if len(ms) > 0 {
			modifiers = append(modifiers, ms...)
		}
	}
	if strings.Index(effect, "精炼+") > 0 {
		for _, splitEffect := range strings.Split(effect, "，") {
			if strings.HasPrefix(splitEffect, "精炼+") {
				if ms, err := b.parseItem(splitEffect); err != nil {
					return false, nil, err
				} else if len(ms) > 0 {
					modifiers = append(modifiers, ms...)
				}
			} else if ms, err := b.parseEffect(splitEffect, rate); err != nil {
				return false, nil, err
			} else if ms != nil {
				modifiers = append(modifiers, ms)
			}
		}
	} else if ms, err := b.parseEffects(effect, rate); err != nil {
		return false, nil, err
	} else if len(ms) > 0 {
		modifiers = append(modifiers, ms...)
	}
	return true, modifiers, nil
}

func (b *Buff) parsePerIntensifyItem(item string) (bool, []model.CharacterModifier, error) {
	effect, rate := "", 1
	if idx := strings.Index(item, "每强化+1"); idx >= 0 {
		rate, effect = 170, strings.TrimPrefix(item[idx+11:], "，")
	} else if idx := strings.Index(item, "每强化10级"); idx >= 0 {
		rate, effect = 17, strings.TrimPrefix(item[idx+14:], "，")
	} else {
		return false, nil, nil
	}
	if modifiers, err := b.parseEffects(effect, rate); err != nil {
		return false, nil, err
	} else {
		return true, modifiers, nil
	}
}

func (b *Buff) cap(item string, prefixes ...string) int {
	for _, prefix := range prefixes {
		if strings.HasPrefix(item, prefix) {
			return len(prefix)
		}
	}
	return -1
}

func (b *Buff) parseRefineWithCondition(refineStr string) (int, string, error) {
	runeArray := []rune(refineStr)
	for idx, char := range runeArray {
		if char < '0' || char > '9' {
			if refine, err := strconv.Atoi(string(runeArray[:idx])); err != nil {
				return -1, "", errors.WithStack(err)
			} else {
				return refine, string(runeArray[idx:]), nil
			}
		}
	}
	if refine, err := strconv.Atoi(refineStr); err != nil {
		return -1, "", errors.WithStack(err)
	} else {
		return refine, "", nil
	}
}

func (b *Buff) parseRefineItem(item string) (bool, []model.CharacterModifier, error) {
	refineStr, effectStr := "", ""
	if cap := b.cap(item, "精炼+", "精炼值+", "精炼等级+", "精炼等级达到+", "精炼至+", "当精炼+", "当精炼到+",
		"武器精炼+", "当盔甲精炼+", "当武器精炼+", "当副手精炼+"); cap <= 0 {
		return false, nil, nil
	} else if idx := strings.Index(item, "时"); idx > 0 {
		refineStr, effectStr = item[cap:idx], strings.TrimPrefix(item[idx+3:], "，")
	} else if idx := strings.Index(item, "，"); idx > 0 {
		refineStr, effectStr = item[cap:idx], item[idx+3:]
	} else {
		return false, nil, errors.Errorf("parseRefineItem: %s", item)
	}
	if strings.Index(refineStr, "、") >= 0 {
		modifiers, err := b.parseRefineItemSplit(refineStr, effectStr, "、")
		return err == nil, modifiers, err
	} else if strings.Index(refineStr, "，") >= 0 {
		modifiers, err := b.parseRefineItemSplit(refineStr, effectStr, "，")
		return err == nil, modifiers, err
	} else if strings.Index(refineStr, ",") >= 0 {
		modifiers, err := b.parseRefineItemSplit(refineStr, effectStr, ",")
		return err == nil, modifiers, err
	} else if refine, condition, err := b.parseRefineWithCondition(refineStr); err != nil {
		return false, nil, errors.Wrapf(err, "parseRefineItem: %s", item)
	} else if modifiers, err := b.parseRefineEffects(refine, effectStr); err != nil {
		return false, nil, errors.Wrapf(err, "parseRefineItem: %s", item)
	} else if len(modifiers) == 0 || len(condition) == 0 {
		return true, modifiers, nil
	} else {
		//TODO 解析条件并限制
		return true, modifiers, nil
	}
}

func (b *Buff) parseRefineItemSplit(refineStr, effectStr, split string) ([]model.CharacterModifier, error) {
	var refines []int
	for _, refineStrSplit := range strings.Split(refineStr, split) {
		if refine, err := strconv.Atoi(refineStrSplit); err != nil {
			return nil, errors.WithStack(err)
		} else {
			refines = append(refines, refine)
		}
	}
	if idx := strings.Index(effectStr, "分别+"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], effectStr[idx+6:], split, refines...)
	} else if idx := strings.Index(effectStr, "分别额外增加+"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], effectStr[idx+18:], split, refines...)
	} else if idx := strings.Index(effectStr, "+"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], effectStr[idx:], split, refines...)
	} else if idx := strings.Index(effectStr, "-"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], effectStr[idx:], split, refines...)
	} else if idx := strings.Index(effectStr, "上升"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], "+"+effectStr[idx+6:], split, refines...)
	} else if idx := strings.Index(effectStr, "减低"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], "-"+effectStr[idx+6:], split, refines...)
	} else if idx := strings.Index(effectStr, "分别提升到"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], "+"+effectStr[idx+15:], split, refines...)
	} else if idx := strings.Index(effectStr, "分别提升"); idx >= 0 {
		return b.parseRefineSplit(effectStr[:idx], "+"+effectStr[idx+12:], split, refines...)
	} else {
		return nil, errors.Errorf("parseRefineItemSplit[%s]: %d%s", split, refines, effectStr)
	}
}

func (b *Buff) parseRefineSplit(base, effectStr, split string, refines ...int) ([]model.CharacterModifier, error) {
	prefix := ""
	if strings.HasPrefix(effectStr, "+") {
		prefix = "+"
	} else if strings.HasPrefix(effectStr, "-") {
		prefix = "-"
	}
	if effects := strings.Split(effectStr, split); len(effects) != len(refines) {
		return nil, errors.Errorf("count mismatch: [%d]%d --> [%d]%s", len(refines), refines, len(effects), effects)
	} else {
		var modifiers []model.CharacterModifier
		for idx, refine := range refines {
			if sub, err := b.parseRefineEffects(refine, base+prefix+strings.TrimPrefix(effects[idx], prefix)); err != nil {
				return nil, err
			} else if len(sub) > 0 {
				modifiers = append(modifiers, sub...)
			}
		}
		return modifiers, nil
	}
}

func (b *Buff) parseRefineEffects(refine int, effectStr string) ([]model.CharacterModifier, error) {
	if modifiers, err := b.parseItem(effectStr); err != nil {
		return nil, err
	} else if len(modifiers) == 0 {
		return nil, nil
	} else {
		// TODO 限制精炼等级
		return modifiers, nil
	}
}

func (b *Buff) parsePetItem(item string) (bool, []model.CharacterModifier, error) {
	if strings.Index(item, "<em>") < 0 {
		return false, nil, nil
	} else if strings.HasPrefix(item, "冒险") {
		return true, nil, nil
	} else if strings.HasPrefix(item, "增加宠物在") && strings.HasSuffix(item, "的打工效率") {
		return true, nil, nil
	} else if strings.Index(item, "对敌方") >= 0 {
		return true, nil, nil
	} else if strings.HasPrefix(item, "增加宠物和主人") {
		modifiers, err := b.parsePetEffects(item[21:], true)
		return err == nil, modifiers, err
	} else if strings.HasPrefix(item, "减少宠物和主人") {
		modifiers, err := b.parsePetEffects(item[21:], false)
		return err == nil, modifiers, err
	} else if strings.HasPrefix(item, "宠物和主人") {
		modifiers, err := b.parsePetEffects(item[15:], true)
		return err == nil, modifiers, err
	} else if strings.HasPrefix(item, "主人和宠物") {
		modifiers, err := b.parsePetEffects(item[15:], true)
		return err == nil, modifiers, err
	} else if strings.HasPrefix(item, "主人") {
		modifiers, err := b.parsePetEffects(item[6:], true)
		return err == nil, modifiers, err
	} else if strings.HasPrefix(item, "增加主人") {
		modifiers, err := b.parsePetEffects(item[12:], true)
		return err == nil, modifiers, err
	} else {
		return true, nil, nil
	}
}

func (b *Buff) parseContinuedItem(item string) (bool, []model.CharacterModifier, error) {
	if strings.Index(item, "持续") < 0 {
		return false, nil, nil
	}

	var effects []string
	runeArray, pos := []rune(item), 0
	for idx, char := range runeArray {
		switch char {
		case '，', '。', '（', '）', ',':
			effects = append(effects, string(runeArray[pos:idx]))
			pos = idx + 1
		case '：':
			if runeArray[idx-1] != 'D' {
				effects = append(effects, string(runeArray[pos:idx]))
				pos = idx + 1
			}
		}
	}
	effects = append(effects, string(runeArray[pos:]))

	if effects, rate, _, _, err := b.parseContinuedTimes(effects); err != nil {
		return false, nil, err
	} else if effects, _, _, _, err := b.parseContinuedTimes(effects); err != nil {
		return false, nil, err
	} else {
		var modifiers []model.CharacterModifier
		for _, effect := range effects {
			if strings.IndexAny(effect, "0123456789") > 0 {
				if modifier, err := b.parseEffect(effect, rate); err != nil {
					return false, nil, err
				} else if modifier != nil {
					modifiers = append(modifiers, modifier)
				}
			}
		}
		return true, modifiers, nil
	}
}

func (b *Buff) parseContinuedTimes(effects []string) ([]string, int, int, int, error) {
	var after []string
	rate, cont, cd := 1, 0, 0
	for _, effect := range effects {
		if idxStart := strings.Index(effect, "持续"); idxStart >= 0 {
			condition, suffer := effect[:idxStart], strings.TrimPrefix(effect[idxStart+6:], "时间")
			add := false
			if strings.HasPrefix(suffer, "增加") {
				add = true
				suffer = strings.TrimPrefix(suffer, "增加")
			}
			if idxEnd := strings.Index(suffer, "秒"); idxEnd > 0 {
				if intVal, err := strconv.Atoi(strings.TrimSpace(suffer[:idxEnd])); err != nil {
					return nil, 0, 0, 0, errors.WithStack(err)
				} else if add {
					cont += intVal
				} else {
					cont = intVal
				}
			} else if idxEnd := strings.Index(suffer, "s"); idxEnd > 0 {
				if intVal, err := strconv.Atoi(suffer[:idxEnd]); err != nil {
					return nil, 0, 0, 0, errors.WithStack(err)
				} else {
					cont = intVal
				}
			} else if idxEnd := strings.Index(suffer, "分钟"); idxEnd > 0 {
				if intVal, err := strconv.Atoi(suffer[:idxEnd]); err != nil {
					return nil, 0, 0, 0, errors.WithStack(err)
				} else {
					cont = intVal * 60
				}
			} else if suffer == "一定时间" {
				cont = -1
			} else {
				after = append(after, effect)
				continue
			}
			if len(condition) > 0 {
				after = append(after, condition)
			}
		} else if idxStart := strings.Index(effect, "CD"); idxStart >= 0 {
			suffer := strings.TrimPrefix(effect[idxStart+2:], "：")
			suffer = strings.TrimSuffix(suffer, "秒")
			suffer = strings.TrimSuffix(suffer, "s")
			if intVal, err := strconv.Atoi(suffer); err != nil {
				return nil, 0, 0, 0, errors.WithStack(err)
			} else {
				cd = intVal
			}
		} else if idxStart := strings.Index(effect, "冷却"); idxStart >= 0 {
			suffer := strings.TrimPrefix(effect[idxStart+6:], "时间")
			suffer = strings.TrimSuffix(suffer, "秒")
			suffer = strings.TrimSuffix(suffer, "s")
			suffer = strings.TrimSpace(suffer)
			del := false
			if strings.HasPrefix(suffer, "减少") {
				del = true
				suffer = strings.TrimPrefix(suffer, "减少")
			}
			if intVal, err := strconv.Atoi(suffer); err != nil {
				return nil, 0, 0, 0, errors.WithStack(err)
			} else if del {
				cd -= intVal
			} else {
				cd = intVal
			}
		} else if idxStart := strings.Index(effect, "叠加"); idxStart >= 0 {
			suffer := effect[idxStart+6:]
			if idxEnd := strings.Index(suffer, "层"); idxEnd > 0 {
				if rateStr := suffer[:idxEnd]; rateStr == "三" {
					rate = 3
				} else if intVal, err := strconv.Atoi(rateStr); err != nil {
					return nil, 0, 0, 0, errors.WithStack(err)
				} else {
					rate = intVal
				}
				continue
			} else if len(suffer) > 0 {
				after = append(after, effect)
			}
		} else if idxStart := strings.Index(effect, "率"); idxStart >= 0 {
			if strings.HasPrefix(effect, "增加") {
				after = append(after, effect)
			} else if suffer := strings.TrimPrefix(effect[idxStart+3:], "使"); len(suffer) > 0 {
				after = append(after, suffer)
			}
		} else if idxStart := strings.Index(effect, "时"); idxStart >= 0 {
			if suffer := effect[idxStart+3:]; len(suffer) > 0 {
				switch []rune(suffer)[0] {
				case '刻', '间':
					after = append(after, effect)
				default:
					after = append(after, suffer)
				}
			}
		} else if len(effect) > 0 {
			after = append(after, effect)
		}
	}
	return after, rate, cont, cd, nil
}

func (b *Buff) parseSkillItem(item string) bool {
	if strings.Index(item, "【") < 0 {
		return false
	} else if strings.HasPrefix(item, "【") || strings.HasPrefix(item, "使【") {
		return true
	} else if strings.HasPrefix(item, "可使用【") || strings.HasPrefix(item, "可以使用") {
		return true
	} else if strings.HasPrefix(item, "可使用技能【") || strings.HasPrefix(item, "技能【") {
		return true
	} else if strings.HasPrefix(item, "获得技能【") || strings.HasPrefix(item, "获得被动技能【") {
		return true
	} else if strings.HasPrefix(item, "习得【") || strings.Index(item, "使用【") >= 0 {
		return true
	} else if strings.Index(item, "施放【") >= 0 || strings.Index(item, "触发【") >= 0 {
		return true
	} else if strings.Index(item, "发动【") >= 0 || strings.Index(item, "施展【") >= 0 {
		return true
	} else if strings.Index(item, "获得【") >= 0 || strings.Index(item, "释放【") >= 0 {
		return true
	} else {
		return true
	}
}

func (b *Buff) parseConditionItem(item string) (bool, []model.CharacterModifier, error) {
	itemNoBracket, _ := b.trimBracket(item)
	condition, effect := "", ""
	if idx := strings.Index(itemNoBracket, "时，"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+6:]
	} else if idx := strings.Index(itemNoBracket, "时,"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+4:]
	} else if idx := strings.Index(itemNoBracket, "时"); idx < 0 {
		return false, nil, nil
	} else if strings.Index(itemNoBracket, "时间") == idx {
		return false, nil, nil
	} else if strings.Index(itemNoBracket, "时刻") == idx {
		return false, nil, nil
	} else if strings.Index(itemNoBracket, "时长") == idx {
		return false, nil, nil
	} else if strings.Index(itemNoBracket, "同时") == idx-3 {
		return false, nil, nil
	} else {
		condition, effect = item[:idx+3], item[idx+3:]
	}
	var modifiers []model.CharacterModifier
	switch condition {
	case "物理攻击时", "普通攻击时", "技能攻击时", "使用技能时", "使用拳刃类武器时", "使用物理伤害技能攻击时", "使用短剑类武器时",
		"普攻时", "使用演奏类技能时", "魔法技能攻击时", "魔法攻击目标时", "远程普通攻击时", "装备来复枪类武器时", "装备弓类型武器时",
		"力量在75以上时", "当诗人或舞娘系职业演奏时", "技能攻击玩家时", "技能攻击目标时", "攻击时", "攻击目标时", "普通攻击暴击时",
		"生命值100%时", "SP值100%时", "自身生命值大于50%时", "击杀目标时", "当装备的职业为剑士系时", "普通攻击攻击目标时",
		"物理技能攻击攻击目标时", "普通攻击，暴击时", "近战攻击时", "任何攻击时", "近战普攻时", "使用物理技能攻击时", "物理技能攻击时",
		"使用锁定单体类技能攻击时", "使用长剑类武器时", "近战职业装备时", "每次魔法技能攻击时", "装备时", "但技能使用时", "技能使用时",
		"法，弓，服，多兰族，悟灵士装备时", "剑，商，盗，超初，忍者，枪手装备时", "主人使用技能时", "主人和宠物普通攻击时", "佩戴时",
		"佩戴者受到伤害时", "使用单体锁定类魔法技能时", "使用普攻或任何技能攻击时", "主人使用普通攻击、释放技能时", "主人击杀目标时",
		"主人攻击时", "主人释放技能时", "使主人在吟唱时", "使用单体锁定类魔法技能击杀魔物时", "使用技能（不含普通攻击）时", "暴击时",
		"使用法系技能时", "使用魔法技能攻击时", "宠物和主人击杀目标时", "宠物和主人攻击时", "宠物和主人普攻时", "当装备短剑时",
		"攻击血量低于自身的目标时", "攻击血量高于自身的目标时", "生命值低于70%时", "生命值高于50%时", "装备拳刃时", "近战职业击杀魔物时",
		"获得转运锦鲤的祝福：魔法技能攻击时", "阿特罗斯卡片触发急速效果时":
	default:
		if idx := strings.Index(condition, "，"); idx > 0 {
			if conditionEffect := condition[:idx]; strings.IndexAny(conditionEffect, "0123456789") > 0 {
				if modifier, err := b.parseEffect(conditionEffect, 1); err != nil {
					return false, nil, err
				} else if modifier != nil {
					modifiers = append(modifiers, modifier)
				}
			}
		}
		if strings.Index(condition, "装备") >= 0 && strings.HasSuffix(condition, "卡片时") {
		} else if strings.Index(condition, "达到") >= 0 || strings.Index(condition, "大于") >= 0 {
		} else if strings.Index(condition, "佩戴者是") >= 0 {
		} else if strings.HasSuffix(condition, "族装备时") || strings.HasSuffix(condition, "系装备时") {
		} else {
			//fmt.Printf("\tparseConditionItem: [%s]%s -- %s\n", condition, effect, item)
			//TODO 解析条件并忽略
			return true, modifiers, nil
		}
	}
	if sub, err := b.parseItem(effect); err != nil {
		return false, nil, err
	} else if len(sub) > 0 {
		modifiers = append(modifiers, sub...)
	}
	//TODO 解析条件并限制
	return true, modifiers, nil
}

func (b *Buff) perQuality(item, prefix string, splits ...string) (bool, string, string) {
	for _, split := range splits {
		if idxSplit := strings.Index(item, split); idxSplit > 0 {
			if idxStart := strings.Index(item, prefix); idxStart < 0 || idxStart >= idxSplit {
				continue
			} else {
				quality, effect := item[idxStart+len(prefix):idxSplit], item[idxSplit+len(split):]
				if r, _ := utf8.DecodeRuneInString(quality); (r >= '0' && r <= '9') || r == '百' {
					return true, quality, effect
				} else {
					continue
				}
			}
		}
	}
	return false, "", ""
}

func (b *Buff) parsePerQualityItem(item string) (bool, []model.CharacterModifier, error) {
	if strings.Index(item, "每") < 0 {
		return false, nil, nil
	} else if match, quality, effect := b.perQuality(item, "每", "会为装备者提供", "额外增加自身", "会额外增加",
		"会额外为其带来", "会为其带来", "额外增加", "额外带来", "额外提高", "可以转换", "增加自身", "增加", "加"); !match {
		//fmt.Printf("\tparsePerQualityItem: %s\n", item)
		return false, nil, nil
	} else {
		return b.parsePerQualityEffects(quality, effect)
	}
}

func (b *Buff) parsePerQualityEffects(quality, effectStr string) (bool, []model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	runeArray, pos := []rune(effectStr), 0
	for idx, char := range runeArray {
		if char == '，' || char == ',' || char == '、' {
			if modifier, err := b.parsePerQualityEffect(quality, string(runeArray[pos:idx])); err != nil {
				return false, nil, err
			} else if modifier != nil {
				modifiers = append(modifiers, modifier)
			}
			pos = idx + 1
		}
	}
	if modifier, err := b.parsePerQualityEffect(quality, string(runeArray[pos:])); err != nil {
		return false, nil, err
	} else if modifier != nil {
		modifiers = append(modifiers, modifier)
	}
	return true, modifiers, nil
}

func (b *Buff) parsePerQualityEffect(qualityStr, effectStr string) (model.CharacterModifier, error) {
	max, effectNoBracket := 0, effectStr
	if after, bracket := b.trimBracket(effectStr); len(bracket) > 0 {
		effectNoBracket = after
		if strings.Index(bracket, "额外") >= 0 {
			switch bracket {
			case "最多额外增加1500点物理攻击":
				max = 1500
			case "额外最多15%":
				max = 15
			default:
				return nil, errors.Errorf("parsePerQualityEffect: [%s]%s -- %s", qualityStr, effectNoBracket, bracket)
			}
		}
	}

	effect := ""
	if idx := strings.Index(effectNoBracket, "点"); idx > 0 {
		effect = fmt.Sprintf("%s+%s", effectNoBracket[idx+3:], effectNoBracket[:idx])
	} else if idx := strings.Index(effectNoBracket, "%"); idx > 0 {
		effect = fmt.Sprintf("%s+%s", effectNoBracket[idx+1:], effectNoBracket[:idx+1])
	} else {
		runeArray := []rune(effectNoBracket)
		for idx, char := range runeArray {
			if char >= '0' && char <= '9' {
				continue
			} else if char == '.' {
				continue
			} else {
				effect = fmt.Sprintf("%s+%s", string(runeArray[idx:]), string(runeArray[:idx]))
				break
			}
		}
		if len(effect) == 0 {
			return nil, errors.Errorf("parsePerQualityEffect: [%s]%s", qualityStr, effectNoBracket)
		}
	}

	if max > 0 {
		return b.parseEffect(effect, max)
	} else {
		quality, num := qualityStr, 1
		if idx := strings.Index(qualityStr, "点"); idx > 0 {
			if qualityVal, err := strconv.Atoi(qualityStr[:idx]); err != nil {
				return nil, errors.WithStack(err)
			} else {
				quality, num = qualityStr[idx+3:], qualityVal
			}
		}
		if rateFn, exist := perQualityRates[quality]; !exist {
			return nil, errors.Errorf("parsePerQualityEffect: [%d][%s] -- %s", num, quality, effect)
		} else {
			return func(character *model.Character) func() {
				if modifier, err := b.parseEffect(effect, rateFn(character, num)); err != nil {
					zap.S().Warnf("%+v", err)
				} else if modifier != nil {
					return modifier(character)
				}
				return func() {}
			}, nil
		}
	}
}

func (b *Buff) parsePetEffects(effectStr string, plus bool) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	runeArray, pos := []rune(effectStr), 0
	for idx, char := range runeArray {
		if char == '，' || char == ',' || char == ';' {
			if modifier, err := b.parsePetEffect(string(runeArray[pos:idx]), plus); err != nil {
				return nil, err
			} else if modifier != nil {
				modifiers = append(modifiers, modifier)
			}
			pos = idx + 1
		}
	}
	if modifier, err := b.parsePetEffect(string(runeArray[pos:]), plus); err != nil {
		return nil, err
	} else if modifier != nil {
		modifiers = append(modifiers, modifier)
	}
	return modifiers, nil
}

func (b *Buff) parsePetEffect(effectStr string, plus bool) (model.CharacterModifier, error) {
	if idxStart := strings.Index(effectStr, "<em>"); idxStart < 0 {
		if effectStr == "并增加5%物理攻击力和5%攻速" {
			if modifiers, err := b.parseEffects("物理攻击力+5%、攻速+5%", 1); err != nil {
				return nil, err
			} else if len(modifiers) > 0 {
				return model.Merge(modifiers...), nil
			}
		}
		return nil, nil
	} else if idxEnd := strings.Index(effectStr, "</em>"); idxEnd < 0 {
		return nil, errors.Errorf("parsePetEffect: %s", effectStr)
	} else {
		prefix, suffix, val, effect := effectStr[:idxStart], effectStr[idxEnd+5:], effectStr[idxStart+4:idxEnd], ""
		suffix = strings.TrimPrefix(suffix, "的")
		if len(prefix) == 0 && len(suffix) > 0 {
			if plus {
				effect = fmt.Sprintf("%s+%s", strings.TrimPrefix(suffix, "点"), val)
			} else {
				effect = fmt.Sprintf("%s-%s", strings.TrimPrefix(suffix, "点"), val)
			}
		} else if len(suffix) == 0 && strings.HasSuffix(prefix, "+") {
			effect = fmt.Sprintf("%s%s", prefix, val)
		} else if len(suffix) == 0 && strings.HasSuffix(prefix, "减少") {
			effect = fmt.Sprintf("%s-%s", strings.TrimSuffix(prefix, "减少"), val)
		} else if len(suffix) == 0 && strings.HasSuffix(prefix, "增加") {
			effect = fmt.Sprintf("%s+%s", strings.TrimSuffix(prefix, "增加"), val)
		} else if len(suffix) > 0 && strings.HasSuffix(prefix, "降低") {
			effect = fmt.Sprintf("%s-%s", suffix, val)
		} else if len(suffix) > 0 && strings.HasSuffix(prefix, "增加") {
			effect = fmt.Sprintf("%s+%s", suffix, val)
		} else if len(suffix) > 0 && strings.HasSuffix(prefix, "提升") {
			effect = fmt.Sprintf("%s+%s", suffix, val)
		} else if prefix == "但移动速度会降低" {
			effect = fmt.Sprintf("移动速度-%s", val)
		} else if prefix == "击杀目标时有" && suffix == "概率使自身物理攻击增加60点" {
			effect = "物理攻击+60"
		} else if prefix == "击杀目标时有" && suffix == "概率额外获得6Zeny" {
			effect = "Zeny+6"
		} else if prefix == "每" {
			if _, modifiers, err := b.parsePerQualityItem(fmt.Sprintf("%s%s%s", prefix, val, suffix)); err != nil {
				return nil, err
			} else if len(modifiers) > 0 {
				return model.Merge(modifiers...), nil
			}
		}
		if len(effect) == 0 {
			return nil, nil
		} else {
			return b.parseEffect(effect, 1)
		}
	}
}

func (b *Buff) parseEffects(effectStr string, rate int) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	runeArray, pos, inBracket := []rune(effectStr), 0, false
	for idx, char := range runeArray {
		switch char {
		case '（', '(':
			inBracket = true
		case '）', ')':
			inBracket = false
		case '、', '，', ',', '和':
			if !inBracket && b.isEndOfDigit(runeArray[idx-1]) {
				if modifier, err := b.parseEffect(string(runeArray[pos:idx]), rate); err != nil {
					return nil, err
				} else if modifier != nil {
					modifiers = append(modifiers, modifier)
				}
				pos = idx + 1
			}
		}
	}
	if modifier, err := b.parseEffect(string(runeArray[pos:]), rate); err != nil {
		return nil, err
	} else if modifier != nil {
		modifiers = append(modifiers, modifier)
	}
	return modifiers, nil
}

func (b *Buff) parseEffect(effectStr string, rate int) (model.CharacterModifier, error) {
	effectNoBracket, _ := b.trimBracket(effectStr)
	runeArray, pos, inBracket, percentage := []rune(effectNoBracket), 0, false, strings.HasSuffix(effectNoBracket, "%")
	for idx, char := range runeArray {
		switch char {
		case '（', '(':
			inBracket = true
		case '）', ')':
			inBracket = false
		case '+', '-':
			if !inBracket && runeArray[idx+1] >= '0' && runeArray[idx+1] <= '9' {
				key, val := string(runeArray[pos:idx]), string(runeArray[idx+1:])
				val = strings.TrimSuffix(val, "%")
				val = strings.TrimSuffix(val, "秒")
				val = strings.TrimSuffix(val, "点")
				if floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err != nil {
					zap.S().Warnf("parseEffect.ParseFloat: %s --- %s", val, effectStr)
				} else {
					if rate > 1 {
						floatVal *= float64(rate)
					}
					if char == '-' {
						floatVal = -floatVal
					}
					return b.findEffect(key, floatVal, percentage)
				}
			}
		}
	}
	return b.findEffect(effectNoBracket, 0, percentage)
}

func (b *Buff) findEffect(key string, val float64, percentage bool) (model.CharacterModifier, error) {
	BuffTotal++
	if strings.Index(key, "恢复") >= 0 {
		BuffIgnore++
		return nil, nil
	} else if strings.Index(key, "消耗") >= 0 {
		BuffIgnore++
		return nil, nil
	} else if strings.Index(key, "目标") >= 0 {
		BuffIgnore++
		return nil, nil
	} else if strings.Index(key, "不包括自身") >= 0 {
		BuffIgnore++
		return nil, nil
	} else if !percentage && val == 0 && strings.HasSuffix(key, "时") {
		BuffIgnore++
		return nil, nil
	} else if strings.HasSuffix(key, "技能倍率") {
		BuffIgnore++
		return nil, nil
	} else if strings.HasPrefix(key, "【") {
		BuffIgnore++
		return nil, nil
	} else if strings.HasPrefix(key, "使【") {
		BuffIgnore++
		return nil, nil
	} else if strings.HasPrefix(key, "使用【") {
		BuffIgnore++
		return nil, nil
	} else if strings.Index(key, "持续") >= 0 {
		BuffIgnore++
		return nil, nil
		//} else if key == "获得基于移动速度的额外物理攻击加成，移动速度每提升1%" {
		//	BuffIgnore++
		//	return nil, errors.Errorf("[%t]%s -- %f", percentage, key, val)
	} else if modifier, exist := b.find(key, val, percentage); !exist {
		BuffUnknown += 1
		item := key
		if percentage {
			item = "[%]" + item
		}
		if oc, ok := Buffs[item]; ok {
			Buffs[item] = oc + 1
		} else {
			Buffs[item] = 1
		}
		return nil, nil
	} else if modifier == nil {
		BuffIgnore++
		return nil, nil
	} else {
		BuffDetected++
		return modifier, nil
	}
}

func (b *Buff) isEndOfDigit(s rune) bool {
	switch s {
	case '%', ' ', '点', '）', ')':
		return true
	default:
		return s >= '0' && s <= '9'
	}
}

func (b *Buff) trimBracket(str string) (string, string) {
	if r, _ := utf8.DecodeLastRuneInString(str); r == '）' {
		if idx := strings.LastIndex(str, "（"); idx >= 0 {
			return str[:idx], str[idx+3 : len(str)-3]
		}
	} else if r == ')' {
		if idx := strings.LastIndex(str, "("); idx >= 0 {
			return str[:idx], str[idx+1 : len(str)-1]
		}
	}
	return str, ""
}
