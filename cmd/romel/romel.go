package main

import (
	"fmt"
	"github.com/dbstarll/game/internal/ro/client"
	"time"
)

func main() {
	api := client.NewRomelApi("sd32rfgfe344edsd")
	if err := getCardList(api); err != nil {
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
