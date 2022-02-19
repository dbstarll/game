package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/model/buff"
	"github.com/dbstarll/game/internal/ro/romel"
	"go.uber.org/zap"
)

type PlayerModel struct {
	CharacterName string        `json:"character-name"`
	Manual        *[]romel.Buff `json:"manual"`
	Union         UnionModel    `json:"union"`
	Rune          *[]romel.Buff `json:"rune"`
}

type UnionModel struct {
	Pray    *[]romel.Buff `json:"pray"`
	Attack  *[]romel.Buff `json:"attack"`
	Defence *[]romel.Buff `json:"defence"`
	Element *[]romel.Buff `json:"element"`
}

func (m *PlayerModel) character() *model.Character {
	character := model.NewCharacter(types.Player, race.Human, nature.Neutral, shape.Medium)

	buff.Quality(9)(character) //B级冒险家属性
	m.apply(character, m.Manual)
	m.apply(character, m.Union.Pray)
	m.apply(character, m.Union.Attack)
	m.apply(character, m.Union.Defence)
	m.apply(character, m.Union.Element)
	m.apply(character, m.Rune)

	return character
}

func (m *PlayerModel) apply(character *model.Character, buffs *[]romel.Buff) {
	if buffs != nil {
		for _, buff := range *buffs {
			if modifiers := buff.Effect(); len(modifiers) == 0 {
				zap.S().Infof("unknown buff: %s", buff)
			} else {
				for _, modifier := range modifiers {
					modifier(character)
				}
			}
		}
	}
}

func (m *PlayerModel) Result() interface{} {
	return &map[string]interface{}{
		"player":    m,
		"character": m.character(),
	}
}
