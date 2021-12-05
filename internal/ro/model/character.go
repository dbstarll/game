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

type CharacterModifier func(character *Character)

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
	return func(character *Character) {
		for _, modifier := range modifiers {
			modifier(character)
		}
	}
}

func Job(job job.Job) CharacterModifier {
	return func(character *Character) {
		character.job = job
	}
}

func AddQuality(quality *Quality) CharacterModifier {
	return func(character *Character) {
		character.quality.Add(quality)
	}
}

func AddLevel(level *Level) CharacterModifier {
	return func(character *Character) {
		character.level.Add(level)
	}
}

func AddGains(magic bool, gains *Gains) CharacterModifier {
	return func(character *Character) {
		character.profits.AddGains(magic, gains)
	}
}

func AddDamage(incr *Damage) CharacterModifier {
	return func(character *Character) {
		character.profits.AddDamage(incr)
	}
}

func AddNatureAttack(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureAttack(incr)
	}
}

func AddRaceDamage(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddRaceDamage(incr)
	}
}

func AddRaceResist(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddRaceResist(incr)
	}
}

func AddShapeDamage(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddShapeDamage(incr)
	}
}

func AddShapeResist(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddShapeResist(incr)
	}
}

func AddNatureDamage(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureDamage(incr)
	}
}

func AddNatureResist(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) {
		character.profits.AddNatureResist(incr)
	}
}

func DetectAttackByPanel(remote bool, expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) {
		character.detectAttackByPanel(false, remote, expectPhysicalPanel)
		character.detectAttackByPanel(true, remote, expectMagicalPanel)
	}
}

func DetectDefenceByPanel(expectPhysicalPanel, expectMagicalPanel float64) CharacterModifier {
	return func(character *Character) {
		character.detectDefenceByPanel(false, expectPhysicalPanel)
		character.detectDefenceByPanel(true, expectMagicalPanel)
	}
}

//素质攻击
func (c *Character) QualityAttack(magic, remote bool) int {
	return c.quality.Attack(magic, remote)
}

//装备攻击
func (c *Character) EquipmentAttack(magic bool) int {
	if magic {
		return c.profits.Attack(magic)
	} else {
		//装备物理攻击 = (装备，强化，附魔，卡片，头饰，祈祷，buff等合计)+ BaseLvAtkRate*人物等级
		return c.profits.Attack(magic) + c.job.BaseLvAtkRate()*c.level.Base
	}
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) int {
	return c.QualityAttack(magic, remote) + c.EquipmentAttack(magic)
}

//面板攻击 = 攻击 * (1 + 攻击%)
func (c *Character) PanelAttack(magic, remote bool) float64 {
	return float64(c.Attack(magic, remote)) * (1 + c.profits.AttackPer(magic)/100)
}

//素质防御
func (c *Character) QualityDefence(magic bool) int {
	return c.quality.Defence(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) int {
	return c.profits.Defence(magic)
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) int {
	return c.QualityDefence(magic) + c.EquipmentDefence(magic)
}

//面板防御 = 防御 * (1 + 防御%)
func (c *Character) PanelDefence(magic bool) float64 {
	return float64(c.Defence(magic)) * (1 + c.profits.DefencePer(magic)/100)
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
	damage = c.finalAttack(target, attack) //最终物攻/最终魔攻
	if attack.magic {
		//TODO *魔防乘数
		damage *= 1 - target.profits.Resist(attack.magic)/100 //*(1-魔伤减免)
		damage += c.profits.Refine(attack.magic)              //+精炼魔攻
		//TODO *技能倍率
		damage *= 1 + c.profits.Damage(attack.magic)                 //*(1+魔伤加成)
		damage *= attack.nature.Restraint(target.nature)             //*属性克制
		damage *= 1 - target.profits.natureResist[attack.nature]/100 //*(1-属性减伤)
		damage -= float64(target.QualityDefence(attack.magic))       //-素质魔防
		damage -= float64(target.QualityDefence(!attack.magic))      //-素质物防/2
	} else if attack.critical { //普攻暴击
		damage *= 1 - target.profits.Resist(attack.magic)/100   //*(1-物伤减免)
		damage += c.profits.Refine(attack.magic)                //+精炼物攻
		damage *= 1.5 + c.profits.CriticalPer(attack.magic)/100 //*(1+暴伤%)
		damage *= 1 + c.profits.Damage(attack.magic)/100        //*(1+物伤加成)
	} else { // 普攻未暴击或技能
		//TODO *物防乘数
		damage *= 1 - target.profits.Resist(attack.magic)/100 //*(1-物伤减免)
		damage += c.profits.Refine(attack.magic)              //+精炼物攻
		//TODO *技能倍率
		damage -= float64(target.QualityDefence(attack.magic)) //-素质物防
		damage *= 1 + c.profits.Damage(attack.magic)/100       //*(1+物伤加成)
	}
	return
}

//最终物攻/最终魔攻
func (c *Character) finalAttack(target *Character, attack *Attack) (damage float64) {
	damage = float64(c.EquipmentAttack(attack.magic)) //装备攻击
	if !attack.skill {
		damage += float64(c.quality.GeneralAttack(attack.magic, attack.remote)) //普攻攻击力
	}
	damage *= 1 + c.profits.AttackPer(attack.magic)/100 //*(1+攻击%)

	if attack.magic {
		damage += float64(c.quality.Int) * c.profits.AttackPer(attack.magic) / 100 //+智力*魔法攻击%
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
	for min, max, current := 0, 100000, c.profits.Attack(magic); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		c.profits.setAttack(magic, current)
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
	c.profits.setAttack(magic, optimumAttack)
	fmt.Printf("detectAttackByPanel[magic=%t]: optimumAttack=%d, optimumPanel=%f\n", magic, optimumAttack, optimumPanel)
	return
}

func (c *Character) detectDefenceByPanel(magic bool, expect float64) (optimumDefence int, optimumPanel float64) {
	for min, max, current := 0, 100000, c.profits.Defence(magic); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		c.profits.setDefence(magic, current)
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
	c.profits.setDefence(magic, optimumDefence)
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
