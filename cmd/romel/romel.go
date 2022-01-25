package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/client"
	"github.com/dbstarll/game/internal/ro/dimension/quality"
	"github.com/dbstarll/game/internal/ro/romel"
	"log"
	"time"
)

func main() {
	//api := client.NewRomelApi("sd32rfgfe344edsd")
	//if err := getMonsterList(api); err != nil {
	//	fmt.Printf("err: %+v\n", err)
	//}

	cnt := make(map[quality.Quality]int)
	if count, err := romel.Equips.Filter(func(item *romel.Equip) error {
		token := item.Rank
		if ov, exist := cnt[token]; exist {
			cnt[token] = ov + 1
		} else {
			cnt[token] = 1
		}
		fmt.Printf("%s: %+v\n", item.Name, item.Rank)
		return nil
	}); err != nil {
		log.Fatalf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
		for token, c := range cnt {
			fmt.Printf("\t%+v=%d\n", token, c)
		}
	}
}

func getCardList(api *client.RomelApi) error {
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

func getHatList(api *client.RomelApi) error {
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

func getEquipList(api *client.RomelApi) error {
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

func getPetList(api *client.RomelApi) error {
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

func getMonsterList(api *client.RomelApi) error {
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
