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
	attack      int     //攻击
	attackPer   float64 //攻击%
	spike       float64 //穿刺
	damage      float64 //伤害%
	refine      float64 //精炼攻击
	critical    float64 //暴击
	criticalPer float64 //暴伤%

	defence    int     //防御
	defencePer float64 //防御%
	resist     float64 //伤害减免%
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
		g.attack += incr.attack
		g.attackPer += incr.attackPer
		g.spike += incr.spike
		g.damage += incr.damage
		g.refine += incr.refine
		g.critical += incr.critical
		g.criticalPer += incr.criticalPer

		g.defence += incr.defence
		g.defencePer += incr.defencePer
		g.resist += incr.resist
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
		return p.magical.attack
	} else {
		return p.physical.attack
	}
}

func (p *Profits) setAttack(magic bool, attack int) {
	if magic {
		p.magical.attack = attack
	} else {
		p.physical.attack = attack
	}
}

//攻击%
func (p *Profits) AttackPer(magic bool) float64 {
	if magic {
		return p.magical.attackPer
	} else {
		return p.physical.attackPer
	}
}

//装备防御
func (p *Profits) Defence(magic bool) int {
	if magic {
		return p.magical.defence
	} else {
		return p.physical.defence
	}
}

func (p *Profits) setDefence(magic bool, defence int) {
	if magic {
		p.magical.defence = defence
	} else {
		p.physical.defence = defence
	}
}

//防御%
func (p *Profits) DefencePer(magic bool) float64 {
	if magic {
		return p.magical.defencePer
	} else {
		return p.physical.defencePer
	}
}

//伤害减免%
func (p *Profits) Resist(magic bool) float64 {
	if magic {
		return p.magical.resist
	} else {
		return p.physical.resist
	}
}

//穿刺
func (p *Profits) Spike(magic bool) float64 {
	if magic {
		return p.magical.spike
	} else {
		return p.physical.spike
	}
}

//伤害%
func (p *Profits) Damage(magic bool) float64 {
	if magic {
		return p.magical.damage
	} else {
		return p.physical.damage
	}
}

//精炼攻击
func (p *Profits) Refine(magic bool) float64 {
	if magic {
		return p.magical.refine
	} else {
		return p.physical.refine
	}
}

//暴击
func (p *Profits) Critical(magic bool) float64 {
	if magic {
		return p.magical.critical
	} else {
		return p.physical.critical
	}
}

//暴伤%
func (p *Profits) CriticalPer(magic bool) float64 {
	if magic {
		return p.magical.criticalPer
	} else {
		return p.physical.criticalPer
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

func (g *Gains) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "attack":
					if err = sub.Decode(&g.attack); err != nil {
						return
					}
				case "attackPer":
					if err = sub.Decode(&g.attackPer); err != nil {
						return
					}
				case "spike":
					if err = sub.Decode(&g.spike); err != nil {
						return
					}
				case "damage":
					if err = sub.Decode(&g.damage); err != nil {
						return
					}
				case "refine":
					if err = sub.Decode(&g.refine); err != nil {
						return
					}
				case "critical":
					if err = sub.Decode(&g.critical); err != nil {
						return
					}
				case "criticalPer":
					if err = sub.Decode(&g.criticalPer); err != nil {
						return
					}
				case "defence":
					if err = sub.Decode(&g.defence); err != nil {
						return
					}
				case "defencePer":
					if err = sub.Decode(&g.defencePer); err != nil {
						return
					}
				case "resist":
					if err = sub.Decode(&g.resist); err != nil {
						return
					}
				default:
					fmt.Printf("missing decode Gains.%s: %+v\n", lastAttr, sub)
				}
			}
		}
	}
	return
}
