package model

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/abnormal"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/race"
	"github.com/dbstarll/game/internal/ro/dimension/shape"
	"github.com/dbstarll/game/internal/ro/dimension/types"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
	"github.com/dbstarll/game/internal/ro/model/attack"
	"github.com/dbstarll/game/internal/ro/model/general"
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
	types         types.Types
	job           job.Job
	nature        nature.Nature
	race          race.Race
	shape         shape.Shape
	level         Level
	Quality       Quality
	Profits       Profits
	detectByPanel DetectByPanel
}

type CharacterModifier func(character *Character) func()

func NewCharacter(types types.Types, race race.Race, nature nature.Nature, shape shape.Shape, modifiers ...CharacterModifier) *Character {
	c := &Character{
		types:  types,
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

func Rate(modifier CharacterModifier, rate func(character *Character) int) CharacterModifier {
	return func(character *Character) func() {
		return func() {

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
		character.Quality.Add(quality)
		return func() {
			character.Quality.Del(quality)
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
		character.Profits.Gains(magic).Add(gains)
		return func() {
			character.Profits.Gains(magic).Del(gains)
		}
	}
}

func AddGeneral(incr *general.General) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.General.Add(incr)
		return func() {
			character.Profits.General.Del(incr)
		}
	}
}

func AddNatureAttack(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddNatureAttack(incr)
		return func() {
			character.Profits.DelNatureAttack(incr)
		}
	}
}

func AddRaceDamage(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddRaceDamage(incr)
		return func() {
			character.Profits.DelRaceDamage(incr)
		}
	}
}

func AddRaceResist(incr *map[race.Race]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddRaceResist(incr)
		return func() {
			character.Profits.DelRaceResist(incr)
		}
	}
}

func AddShapeDamage(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddShapeDamage(incr)
		return func() {
			character.Profits.DelShapeDamage(incr)
		}
	}
}

func AddShapeResist(incr *map[shape.Shape]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddShapeResist(incr)
		return func() {
			character.Profits.DelShapeResist(incr)
		}
	}
}

func AddNatureDamage(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddNatureDamage(incr)
		return func() {
			character.Profits.DelNatureDamage(incr)
		}
	}
}

func AddNatureResist(incr *map[nature.Nature]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddNatureResist(incr)
		return func() {
			character.Profits.DelNatureResist(incr)
		}
	}
}

func AddAbnormalResist(incr *map[abnormal.Abnormal]float64) CharacterModifier {
	return func(character *Character) func() {
		character.Profits.AddAbnormalResist(incr)
		return func() {
			character.Profits.DelAbnormalResist(incr)
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

func (c *Character) Job() job.Job {
	return c.job
}

func (c *Character) Apply(modifiers ...CharacterModifier) func() {
	return Merge(modifiers...)(c)
}

//素质攻击
func (c *Character) QualityAttack(magic, remote bool) int {
	return c.Quality.Attack(magic, remote)
}

//装备攻击
func (c *Character) EquipmentAttack(magic bool) (atk float64) {
	atk = c.Profits.Gains(magic).Attack
	if !magic {
		//装备物理攻击 = (装备，强化，附魔，卡片，头饰，祈祷，buff等合计)+ BaseLvAtkRate*人物等级
		atk += float64(c.job.BaseLvAtkRate() * c.level.Base)
	}
	return
}

//攻击 = 素质攻击 + 装备攻击
func (c *Character) Attack(magic, remote bool) float64 {
	return float64(c.QualityAttack(magic, remote)) + c.EquipmentAttack(magic)
}

//面板攻击 = 攻击 * (1 + 攻击%)
func (c *Character) PanelAttack(magic, remote bool) float64 {
	return c.Attack(magic, remote) * (1 + c.Profits.Gains(magic).AttackPer/100)
}

//素质防御
func (c *Character) QualityDefence(magic bool) int {
	return c.Quality.Defence(magic)
}

//装备防御
func (c *Character) EquipmentDefence(magic bool) float64 {
	return c.Profits.Gains(magic).Defence
}

//防御 = 素质防御 + 装备防御
func (c *Character) Defence(magic bool) float64 {
	return float64(c.QualityDefence(magic)) + c.EquipmentDefence(magic)
}

//面板防御 = 防御 * (1 + 防御%)
func (c *Character) PanelDefence(magic bool) float64 {
	return c.Defence(magic) * (1 + c.Profits.Gains(magic).DefencePer/100)
}

func (c *Character) AttackWithWeapon(weapon weapon.Weapon) *attack.Attack {
	return attack.UseWeapon(c.job, weapon)
}

func (c *Character) SkillDamageRate(target *Character, magic bool, skillNature nature.Nature) (rate float64) {
	gains, targetGains := c.Profits.Gains(magic), target.Profits.Gains(magic)

	//finalAttack
	rate = 1 + c.Profits.shapeDamage[target.shape]/100 - target.Profits.shapeResist[c.shape]/100 //*(1+体型增伤%-体型减伤%)
	rate *= skillNature.Restraint(target.nature)                                                 //*属性克制
	rate *= 1 + c.Profits.natureAttack[skillNature]/100                                          //*(1+属性攻击%)
	rate *= 1 + c.Profits.natureDamage[target.nature]/100                                        //*(1+属性魔物增伤%)
	rate *= 1 - target.Profits.natureResist[skillNature]/100                                     //*(1-属性减伤%)
	//rate *= 1 + c.Profits.raceDamage[target.race]/100 - target.Profits.raceResist[c.race]/100    //*(1+种族增伤%-种族减伤%)
	rate *= 1 + c.Profits.raceDamage[target.race]/100 //*(1+种族增伤%)
	rate *= 1 - target.Profits.raceResist[c.race]/100 //*(1-种族减伤%)
	if target.types.IsPlayer() {
		//TODO *(1+玩家增伤%)
	} else if target.types.IsBoss() {
		rate *= 1 + c.Profits.General.MVP/100 //*(1+MVP增伤%)
	} else {
		rate *= 1 + c.Profits.General.NoMVP/100 //*(1+普通魔物增伤%)
	}

	//baseDamage
	rate *= 1 + c.Profits.weaponSpikes()/100 + gains.Spike/100 - targetGains.Resist/100 //*(1+装备穿刺%+穿刺%-伤害减免%)
	rate *= 1 + gains.Damage/100 + gains.NearDamage/100                                 //*(1+伤害加成%+近战伤害%)
	rate *= 1 + c.Profits.General.Skill/100                                             //*(1+技能伤害加成%)

	//finalDamage

	return
}

//最终伤害
func (c *Character) FinalDamage(target *Character, attack *attack.Attack) (damage float64) {
	//最终伤害 = 基础伤害 * (1+元素加伤) * 状态加伤 * (1+真实伤害)
	damage = c.baseDamage(target, attack) //基础伤害
	// TODO *状态加伤
	// TODO *(1+真实伤害)
	if attack.GetWeapon() == weapon.Rifle && attack.IsOrdinary() {
		damage *= 2 //来复枪伤害翻倍
	}
	return
}

//基础伤害
func (c *Character) baseDamage(target *Character, attack *attack.Attack) (damage float64) {
	gains, targetGains := c.Profits.Gains(attack.IsMagic()), target.Profits.Gains(attack.IsMagic())
	damage = c.finalAttack(target, attack) //最终物攻/最终魔攻
	// 物理最终减伤系数=1-(1+装备穿刺+物理穿刺-物理减免)*物理技能伤害减免
	damage *= 1 + c.Profits.weaponSpikes()/100 + gains.Spike/100 - targetGains.Resist/100 //*(1+装备穿刺%+穿刺%-伤害减免%)
	if attack.IsMagic() {
		damage *= target.defenceMultiplier(c, attack)                     //*魔防乘数
		damage += gains.Refine                                            //+精炼魔攻
		damage *= attack.SkillRate()                                      //*技能倍率
		damage *= 1 + gains.Damage                                        //*(1+魔伤加成)
		damage *= attack.GetNature().Restraint(target.nature)             //*属性克制
		damage *= 1 - target.Profits.natureResist[attack.GetNature()]/100 //*(1-属性减伤)
		damage -= float64(target.QualityDefence(true))                    //-素质魔防
		damage -= float64(target.QualityDefence(false))                   //-素质物防/2
	} else {
		if attack.IsCritical() && attack.IsOrdinary() { //普攻暴击
			damage += gains.Refine
			damage *= 1.5 + c.Profits.General.CriticalDamage/100 - target.Profits.General.CriticalDamageResist/100 //*(1+暴伤%-爆伤减免%)
		} else { // 普攻未暴击或技能
			damage *= target.defenceMultiplier(c, attack)   //*物防乘数
			damage += gains.Refine                          //+精炼物攻
			damage *= attack.SkillRate()                    //*技能倍率
			damage -= float64(target.QualityDefence(false)) //-素质物防
		}
		if attack.IsRemote() {
			damage *= 1 + gains.Damage/100 + gains.RemoteDamage/100 - targetGains.RemoteResist/100 //*(1+物伤加成%+远程物理伤害%-远程伤害减免%)
		} else {
			damage *= 1 + gains.Damage/100 + gains.NearDamage/100 - targetGains.NearResist/100 //*(1+物伤加成%+近战物理伤害%-近战伤害减免%)
		}
		if attack.IsOrdinary() {
			damage *= 1 + c.Profits.General.OrdinaryDamage/100 - target.Profits.General.OrdinaryResist/100 //*(1+普攻伤害加成%-普攻伤害减免%)
		} else {
			damage *= 1 + c.Profits.General.Skill/100 - target.Profits.General.SkillResist/100 //*(1+技能伤害加成%-技能伤害减免%)
		}
	}

	return
}

//最终物攻/最终魔攻
func (c *Character) finalAttack(target *Character, attack *attack.Attack) (damage float64) {
	damage = c.EquipmentAttack(attack.IsMagic()) //装备攻击
	if attack.IsOrdinary() {
		if c.job >= job.Archer && c.job <= job.Hunter4 {
			damage += float64(c.Quality.OrdinaryAttack(attack.IsMagic(), attack.IsRemote())) //素质普攻攻击力
			damage += float64(c.Profits.General.Ordinary)                                    //TODO 这里存疑普攻攻击力
		} else if attack.IsCritical() {
			damage += float64(c.Quality.OrdinaryAttack(attack.IsMagic(), attack.IsRemote())) //素质普攻攻击力
		}
		if c.job >= job.Hunter2 && c.job <= job.Hunter4 {
			damage += float64(200 + c.Quality.Dex*5) //猎人进阶二转技能：元素箭矢20级被动效果
		}
	}
	damage *= 1 + c.Profits.Gains(attack.IsMagic()).AttackPer/100 //*(1+攻击%)

	if attack.IsMagic() {
		damage += float64(c.Quality.Int) * c.Profits.Gains(true).AttackPer / 100 //+智力*魔法攻击%
	} else {
		damage *= attack.GetWeapon().Restraint(target.shape)                                            //*武器体型修正
		damage *= 1 + c.Profits.shapeDamage[target.shape]/100 - target.Profits.shapeResist[c.shape]/100 //*(1+体型增伤%-体型减伤%)
		damage *= attack.GetNature().Restraint(target.nature)                                           //*属性克制
		damage *= 1 + c.Profits.natureAttack[attack.GetNature()]/100                                    //*(1+属性攻击%)
		damage *= 1 + c.Profits.natureDamage[target.nature]/100                                         //*(1+属性魔物增伤%)
		damage *= 1 - target.Profits.natureResist[attack.GetNature()]/100                               //*(1-属性减伤%)
	}
	damage += float64(c.QualityAttack(attack.IsMagic(), attack.IsRemote())) //+素质攻击
	//damage *= 1 + c.Profits.raceDamage[target.race]/100 - target.Profits.raceResist[c.race]/100 //*(1+种族增伤%-种族减伤%)
	damage *= 1 + c.Profits.raceDamage[target.race]/100 //*(1+种族增伤%)
	damage *= 1 - target.Profits.raceResist[c.race]/100 //*(1-种族减伤%)
	if target.types.IsPlayer() {
		//TODO *(1+玩家增伤%)
	} else if target.types.IsBoss() {
		damage *= 1 + c.Profits.General.MVP/100 //*(1+MVP增伤%)
	} else {
		damage *= 1 + c.Profits.General.NoMVP/100 //*(1+普通魔物增伤%)
	}
	return
}

//最终物防/最终魔防
func (c *Character) finalDefence(target *Character, attack *attack.Attack) (defence float64) {
	gain, targetGain := c.Profits.Gains(attack.IsMagic()), target.Profits.Gains(attack.IsMagic())
	defence = c.EquipmentDefence(attack.IsMagic()) //装备防御
	if attack.IsMagic() {
		//TODO 最终魔防
	} else {
		defence *= 1 + gain.DefencePer/100 - targetGain.Ignore/100 //*(1+物理防御%-忽视物防%)
	}
	return
}

//物防乘数/魔防乘数
func (c *Character) defenceMultiplier(target *Character, attack *attack.Attack) (rate float64) {
	if finalDefence := c.finalDefence(target, attack); finalDefence <= 0 {
		rate = 1
	} else if attack.IsMagic() {
		//TODO 魔防乘数
		rate = 1
	} else {
		//物防乘数 = (4000+最终物防)/(4000+最终物防*10)
		rate = (4000 + finalDefence) / (4000 + finalDefence*10)
	}
	return
}

func (c *Character) detectAttackByPanel(magic, remote bool, expect float64) (optimumAttack int, optimumPanel float64) {
	gains := c.Profits.Gains(magic)
	for min, max, current := 0, 100000, int(gains.Attack); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		gains.Attack = float64(current)
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
	gains.Attack = float64(optimumAttack)
	fmt.Printf("detectAttackByPanel[magic=%t]: optimumAttack=%d, optimumPanel=%f\n", magic, optimumAttack, optimumPanel)
	return
}

func (c *Character) detectDefenceByPanel(magic bool, expect float64) (optimumDefence int, optimumPanel float64) {
	gains := c.Profits.Gains(magic)
	for min, max, current := 0, 100000, int(gains.Defence); ; current = int(math.Floor(float64(min+max)/2.0 + 0.5)) {
		gains.Defence = float64(current)
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
	gains.Defence = float64(optimumDefence)
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
				case "types":
					if err = sub.Decode(&c.types); err != nil {
						return
					}
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
					if err = sub.Decode(&c.Quality); err != nil {
						return
					}
				case "profits":
					if err = sub.Decode(&c.Profits); err != nil {
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
