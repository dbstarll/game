package general

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

//通用增益
type General struct {
	Critical             int     `json:"critical"`               //暴击
	CriticalDamage       float64 `json:"critical_damage"`        //暴伤%
	CriticalResist       int     `json:"critical_resist"`        //暴击防护
	CriticalDamageResist float64 `json:"critical_damage_resist"` //爆伤减免%
	Ordinary             int     `json:"ordinary"`               //普攻攻击力
	OrdinaryDamage       float64 `json:"ordinary_damage"`        //普攻伤害加成%
	OrdinaryResist       float64 `json:"ordinary_resist"`        //普攻伤害减免%
	Skill                float64 `json:"skill"`                  //技能伤害加成%
	SkillResist          float64 `json:"skill_resist"`           //技能伤害减免%
	MVP                  float64 `json:"mvp"`                    //MVP增伤%
	MVPResist            float64 `json:"mvp_resist"`             //MVP减伤%
	NoMVP                float64 `json:"no_mvp"`                 //普通魔物增伤%
	NoMVPResist          float64 `json:"no_mvp_resist"`          //普通魔物减伤%
	AttackSpeed          float64 `json:"attack_speed"`           //攻击速度%
	MoveSpeed            float64 `json:"move_speed"`             //移动速度%

	// TODO 以下为增加的属性
	Hp                          int     `json:"hp"`                             //生命上限
	HpPer                       float64 `json:"hp_per"`                         //生命上限%
	Sp                          int     `json:"sp"`                             //魔法上限
	SpPer                       float64 `json:"sp_per"`                         //魔法上限%
	SpCost                      float64 `json:"sp_cost"`                        //SP消耗%
	Hit                         int     `json:"hit"`                            //命中
	Dodge                       int     `json:"dodge"`                          //闪避
	DodgePer                    float64 `json:"dodge_per"`                      //闪避%
	SingPerFixed                float64 `json:"sing_per_fixed"`                 //固定吟唱时间%
	SingPerElasticity           float64 `json:"sing_per_elasticity"`            //可变吟唱时间%
	SingElasticity              float64 `json:"sing_elasticity"`                //可变吟唱时间
	SkillCooling                float64 `json:"skill_cooling"`                  //技能冷却%
	SkillDelay                  float64 `json:"skill_delay"`                    //技能延迟%
	Cure                        float64 `json:"cure"`                           //治疗加成%
	Cured                       float64 `json:"cured"`                          //受治疗加成%
	MagicOrdinaryCriticalRate   float64 `json:"magic_ordinary_critical_rate"`   //法术普攻暴击概率%
	MagicOrdinaryCriticalDamage float64 `json:"magic_ordinary_critical_damage"` //法术普攻暴击伤害%
	BaseExp                     float64 `json:"base_exp"`                       //击杀魔物Base经验%
	JobExp                      float64 `json:"job_exp"`                        //击杀魔物Job经验%
	Player                      float64 `json:"player"`                         //玩家增伤%
	PlayerResist                float64 `json:"player_resist"`                  //玩家减伤%
	Zeny                        int     `json:"zeny"`                           //Zeny
	Bag                         int     `json:"bag"`                            //包包格子
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
		d.NoMVP += incr.NoMVP
		d.NoMVPResist += incr.NoMVPResist
		d.AttackSpeed += incr.AttackSpeed
		d.MoveSpeed += incr.MoveSpeed

		d.Cure += incr.Cure
		d.Cured += incr.Cured
		d.SkillDelay += incr.SkillDelay
		d.SkillCooling += incr.SkillCooling
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
		d.NoMVP -= incr.NoMVP
		d.NoMVPResist -= incr.NoMVPResist
		d.AttackSpeed -= incr.AttackSpeed
		d.MoveSpeed -= incr.MoveSpeed

		d.Cure -= incr.Cure
		d.Cured -= incr.Cured
		d.SkillDelay -= incr.SkillDelay
		d.SkillCooling -= incr.SkillCooling
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
				case "noMvp":
					if err = sub.Decode(&d.NoMVP); err != nil {
						return
					}
				case "noMvpResist":
					if err = sub.Decode(&d.NoMVPResist); err != nil {
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
