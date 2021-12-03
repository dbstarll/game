package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/model"
)

//最终伤害 = 基础伤害 * 元素加伤 * 状态加伤 *(1+真实伤害)
//
//*面板物理攻击 = 物理攻击 * (1+物理攻击%)
//
//魔法攻击 = 素质魔法攻击 + 装备魔法攻击
//*面板魔法攻击 = 魔法攻击 * (1+魔法攻击%)
func main() {
	player := model.NewPlayer(job.Crusader4,
		model.AddQuality(&model.Quality{
			Vit: 402,
		}), model.AddGains(false, &model.Gains{
			Defence:    3246,
			DefencePer: 112.5,
		}))
	fmt.Printf("%+v\n", player.PanelDefence(false))
}
