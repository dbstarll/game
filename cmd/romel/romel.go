package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/client"
	"time"
)

func main() {
	api := client.NewRomelApi("sd32rfgfe344edsd")
	if err := getMonsterList(api); err != nil {
		fmt.Printf("err: %+v\n", err)
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
