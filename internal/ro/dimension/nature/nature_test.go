package nature

import "testing"

func TestNature_Restraint(t *testing.T) {
	tests := []struct {
		name    string
		n       Nature
		defence Nature
		want    float64
	}{
		{
			name:    "Earth->Wind",
			n:       Earth,
			defence: Wind,
			want:    2,
		},
		{
			name:    "Poison->Ghost",
			n:       Poison,
			defence: Ghost,
			want:    0.5,
		},
		{
			name:    "Poison->Unlimited",
			n:       Poison,
			defence: Unlimited,
			want:    1,
		},
		{
			name:    "Unknown->Earth",
			n:       13,
			defence: Earth,
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Restraint(tt.defence); got != tt.want {
				t.Errorf("Restraint() = %v, want %v", got, tt.want)
			}
		})
	}
}
