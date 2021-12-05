package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Player struct {
	*Character
}

func LoadPlayerFromYaml(name string, remote bool) (*Player, error) {
	yamlFile := fmt.Sprintf("configs/player/%s.yaml", name)
	if data, err := ioutil.ReadFile(yamlFile); err != nil {
		return nil, errors.WithStack(err)
	} else {
		player := &Player{}
		if err := yaml.Unmarshal(data, player); err != nil {
			return nil, errors.WithStack(err)
		} else {
			player.Character.nature = nature.Neutral
			player.Character.race = race.Human
			player.Character.shape = shape.Medium
			if player.detectByPanel.Atk > 0 {
				player.detectAttackByPanel(false, remote, player.detectByPanel.Atk)
			}
			if player.detectByPanel.MAtk > 0 {
				player.detectAttackByPanel(true, remote, player.detectByPanel.MAtk)
			}
			if player.detectByPanel.Def > 0 {
				player.detectDefenceByPanel(false, player.detectByPanel.Def)
			}
			if player.detectByPanel.MDef > 0 {
				player.detectDefenceByPanel(true, player.detectByPanel.MDef)
			}
			return player, nil
		}
	}
}

func NewPlayer(job job.Job, modifiers ...CharacterModifier) *Player {
	return &Player{
		Character: NewCharacter(race.Human, nature.Neutral, shape.Medium, append([]CharacterModifier{Job(job)}, modifiers...)...),
	}
}

func (p *Player) SkillDamageRate(target *Monster, magic bool, skillNature nature.Nature) (rate float64) {
	rate = p.Character.SkillDamageRate(target.Character, magic, skillNature)
	if target.types.IsBoss() {
		rate *= 1 + p.profits.damage.MVP/100 //*(1+MVP增伤%)
	}
	return
}

func (p *Player) SkillEarth() (damage float64) {
	damage = float64(p.quality.Vit*p.quality.Vit) *
		p.PanelDefence(false) / 10000 *
		9.6 //基础技能倍率
	damage *= 1 + 10.0/100 //*(1+守护之盾技能增伤%)
	damage *= 1 + 16.2/100 //*(1+铁蹄直驱符文增伤%)
	return
}

func (p *Player) GeneralAttack(target *Monster, attack *Attack) (damage float64) {
	magic, remote, nature := attack.magic, attack.remote, attack.nature
	//最终物攻
	damage = float64(p.EquipmentAttack(magic) + p.quality.GeneralAttack(magic, remote)) //装备攻击
	damage *= 1 + p.Character.profits.AttackPer(magic)/100                              //*(1+攻击%)
	// TODO *武器体型修正
	damage *= 1 + p.Character.profits.shapeDamage[target.shape]/100 - target.profits.shapeResist[p.shape]/100 //*(1+体型增伤%-体型减伤%)
	damage *= nature.Restraint(target.nature)                                                                 //*属性克制
	damage *= 1 + p.Character.profits.natureDamage[target.nature]/100                                         //*(1+属性魔物增伤%)
	damage *= 1 - target.profits.natureResist[nature]/100                                                     //*(1-属性减伤%)
	damage += float64(p.QualityAttack(magic, remote))                                                         //+素质攻击
	damage *= 1 + p.Character.profits.raceDamage[target.race]/100 - target.profits.raceResist[p.race]/100     //*(1+种族增伤%-种族减伤%)

	//基础伤害
	//普攻未暴击: 基础伤害 = ((最终物攻 * 物防乘数 *物理最终减伤值+精炼物攻)*技能倍率 -素质物防)
	//普攻暴击: 基础伤害 = (最终物攻*物理最终减伤值+精炼物攻)*(1+暴伤%)
	//最终物防 = (物理防御-素质物理防御)*(1+物理防御%-忽视物防%)
	//物防乘数 = (4000+最终物防)/(4000+最终物防*10)
	damage += p.Character.profits.Refine(magic) //+精炼物攻

	damage *= 1 + p.Character.profits.Damage(magic)/100                                                  //*(1+伤害加成%)
	damage *= 1 + p.Character.profits.natureAttack[nature]/100 - target.profits.natureResist[nature]/100 //*(1+属性攻击%-属性减伤%)
	if target.types.IsBoss() {
		damage *= 1 + p.profits.damage.MVP/100 //*(1+MVP增伤%)
	}

	return
}

//最终伤害
func (p *Player) FinalDamage(target *Monster, attack *Attack) (damage float64) {
	//最终伤害 = 基础伤害 * (1+元素加伤) * (1+MVP增伤%) * 状态加伤 * (1+真实伤害)
	damage = p.baseDamage(target.Character, attack)         //基础伤害
	damage *= 1 + p.profits.natureAttack[attack.nature]/100 //*(1+属性攻击%)
	if target.types.IsBoss() {
		damage *= 1 + p.profits.damage.MVP/100 //*(1+MVP增伤%)
	}
	// TODO *状态加伤
	// TODO *(1+真实伤害)
	return
}

func (p *Player) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "character":
					p.Character = &Character{}
					if err = sub.Decode(p.Character); err != nil {
						return
					}
				default:
					fmt.Printf("missing decode Player.%s: %+v\n", lastAttr, sub)
				}
			}
		}
	}
	return
}
