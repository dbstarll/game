package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension"
	"github.com/dbstarll/game/internal/ro/dimension/job"
)

//最终伤害 = 基础伤害 * 元素加伤 * 状态加伤 *(1+真实伤害)
//
//*面板物理攻击 = 物理攻击 * (1+物理攻击%)
//
//魔法攻击 = 素质魔法攻击 + 装备魔法攻击
//*面板魔法攻击 = 魔法攻击 * (1+魔法攻击%)
func main() {
	player := dimension.NewPlayer(job.Crusader4,
		dimension.AddQuality(&dimension.Quality{
			Str: 0,
			Agi: 0,
			Vit: 0,
			Int: 0,
			Dex: 0,
			Luk: 0,
		}), dimension.AddLevel(&dimension.Level{
			Base: 170,
			Job:  70,
		}), dimension.AddGains(false, &dimension.Gains{
			Attack:  0,
			Defence: 0,
		}))
	fmt.Printf("%+v\n", player.Defence(false))
}
