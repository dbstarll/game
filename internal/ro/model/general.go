package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type General struct {
	Critical             float64 //暴击
	CriticalDamage       float64 //暴伤%
	CriticalResist       float64 //暴击防护
	CriticalDamageResist float64 //爆伤减免%
	Ordinary             float64 //普攻攻击力
	OrdinaryDamage       float64 //普攻伤害加成%
	OrdinaryResist       float64 //普攻伤害减免%
	Skill                float64 //技能伤害加成%
	SkillResist          float64 //技能伤害减免%
	MVP                  float64 //MVP增伤%
	MVPResist            float64 //MVP减伤%
	AttackSpeed          float64 //攻击速度%
	MoveSpeed            float64 //移动速度%
}

func (d *General) Add(incr *General) {
	if incr != nil {
		d.Critical += incr.Critical
		d.CriticalDamage += incr.CriticalDamage
		d.CriticalResist += incr.CriticalResist
		d.CriticalDamageResist += incr.CriticalDamageResist
		d.Ordinary += incr.Ordinary
		d.OrdinaryDamage += incr.OrdinaryDamage
		d.OrdinaryResist += incr.OrdinaryResist
		d.Skill += incr.Skill
		d.SkillResist += incr.SkillResist
		d.MVP += incr.MVP
		d.MVPResist += incr.MVPResist
		d.AttackSpeed += incr.AttackSpeed
		d.MoveSpeed += incr.MoveSpeed
	}
}

func (d *General) Del(incr *General) {
	if incr != nil {
		d.Critical -= incr.Critical
		d.CriticalDamage -= incr.CriticalDamage
		d.CriticalResist -= incr.CriticalResist
		d.CriticalDamageResist -= incr.CriticalDamageResist
		d.Ordinary -= incr.Ordinary
		d.OrdinaryDamage -= incr.OrdinaryDamage
		d.OrdinaryResist -= incr.OrdinaryResist
		d.Skill -= incr.Skill
		d.SkillResist -= incr.SkillResist
		d.MVP -= incr.MVP
		d.MVPResist -= incr.MVPResist
		d.AttackSpeed -= incr.AttackSpeed
		d.MoveSpeed -= incr.MoveSpeed
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
					if err = sub.Decode(&d.Critical); err != nil {
						return
					}
				case "criticalDamage":
					if err = sub.Decode(&d.CriticalDamage); err != nil {
						return
					}
				case "criticalResist":
					if err = sub.Decode(&d.CriticalResist); err != nil {
						return
					}
				case "criticalDamageResist":
					if err = sub.Decode(&d.CriticalDamageResist); err != nil {
						return
					}
				case "ordinary":
					if err = sub.Decode(&d.Ordinary); err != nil {
						return
					}
				case "ordinaryDamage":
					if err = sub.Decode(&d.OrdinaryDamage); err != nil {
						return
					}
				case "ordinaryResist":
					if err = sub.Decode(&d.OrdinaryResist); err != nil {
						return
					}
				case "skill":
					if err = sub.Decode(&d.Skill); err != nil {
						return
					}
				case "skillResist":
					if err = sub.Decode(&d.SkillResist); err != nil {
						return
					}
				case "mvp":
					if err = sub.Decode(&d.MVP); err != nil {
						return
					}
				case "mvpResist":
					if err = sub.Decode(&d.MVPResist); err != nil {
						return
					}
				case "attackSpeed":
					if err = sub.Decode(&d.AttackSpeed); err != nil {
						return
					}
				case "moveSpeed":
					if err = sub.Decode(&d.MoveSpeed); err != nil {
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
