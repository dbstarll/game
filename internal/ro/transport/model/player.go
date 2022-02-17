package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/model"
	"github.com/dbstarll/game/internal/ro/romel"
	"go.uber.org/zap"
)

type PlayerModel struct {
	CharacterName string        `json:"character-name"`
	Manual        *[]romel.Buff `json:"manual"`
}

func (m *PlayerModel) character() *model.Character {
	character := model.NewCharacter(types.Player, race.Human, nature.Neutral, shape.Medium)

	if m.Manual != nil {
		for _, buff := range *m.Manual {
			if modifiers := buff.Effect(); len(modifiers) == 0 {
				zap.S().Infof("unknown buff: %s", buff)
			} else {
				for _, modifier := range modifiers {
					modifier(character)
				}
			}
		}
	}

	return character
}

func (m *PlayerModel) Result() interface{} {
	return &map[string]interface{}{
		"player":    m,
		"character": m.character(),
	}
}
