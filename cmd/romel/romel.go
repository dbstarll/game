package main

import (
	"fmt"
	_ "github.com/dbstarll/game/internal/logger"
	"github.com/dbstarll/game/internal/ro/romel"
	"sort"
	"time"
)

type BuffItem struct {
	name  string
	count int
}

func main() {
	//updateApi()
	detectBuffEffect()
}

func detectBuffEffect() {
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
		if items[i].count < items[j].count {
			return false
		} else if items[i].count > items[j].count {
			return true
		} else {
			return items[i].name < items[j].name
		}
	})
	for idx, item := range items {
		fmt.Printf("%d.占比：%2.4f%% - [%d]%s\n", idx, 100*float64(item.count)/float64(romel.BuffUnknown), item.count, item.name)
	}
}

func updateApi() {
	api := romel.NewRomelApi("sd32rfgfe344edsd")
	if err := getCardList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	if err := getHatList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	if err := getEquipList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	if err := getPetList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
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
