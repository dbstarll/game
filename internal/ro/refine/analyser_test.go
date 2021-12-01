package refine

import "testing"

func TestAnalyse_analyse(t *testing.T) {
	analyser := NewAnalyse()
	//副手
	if _, err := analyser.play("data/refine1.log", 4, &Resource{zeny: 905725339, al: 7559, alPlus: 77, fire: 234}, false); err != nil {
		t.Fatal(err)
	}
	//盔甲
	//if _, err := analyser.play("data/refine2.log", 4, &Resource{zeny: 885990627, al: 6340, alPlus: 234, fire: 162}, false); err != nil {
	//	t.Fatal(err)
	//}
	if err := analyser.analyse(6); err != nil {
		t.Fatal(err)
	}
	//5.09% - XXOXX-O: 65.38%, 26
	//2.15% - OXXXO-O: 63.64%, 11
	//3.91% - XXOOX-O: 60.00%, 20

	//4.82% - XXXXO-O: 72.73%, 11
	//4.82% - OXXXX-O: 63.64%, 11
	//4.39% - XOOOX-O: 60.00%, 10

}
