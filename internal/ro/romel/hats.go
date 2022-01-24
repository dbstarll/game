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

func (e *hats) add(item map[string]interface{}, data []byte) error {
	hat := &Hat{}
	if err := json.Unmarshal(data, hat); err != nil {
		return errors.WithStack(err)
	} else {
		e.ids[hat.Id] = hat
		e.names[hat.Name] = hat
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

func (e *hats) Size() int {
	return len(e.ids)
}

func (e *hats) Get(name string) *Hat {
	return e.names[name]
}

func (e *hats) Filter(filter *Hat, fn func(*Hat) error) (int, error) {
	if filter == nil {
		filter = &Hat{}
	}
	count := 0
	for _, hat := range e.ids {
		if filter.Rank > 0 && filter.Rank != hat.Rank {
			continue
		} else if filter.Position > position.Unlimited && filter.Position != hat.Position {
			continue
		} else if filter.IsCompose >= 0 && filter.IsCompose != hat.IsCompose {
			continue
		} else if len(filter.Name) > 0 && strings.Index(hat.Name, filter.Name) < 0 {
			continue
		} else if len(filter.Buff) > 0 && strings.Index(hat.Buff, filter.Buff) < 0 {
			continue
		} else if len(filter.AdventureBuff) > 0 && strings.Index(hat.AdventureBuff, filter.AdventureBuff) < 0 {
			continue
		} else if len(filter.StorageBuff) > 0 && strings.Index(hat.StorageBuff, filter.StorageBuff) < 0 {
			continue
		} else if filter.StorageRefineBuff != nil {
			if hat.StorageRefineBuff == nil {
				continue
			} else {
				match, testCount := false, 0
				for _, frb := range *filter.StorageRefineBuff {
					if len(frb.Buff) > 0 {
						testCount++
						for _, rb := range *hat.StorageRefineBuff {
							if strings.Index(rb.Buff, frb.Buff) >= 0 {
								match = true
								break
							}
						}
						if match {
							break
						}
					}
				}
				if testCount > 0 && !match {
					continue
				}
			}
		}

		count++
		if err := fn(hat); err != nil {
			return 0, err
		}
	}
	return count, nil
}
