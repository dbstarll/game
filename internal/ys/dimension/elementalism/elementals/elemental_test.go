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
				tests = append(tests, test{name: "非附魔元素可以被附魔元素附魔", owner: from, infusion: to, want: to, twoWay: false})
			} else if !to.CanInfusion() {
				tests = append(tests, test{name: "不可附魔元素不能改变原元素", owner: from, infusion: to, want: from, twoWay: false})
			} else if from == to {
				tests = append(tests, test{name: "同元素附魔", owner: from, infusion: from, want: from, twoWay: false})
			}
		}
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
	type test struct {
		name     string
		trigger  Elemental
		attached Elemental
		want     *reactions.React
		twoWay   bool
	}
	all := make(map[string]bool)
	for _, from := range append(Elementals, -1, 1000) {
		for _, to := range append(Elementals, -1, 1000) {
			all[fmt.Sprintf("%s -> %s", from, to)] = false
		}
	}
	tests := []test{
		// 增幅反应
		{name: "火水蒸发", trigger: Fire, attached: Water, want: reactions.NewReact(reactions.Vaporize, 1.5)},
		{name: "水火蒸发", trigger: Water, attached: Fire, want: reactions.NewReact(reactions.Vaporize, 2)},
		{name: "火冰融化", trigger: Fire, attached: Ice, want: reactions.NewReact(reactions.Melt, 2)},
		{name: "冰火融化", trigger: Ice, attached: Fire, want: reactions.NewReact(reactions.Melt, 1.5)},
		// 剧变反应
		{name: "火雷超载", trigger: Fire, attached: Electric, want: reactions.NewReact(reactions.Overload, 2), twoWay: true},
		{name: "冰雷超导", trigger: Ice, attached: Electric, want: reactions.NewReact(reactions.Superconduct, 0.5), twoWay: true},
		{name: "水雷感电", trigger: Water, attached: Electric, want: reactions.NewReact(reactions.ElectroCharged, 1.2), twoWay: true},
		{name: "水草绽放", trigger: Water, attached: Grass, want: reactions.NewReact(reactions.Bloom, 2), twoWay: true},
		{name: "火草燃烧", trigger: Fire, attached: Grass, want: reactions.NewReact(reactions.Burn, 0.25), twoWay: true},
		{name: "水冰冻结", trigger: Water, attached: Ice, want: reactions.NewReact(reactions.Frozen, 0), twoWay: true},
		// 激化反应
		{name: "草雷激化", trigger: Grass, attached: Electric, want: reactions.NewReact(reactions.Catalyze, 0), twoWay: true},
		// 无反应
		{name: "草冰无反应", trigger: Grass, attached: Ice, want: nil, twoWay: true},
	}
	for _, from := range append(Elementals, -1, 1000) {
		for _, to := range append(Elementals, -1, 1000) {
			if !from.IsValid() || !to.IsValid() {
				tests = append(tests, test{name: "无效元素无反应", trigger: from, attached: to, want: nil})
			} else if from == Physical || to == Physical {
				tests = append(tests, test{name: "物理无反应", trigger: from, attached: to, want: nil})
			} else if from == to {
				tests = append(tests, test{name: "同元素无反应", trigger: from, attached: to, want: nil})
			} else if from == Wind {
				if to.CanInfusion() {
					tests = append(tests, test{name: fmt.Sprintf("风%s扩散", to), trigger: from, attached: to, want: reactions.NewReact(reactions.Swirl, 0.6), twoWay: true})
				} else {
					tests = append(tests, test{name: fmt.Sprintf("风%s无反应", to), trigger: from, attached: to, want: nil, twoWay: true})
				}
			} else if from == Earth {
				if to.CanInfusion() {
					tests = append(tests, test{name: fmt.Sprintf("岩%s结晶", to), trigger: from, attached: to, want: reactions.NewReact(reactions.Crystallize, 0), twoWay: true})
				} else {
					tests = append(tests, test{name: fmt.Sprintf("岩%s无反应", to), trigger: from, attached: to, want: nil, twoWay: true})
				}
			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.trigger.Reaction(tt.attached); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s.Reaction(%s) = %v, want %v", tt.trigger, tt.attached, got, tt.want)
			} else {
				delete(all, fmt.Sprintf("%s -> %s", tt.trigger, tt.attached))
			}
			if tt.twoWay {
				if got := tt.attached.Reaction(tt.trigger); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s.Reaction(%s) = %v, want %v", tt.attached, tt.trigger, got, tt.want)
				} else {
					delete(all, fmt.Sprintf("%s -> %s", tt.attached, tt.trigger))
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
