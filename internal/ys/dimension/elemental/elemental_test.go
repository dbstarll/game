package elemental

import "testing"

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
		{name: "风元素附魔会被水火冰雷任何一种元素覆盖", a: Wind, b: Water, want: Water},
		{name: "风元素附魔会被水火冰雷任何一种元素覆盖", a: Wind, b: Fire, want: Fire},
		{name: "风元素附魔会被水火冰雷任何一种元素覆盖", a: Wind, b: Ice, want: Ice},
		{name: "风元素附魔会被水火冰雷任何一种元素覆盖", a: Wind, b: Electric, want: Electric},
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
