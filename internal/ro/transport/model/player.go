package model

import "github.com/dbstarll/game/internal/ro/romel"

type PlayerModel struct {
	CharacterName string        `json:"character-name"`
	Manual        *[]romel.Buff `json:"manual"`
}
