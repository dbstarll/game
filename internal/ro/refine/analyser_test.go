package refine

import "testing"

func TestAnalyse_analyse(t *testing.T) {
	analyser := NewAnalyse()
	//if err := analyser.play("refine1.log", 3, 4, &Resource{zeny: 905725339, al: 7559, alPlus: 77, fire: 234}); err != nil {
	//	t.Fatal(err)
	//}
	if _, _, err := analyser.play("data/refine2.log", 4, &Resource{zeny: 885990627, al: 6340, alPlus: 234, fire: 162}); err != nil {
		t.Fatal(err)
	}
}
