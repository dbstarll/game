package model

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/elementals"
	"github.com/dbstarll/game/internal/ys/dimension/weaponType"
	"github.com/dbstarll/game/internal/ys/model/buff"
	"time"
)

var CharacterFactory神里绫华 = func(normal, skill, burst, constellation int) *Character {
	return NewCharacter(5, elementals.Ice, weaponType.Sword,
		BaseCharacter(90, 12858, 342, 784, buff.AddCriticalDamage(38.4)),
		TalentsTemplateModifier(NewTalentsTemplate(
			&NormalAttack{name: "神里流·倾", lv: 11, charged: ChargedAttack{stamina: 20}},
			&ElementalSkill{name: "神里流·冰华", lv: 13, cd: time.Second * 10},
			&ElementalBurst{name: "神里流·霜灭", lv: 14, cd: time.Second * 20, infusionDuration: time.Second * 5, energyCost: 80}).
			addNormalAttacks(
				&NormalAttack{lv: 1, hits: []float64{45.7, 48.7, 62.6, 22.7 * 3, 78.2}, charged: ChargedAttack{hits: []float64{55.1, 55.1, 55.1}}, plunge: PlungeAttack{64, 128, 160}},
				&NormalAttack{lv: 2, hits: []float64{49.5, 52.7, 67.7, 24.5 * 3, 84.6}, charged: ChargedAttack{hits: []float64{59.6, 59.6, 59.6}}, plunge: PlungeAttack{69, 138, 173}},
				&NormalAttack{lv: 3, hits: []float64{53.2, 56.6, 72.8, 26.2 * 3, 90.9}, charged: ChargedAttack{hits: []float64{64.1, 64.1, 64.1}}, plunge: PlungeAttack{74, 149, 186}},
				&NormalAttack{lv: 4, hits: []float64{58.5, 62.3, 80.1, 29.0 * 3, 100.0}, charged: ChargedAttack{hits: []float64{70.5, 70.5, 70.5}}, plunge: PlungeAttack{82, 164, 204}},
				&NormalAttack{lv: 5, hits: []float64{62.2, 66.2, 85.2, 30.8 * 3, 106.4}, charged: ChargedAttack{hits: []float64{75.0, 75.0, 75.0}}, plunge: PlungeAttack{87, 174, 217}},
				&NormalAttack{lv: 6, hits: []float64{66.5, 70.8, 91.0, 32.9 * 3, 113.6}, charged: ChargedAttack{hits: []float64{80.1, 80.1, 80.1}}, plunge: PlungeAttack{93, 186, 232}},
				&NormalAttack{lv: 7, hits: []float64{72.3, 77.0, 99.0, 35.8 * 3, 123.6}, charged: ChargedAttack{hits: []float64{87.2, 87.2, 87.2}}, plunge: PlungeAttack{101, 202, 253}},
				&NormalAttack{lv: 8, hits: []float64{78.2, 83.2, 107.0, 38.7 * 3, 133.6}, charged: ChargedAttack{hits: []float64{94.2, 94.2, 94.2}}, plunge: PlungeAttack{109, 219, 273}},
				&NormalAttack{lv: 9, hits: []float64{84.0, 89.4, 115.1, 41.6 * 3, 143.6}, charged: ChargedAttack{hits: []float64{101.3, 101.3, 101.3}}, plunge: PlungeAttack{117, 235, 293}},
				&NormalAttack{lv: 10, hits: []float64{90.4, 96.2, 123.8, 44.8 * 3, 154.6}, charged: ChargedAttack{hits: []float64{109.0, 109.0, 109.0}}, plunge: PlungeAttack{126, 253, 316}},
				&NormalAttack{lv: 11, hits: []float64{96.8, 103.0, 132.5, 47.9 * 3, 165.5}, charged: ChargedAttack{hits: []float64{116.7, 116.7, 116.7}}, plunge: PlungeAttack{135, 271, 338}},
			).
			addElementalSkills(
				&ElementalSkill{lv: 1, dmgs: map[string]float64{"技能": 239}},
				&ElementalSkill{lv: 2, dmgs: map[string]float64{"技能": 257}},
				&ElementalSkill{lv: 3, dmgs: map[string]float64{"技能": 275}},
				&ElementalSkill{lv: 4, dmgs: map[string]float64{"技能": 299}},
				&ElementalSkill{lv: 5, dmgs: map[string]float64{"技能": 317}},
				&ElementalSkill{lv: 6, dmgs: map[string]float64{"技能": 335}},
				&ElementalSkill{lv: 7, dmgs: map[string]float64{"技能": 359}},
				&ElementalSkill{lv: 8, dmgs: map[string]float64{"技能": 383}},
				&ElementalSkill{lv: 9, dmgs: map[string]float64{"技能": 407}},
				&ElementalSkill{lv: 10, dmgs: map[string]float64{"技能": 431}},
				&ElementalSkill{lv: 11, dmgs: map[string]float64{"技能": 454}},
				&ElementalSkill{lv: 12, dmgs: map[string]float64{"技能": 478}},
				&ElementalSkill{lv: 13, dmgs: map[string]float64{"技能": 508}},
			).
			addElementalBursts(
				&ElementalBurst{lv: 1, dmgs: map[string]float64{"切割": 112, "绽放": 168}},
				&ElementalBurst{lv: 2, dmgs: map[string]float64{"切割": 121, "绽放": 181}},
				&ElementalBurst{lv: 3, dmgs: map[string]float64{"切割": 129, "绽放": 194}},
				&ElementalBurst{lv: 4, dmgs: map[string]float64{"切割": 140, "绽放": 211}},
				&ElementalBurst{lv: 5, dmgs: map[string]float64{"切割": 149, "绽放": 223}},
				&ElementalBurst{lv: 6, dmgs: map[string]float64{"切割": 157, "绽放": 236}},
				&ElementalBurst{lv: 7, dmgs: map[string]float64{"切割": 168, "绽放": 253}},
				&ElementalBurst{lv: 8, dmgs: map[string]float64{"切割": 180, "绽放": 270}},
				&ElementalBurst{lv: 9, dmgs: map[string]float64{"切割": 191, "绽放": 286}},
				&ElementalBurst{lv: 10, dmgs: map[string]float64{"切割": 202, "绽放": 303}},
				&ElementalBurst{lv: 11, dmgs: map[string]float64{"切割": 213, "绽放": 320}},
				&ElementalBurst{lv: 12, dmgs: map[string]float64{"切割": 225, "绽放": 337}},
				&ElementalBurst{lv: 13, dmgs: map[string]float64{"切割": 239, "绽放": 358}},
				&ElementalBurst{lv: 14, dmgs: map[string]float64{"切割": 253, "绽放": 379}},
			).check()),
	).Talents(normal, skill, burst)
}
