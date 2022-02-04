package romel

import (
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Buff struct {
	buff      string
	modifiers []model.CharacterModifier
}

var (
	BuffTotal    = 0
	BuffUnknown  = 0
	BuffError    = 0
	BuffIgnore   = 0
	BuffDetected = 0
	Buffs        = make(map[string]int)
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
				return nil, err
			} else if len(ms) > 0 {
				modifiers = append(modifiers, ms...)
			}
		}
	}
	return modifiers, nil
}

func (b *Buff) parseItem(item string) ([]model.CharacterModifier, error) {
	BuffTotal += 1
	if match, modifiers := b.parseCombineItem(item); match {
		return modifiers, nil
	} else if match, modifiers := b.parsePvpItem(item); match {
		return modifiers, nil
	} else if match, modifiers := b.parseTowerItem(item); match {
		return modifiers, nil
	} else if match, modifiers, err := b.parsePerRefineItem(item); err != nil {
		return nil, err
	} else if match {
		return modifiers, nil
	} else if match, modifiers, err := b.parseRefineItem(item); err != nil {
		return nil, err
	} else if match {
		return modifiers, nil
	} else if match, modifiers, err := b.parseConditionItem(item); err != nil {
		return nil, err
	} else if match {
		return modifiers, nil
	} else if match, modifiers, err := b.parseEffects(item, 1); err != nil {
		BuffError++
		//TODO 处理异常
		return nil, nil
	} else if match {
		BuffDetected++
		return modifiers, nil
	} else {
		BuffUnknown += 1
		if oc, ok := Buffs[item]; ok {
			Buffs[item] = oc + 1
		} else {
			Buffs[item] = 1
		}
		return modifiers, nil
	}
}

func (b *Buff) parseCombineItem(item string) (bool, []model.CharacterModifier) {
	// 忽略装备组合增益
	if strings.Index(item, "】+【") > 0 || strings.Index(item, "）+【") > 0 {
		BuffIgnore++
		return true, nil
	} else {
		return false, nil
	}
}

func (b *Buff) parsePvpItem(item string) (bool, []model.CharacterModifier) {
	// 忽略竞技场模式
	if strings.HasPrefix(item, "竞技场模式") {
		BuffIgnore++
		return true, nil
	} else {
		return false, nil
	}
}

func (b *Buff) parseTowerItem(item string) (bool, []model.CharacterModifier) {
	// 忽略竞技场模式
	if strings.HasPrefix(item, "达纳托斯之塔") {
		BuffIgnore++
		return true, nil
	} else {
		return false, nil
	}
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
			condition = strings.TrimPrefix(condition, "精炼+10时，")
		} else {
			rate = 15
		}
	}
	var modifiers []model.CharacterModifier
	if len(condition) > 0 {
		if _, ms, err := b.parseEffects(condition, 1); err != nil {
			return false, nil, err
		} else if len(ms) > 0 {
			modifiers = append(modifiers, ms...)
		}
	}
	if _, ms, err := b.parseEffects(effect, rate); err != nil {
		return false, nil, err
	} else if len(ms) > 0 {
		modifiers = append(modifiers, ms...)
	}
	return true, modifiers, nil
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
	if cap := b.cap(item, "精炼+", "精炼值+", "精炼等级+", "精炼等级达到+", "精炼至+", "当精炼+", "当精炼到+", "武器精炼+", "当盔甲精炼+", "当武器精炼+", "当副手精炼+"); cap <= 0 {
		return false, nil, nil
	} else if idx := strings.Index(item, "时"); idx < 0 {
		return false, nil, nil
	} else if refineStr, effectStr := item[cap:idx], strings.TrimPrefix(item[idx+3:], "，"); strings.Index(refineStr, "、") >= 0 {
		return b.parseRefineItemSplit(refineStr, effectStr, "、")
	} else if strings.Index(refineStr, "，") >= 0 {
		return b.parseRefineItemSplit(refineStr, effectStr, "，")
	} else if strings.Index(refineStr, ",") >= 0 {
		return b.parseRefineItemSplit(refineStr, effectStr, ",")
	} else if refine, condition, err := b.parseRefineWithCondition(refineStr); err != nil {
		return false, nil, errors.Wrapf(err, "parseRefineItem: %s", item)
	} else if match, modifiers, err := b.parseRefineEffects(refine, effectStr); err != nil {
		return false, nil, errors.Wrapf(err, "parseRefineItem: %s", item)
	} else if !match {
		return false, nil, nil
	} else if len(modifiers) == 0 || len(condition) == 0 {
		return true, modifiers, nil
	} else {
		//TODO 解析条件并限制
		return true, modifiers, nil
	}
}

func (b *Buff) parseRefineItemSplit(refineStr, effectStr, split string) (bool, []model.CharacterModifier, error) {
	var refines []int
	for _, refineStrSplit := range strings.Split(refineStr, split) {
		if refine, err := strconv.Atoi(refineStrSplit); err != nil {
			return false, nil, errors.WithStack(err)
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
		return false, nil, errors.Errorf("parseRefineItemSplit[%s] -- %d%s\n", split, refines, effectStr)
	}
}

func (b *Buff) parseRefineSplit(base, effectStr, split string, refines ...int) (bool, []model.CharacterModifier, error) {
	prefix := ""
	if strings.HasPrefix(effectStr, "+") {
		prefix = "+"
	} else if strings.HasPrefix(effectStr, "-") {
		prefix = "-"
	}
	if effects := strings.Split(effectStr, split); len(effects) != len(refines) {
		return false, nil, errors.Errorf("count mismatch: [%d]%d --> [%d]%s", len(refines), refines, len(effects), effects)
	} else {
		var modifiers []model.CharacterModifier
		for idx, refine := range refines {
			if match, sub, err := b.parseRefineEffects(refine, base+prefix+strings.TrimPrefix(effects[idx], prefix)); err != nil {
				return false, nil, err
			} else if !match {
				return false, nil, nil
			} else if len(sub) > 0 {
				modifiers = append(modifiers, sub...)
			}
		}
		return true, modifiers, nil
	}
}

func (b *Buff) parseRefineEffects(refine int, effectStr string) (bool, []model.CharacterModifier, error) {
	if _, modifiers, err := b.parseEffects(effectStr, 1); err != nil {
		return false, nil, err
	} else if len(modifiers) == 0 {
		return true, nil, nil
	} else {
		// TODO 限制精炼等级
		return true, modifiers, nil
	}
}

func (b *Buff) parseConditionItem(item string) (bool, []model.CharacterModifier, error) {
	condition, effect := "", ""
	if idx := strings.Index(item, "时，"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+6:]
	} else if idx := strings.Index(item, "时,"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+4:]
	} else {
		return false, nil, nil
	}
	var modifiers []model.CharacterModifier
	if _, sub, err := b.parseEffects(condition, 1); err != nil {
		return false, nil, err
	} else if len(sub) > 0 {
		modifiers = append(modifiers, sub...)
	}
	switch condition {
	case "物理攻击时", "普通攻击时", "技能攻击时", "使用技能时", "使用拳刃类武器时", "使用物理伤害技能攻击时", "使用短剑类武器时",
		"使用演奏类技能时", "魔法技能攻击时", "魔法攻击目标时", "远程普通攻击时", "装备来复枪类武器时", "装备弓类型武器时",
		"力量在75以上时", "当诗人或舞娘系职业演奏时", "技能攻击玩家时", "技能攻击目标时", "攻击时", "攻击目标时", "普通攻击暴击时",
		"生命值100%时", "SP值100%时", "自身生命值大于50%时", "击杀目标时":
	case "普攻伤害+7.5%，当敏捷大于180时":
	default:
		if strings.Index(condition, "装备") >= 0 && strings.HasSuffix(condition, "卡片时") {
		} else if strings.Index(condition, "达到") >= 0 {
		} else if strings.Index(condition, "佩戴者是") >= 0 {
		} else if strings.HasSuffix(condition, "族装备时") || strings.HasSuffix(condition, "系装备时") {
		} else {
			//TODO 其他需要解析的条件
			BuffIgnore++
			return true, modifiers, nil
		}
	}
	if _, sub, err := b.parseEffects(effect, 1); err != nil {
		return false, nil, err
	} else if len(sub) > 0 {
		modifiers = append(modifiers, sub...)
	}
	return true, modifiers, nil
}

func (b *Buff) parseEffects(effectStr string, rate int) (bool, []model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	runeArray, pos := []rune(effectStr), 0
	for idx, char := range runeArray {
		if (char == '、' || char == '，' || char == ',') && b.isEndOfDigit(runeArray[idx-1]) {
			if match, modifier, err := b.parseEffect(string(runeArray[pos:idx]), rate); err != nil {
				//TODO 处理异常
				return false, nil, nil
			} else if !match {
				return false, modifiers, nil
			} else if modifier != nil {
				modifiers = append(modifiers, modifier)
			}
			pos = idx + 1
		}
	}
	if match, modifier, err := b.parseEffect(string(runeArray[pos:]), rate); err != nil {
		//TODO 处理异常
		return false, nil, nil
	} else if !match {
		return false, modifiers, nil
	} else if modifier != nil {
		modifiers = append(modifiers, modifier)
	}
	return true, modifiers, nil
}

func (b *Buff) parseEffect(effectStr string, rate int) (bool, model.CharacterModifier, error) {
	runeArray, pos, percentage := []rune(effectStr), 0, strings.HasSuffix(effectStr, "%")
	for idx, char := range runeArray {
		if char == '+' || char == '-' {
			key, val := string(runeArray[pos:idx]), strings.TrimSuffix(string(runeArray[idx+1:]), "%")
			if floatVal, err := strconv.ParseFloat(val, 64); err != nil {
				return false, nil, errors.WithStack(err)
			} else {
				if rate > 1 {
					floatVal *= float64(rate)
				}
				if char == '-' {
					floatVal = -floatVal
				}
				if modifier, exist := b.find(key, floatVal, percentage); !exist {
					return false, nil, nil
				} else {
					return true, modifier, nil
				}
			}
		}
	}
	return false, nil, nil
}

func (b *Buff) isEndOfDigit(s rune) bool {
	return s == '%' || (s >= '0' && s <= '9')
}

func (b *Buff) inBrackets(str string) (string, []string) {
	var brackets []string
	after, pos, runeArray := "", 0, []rune(str)
	for idx, char := range runeArray {
		if char == '（' {
			after += string(runeArray[pos:idx])
			pos = idx + 1
		} else if char == '）' {
			brackets = append(brackets, string(runeArray[pos:idx]))
			pos = idx + 1
		}
	}
	after += string(runeArray[pos:])
	return after, brackets
}
