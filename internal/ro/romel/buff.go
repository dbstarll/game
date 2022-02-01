package romel

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
)

type Buff string

var (
	BuffTotal    = 0
	BuffUnknown  = 0
	BuffError    = 0
	BuffIgnore   = 0
	BuffDetected = 0
	Buffs        = make(map[string]int)
)

func (b Buff) Contains(o Buff) bool {
	return strings.Index(string(b), string(o)) >= 0
}

func (b Buff) Empty() bool {
	return len(string(b)) == 0
}

func (b Buff) Size() int {
	return len(b.Items())
}

func (b Buff) Items() []string {
	return strings.Split(string(b), "\n")
}

func (b Buff) Effect() ([]model.CharacterModifier, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var modifiers []model.CharacterModifier
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "达纳托斯之塔") {
			continue
		}
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
	} else if modifiers, err := b.resolveEffects(item, 1); err != nil {
		zap.S().Warnf("resolveItem: %s", item)
		return nil, nil
	} else {
		return modifiers, nil
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
		} else if modifiers, err := b.resolveRefine(effectStr, refine); err != nil {
			zap.S().Warnf("resolveRefineItem: [%d]%s || %s", refine, effectStr, item)
			return false, nil, nil
		} else {
			return true, modifiers, nil
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
	if modifiers, err := b.resolveEffects(effectStr, 1); err != nil {
		return nil, err
	} else if len(modifiers) == 0 {
		return nil, nil
	} else {
		// TODO 限制精炼等级
		//fmt.Printf("resolveRefine: %d --> %s\n", refine, effectStr)
		return modifiers, nil
	}
}

func (b Buff) resolveEffects(effectStr string, rate int) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	pos, runeArray := 0, []rune(effectStr)
	for idx, char := range runeArray {
		if (char == '、' || char == '，') && b.isEndOfDigit(runeArray[idx-1]) {
			if effect := string(runeArray[pos:idx]); strings.HasSuffix(effect, "精炼+1至+6") {
				if effect == "精炼+1至+6" {
					continue
				} else if sub, err := b.resolveEffect(strings.TrimSuffix(effect, "，精炼+1至+6"), rate); err != nil {
					return nil, err
				} else if sub != nil {
					modifiers = append(modifiers, sub)
				}
			} else if strings.Index(effect, "每精炼+1") >= 0 {
				break
			} else if strings.Index(effect, "精炼每+1") >= 0 {
				break
			} else if sub, err := b.resolveEffect(effect, rate); err != nil {
				return nil, err
			} else if sub != nil {
				modifiers = append(modifiers, sub)
			}
			pos = idx + 1
		}
	}
	if effect := string(runeArray[pos:]); strings.Index(effect, "每精炼+1") >= 0 || strings.Index(effect, "精炼每+1") >= 0 {
		return b.resolvePerRefineItem(effect)
	} else if sub, err := b.resolveEffect(effect, rate); err != nil {
		return nil, err
	} else if sub != nil {
		modifiers = append(modifiers, sub)
	}
	return modifiers, nil
}

func (b Buff) resolvePerRefineItem(effectStr string) ([]model.CharacterModifier, error) {
	condition, effect := "", ""
	if idx := strings.Index(effectStr, "每精炼+1时"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+14:], "，")
	} else if idx := strings.Index(effectStr, "盔甲每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+17:], "，")
	} else if idx := strings.Index(effectStr, "副手每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+17:], "，")
	} else if idx := strings.Index(effectStr, "头饰每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+17:], "，")
	} else if idx := strings.Index(effectStr, "装备这张卡片的武器每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+38:], "，")
	} else if idx := strings.Index(effectStr, "每精炼+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+11:], "，")
	} else if idx := strings.Index(effectStr, "精炼每+1时"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+14:], "，")
	} else if idx := strings.Index(effectStr, "精炼每+1"); idx >= 0 {
		condition, effect = strings.TrimSuffix(effectStr[:idx], "，"), strings.TrimPrefix(effectStr[idx+11:], "，")
	}
	if condition == "" {
		return b.resolveEffects(effect, 15)
	} else if condition == "精炼+1至+6" {
		return b.resolveEffects(effect, 6)
	} else if condition == "精炼+8开始" {
		return b.resolveEffects(effect, 8)
	} else if condition == "精炼+10开始" {
		return b.resolveEffects(effect, 6)
	} else if condition == "精炼+5开始，再" {
		return b.resolveEffects(effect, 10)
	} else if condition == "精炼+8开始，再" {
		return b.resolveEffects(effect, 7)
	} else if condition == "精炼+5以后，再" {
		return b.resolveEffects(effect, 10)
	} else if condition == "精炼+10以后，再" {
		return b.resolveEffects(effect, 5)
	} else if condition == "使用演奏类技能时" {
		return b.resolveEffects(effect, 15)
	} else {
		return b.resolveEffects(effect, 1)
	}
}

func (b Buff) isEndOfDigit(s rune) bool {
	return s == '%' || s == '）' || (s >= '0' && s <= '9')
}

func (b Buff) inBrackets(str string) (string, []string) {
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

func (b Buff) resolveEffect(effectStr string, rate int) (model.CharacterModifier, error) {
	rateStr := ""
	if strings.Index(effectStr, "（总计") > 0 {
		if afterEffect, brackets := b.inBrackets(effectStr); len(brackets) > 0 {
			effectStr = afterEffect
			rateStr = brackets[len(brackets)-1]
		}
	}
	runeArray, pos, condition, percentage := []rune(effectStr), 0, "", strings.HasSuffix(effectStr, "%")
	for idx, char := range runeArray {
		if char == '时' && runeArray[idx+1] != '间' && runeArray[idx+1] != '长' && runeArray[idx-1] != '同' {
			condition = string(runeArray[:idx+1])
			if runeArray[idx+1] == '，' || runeArray[idx+1] == '：' || runeArray[idx+1] == ',' {
				pos = idx + 2
			} else {
				pos = idx + 1
			}
		} else if char == '+' || char == '-' {
			BuffTotal++
			key, val := string(runeArray[pos:idx]), strings.TrimSuffix(string(runeArray[idx+1:]), "%")
			if strings.Index(key, "竞技场模式") >= 0 || strings.Index(condition, "竞技场模式") >= 0 {
				BuffIgnore++
				//忽略竞技场模式加成
				return nil, nil
			} else if strings.Index(key, "达纳托斯之塔") >= 0 || strings.Index(condition, "达纳托斯之塔") >= 0 {
				BuffIgnore++
				//忽略达纳托斯之塔加成
				return nil, nil
			} else if strings.Index(key, "【") >= 0 {
				//过滤掉技能
				BuffIgnore++
				return nil, nil
			} else if strings.Index(key, "恢复") >= 0 {
				BuffIgnore++
				//忽略"生命自然恢复", "SP恢复", "Sp恢复", "魔法恢复", "生命恢复", "Hp恢复"
				return nil, nil
			} else if floatVal, err := strconv.ParseFloat(val, 64); err != nil {
				BuffError++
				zap.S().Warnf("resolveEffect: [%t]%s[%s]%s - %s || %s", percentage, key, string(char), val, condition, effectStr)
				return nil, errors.Wrapf(err, "resolveEffect: [%t]%s[%s]%s - %s || %s", percentage, key, string(char), val, condition, effectStr)
			} else {
				if rate > 1 {
					floatVal *= float64(rate)
				} else if len(rateStr) > 0 {
					val = strings.TrimSuffix(strings.TrimPrefix(rateStr, "总计"), "%")
					if floatVal2, err := strconv.ParseFloat(val, 64); err != nil {
						return nil, errors.Wrapf(err, "resolveEffect: [%t]%s[%s]%s - %s || %s", percentage, key, string(char), val, condition, effectStr)
					} else {
						floatVal = floatVal2
					}
				}
				if char == '-' {
					floatVal = -floatVal
				}
				if m, exist := b.find(key, floatVal, percentage); exist {
					BuffDetected++
					return m, nil
				} else {
					BuffUnknown++
					if oc, exist := Buffs[key]; exist {
						Buffs[key] = oc + 1
					} else {
						Buffs[key] = 1
					}
					//fmt.Printf("\tresolveEffect: [%t]%s[%s]%f - %s || %s\n", percentage, key, string(char), floatVal, condition, effectStr)
					log.Printf("resolveEffect: [%t]%s[%s]%f - %s || %s", percentage, key, string(char), floatVal, condition, effectStr)
					return nil, nil
				}
			}
		}
	}
	fmt.Printf("\tresolveEffect: [%t] - %s || %s\n", percentage, condition, effectStr)
	log.Printf("resolveEffect: [%t] - %s || %s", percentage, condition, effectStr)
	return nil, nil
}
