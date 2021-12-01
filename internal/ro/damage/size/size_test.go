package size

import "testing"

func TestSize_String(t *testing.T) {
	tests := []struct {
		name string
		s    Size
		want string
	}{
		{
			name: "unlimited",
			s:    Unlimited,
			want: "不限",
		},
		{
			name: "small",
			s:    Small,
			want: "小型",
		},
		{
			name: "medium",
			s:    Medium,
			want: "中型",
		},
		{
			name: "large",
			s:    Large,
			want: "大型",
		},
		{
			name: "unknown",
			s:    5,
			want: "未知",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
