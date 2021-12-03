package dimension

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

type Character struct {
	_nature          nature.Nature
	_race            race.Race
	_shape           shape.Shape
	level            Level
	quality          Quality
	equipmentProfits Profits
}
