package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"gopkg.in/yaml.v3"
)

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	Attack      int     //攻击
	AttackPer   float64 //攻击%
	Spike       float64 //穿刺
	Damage      float64 //伤害%
	Refine      float64 //精炼攻击
	Critical    float64 //暴击
	CriticalPer float64 //暴伤%

	Defence    int
	DefencePer float64
	Resist     float64 //伤害减免%
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
	raceResist   map[race.Race]float64     //种族减伤%
	shapeDamage  map[shape.Shape]float64   //体型增伤%
	shapeResist  map[shape.Shape]float64   //体型减伤%
	natureDamage map[nature.Nature]float64 //属性增伤%
	natureResist map[nature.Nature]float64 //属性减伤%
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.Attack += incr.Attack
		g.AttackPer += incr.AttackPer
		g.Spike += incr.Spike
		g.Damage += incr.Damage
		g.Refine += incr.Refine
		g.Critical += incr.Critical
		g.CriticalPer += incr.CriticalPer

		g.Defence += incr.Defence
		g.DefencePer += incr.DefencePer
		g.Resist += incr.Resist
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

func (p *Profits) AddRaceResist(incr *map[race.Race]float64) {
	if incr != nil {
		if p.raceResist == nil {
			p.raceResist = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.raceResist[n]; exist {
				p.raceResist[n] = ov + v
			} else {
				p.raceResist[n] = v
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

func (p *Profits) AddShapeResist(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.shapeResist == nil {
			p.shapeResist = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.shapeResist[n]; exist {
				p.shapeResist[n] = ov + v
			} else {
				p.shapeResist[n] = v
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

func (p *Profits) AddNatureResist(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureResist == nil {
			p.natureResist = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureResist[n]; exist {
				p.natureResist[n] = ov + v
			} else {
				p.natureResist[n] = v
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

func (p *Profits) setAttack(magic bool, attack int) {
	if magic {
		p.magical.Attack = attack
	} else {
		p.physical.Attack = attack
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

//伤害减免%
func (p *Profits) Resist(magic bool) float64 {
	if magic {
		return p.magical.Resist
	} else {
		return p.physical.Resist
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

//精炼攻击
func (p *Profits) Refine(magic bool) float64 {
	if magic {
		return p.magical.Refine
	} else {
		return p.physical.Refine
	}
}

//暴击
func (p *Profits) Critical(magic bool) float64 {
	if magic {
		return p.magical.Critical
	} else {
		return p.physical.Critical
	}
}

//暴伤%
func (p *Profits) CriticalPer(magic bool) float64 {
	if magic {
		return p.magical.CriticalPer
	} else {
		return p.physical.CriticalPer
	}
}

func (p *Profits) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	rate = 1 + p.damage.Skill/100                                         //*(1+技能伤害加成%)
	rate *= 1 + p.natureDamage[target.nature]/100                         //*(1+属性魔物增伤%)
	rate *= 1 - target.profits.natureResist[skillNature]/100              //*(1-属性减伤%)
	rate *= 1 + p.natureAttack[skillNature]/100                           //*(1+属性攻击%)
	rate *= 1 + p.Damage(magic)/100                                       //*(1+伤害加成%)
	rate *= 1.027 + p.Spike(magic)/100 - target.profits.Resist(magic)/100 //*(1.027+穿刺%-伤害减免%)
	return
}

func (p *Profits) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "physical":
					if err = sub.Decode(&p.physical); err != nil {
						return
					}
				case "magical":
					if err = sub.Decode(&p.magical); err != nil {
						return
					}
				case "damage":
					if err = sub.Decode(&p.damage); err != nil {
						return
					}
				case "natureAttack":
					if err = sub.Decode(&p.natureAttack); err != nil {
						return
					}
				case "raceDamage":
					if err = sub.Decode(&p.raceDamage); err != nil {
						return
					}
				case "raceResist":
					if err = sub.Decode(&p.raceResist); err != nil {
						return
					}
				case "shapeDamage":
					if err = sub.Decode(&p.shapeDamage); err != nil {
						return
					}
				case "shapeResist":
					if err = sub.Decode(&p.shapeResist); err != nil {
						return
					}
				case "natureDamage":
					if err = sub.Decode(&p.natureDamage); err != nil {
						return
					}
				case "natureResist":
					if err = sub.Decode(&p.natureResist); err != nil {
						return
					}
				default:
					fmt.Printf("missing decode Profits.%s: %+v\n", lastAttr, sub)
				}
			}
		}
	}
	return
}
