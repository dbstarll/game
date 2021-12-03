package dimension

import (
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

type Player struct {
	job job.Job
	*Character
}

func NewPlayer(job job.Job, modifiers ...CharacterModifier) *Player {
	return &Player{
		job:       job,
		Character: NewCharacter(nature.Neutral, race.Human, shape.Medium, modifiers...),
	}
}

//装备攻击
func (p *Player) EquipmentAttack(magic bool) int {
	if magic {
		return p.Character.EquipmentAttack(magic)
	} else {
		//装备物理攻击 = (装备，强化，附魔，卡片，头饰，祈祷，buff等合计)+ BaseLvAtkRate*人物等级
		return p.Character.EquipmentAttack(magic) + p.job.BaseLvAtkRate()*p.level.Base
	}
}
