package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	Attack  int //攻击
	Defence int //防御

	Spike float64 //穿刺

	AttackPer    float64 //攻击%
	DefencePer   float64 //防御%
	Damage       float64 //伤害%
	NearDamage   float64 //近战伤害%
	RemoteDamage float64 //远程伤害%
	Ignore       float64 //忽视防御%
	Resist       float64 //伤害减免%
	NearResist   float64 //近战伤害减免%
	RemoteResist float64 //远程伤害减免%
	Refine       float64 //精炼攻击
	RefineResist float64 //精炼伤害减免%
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.Attack += incr.Attack
		g.Defence += incr.Defence

		g.Spike += incr.Spike

		g.AttackPer += incr.AttackPer
		g.DefencePer += incr.DefencePer
		g.Damage += incr.Damage
		g.NearDamage += incr.NearDamage
		g.RemoteDamage += incr.RemoteDamage
		g.Ignore += incr.Ignore
		g.Resist += incr.Resist
		g.NearResist += incr.NearResist
		g.RemoteResist += incr.RemoteResist
		g.Refine += incr.Refine
		g.RefineResist += incr.RefineResist
	}
}

func (g *Gains) Del(incr *Gains) {
	if incr != nil {
		g.Attack -= incr.Attack
		g.Defence -= incr.Defence

		g.Spike -= incr.Spike

		g.AttackPer -= incr.AttackPer
		g.DefencePer -= incr.DefencePer
		g.Damage -= incr.Damage
		g.NearDamage -= incr.NearDamage
		g.RemoteDamage -= incr.RemoteDamage
		g.Ignore -= incr.Ignore
		g.Resist -= incr.Resist
		g.NearResist -= incr.NearResist
		g.RemoteResist -= incr.RemoteResist
		g.Refine -= incr.Refine
		g.RefineResist -= incr.RefineResist
	}
}

func (g *Gains) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "spike":
					if err = sub.Decode(&g.Spike); err != nil {
						return
					}
				case "attackPer":
					if err = sub.Decode(&g.AttackPer); err != nil {
						return
					}
				case "defencePer":
					if err = sub.Decode(&g.DefencePer); err != nil {
						return
					}
				case "damage":
					if err = sub.Decode(&g.Damage); err != nil {
						return
					}
				case "nearDamage":
					if err = sub.Decode(&g.NearDamage); err != nil {
						return
					}
				case "remoteDamage":
					if err = sub.Decode(&g.RemoteDamage); err != nil {
						return
					}
				case "ignore":
					if err = sub.Decode(&g.Ignore); err != nil {
						return
					}
				case "resist":
					if err = sub.Decode(&g.Resist); err != nil {
						return
					}
				case "nearResist":
					if err = sub.Decode(&g.NearResist); err != nil {
						return
					}
				case "remoteResist":
					if err = sub.Decode(&g.RemoteResist); err != nil {
						return
					}
				case "refine":
					if err = sub.Decode(&g.Refine); err != nil {
						return
					}
				case "refineResist":
					if err = sub.Decode(&g.RefineResist); err != nil {
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
