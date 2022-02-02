package romel

import (
	"encoding/json"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/dbstarll/game/internal/ro/dimension/quality"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var Cards *cards

type cards struct {
	ids   map[string]*Card
	names map[string]*Card
}

type Card struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	Quality       quality.Quality   `json:"quality"`
	Position      position.Position `json:"position"`
	Buff          *Buff             `json:"buff"`
	AdventureBuff *Buff             `json:"adventureBuff"`
	StorageBuff   *Buff             `json:"storageBuff"`
	IsCompose     int               `json:"isCompose"`
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

func (c *cards) Filter(fn func(*Card) error, filterFn ...func(filter *Card)) (int, error) {
	count, filters := 0, make([]*Card, len(filterFn))
	for idx, f := range filterFn {
		filters[idx] = &Card{IsCompose: -1}
		f(filters[idx])
	}
	for _, card := range c.ids {
		if card.matchAny(filters...) {
			count++
			if err := fn(card); err != nil {
				return 0, err
			}
		}
	}
	return count, nil
}

func (c *Card) matchAny(filters ...*Card) bool {
	for _, filter := range filters {
		if c.match(filter) {
			return true
		}
	}
	return len(filters) == 0
}

func (c *Card) match(filter *Card) bool {
	if filter.Quality > quality.Unlimited && filter.Quality != c.Quality {
		return false
	} else if filter.Position > position.Unlimited && filter.Position != c.Position {
		return false
	} else if filter.IsCompose >= 0 && filter.IsCompose != c.IsCompose {
		return false
	} else if len(filter.Name) > 0 && strings.Index(c.Name, filter.Name) < 0 {
		return false
	} else if !c.Buff.Contains(filter.Buff) {
		return false
	} else if !c.AdventureBuff.Contains(filter.AdventureBuff) {
		return false
	} else if !c.StorageBuff.Contains(filter.StorageBuff) {
		return false
	} else {
		return true
	}
}
