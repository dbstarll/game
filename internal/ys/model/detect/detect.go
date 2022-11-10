package detect

import (
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/model/action"
	"github.com/dbstarll/game/internal/ys/model/attr"
	"github.com/dbstarll/game/internal/ys/model/character"
	"github.com/dbstarll/game/internal/ys/model/enemy"
	"sort"
)

type Profit struct {
	Name  string
	Value float64
}

type FinalDamage func(character *character.Character, enemy *enemy.Enemy, action *action.Action, debug bool, finalModifiers ...attr.AttributeModifier) float64

var (
	baseDetects = initBaseDetects(map[string]*attr.Modifier{})
)

func initBaseDetects(detects map[string]*attr.Modifier) map[string]*attr.Modifier {
	for _, entry := range entry.Entries {
		if rate, fn := entry.Multiple(); rate == 0 || fn == nil {
			continue
		} else if ratio, exist := entry.Secondary(); exist {
			detects[entry.String()] = attr.NewCharacterModifier(fn(3.89 * rate / ratio))
		} else {
			detects[entry.String()] = attr.NewCharacterModifier(fn(3.89 * rate))
		}
	}
	return detects
}

func ProfitDetect(character *character.Character, enemy *enemy.Enemy, action *action.Action, baseDetect bool, fn FinalDamage, customDetects map[string]*attr.Modifier, finalModifiers ...attr.AttributeModifier) []*Profit {
	base := fn(character, enemy, action, false, finalModifiers...)
	var profits []*Profit
	if baseDetect {
		for name, modifier := range baseDetects {
			cancel := modifier.Apply(character, enemy, action)
			value := fn(character, enemy, action, false, finalModifiers...)
			if value != base {
				profits = append(profits, &Profit{
					Name:  name,
					Value: 100 * (value - base) / base,
				})
			}
			cancel()
		}
	}
	for name, modifier := range customDetects {
		cancel := modifier.Apply(character, enemy, action)
		value := fn(character, enemy, action, false, finalModifiers...)
		profits = append(profits, &Profit{
			Name:  name,
			Value: 100 * (value - base) / base,
		})
		cancel()
	}
	sort.Slice(profits, func(i, j int) bool {
		if profits[i].Value < profits[j].Value {
			return false
		} else if profits[i].Value > profits[j].Value {
			return true
		} else {
			return profits[i].Name < profits[j].Name
		}
	})
	return profits
}
