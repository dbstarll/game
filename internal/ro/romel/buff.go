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
	runeArray, percentage := []rune(effectStr), strings.HasSuffix(effectStr, "%")
	for idx, char := range runeArray {
		if char == '+' || char == '-' {
			key, val := string(runeArray[:idx]), strings.TrimSuffix(string(runeArray[idx+1:]), "%")
			if floatVal, err := strconv.ParseFloat(val, 64); err != nil {
				zap.S().Warnf("resolveEffect: [%t]%s[%s]%s || %s", percentage, key, string(char), val, effectStr)
			} else {
				if char == '-' {
					floatVal = -floatVal
				}
				if m, exist := b.find(key, floatVal, percentage); exist {
					return m, nil
				} else if strings.HasSuffix(key, "恢复") {
					//忽略"生命自然恢复", "SP恢复", "Sp恢复", "魔法恢复", "生命恢复", "Hp恢复"
					return nil, nil
				} else if strings.Index(key, "【") < 0 {
					//过滤掉技能
					fmt.Printf("\tresolveEffect: [%t]%s[%s]%f || %s\n", percentage, key, string(char), floatVal, effectStr)
					log.Printf("resolveEffect: [%t]%s[%s]%f || %s", percentage, key, string(char), floatVal, effectStr)
				}
			}
			break
		}
	}
	return nil, nil
}
