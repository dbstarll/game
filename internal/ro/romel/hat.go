package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

var Hats *hats

type hats struct {
	ids   map[string]*Hat
	names map[string]*Hat
}

type Hat struct {
	Id                string       `json:"id"`
	Name              string       `json:"name"`
	Rank              int          `json:"rank"`
	Position          int          `json:"position"`
	Effect            string       `json:"effect"`
	Buff              string       `json:"buff"`
	AdventureBuff     string       `json:"adventureBuff"`
	StorageBuff       string       `json:"storageBuff"`
	StorageRefineBuff []RefineBuff `json:"storageRefineBuff"`
	IsCompose         int          `json:"isCompose"`
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

func (e *hats) Get(name string) (*Hat, bool) {
	hat, found := e.names[name]
	return hat, found
}
