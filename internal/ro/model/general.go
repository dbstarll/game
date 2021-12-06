package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type General struct {
	critical             float64 //暴击
	criticalDamage       float64 //暴伤%
	criticalResist       float64 //暴击防护
	criticalDamageResist float64 //爆伤减免%
	ordinary             float64 //普攻伤害加成%
	ordinaryResist       float64 //普攻伤害减免%
	skill                float64 //技能伤害加成%
	skillResist          float64 //技能伤害减免%
	mvp                  float64 //MVP增伤%
	mvpResist            float64 //MVP减伤%
}

func (d *General) Add(incr *General) {
	if incr != nil {
		d.critical += incr.critical
		d.criticalDamage += incr.criticalDamage
		d.criticalResist += incr.criticalResist
		d.criticalDamageResist += incr.criticalDamageResist
		d.ordinary += incr.ordinary
		d.ordinaryResist += incr.ordinaryResist
		d.skill += incr.skill
		d.skillResist += incr.skillResist
		d.mvp += incr.mvp
		d.mvpResist += incr.mvpResist
	}
}

func (d *General) Del(incr *General) {
	if incr != nil {
		d.critical -= incr.critical
		d.criticalDamage -= incr.criticalDamage
		d.criticalResist -= incr.criticalResist
		d.criticalDamageResist -= incr.criticalDamageResist
		d.ordinary -= incr.ordinary
		d.ordinaryResist -= incr.ordinaryResist
		d.skill -= incr.skill
		d.skillResist -= incr.skillResist
		d.mvp -= incr.mvp
		d.mvpResist -= incr.mvpResist
	}
}

func (d *General) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "critical":
					if err = sub.Decode(&d.critical); err != nil {
						return
					}
				case "criticalDamage":
					if err = sub.Decode(&d.criticalDamage); err != nil {
						return
					}
				case "criticalResist":
					if err = sub.Decode(&d.criticalResist); err != nil {
						return
					}
				case "criticalDamageResist":
					if err = sub.Decode(&d.criticalDamageResist); err != nil {
						return
					}
				case "ordinary":
					if err = sub.Decode(&d.ordinary); err != nil {
						return
					}
				case "ordinaryResist":
					if err = sub.Decode(&d.ordinaryResist); err != nil {
						return
					}
				case "skill":
					if err = sub.Decode(&d.skill); err != nil {
						return
					}
				case "skillResist":
					if err = sub.Decode(&d.skillResist); err != nil {
						return
					}
				case "mvp":
					if err = sub.Decode(&d.mvp); err != nil {
						return
					}
				case "mvpResist":
					if err = sub.Decode(&d.mvpResist); err != nil {
						return
					}
				default:
					fmt.Printf("missing decode General.%s: %+v\n", lastAttr, sub)
				}
			}
		}
	}
	return
}
