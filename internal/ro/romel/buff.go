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
		} else if ms, err := b.parseEffects(condition, 1); err != nil {
			return false, nil, err
		} else if len(ms) > 0 {
			modifiers = append(modifiers, ms...)
		}
	}
	if ms, err := b.parseEffects(effect, rate); err != nil {
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
	if modifiers, err := b.parseEffects(effectStr, 1); err != nil {
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
	condition, effect := "", ""
	if idx := strings.Index(item, "时，"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+6:]
	} else if idx := strings.Index(item, "时,"); idx > 0 {
		condition, effect = item[:idx+3], item[idx+4:]
	} else if idx := strings.Index(item, "时"); idx < 0 {
		return false, nil, nil
	} else if strings.Index(item, "时间") == idx {
		return false, nil, nil
	} else if strings.Index(item, "时刻") == idx {
		return false, nil, nil
	} else if strings.Index(item, "时长") == idx {
		return false, nil, nil
	} else if strings.Index(item, "同时") == idx-3 {
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
		if sub, err := b.parseEffects(condition, 1); err != nil {
			return false, nil, err
		} else if len(sub) > 0 {
			modifiers = append(modifiers, sub...)
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
	if sub, err := b.parseEffects(effect, 1); err != nil {
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
				if r, _ := utf8.DecodeRuneInString(quality); r >= '0' && r <= '9' {
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
	} else if match, quality, effect := b.perQuality(item, "每", "会为装备者提供", "额外增加自身",
		"额外增加", "额外带来", "额外提高", "增加自身", "增加", "加"); !match {
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
	max, effectNoBrackets := 0, effectStr
	if after, brackets := b.inBrackets(effectStr); len(brackets) > 0 {
		effectNoBrackets = after
		if strings.Index(brackets[0], "额外") >= 0 {
			switch brackets[0] {
			case "最多额外增加1500点物理攻击":
				max = 1500
			case "额外最多15%":
				max = 15
			default:
				return nil, errors.Errorf("parsePerQualityEffect: [%s]%s -- %s", qualityStr, effectNoBrackets, brackets)
			}
		}
	}

	effect := ""
	if idx := strings.Index(effectNoBrackets, "点"); idx > 0 {
		effect = fmt.Sprintf("%s+%s", effectNoBrackets[idx+3:], effectNoBrackets[:idx])
	} else if idx := strings.Index(effectNoBrackets, "%"); idx > 0 {
		effect = fmt.Sprintf("%s+%s", effectNoBrackets[idx+1:], effectNoBrackets[:idx+1])
	} else {
		return nil, errors.Errorf("parsePerQualityEffect: [%s]%s", qualityStr, effectNoBrackets)
	}

	if max > 0 {
		_, modifier, err := b.parseEffect(effect, max)
		return modifier, err
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
				if _, modifier, err := b.parseEffect(effect, rateFn(character, num)); err != nil {
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
		} else if _, modifier, err := b.parseEffect(effect, 1); err != nil {
			return nil, err
		} else {
			return modifier, nil
		}
	}
}

func (b *Buff) parseEffects(effectStr string, rate int) ([]model.CharacterModifier, error) {
	var modifiers []model.CharacterModifier
	runeArray, pos := []rune(effectStr), 0
	for idx, char := range runeArray {
		if (char == '、' || char == '，' || char == ',') && b.isEndOfDigit(runeArray[idx-1]) {
			if match, modifier, err := b.parseEffect(string(runeArray[pos:idx]), rate); err != nil {
				//TODO 处理异常
				return nil, nil
			} else if !match {
				return modifiers, nil
			} else if modifier != nil {
				modifiers = append(modifiers, modifier)
			}
			pos = idx + 1
		}
	}
	if match, modifier, err := b.parseEffect(string(runeArray[pos:]), rate); err != nil {
		//TODO 处理异常
		return nil, nil
	} else if !match {
		return modifiers, nil
	} else if modifier != nil {
		modifiers = append(modifiers, modifier)
	}
	return modifiers, nil
}

func (b *Buff) parseEffect(effectStr string, rate int) (bool, model.CharacterModifier, error) {
	runeArray, pos, percentage := []rune(effectStr), 0, strings.HasSuffix(effectStr, "%")
	for idx, char := range runeArray {
		if char == '+' || char == '-' {
			key, val := string(runeArray[pos:idx]), strings.TrimSuffix(string(runeArray[idx+1:]), "%")
			if floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err != nil {
				return false, nil, errors.WithStack(err)
			} else {
				if rate > 1 {
					floatVal *= float64(rate)
				}
				if char == '-' {
					floatVal = -floatVal
				}
				if modifier, exist := b.find(key, floatVal, percentage); !exist {
					return true, nil, nil
				} else if modifier == nil {
					return true, nil, nil
				} else {

					//BuffUnknown += 1
					//if oc, ok := Buffs[item]; ok {
					//	Buffs[item] = oc + 1
					//} else {
					//	Buffs[item] = 1
					//}

					return true, modifier, nil
				}
			}
		}
	}
	if modifier, exist := b.find(effectStr, 0, percentage); !exist {
		return false, nil, nil
	} else if modifier == nil {
		return false, nil, nil
	} else {
		return true, modifier, nil
	}
}

func (b *Buff) isEndOfDigit(s rune) bool {
	return s == '%' || s == ' ' || (s >= '0' && s <= '9')
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
