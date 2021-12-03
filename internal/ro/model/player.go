package model

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

func (p *Player) SkillDamageRate(target *Monster, magic bool, skillNature nature.Nature) (rate float64) {
	rate = p.Character.SkillDamageRate(target.Character, magic, skillNature)
	if target.types.IsBoss() {
		rate *= 1 + p.profits.damage.MVP/100 //*(1+MVP增伤%)
	}
	return
}

func (p *Player) SkillEarth() (damage float64) {
	damage = float64(p.quality.Vit*p.quality.Vit) *
		p.PanelDefence(false) / 10000 *
		9.6 //基础技能倍率
	damage *= 1 + 10.0/100 //*(1+守护之盾技能增伤%)
	damage *= 1 + 16.2/100 //*(1+铁蹄直驱符文增伤%)
	return
}
