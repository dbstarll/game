package elementals

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions"
	"reflect"
	"testing"
)

func TestElemental_Infusion(t *testing.T) {
	type test struct {
		name     string
		owner    Elemental
		infusion Elemental
		want     Elemental
		twoWay   bool
	}
	all := make(map[string]bool)
	for _, from := range append(Elementals, -1, 1000) {
		for _, to := range append(Elementals, -1, 1000) {
			all[fmt.Sprintf("%s -> %s", from, to)] = false
		}
	}

	tests := []test{
		{name: "火水附魔，水伤", owner: Fire, infusion: Water, want: Water, twoWay: true},
		{name: "火雷附魔，火伤", owner: Fire, infusion: Electric, want: Fire, twoWay: true},
		{name: "火冰附魔，火伤", owner: Fire, infusion: Ice, want: Fire, twoWay: true},
		{name: "水雷附魔，水伤", owner: Water, infusion: Electric, want: Water, twoWay: true},
		{name: "水冰附魔，冰伤", owner: Water, infusion: Ice, want: Ice, twoWay: true},
		{name: "雷冰附魔，冰伤", owner: Electric, infusion: Ice, want: Ice, twoWay: true},
	}
	for _, from := range append(Elementals, -1, 1000) {
		for _, to := range append(Elementals, -1, 1000) {
			if !from.CanInfusion() && to.CanInfusion() {
				tests = append(tests, []test{
					{name: "非附魔元素可以被附魔元素附魔", owner: from, infusion: to, want: to, twoWay: false},
				}...)
			} else if !to.CanInfusion() {
				tests = append(tests, []test{
					{name: "不可附魔元素不能改变原元素", owner: from, infusion: to, want: from, twoWay: false},
				}...)
			} else if from == to {
				tests = append(tests, []test{
					{name: "同元素附魔", owner: from, infusion: from, want: from, twoWay: false},
				}...)
			}
		}
		//	tests = append(tests, []test{
		//		{name: "物理可以被任何元素附魔", owner: Physical, infusion: e, want: e, twoWay: true},
		//		{name: "相同附魔无变化", owner: e, infusion: e, want: e, twoWay: true},
		//		{name: "未知[-1]可以被任何元素附魔", owner: -1, infusion: e, want: e, twoWay: true},
		//		{name: "未知[1000]可以被任何元素附魔", owner: 1000, infusion: e, want: e, twoWay: true},
		//	}...)
		//	if e.CanInfusion() {
		//		tests = append(tests, []test{
		//			{name: "岩元素可以被任何元素附魔", owner: Earth, infusion: e, want: e, twoWay: true},
		//			{name: "草元素可以被任何元素附魔", owner: Grass, infusion: e, want: e, twoWay: true},
		//		}...)
		//	}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.owner.Infusion(tt.infusion); got != tt.want {
				t.Errorf("%s.Infusion(%s) = %v, want %v", tt.owner, tt.infusion, got, tt.want)
			} else {
				delete(all, fmt.Sprintf("%s -> %s", tt.owner, tt.infusion))
			}
			if tt.twoWay {
				if got := tt.infusion.Infusion(tt.owner); got != tt.want {
					t.Errorf("%s.Infusion(%s) = %v, want %v", tt.infusion, tt.owner, got, tt.want)
				} else {
					delete(all, fmt.Sprintf("%s -> %s", tt.infusion, tt.owner))
				}
			}
		})
	}

	if len(all) > 0 {
		t.Errorf("未测试：%d", len(all))
		for k, _ := range all {
			t.Logf("\t场景：%s", k)
		}
	}
}

func TestElemental_Reaction(t *testing.T) {
	tests := []struct {
		name     string
		trigger  Elemental
		attached Elemental
		want     *reactions.React
	}{
		// 增幅反应
		{name: "火水蒸发", trigger: Fire, attached: Water, want: reactions.NewReact(reactions.Vaporize, 1.5)},
		{name: "水火蒸发", trigger: Water, attached: Fire, want: reactions.NewReact(reactions.Vaporize, 2)},
		{name: "火冰融化", trigger: Fire, attached: Ice, want: reactions.NewReact(reactions.Melt, 2)},
		{name: "冰火融化", trigger: Ice, attached: Fire, want: reactions.NewReact(reactions.Melt, 1.5)},
		// 剧变反应
		{name: "火雷超载", trigger: Fire, attached: Electric, want: reactions.NewReact(reactions.Overload, 2)},
		{name: "雷火超载", trigger: Electric, attached: Fire, want: reactions.NewReact(reactions.Overload, 2)},
		{name: "冰雷超导", trigger: Ice, attached: Electric, want: reactions.NewReact(reactions.Superconduct, 0.5)},
		{name: "雷冰超导", trigger: Electric, attached: Ice, want: reactions.NewReact(reactions.Superconduct, 0.5)},
		{name: "水雷感电", trigger: Water, attached: Electric, want: reactions.NewReact(reactions.ElectroCharged, 1.2)},
		{name: "雷水感电", trigger: Electric, attached: Water, want: reactions.NewReact(reactions.ElectroCharged, 1.2)},
		{name: "火风扩散", trigger: Fire, attached: Wind, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "风火扩散", trigger: Wind, attached: Fire, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "水风扩散", trigger: Water, attached: Wind, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "风水扩散", trigger: Wind, attached: Water, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "冰风扩散", trigger: Ice, attached: Wind, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "风冰扩散", trigger: Wind, attached: Ice, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "雷风扩散", trigger: Electric, attached: Wind, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "风雷扩散", trigger: Wind, attached: Electric, want: reactions.NewReact(reactions.Swirl, 0.6)},
		{name: "水草绽放", trigger: Water, attached: Grass, want: reactions.NewReact(reactions.Bloom, 2)},
		{name: "草水绽放", trigger: Grass, attached: Water, want: reactions.NewReact(reactions.Bloom, 2)},
		{name: "火草燃烧", trigger: Fire, attached: Grass, want: reactions.NewReact(reactions.Burn, 0.25)},
		{name: "草火燃烧", trigger: Grass, attached: Fire, want: reactions.NewReact(reactions.Burn, 0.25)},
		{name: "水冰冻结", trigger: Water, attached: Ice, want: reactions.NewReact(reactions.Frozen, 0)},
		{name: "冰水冻结", trigger: Ice, attached: Water, want: reactions.NewReact(reactions.Frozen, 0)},

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
