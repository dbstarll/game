package romel

import (
	"encoding/json"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var Hats *hats

type hats struct {
	ids   map[string]*Hat
	names map[string]*Hat
}

type Hat struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Rank              int               `json:"rank"`
	Position          position.Position `json:"position"`
	Buff              string            `json:"buff"`
	AdventureBuff     string            `json:"adventureBuff"`
	StorageBuff       string            `json:"storageBuff"`
	StorageRefineBuff *[]RefineBuff     `json:"storageRefineBuff"`
	IsCompose         int               `json:"isCompose"`
}

type RefineBuff struct {
	Lv   int    `json:"lv"`
	Buff string `json:"buff"`
}

func init() {
	if hats, err := loadHats(); err != nil {
		log.Fatalf("load hats failed: %+v", err)
	} else {
		Hats = hats
	}
}

func loadHats() (*hats, error) {
	root, hats := RootRomel+"hat", &hats{
		ids:   make(map[string]*Hat),
		names: make(map[string]*Hat),
	}
	if err := iterate(root, func(item map[string]interface{}, data []byte) error {
		return hats.add(item, data)
	}); err != nil {
		return nil, err
	} else {
		log.Printf("load %d hats from %s", hats.Size(), root)
		return hats, nil
	}
}

func (h *hats) add(item map[string]interface{}, data []byte) error {
	hat := &Hat{}
	if err := json.Unmarshal(data, hat); err != nil {
		return errors.WithStack(err)
	} else {
		h.ids[hat.Id] = hat
		h.names[hat.Name] = hat
		delete(item, "id")
		delete(item, "name")
		delete(item, "icon")
		delete(item, "rank")
		delete(item, "position")
		delete(item, "effect")
		delete(item, "buff")
		delete(item, "adventureBuff")
		delete(item, "storageBuff")
		delete(item, "storageRefineBuff")
		delete(item, "isCompose")
		if len(item) > 0 {
			for k, v := range item {
				log.Printf("unknown hat property: %s=%+v", k, v)
			}
		}
		return nil
	}
}

func (h *hats) Size() int {
	return len(h.ids)
}

func (h *hats) Get(name string) *Hat {
	return h.names[name]
}

func (h *hats) Filter(fn func(*Hat) error, filterFn ...func(filter *Hat)) (int, error) {
	count, filters := 0, make([]*Hat, len(filterFn))
	for idx, f := range filterFn {
		filters[idx] = &Hat{IsCompose: -1}
		f(filters[idx])
	}
	for _, hat := range h.ids {
		if hat.matchAny(filters...) {
			count++
			if err := fn(hat); err != nil {
				return 0, err
			}
		}
	}
	return count, nil
}

func (h *Hat) matchAny(filters ...*Hat) bool {
	for _, filter := range filters {
		if h.match(filter) {
			return true
		}
	}
	return len(filters) == 0
}

func (h *Hat) match(filter *Hat) bool {
	if filter.Rank > 0 && filter.Rank != h.Rank {
		return false
	} else if filter.Position > position.Unlimited && filter.Position != h.Position {
		return false
	} else if filter.IsCompose >= 0 && filter.IsCompose != h.IsCompose {
		return false
	} else if len(filter.Name) > 0 && strings.Index(h.Name, filter.Name) < 0 {
		return false
	} else if len(filter.Buff) > 0 && strings.Index(h.Buff, filter.Buff) < 0 {
		return false
	} else if len(filter.AdventureBuff) > 0 && strings.Index(h.AdventureBuff, filter.AdventureBuff) < 0 {
		return false
	} else if len(filter.StorageBuff) > 0 && strings.Index(h.StorageBuff, filter.StorageBuff) < 0 {
		return false
	} else if filter.StorageRefineBuff != nil {
		if h.StorageRefineBuff == nil {
			return false
		} else {
			match, testCount := false, 0
			for _, filterBuff := range *filter.StorageRefineBuff {
				if len(filterBuff.Buff) > 0 {
					testCount++
					for _, buff := range *h.StorageRefineBuff {
						if strings.Index(buff.Buff, filterBuff.Buff) >= 0 {
							match = true
							break
						}
					}
					if match {
						break
					}
				}
			}
			return testCount == 0 || match
		}
	} else {
		return true
	}
}
