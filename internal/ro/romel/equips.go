package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

var Equips *equips

type equips struct {
	ids   map[string]*Equip
	names map[string]*Equip
}

type Equip struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Rank       int      `json:"rank"`
	Job        []string `json:"job"`
	Position   int      `json:"position"`
	Effect     string   `json:"effect"`
	Buff       string   `json:"buff"`
	IsCompose  int      `json:"isCompose"`
	IsUpgrade  int      `json:"isUpgrade"`
	IsHigh     int      `json:"isHigh"`
	CanSlot    int      `json:"canSlot"`
	CanUpgrade int      `json:"canUpgrade"`
	RandomBuff string   `json:"randomBuff"`
}

func init() {
	if equips, err := loadEquips(); err != nil {
		log.Fatalf("load equips failed: %+v", err)
	} else {
		Equips = equips
	}
}

func loadEquips() (*equips, error) {
	root, equips := RootRomel+"equip", &equips{
		ids:   make(map[string]*Equip),
		names: make(map[string]*Equip),
	}
	if err := iterate(root, func(item map[string]interface{}, data []byte) error {
		return equips.add(item, data)
	}); err != nil {
		return nil, err
	} else {
		log.Printf("load %d equips from %s", equips.Size(), root)
		return equips, nil
	}
}

func (e *equips) add(item map[string]interface{}, data []byte) error {
	equip := &Equip{}
	if err := json.Unmarshal(data, equip); err != nil {
		return errors.WithStack(err)
	} else {
		e.ids[equip.Id] = equip
		e.names[equip.Name] = equip
		delete(item, "id")
		delete(item, "name")
		delete(item, "icon")
		delete(item, "rank")
		delete(item, "job")
		delete(item, "position")
		delete(item, "effect")
		delete(item, "buff")
		delete(item, "isCompose")
		delete(item, "isUpgrade")
		delete(item, "isHigh")
		delete(item, "canSlot")
		delete(item, "canUpgrade")
		delete(item, "randomBuff")
		if len(item) > 0 {
			for k, v := range item {
				log.Printf("unknown equip property: %s=%+v", k, v)
			}
		}
		return nil
	}
}

func (e *equips) Size() int {
	return len(e.ids)
}

func (e *equips) Get(name string) (*Equip, bool) {
	equip, found := e.names[name]
	return equip, found
}
