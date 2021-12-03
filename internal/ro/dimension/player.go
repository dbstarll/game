package dimension

import "github.com/dbstarll/game/internal/ro/dimension/job"

type Player struct {
	Character
	_job job.Job
}

//装备攻击
func (p *Player) EquipmentAttack(magic bool) int {
	if magic {
		return p.Character.EquipmentAttack(magic)
	} else {
		//装备物理攻击 = (装备，强化，附魔，卡片，头饰，祈祷，buff等合计)+ BaseLvAtkRate*人物等级
		return p.Character.EquipmentAttack(magic) + p._job.BaseLvAtkRate()*p.level.base
	}
}
