package entry

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
	"reflect"
	"testing"
)

func TestEntry_Primary(t *testing.T) {
	tests := []struct {
		name string
		pos  position.Position
		want []Entry
	}{
		{name: "生之花主词条", pos: position.FlowerOfLife, want: []Entry{Hp}},
		{name: "死之羽主词条", pos: position.PlumeOfDeath, want: []Entry{Atk}},
		{name: "时之沙主词条", pos: position.SandsOfEon, want: []Entry{HpPercentage, AtkPercentage, DefPercentage, ElementalMastery, EnergyRecharge}},
		{name: "空之杯主词条", pos: position.GobletOfEonothem, want: []Entry{
			HpPercentage, AtkPercentage, DefPercentage, ElementalMastery, PhysicalDamageBonus, FireDamageBonus,
			WaterDamageBonus, GrassDamageBonus, ElectricDamageBonus, WindDamageBonus, IceDamageBonus, EarthDamageBonus,
		}},
		{name: "理之冠主词条", pos: position.CircletOfLogos, want: []Entry{HpPercentage, AtkPercentage, DefPercentage, ElementalMastery, CriticalRate, CriticalDamage, HealingBonus}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []Entry
			for _, ent := range Entries {
				if ent.Primary(tt.pos) {
					got = append(got, ent)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Primary(%s) = %v, want %v", tt.pos, got, tt.want)
			}
		})
	}
}

func TestEntry_Secondary(t *testing.T) {
	type test struct {
		name  string
		e     Entry
		rate  float64
		exist bool
	}
	tests := []test{
		{name: "副词条生命值", e: Hp, rate: 2, exist: true},
		{name: "副词条生命值%", e: HpPercentage, rate: 1, exist: true},
		{name: "副词条攻击力", e: Atk, rate: 2, exist: true},
		{name: "副词条攻击力%", e: AtkPercentage, rate: 1, exist: true},
		{name: "副词条防御力", e: Def, rate: 1, exist: true},
		{name: "副词条防御力%", e: DefPercentage, rate: 1, exist: true},
		{name: "副词条元素精通", e: ElementalMastery, rate: 1, exist: true},
		{name: "副词条暴击率", e: CriticalRate, rate: 1, exist: true},
		{name: "副词条暴击伤害", e: CriticalDamage, rate: 1, exist: true},
		{name: "副词条治疗加成", e: HealingBonus, rate: 0, exist: false},
		{name: "副词条元素充能效率", e: EnergyRecharge, rate: 1, exist: true},
	}
	for _, ent := range Entries {
		if ent >= PhysicalDamageBonus && ent <= EarthDamageBonus {
			tests = append(tests, test{name: fmt.Sprintf("副词条%s", ent), e: ent, rate: 0, exist: false})
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.Secondary()
			if got != tt.rate {
				t.Errorf("%s.Secondary() got = %v, want %v", tt.e, got, tt.rate)
			}
			if got1 != tt.exist {
				t.Errorf("%s.Secondary() got1 = %v, want %v", tt.e, got1, tt.exist)
			}
		})
	}
}
