package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

//装备，强化，附魔，卡片，头饰，祈祷，buff等合计
type Gains struct {
	attack  int //攻击
	defence int //防御

	spike float64 //穿刺

	attackPer    float64 //攻击%
	defencePer   float64 //防御%
	damage       float64 //伤害%
	ignore       float64 //忽视防御%
	resist       float64 //伤害减免%
	nearResist   float64 //近战伤害减免%
	remoteResist float64 //远程伤害减免%
	refine       float64 //精炼攻击
	refineResist float64 //精炼伤害减免%
}

func (g *Gains) Add(incr *Gains) {
	if incr != nil {
		g.attack += incr.attack
		g.defence += incr.defence

		g.spike += incr.spike

		g.attackPer += incr.attackPer
		g.defencePer += incr.defencePer
		g.damage += incr.damage
		g.ignore += incr.ignore
		g.resist += incr.resist
		g.nearResist += incr.nearResist
		g.remoteResist += incr.remoteResist
		g.refine += incr.refine
		g.refineResist += incr.refineResist
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
					if err = sub.Decode(&g.spike); err != nil {
						return
					}
				case "attackPer":
					if err = sub.Decode(&g.attackPer); err != nil {
						return
					}
				case "defencePer":
					if err = sub.Decode(&g.defencePer); err != nil {
						return
					}
				case "damage":
					if err = sub.Decode(&g.damage); err != nil {
						return
					}
				case "ignore":
					if err = sub.Decode(&g.ignore); err != nil {
						return
					}
				case "resist":
					if err = sub.Decode(&g.resist); err != nil {
						return
					}
				case "nearResist":
					if err = sub.Decode(&g.nearResist); err != nil {
						return
					}
				case "remoteResist":
					if err = sub.Decode(&g.remoteResist); err != nil {
						return
					}
				case "refine":
					if err = sub.Decode(&g.refine); err != nil {
						return
					}
				case "refineResist":
					if err = sub.Decode(&g.refineResist); err != nil {
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
