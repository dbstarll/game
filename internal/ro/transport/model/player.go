package model

import "github.com/dbstarll/game/internal/ro/romel"

type PlayerModel struct {
	Manual *[]romel.Buff `json:"manual"`
}
