package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

var Cards *cards

type cards struct {
	ids   map[string]*Card
	names map[string]*Card
}

type Card struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Quality       int    `json:"quality"`
	Position      int    `json:"position"`
	Buff          string `json:"buff"`
	AdventureBuff string `json:"adventureBuff"`
	StorageBuff   string `json:"storageBuff"`
	IsCompose     int    `json:"isCompose"`
}

func init() {
	if cards, err := loadCards(); err != nil {
		log.Fatalf("load cards failed: %+v", err)
	} else {
		Cards = cards
	}
}

func loadCards() (*cards, error) {
	root, cards := RootRomel+"card", &cards{
		ids:   make(map[string]*Card),
		names: make(map[string]*Card),
	}
	if err := iterate(root, func(item map[string]interface{}, data []byte) error {
		return cards.add(item, data)
	}); err != nil {
		return nil, err
	} else {
		log.Printf("load %d cards from %s", cards.Size(), root)
		return cards, nil
	}
}

func (c *cards) add(item map[string]interface{}, data []byte) error {
	card := &Card{}
	if err := json.Unmarshal(data, card); err != nil {
		return errors.WithStack(err)
	} else {
		c.ids[card.Id] = card
		c.names[card.Name] = card
		delete(item, "id")
		delete(item, "name")
		delete(item, "icon")
		delete(item, "picture")
		delete(item, "quality")
		delete(item, "position")
		delete(item, "buff")
		delete(item, "adventureBuff")
		delete(item, "storageBuff")
		delete(item, "isCompose")
		if len(item) > 0 {
			for k, v := range item {
				log.Printf("unknown card property: %s=%+v", k, v)
			}
		}
		return nil
	}
}

func (c *cards) Size() int {
	return len(c.ids)
}

func (c *cards) Get(name string) *Card {
	return c.names[name]
}
