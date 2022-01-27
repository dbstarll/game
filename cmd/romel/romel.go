package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/client"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/dbstarll/game/internal/ro/romel"
	"go.uber.org/zap"
	"log"
	"time"
)

func init() {
	if logger, err := zap.NewDevelopment(); err != nil {
		log.Fatalf("%+v", err)
	} else {
		zap.RedirectStdLog(logger)
	}
}

func main() {
	//api := client.NewRomelApi("sd32rfgfe344edsd")
	//if err := getMonsterList(api); err != nil {
	//	fmt.Printf("err: %+v\n", err)
	//}

	cnt := make(map[position.Position]int)
	if count, err := romel.Hats.Filter(func(item *romel.Hat) error {
		token := item.Position
		if ov, exist := cnt[token]; exist {
			cnt[token] = ov + 1
		} else {
			cnt[token] = 1
		}
		if _, err := item.Buff.Effect(); err != nil {
			return err
		} else if _, err := item.AdventureBuff.Effect(); err != nil {
			return err
		} else if _, err := item.StorageBuff.Effect(); err != nil {
			return err
		} else if item.StorageRefineBuff != nil {
			for _, rb := range *item.StorageRefineBuff {
				if _, err := rb.Buff.Effect(); err != nil {
					return err
				}
			}
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		log.Fatalf("%+v", err)
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
		if _, err := item.Effect.Effect(); err != nil {
			return err
		} else if _, err := item.Buff.Effect(); err != nil {
			return err
		} else if _, err := item.RandomBuff.Effect(); err != nil {
			return err
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		log.Fatalf("%+v", err)
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
		if _, err := item.Buff.Effect(); err != nil {
			return err
		} else if _, err := item.AdventureBuff.Effect(); err != nil {
			return err
		} else if _, err := item.StorageBuff.Effect(); err != nil {
			return err
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		log.Fatalf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
		for token, c := range cnt {
			fmt.Printf("\t%+v=%d\n", token, c)
		}
	}

	if count, err := romel.Pets.Filter(func(item *romel.Pet) error {
		if _, err := item.AdventureBuff.Effect(); err != nil {
			return err
		} else {
			for _, skill := range *item.Skill {
				if _, err := skill.Intro.Effect(); err != nil {
					return err
				}
			}
		}
		//fmt.Printf("%s: %s\n", item.Name, item.Position)
		return nil
	}); err != nil {
		log.Fatalf("%+v", err)
	} else {
		fmt.Printf("count: %d\n", count)
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
