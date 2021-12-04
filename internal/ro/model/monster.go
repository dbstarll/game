package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
)

type Monster struct {
	types types.Types
	*Character
}

func NewMonster(types types.Types, race race.Race, nature nature.Nature, shape shape.Shape, modifiers ...CharacterModifier) *Monster {
	return &Monster{
		types:     types,
		Character: NewCharacter(race, nature, shape, modifiers...),
	}
}
