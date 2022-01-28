package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/abnormal"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/model/general"
	"gopkg.in/yaml.v3"
)

type Refine struct {
	Weapon int //武器精炼等级
	Ring1  int //饰品1精炼等级
	Ring2  int //饰品2精炼等级
}

type Profits struct {
	physical       Gains
	magical        Gains
	general        general.General
	refine         Refine
	natureAttack   map[nature.Nature]float64     //属性攻击%
	raceDamage     map[race.Race]float64         //种族增伤%
	raceResist     map[race.Race]float64         //种族减伤%
	shapeDamage    map[shape.Shape]float64       //体型增伤%
	shapeResist    map[shape.Shape]float64       //体型减伤%
	natureDamage   map[nature.Nature]float64     //属性增伤%
	natureResist   map[nature.Nature]float64     //属性减伤%
	abnormalResist map[abnormal.Abnormal]float64 //异常状态抵抗%
}

func (p *Profits) gains(magic bool) *Gains {
	if magic {
		return &p.magical
	} else {
		return &p.physical
	}
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

func (p *Profits) DelNatureAttack(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureAttack == nil {
			p.natureAttack = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureAttack[n]; exist {
				p.natureAttack[n] = ov - v
			} else {
				p.natureAttack[n] = -v
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

func (p *Profits) DelRaceDamage(incr *map[race.Race]float64) {
	if incr != nil {
		if p.raceDamage == nil {
			p.raceDamage = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.raceDamage[n]; exist {
				p.raceDamage[n] = ov - v
			} else {
				p.raceDamage[n] = -v
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

func (p *Profits) DelRaceResist(incr *map[race.Race]float64) {
	if incr != nil {
		if p.raceResist == nil {
			p.raceResist = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.raceResist[n]; exist {
				p.raceResist[n] = ov - v
			} else {
				p.raceResist[n] = -v
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

func (p *Profits) DelShapeDamage(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.shapeDamage == nil {
			p.shapeDamage = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.shapeDamage[n]; exist {
				p.shapeDamage[n] = ov - v
			} else {
				p.shapeDamage[n] = -v
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

func (p *Profits) DelShapeResist(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.shapeResist == nil {
			p.shapeResist = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.shapeResist[n]; exist {
				p.shapeResist[n] = ov - v
			} else {
				p.shapeResist[n] = -v
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

func (p *Profits) DelNatureDamage(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureDamage == nil {
			p.natureDamage = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureDamage[n]; exist {
				p.natureDamage[n] = ov - v
			} else {
				p.natureDamage[n] = -v
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

func (p *Profits) DelNatureResist(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.natureResist == nil {
			p.natureResist = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.natureResist[n]; exist {
				p.natureResist[n] = ov - v
			} else {
				p.natureResist[n] = -v
			}
		}
	}
}

func (p *Profits) AddAbnormalResist(incr *map[abnormal.Abnormal]float64) {
	if incr != nil {
		if p.abnormalResist == nil {
			p.abnormalResist = make(map[abnormal.Abnormal]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.abnormalResist[n]; exist {
				p.abnormalResist[n] = ov + v
			} else {
				p.abnormalResist[n] = v
			}
		}
	}
}

func (p *Profits) DelAbnormalResist(incr *map[abnormal.Abnormal]float64) {
	if incr != nil {
		if p.abnormalResist == nil {
			p.abnormalResist = make(map[abnormal.Abnormal]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.abnormalResist[n]; exist {
				p.abnormalResist[n] = ov - v
			} else {
				p.abnormalResist[n] = -v
			}
		}
	}
}

func (p *Profits) weaponSpikes() float64 {
	return p.refine.weaponSpikes()
}

func (r *Refine) weaponSpikes() float64 {
	return r.weaponSpike(r.Weapon) + r.weaponSpike(r.Ring1) + r.weaponSpike(r.Ring2)
}

func (r *Refine) weaponSpike(lvl int) float64 {
	switch {
	case lvl >= 15:
		return 3.6
	case lvl >= 10:
		return 1.8
	case lvl >= 5:
		return 0.9
	default:
		return 0.0
	}
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
				case "general":
					if err = sub.Decode(&p.general); err != nil {
						return
					}
				case "refine":
					if err = sub.Decode(&p.refine); err != nil {
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
