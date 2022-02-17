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
	Weapon int `json:"weapon"` //武器精炼等级
	Ring1  int `json:"ring1"`  //饰品1精炼等级
	Ring2  int `json:"ring2"`  //饰品2精炼等级
}

type Profits struct {
	Physical       Gains                         `json:"physical"`
	Magical        Gains                         `json:"magical"`
	General        general.General               `json:"general"`
	Refine         Refine                        `json:"refine"`
	NatureAttack   map[nature.Nature]float64     `json:"nature_attack"`   //属性攻击%
	RaceDamage     map[race.Race]float64         `json:"race_damage"`     //种族增伤%
	RaceResist     map[race.Race]float64         `json:"race_resist"`     //种族减伤%
	ShapeDamage    map[shape.Shape]float64       `json:"shape_damage"`    //体型增伤%
	ShapeResist    map[shape.Shape]float64       `json:"shape_resist"`    //体型减伤%
	NatureDamage   map[nature.Nature]float64     `json:"nature_damage"`   //属性增伤%
	NatureResist   map[nature.Nature]float64     `json:"nature_resist"`   //属性减伤%
	AbnormalResist map[abnormal.Abnormal]float64 `json:"abnormal_resist"` //异常状态抵抗%
}

func (p *Profits) Gains(magic bool) *Gains {
	if magic {
		return &p.Magical
	} else {
		return &p.Physical
	}
}

func (p *Profits) AddNatureAttack(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureAttack == nil {
			p.NatureAttack = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureAttack[n]; exist {
				p.NatureAttack[n] = ov + v
			} else {
				p.NatureAttack[n] = v
			}
		}
	}
}

func (p *Profits) DelNatureAttack(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureAttack == nil {
			p.NatureAttack = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureAttack[n]; exist {
				p.NatureAttack[n] = ov - v
			} else {
				p.NatureAttack[n] = -v
			}
		}
	}
}

func (p *Profits) AddRaceDamage(incr *map[race.Race]float64) {
	if incr != nil {
		if p.RaceDamage == nil {
			p.RaceDamage = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.RaceDamage[n]; exist {
				p.RaceDamage[n] = ov + v
			} else {
				p.RaceDamage[n] = v
			}
		}
	}
}

func (p *Profits) DelRaceDamage(incr *map[race.Race]float64) {
	if incr != nil {
		if p.RaceDamage == nil {
			p.RaceDamage = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.RaceDamage[n]; exist {
				p.RaceDamage[n] = ov - v
			} else {
				p.RaceDamage[n] = -v
			}
		}
	}
}

func (p *Profits) AddRaceResist(incr *map[race.Race]float64) {
	if incr != nil {
		if p.RaceResist == nil {
			p.RaceResist = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.RaceResist[n]; exist {
				p.RaceResist[n] = ov + v
			} else {
				p.RaceResist[n] = v
			}
		}
	}
}

func (p *Profits) DelRaceResist(incr *map[race.Race]float64) {
	if incr != nil {
		if p.RaceResist == nil {
			p.RaceResist = make(map[race.Race]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.RaceResist[n]; exist {
				p.RaceResist[n] = ov - v
			} else {
				p.RaceResist[n] = -v
			}
		}
	}
}

func (p *Profits) AddShapeDamage(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.ShapeDamage == nil {
			p.ShapeDamage = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.ShapeDamage[n]; exist {
				p.ShapeDamage[n] = ov + v
			} else {
				p.ShapeDamage[n] = v
			}
		}
	}
}

func (p *Profits) DelShapeDamage(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.ShapeDamage == nil {
			p.ShapeDamage = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.ShapeDamage[n]; exist {
				p.ShapeDamage[n] = ov - v
			} else {
				p.ShapeDamage[n] = -v
			}
		}
	}
}

func (p *Profits) AddShapeResist(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.ShapeResist == nil {
			p.ShapeResist = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.ShapeResist[n]; exist {
				p.ShapeResist[n] = ov + v
			} else {
				p.ShapeResist[n] = v
			}
		}
	}
}

func (p *Profits) DelShapeResist(incr *map[shape.Shape]float64) {
	if incr != nil {
		if p.ShapeResist == nil {
			p.ShapeResist = make(map[shape.Shape]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.ShapeResist[n]; exist {
				p.ShapeResist[n] = ov - v
			} else {
				p.ShapeResist[n] = -v
			}
		}
	}
}

func (p *Profits) AddNatureDamage(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureDamage == nil {
			p.NatureDamage = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureDamage[n]; exist {
				p.NatureDamage[n] = ov + v
			} else {
				p.NatureDamage[n] = v
			}
		}
	}
}

func (p *Profits) DelNatureDamage(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureDamage == nil {
			p.NatureDamage = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureDamage[n]; exist {
				p.NatureDamage[n] = ov - v
			} else {
				p.NatureDamage[n] = -v
			}
		}
	}
}

func (p *Profits) AddNatureResist(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureResist == nil {
			p.NatureResist = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureResist[n]; exist {
				p.NatureResist[n] = ov + v
			} else {
				p.NatureResist[n] = v
			}
		}
	}
}

func (p *Profits) DelNatureResist(incr *map[nature.Nature]float64) {
	if incr != nil {
		if p.NatureResist == nil {
			p.NatureResist = make(map[nature.Nature]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.NatureResist[n]; exist {
				p.NatureResist[n] = ov - v
			} else {
				p.NatureResist[n] = -v
			}
		}
	}
}

func (p *Profits) AddAbnormalResist(incr *map[abnormal.Abnormal]float64) {
	if incr != nil {
		if p.AbnormalResist == nil {
			p.AbnormalResist = make(map[abnormal.Abnormal]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.AbnormalResist[n]; exist {
				p.AbnormalResist[n] = ov + v
			} else {
				p.AbnormalResist[n] = v
			}
		}
	}
}

func (p *Profits) DelAbnormalResist(incr *map[abnormal.Abnormal]float64) {
	if incr != nil {
		if p.AbnormalResist == nil {
			p.AbnormalResist = make(map[abnormal.Abnormal]float64)
		}
		for n, v := range *incr {
			if ov, exist := p.AbnormalResist[n]; exist {
				p.AbnormalResist[n] = ov - v
			} else {
				p.AbnormalResist[n] = -v
			}
		}
	}
}

func (p *Profits) weaponSpikes() float64 {
	return p.Refine.weaponSpikes()
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
					if err = sub.Decode(&p.Physical); err != nil {
						return
					}
				case "magical":
					if err = sub.Decode(&p.Magical); err != nil {
						return
					}
				case "general":
					if err = sub.Decode(&p.General); err != nil {
						return
					}
				case "refine":
					if err = sub.Decode(&p.Refine); err != nil {
						return
					}
				case "natureAttack":
					if err = sub.Decode(&p.NatureAttack); err != nil {
						return
					}
				case "raceDamage":
					if err = sub.Decode(&p.RaceDamage); err != nil {
						return
					}
				case "raceResist":
					if err = sub.Decode(&p.RaceResist); err != nil {
						return
					}
				case "shapeDamage":
					if err = sub.Decode(&p.ShapeDamage); err != nil {
						return
					}
				case "shapeResist":
					if err = sub.Decode(&p.ShapeResist); err != nil {
						return
					}
				case "natureDamage":
					if err = sub.Decode(&p.NatureDamage); err != nil {
						return
					}
				case "natureResist":
					if err = sub.Decode(&p.NatureResist); err != nil {
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
