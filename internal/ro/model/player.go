package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
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
			player.types = types.Player
			player.nature = nature.Neutral
			player.race = race.Human
			player.shape = shape.Medium
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

func LoadMonsterFromYaml(name string) (*Character, error) {
	yamlFile := fmt.Sprintf("configs/monster/%s.yaml", name)
	if data, err := ioutil.ReadFile(yamlFile); err != nil {
		return nil, errors.WithStack(err)
	} else {
		monster := &Character{}
		if err := yaml.Unmarshal(data, monster); err != nil {
			return nil, errors.WithStack(err)
		} else {
			return monster, nil
		}
	}
}

func (p *Player) SkillEarth() (damage float64) {
	damage = float64(p.Quality.Vit*p.Quality.Vit) *
		p.PanelDefence(false) / 10000 *
		9.6 //基础技能倍率
	damage *= 1 + 10.0/100 //*(1+守护之盾技能增伤%)
	damage *= 1 + 16.2/100 //*(1+铁蹄直驱符文增伤%)
	return
}

func (p *Player) SkillEarth2() (damage float64) {
	damage = float64(p.Quality.Vit*p.Quality.Vit) *
		p.PanelDefence(false) / 10000 *
		(9.6 + 0.46*(p.Profits.General.MoveSpeed+100)/10) //基础技能倍率 + 冲锋领袖符文,每提高10%%移动速度，【大地猛击】技能倍率 +10% ~ 50%
	damage *= 1 + 10.0/100 //*(1+守护之盾技能增伤%)
	damage *= 1 + 16.2/100 //*(1+铁蹄直驱符文增伤%)
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
