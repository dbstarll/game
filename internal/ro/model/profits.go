package model

import (
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
)

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	Attack    int     //攻击
	AttackPer float64 //攻击%
	Spike     float64 //穿刺
	Damage    float64 //伤害%

	Defence    int
	DefencePer float64
}

type Damage struct {
	Skill float64 //技能伤害加成%
	MVP   float64 //MVP增伤%
}

type Profits struct {
	physical     Gains
	magical      Gains
	damage       Damage
	natureAttack map[nature.Nature]float64 //属性攻击%
	raceDamage   map[race.Race]float64     //种族增伤%
	shapeDamage  map[shape.Shape]float64   //体型增伤%
	natureDamage map[nature.Nature]float64 //属性增伤%
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.Attack += incr.Attack
		g.AttackPer += incr.AttackPer
		g.Spike += incr.Spike
		g.Damage += incr.Damage

		g.Defence += incr.Defence
		g.DefencePer += incr.DefencePer
	}
}

func (d *Damage) Add(incr *Damage) {
	if incr != nil {
		d.Skill += incr.Skill
		d.MVP += incr.MVP
	}
}

func (p *Profits) AddGains(magic bool, incr *Gains) {
	if magic {
		p.magical.Add(incr)
	} else {
		p.physical.Add(incr)
	}
}

func (p *Profits) AddDamage(incr *Damage) {
	p.damage.Add(incr)
}

func (p *Profits) AddNatureAttack(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureAttack == nil {
			p.natureAttack = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureAttack[n]; exist {
				p.natureAttack[n] = ov + v
			} else {
				p.natureAttack[n] = v
			}
		}
	}
}

func (p *Profits) AddRaceDamage(incr *map[race.Race]float64) {
	if incr != nil {
		if p.raceDamage == nil {
			p.raceDamage = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.raceDamage[n]; exist {
				p.raceDamage[n] = ov + v
			} else {
				p.raceDamage[n] = v
			}
		}
	}
}

func (p *Profits) AddShapeDamage(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.shapeDamage == nil {
			p.shapeDamage = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.shapeDamage[n]; exist {
				p.shapeDamage[n] = ov + v
			} else {
				p.shapeDamage[n] = v
			}
		}
	}
}

func (p *Profits) AddNatureDamage(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureDamage == nil {
			p.natureDamage = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureDamage[n]; exist {
				p.natureDamage[n] = ov + v
			} else {
				p.natureDamage[n] = v
			}
		}
	}
}

//装备攻击
func (p *Profits) Attack(magic bool) int {
	if magic {
		return p.magical.Attack
	} else {
		return p.physical.Attack
	}
}

//攻击%
func (p *Profits) AttackPer(magic bool) float64 {
	if magic {
		return p.magical.AttackPer
	} else {
		return p.physical.AttackPer
	}
}

//装备防御
func (p *Profits) Defence(magic bool) int {
	if magic {
		return p.magical.Defence
	} else {
		return p.physical.Defence
	}
}

func (p *Profits) setDefence(magic bool, defence int) {
	if magic {
		p.magical.Defence = defence
	} else {
		p.physical.Defence = defence
	}
}

//防御%
func (p *Profits) DefencePer(magic bool) float64 {
	if magic {
		return p.magical.DefencePer
	} else {
		return p.physical.DefencePer
	}
}

//穿刺
func (p *Profits) Spike(magic bool) float64 {
	if magic {
		return p.magical.Spike
	} else {
		return p.physical.Spike
	}
}

//伤害%
func (p *Profits) Damage(magic bool) float64 {
	if magic {
		return p.magical.Damage
	} else {
		return p.physical.Damage
	}
}

func (p *Profits) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	rate = 1 + p.damage.Skill/100                 //*(1+技能伤害加成%)
	rate *= 1 + p.shapeDamage[target.shape]/100   //*(1+体型增伤%)
	rate *= 1 + p.raceDamage[target.race]/100     //*(1+种族增伤%)
	rate *= 1 + p.natureDamage[target.nature]/100 //*(1+属性魔物增伤%)
	rate *= 1 + p.natureAttack[skillNature]/100   //*(1+属性攻击%)
	rate *= 1 + p.Damage(magic)/100               //*(1+伤害加成%)
	rate *= 1.027 + p.Spike(magic)/100            //*(1.027+穿刺%)
	return
}
