package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/dbstarll/game/internal/ro/romel"
	"go.uber.org/zap"
	"sort"
	"time"
)

type BuffItem struct {
	name  string
	count int
}

func main() {
	//updateApi()
	detectBuffEffect2()
}

func detectBuffEffect2() {
	fmt.Printf("Buff Total: %d\n", romel.BuffTotal)
	fmt.Printf("\t[%2.2f%%]Detected: %d\n", 100*float64(romel.BuffDetected)/float64(romel.BuffTotal), romel.BuffDetected)
	fmt.Printf("\t[%2.2f%%]Unknown: %d\n", 100*float64(romel.BuffUnknown)/float64(romel.BuffTotal), romel.BuffUnknown)
	fmt.Printf("\t[%2.2f%%]Ignore: %d\n", 100*float64(romel.BuffIgnore)/float64(romel.BuffTotal), romel.BuffIgnore)
	fmt.Printf("\t[%2.2f%%]Error: %d\n", 100*float64(romel.BuffError)/float64(romel.BuffTotal), romel.BuffError)
	var items []*BuffItem
	for k, v := range romel.Buffs {
		items = append(items, &BuffItem{
			name:  k,
			count: v,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		//if items[i].count < items[j].count {
		//	return false
		//} else if items[i].count > items[j].count {
		//	return true
		//} else {
		return items[i].name < items[j].name
		//}
	})
	for idx, item := range items {
		if idx > 20 {
			//break
		}
		//fmt.Printf("占比：%2.4f%% - [%d]%s\n", 100*float64(item.count)/float64(romel.BuffUnknown), item.count, item.name)
		fmt.Printf("%s\n", item.name)
	}
}

func detectBuffEffect() {
	cnt := make(map[position.Position]int)
	if count, err := romel.Hats.Filter(func(item *romel.Hat) error {
		token := item.Position
		if ov, exist := cnt[token]; exist {
			cnt[token] = ov + 1
		} else {
			cnt[token] = 1
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		zap.S().Errorf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
		for token, c := range cnt {
			fmt.Printf("\t%+v=%d\n", token, c)
		}
	}

	cnt = make(map[position.Position]int)
	if count, err := romel.Equips.Filter(func(item *romel.Equip) error {
		token := item.Position
		if ov, exist := cnt[token]; exist {
			cnt[token] = ov + 1
		} else {
			cnt[token] = 1
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		zap.S().Errorf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
		for token, c := range cnt {
			fmt.Printf("\t%+v=%d\n", token, c)
		}
	}

	cnt = make(map[position.Position]int)
	if count, err := romel.Cards.Filter(func(item *romel.Card) error {
		token := item.Position
		if ov, exist := cnt[token]; exist {
			cnt[token] = ov + 1
		} else {
			cnt[token] = 1
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		zap.S().Errorf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
		for token, c := range cnt {
			fmt.Printf("\t%+v=%d\n", token, c)
		}
	}

	if count, err := romel.Pets.Filter(func(item *romel.Pet) error {
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		zap.S().Errorf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
	}
}

func updateApi() {
	api := romel.NewRomelApi("sd32rfgfe344edsd")
	if err := getMonsterList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
}

func getCardList(api *romel.RomelApi) error {
	if result, err := api.GetCardList(1); err != nil {
		return err
	} else {
		for page := 2; page <= result.Data.PageCount; page++ {
			time.Sleep(time.Second * 10)
			if _, err := api.GetCardList(page); err != nil {
				return err
			}
		}
		return nil
	}
}

func getHatList(api *romel.RomelApi) error {
	if result, err := api.GetHatList(1); err != nil {
		return err
	} else {
		for page := 2; page <= result.Data.PageCount; page++ {
			time.Sleep(time.Second * 10)
			if _, err := api.GetHatList(page); err != nil {
				return err
			}
		}
		return nil
	}
}

func getEquipList(api *romel.RomelApi) error {
	if result, err := api.GetEquipList(1); err != nil {
		return err
	} else {
		for page := 2; page <= result.Data.PageCount; page++ {
			time.Sleep(time.Second * 10)
			if _, err := api.GetEquipList(page); err != nil {
				return err
			}
		}
		return nil
	}
}

func getPetList(api *romel.RomelApi) error {
	if result, err := api.GetPetList(1); err != nil {
		return err
	} else {
		for page := 2; page <= result.Data.PageCount; page++ {
			time.Sleep(time.Second * 10)
			if _, err := api.GetPetList(page); err != nil {
				return err
			}
		}
		return nil
	}
}

func getMonsterList(api *romel.RomelApi) error {
	if result, err := api.GetMonsterList(1); err != nil {
		return err
	} else {
		for page := 2; page <= result.Data.PageCount; page++ {
			time.Sleep(time.Second * 10)
			if _, err := api.GetMonsterList(page); err != nil {
				return err
			}
		}
		return nil
	}
}
