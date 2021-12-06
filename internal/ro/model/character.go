package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
	"gopkg.in/yaml.v3"
	"math"
)

type DetectByPanel struct {
	Atk  float64
	Def  float64
	MAtk float64
	MDef float64
}

type Character struct {
	job           job.Job
	nature        nature.Nature
	race          race.Race
	shape         shape.Shape
	level         Level
	quality       Quality
	profits       Profits
	detectByPanel DetectByPanel
}

type CharacterModifier func(character *Character) func()

func NewCharacter(race race.Race, nature nature.Nature, shape shape.Shape, modifiers ...CharacterModifier) *Character {
	c := &Character{
		nature: nature,
		race:   race,
		shape:  shape,
	}
	for _, modifier := range modifiers {
		modifier(c)
	}
	return c
}

func Merge(modifiers ...CharacterModifier) CharacterModifier {
	return func(character *Character) func() {
		size := len(modifiers)
		cancelList := make([]func(), size)
		for idx, modifier := range modifiers {
			cancelList[size-idx-1] = modifier(character)
		}
		return func() {
			for _, cancel := range cancelList {
				cancel()
			}
		}
	}
}

func Job(job job.Job) CharacterModifier {
	return func(character *Character) func() {
		oldJob := character.job
		character.job = job
		return func() {
			character.job = oldJob
		}
	}
}

func AddQuality(quality *Quality) CharacterModifier {
	return func(character *Character) func() {
		character.quality.Add(quality)
		return func() {
			character.quality.Del(quality)
		}
	}
}

func AddLevel(level *Level) CharacterModifier {
	return func(character *Character) func() {
		character.level.Add(level)
		return func() {
			character.level.Del(level)
		}
	}
}

func AddGains(magic bool, gains *Gains) CharacterModifier {
	return func(character *Character) func() {
		character.profits.gains(magic).Add(gains)
		return func() {
			character.profits.gains(magic).Del(gains)
		}
	}
}

func AddGeneral(incr *General) CharacterModifier {
	return func(character *Character) func() {
		character.profits.general.Add(incr)
		return func() {
			character.profits.general.Del(incr)
		}
	}
}

func AddNatureAttack(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddNatureAttack(incr)
		return func() {
			character.profits.DelNatureAttack(incr)
		}
	}
}

func AddRaceDamage(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddRaceDamage(incr)
		return func() {
			character.profits.DelRaceDamage(incr)
		}
	}
}

func AddRaceResist(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddRaceResist(incr)
		return func() {
			character.profits.DelRaceResist(incr)
		}
	}
}

func AddShapeDamage(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddShapeDamage(incr)
		return func() {
			character.profits.DelShapeDamage(incr)
		}
	}
}

func AddShapeResist(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddShapeResist(incr)
		return func() {
			character.profits.DelShapeResist(incr)
		}
	}
}

func AddNatureDamage(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddNatureDamage(incr)
		return func() {
			character.profits.DelNatureDamage(incr)
		}
	}
}

func AddNatureResist(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.profits.AddNatureResist(incr)
		return func() {
			character.profits.DelNatureResist(incr)
		}
	}
}

func DetectAttackByPanel(remote bool, expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) func() {
		character.detectAttackByPanel(false, remote, expectPhysicalPanel)
		character.detectAttackByPanel(true, remote, expectMagicalPanel)
		return func() {}
	}
}

func DetectDefenceByPanel(expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) func() {
		character.detectDefenceByPanel(false, expectPhysicalPanel)
		character.detectDefenceByPanel(true, expectMagicalPanel)
		return func() {}
	}
}

//素质攻击
func (c *Character) QualityAttack(magic, remote bool) int {
	return c.quality.Attack(magic, remote)
}

//装备攻击
func (c *Character) EquipmentAttack(magic bool) (atk int) {
	atk = c.profits.gains(magic).Attack
	if !magic {
		//装备物理攻击 = (装备，强化，附魔，卡片，头饰，祈祷，buff等合计)+ BaseLvAtkRate*人物等级
		atk += c.job.BaseLvAtkRate() * c.level.Base
	}
	return
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) int {
	return c.QualityAttack(magic, remote) + c.EquipmentAttack(magic)
}

//面板攻击 = 攻击 * (1 + 攻击%)
func (c *Character) PanelAttack(magic, remote bool) float64 {
	return float64(c.Attack(magic, remote)) * (1 + c.profits.gains(magic).AttackPer/100)
}

//素质防御
func (c *Character) QualityDefence(magic bool) int {
	return c.quality.Defence(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) int {
	return c.profits.gains(magic).Defence
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) int {
	return c.QualityDefence(magic) + c.EquipmentDefence(magic)
}

//面板防御 = 防御 * (1 + 防御%)
func (c *Character) PanelDefence(magic bool) float64 {
	return float64(c.Defence(magic)) * (1 + c.profits.gains(magic).DefencePer/100)
}

func (c *Character) AttackWithWeapon(weapon weapon.Weapon) *Attack {
	return &Attack{
		weapon: weapon,
		magic:  weapon.IsMagic(c.job),
		remote: weapon.IsRemote(c.job),
		nature: nature.Neutral,
	}
}

func (c *Character) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	rate = c.profits.SkillDamageRate(target, magic, skillNature)
	rate *= 1 + c.profits.shapeDamage[target.shape]/100 - target.profits.shapeResist[c.shape]/100 //*(1+体型增伤%-体型减伤%)
	//TODO 种族减伤待验证，是加算还是乘算
	//rate *= 1 + c.profits.raceDamage[target.race]/100 - target.profits.raceResist[c.race]/100 //*(1+种族增伤%-种族减伤%)
	rate *= 1 + c.profits.raceDamage[target.race]/100 //*(1+种族增伤%)
	rate *= 1 - target.profits.raceResist[c.race]/100 //*(1-种族减伤%)
	rate *= skillNature.Restraint(target.nature)      //*属性克制
	return
}

//基础伤害
func (c *Character) baseDamage(target *Character, attack *Attack) (damage float64) {
	gains, targetGains := c.profits.gains(attack.magic), target.profits.gains(attack.magic)
	damage = c.finalAttack(target, attack) //最终物攻/最终魔攻
	if attack.magic {
		//TODO *魔防乘数
		damage *= 1 - targetGains.Resist/100 //*(1-魔伤减免)
		damage += gains.Refine               //+精炼魔攻
		//TODO *技能倍率
		damage *= 1 + gains.Damage                                   //*(1+魔伤加成)
		damage *= attack.nature.Restraint(target.nature)             //*属性克制
		damage *= 1 - target.profits.natureResist[attack.nature]/100 //*(1-属性减伤)
		damage -= float64(target.QualityDefence(attack.magic))       //-素质魔防
		damage -= float64(target.QualityDefence(!attack.magic))      //-素质物防/2
	} else if attack.critical { //普攻暴击
		damage *= 1 - targetGains.Resist/100                 //*(1-物伤减免)
		damage += gains.Refine                               //+精炼物攻
		damage *= 1.5 + c.profits.general.CriticalDamage/100 //*(1+暴伤%)
		damage *= 1 + gains.Damage/100                       //*(1+物伤加成)
	} else { // 普攻未暴击或技能
		//TODO *物防乘数
		damage *= 1 - targetGains.Resist/100 //*(1-物伤减免)
		damage += gains.Refine               //+精炼物攻
		//TODO *技能倍率
		damage -= float64(target.QualityDefence(attack.magic)) //-素质物防
		damage *= 1 + gains.Damage/100                         //*(1+物伤加成)
	}
	return
}

//最终物攻/最终魔攻
func (c *Character) finalAttack(target *Character, attack *Attack) (damage float64) {
	damage = float64(c.EquipmentAttack(attack.magic)) //装备攻击
	if !attack.skill {
		damage += float64(c.quality.OrdinaryAttack(attack.magic, attack.remote)) //素质普攻攻击力
		damage += c.profits.general.Ordinary                                     //普攻攻击力
	}
	damage *= 1 + c.profits.gains(attack.magic).AttackPer/100 //*(1+攻击%)

	if attack.magic {
		damage += float64(c.quality.Int) * c.profits.gains(attack.magic).AttackPer / 100 //+智力*魔法攻击%
	} else {
		damage *= attack.weapon.Restraint(target.shape)                                                 //*武器体型修正
		damage *= 1 + c.profits.shapeDamage[target.shape]/100 - target.profits.shapeResist[c.shape]/100 //*(1+体型增伤%-体型减伤%)
		damage *= attack.nature.Restraint(target.nature)                                                //*属性克制
		damage *= 1 + c.profits.natureDamage[target.nature]/100                                         //*(1+属性魔物增伤%)
		damage *= 1 - target.profits.natureResist[attack.nature]/100                                    //*(1-属性减伤%)
	}
	damage += float64(c.QualityAttack(attack.magic, attack.remote))                             //+素质攻击
	damage *= 1 + c.profits.raceDamage[target.race]/100 - target.profits.raceResist[c.race]/100 //*(1+种族增伤%-种族减伤%)
	return
}

func (c *Character) detectAttackByPanel(magic, remote bool, expect float64) (optimumAttack int, optimumPanel float64) {
	gains := c.profits.gains(magic)
	for min, max, current := 0, 100000, gains.Attack; ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		gains.Attack = current
		actual := c.PanelAttack(magic, remote)

		if actual >= expect && math.Abs(actual-expect) < math.Abs(optimumPanel-expect) {
			optimumAttack, optimumPanel = current, actual
		}
		if actual > expect {
			if max == current && max-min == 1 {
				break
			} else {
				max = current
			}
		} else if min == current && max-min == 1 {
			break
		} else {
			min = current
		}
	}
	gains.Attack = optimumAttack
	fmt.Printf("detectAttackByPanel[magic=%t]: optimumAttack=%d, optimumPanel=%f\n", magic, optimumAttack, optimumPanel)
	return
}

func (c *Character) detectDefenceByPanel(magic bool, expect float64) (optimumDefence int, optimumPanel float64) {
	gains := c.profits.gains(magic)
	for min, max, current := 0, 100000, gains.Defence; ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		gains.Defence = current
		actual := c.PanelDefence(magic)

		if actual >= expect && math.Abs(actual-expect) < math.Abs(optimumPanel-expect) {
			optimumDefence, optimumPanel = current, actual
		}
		if actual > expect {
			if max == current && max-min == 1 {
				break
			} else {
				max = current
			}
		} else if min == current && max-min == 1 {
			break
		} else {
			min = current
		}
	}
	gains.Defence = optimumDefence
	fmt.Printf("detectDefenceByPanel[magic=%t]: optimumDefence=%d, optimumPanel=%f\n", magic, optimumDefence, optimumPanel)
	return
}

func (c *Character) UnmarshalYAML(value *yaml.Node) (err error) {
	if value.Kind == yaml.MappingNode {
		var lastAttr string
		for idx, sub := range value.Content {
			if sub.Kind == yaml.ScalarNode && idx%2 == 0 {
				lastAttr = sub.Value
			} else {
				switch lastAttr {
				case "job":
					if err = sub.Decode(&c.job); err != nil {
						return
					}
				case "nature":
					if err = sub.Decode(&c.nature); err != nil {
						return
					}
				case "race":
					if err = sub.Decode(&c.race); err != nil {
						return
					}
				case "shape":
					if err = sub.Decode(&c.shape); err != nil {
						return
					}
				case "level":
					if err = sub.Decode(&c.level); err != nil {
						return
					}
				case "quality":
					if err = sub.Decode(&c.quality); err != nil {
						return
					}
				case "profits":
					if err = sub.Decode(&c.profits); err != nil {
						return
					}
				case "detectByPanel":
					if err = sub.Decode(&c.detectByPanel); err != nil {
						return
					}
				default:
					fmt.Printf("missing decode Character.%s: %+v\n", lastAttr, sub)
				}
			}
		}
	}
	return
}
