package elemental

import (
	"github.com/dbstarll/game/internal/ys/dimension/reaction"
	"reflect"
	"testing"
)

func TestElemental_Infusion(t *testing.T) {
	tests := []struct {
		name string
		a    Elemental
		b    Elemental
		want Elemental
	}{
		{name: "火冰附魔，火伤", a: Fire, b: Ice, want: Fire},
		{name: "火雷附魔，火伤", a: Fire, b: Electric, want: Fire},
		{name: "冰水附魔，冰伤", a: Ice, b: Water, want: Ice},
		{name: "冰雷附魔，冰伤", a: Ice, b: Electric, want: Ice},
		{name: "水火附魔，水伤", a: Water, b: Fire, want: Water},
		{name: "水雷附魔，水伤", a: Water, b: Electric, want: Water},
		{name: "风水附魔，水伤", a: Wind, b: Water, want: Water},
		{name: "风火附魔，火伤", a: Wind, b: Fire, want: Fire},
		{name: "风冰附魔，冰伤", a: Wind, b: Ice, want: Ice},
		{name: "风雷附魔，雷伤", a: Wind, b: Electric, want: Electric},
	}
	for _, e := range Elements {
		tests = append(tests, []struct {
			name string
			a    Elemental
			b    Elemental
			want Elemental
		}{
			{name: "物理可以被任何元素附魔", a: Physical, b: e, want: e},
			{name: "相同附魔无变化", a: e, b: e, want: e},
		}...)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Infusion(tt.b); got != tt.want {
				t.Errorf("Infusion() = %v, want %v", got, tt.want)
			}
			if got := tt.b.Infusion(tt.a); got != tt.want {
				t.Errorf("Infusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElemental_Reaction(t *testing.T) {
	tests := []struct {
		name     string
		trigger  Elemental
		attached Elemental
		want     *reaction.Factor
	}{
		// 增幅反应
		{name: "火水蒸发", trigger: Fire, attached: Water, want: reaction.NewFactor(reaction.Vaporize, 1.5)},
		{name: "水火蒸发", trigger: Water, attached: Fire, want: reaction.NewFactor(reaction.Vaporize, 2)},
		{name: "火冰融化", trigger: Fire, attached: Ice, want: reaction.NewFactor(reaction.Melt, 2)},
		{name: "冰火融化", trigger: Ice, attached: Fire, want: reaction.NewFactor(reaction.Melt, 1.5)},
		// 剧变反应
		{name: "火雷超载", trigger: Fire, attached: Electric, want: reaction.NewFactor(reaction.Overload, 2)},
		{name: "雷火超载", trigger: Electric, attached: Fire, want: reaction.NewFactor(reaction.Overload, 2)},
		{name: "冰雷超导", trigger: Ice, attached: Electric, want: reaction.NewFactor(reaction.Superconduct, 0.5)},
		{name: "雷冰超导", trigger: Electric, attached: Ice, want: reaction.NewFactor(reaction.Superconduct, 0.5)},
		{name: "水雷感电", trigger: Water, attached: Electric, want: reaction.NewFactor(reaction.ElectroCharged, 1.2)},
		{name: "雷水感电", trigger: Electric, attached: Water, want: reaction.NewFactor(reaction.ElectroCharged, 1.2)},
		{name: "火风扩散", trigger: Fire, attached: Wind, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "风火扩散", trigger: Wind, attached: Fire, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "水风扩散", trigger: Water, attached: Wind, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "风水扩散", trigger: Wind, attached: Water, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "冰风扩散", trigger: Ice, attached: Wind, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "风冰扩散", trigger: Wind, attached: Ice, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "雷风扩散", trigger: Electric, attached: Wind, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "风雷扩散", trigger: Wind, attached: Electric, want: reaction.NewFactor(reaction.Swirl, 0.6)},
		{name: "水草绽放", trigger: Water, attached: Grass, want: reaction.NewFactor(reaction.Bloom, 2)},
		{name: "草水绽放", trigger: Grass, attached: Water, want: reaction.NewFactor(reaction.Bloom, 2)},
		{name: "火草燃烧", trigger: Fire, attached: Grass, want: reaction.NewFactor(reaction.Burn, 0.25)},
		{name: "草火燃烧", trigger: Grass, attached: Fire, want: reaction.NewFactor(reaction.Burn, 0.25)},
		{name: "水冰冻结", trigger: Water, attached: Ice, want: reaction.NewFactor(reaction.Frozen, 0)},
		{name: "冰水冻结", trigger: Ice, attached: Water, want: reaction.NewFactor(reaction.Frozen, 0)},

		// 碎冰1.5：烈绽放3：超绽放3
		//Shattered                      // 碎冰
		//Crystallize                    // 结晶
		//Hyperbloom                     // 超绽放
		//Burgeon                        // 烈绽放
		//Catalyze                       // 激化
		//Quicken                        // 原激化
		//Aggravate                      // 超激化
		//Spread                         // 蔓激化
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.trigger.Reaction(tt.attached); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
