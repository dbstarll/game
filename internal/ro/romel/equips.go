package romel

import (
	"encoding/json"
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/position"
	"github.com/dbstarll/game/internal/ro/dimension/quality"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var Equips *equips

type equips struct {
	ids   map[string]*Equip
	names map[string]*Equip
}

type Equip struct {
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Rank       quality.Quality   `json:"rank"`
	Job        *[]job.Job        `json:"job"`
	Position   position.Position `json:"position"`
	Effect     string            `json:"effect"`
	Buff       string            `json:"buff"`
	IsCompose  int               `json:"isCompose"`
	IsUpgrade  int               `json:"isUpgrade"`
	IsHigh     int               `json:"isHigh"`
	CanSlot    int               `json:"canSlot"`
	CanUpgrade int               `json:"canUpgrade"`
	RandomBuff string            `json:"randomBuff"`
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

func (e *equips) Get(name string) *Equip {
	return e.names[name]
}

func (e *equips) Filter(fn func(*Equip) error, filterFn ...func(filter *Equip)) (int, error) {
	count, filters := 0, make([]*Equip, len(filterFn))
	for idx, f := range filterFn {
		filters[idx] = &Equip{IsCompose: -1, IsUpgrade: -1, IsHigh: -1, CanSlot: -1, CanUpgrade: -1}
		f(filters[idx])
	}
	for _, equip := range e.ids {
		if equip.matchAny(filters...) {
			count++
			if err := fn(equip); err != nil {
				return 0, err
			}
		}
	}
	return count, nil
}

func (e *Equip) matchAny(filters ...*Equip) bool {
	for _, filter := range filters {
		if e.match(filter) {
			return true
		}
	}
	return len(filters) == 0
}

func (e *Equip) match(filter *Equip) bool {
	if filter.Rank > quality.Unlimited && filter.Rank != e.Rank {
		return false
	} else if filter.Position > position.Unlimited && filter.Position != e.Position {
		return false
	} else if filter.IsCompose >= 0 && filter.IsCompose != e.IsCompose {
		return false
	} else if filter.IsHigh >= 0 && filter.IsHigh != e.IsHigh {
		return false
	} else if filter.IsUpgrade >= 0 && filter.IsUpgrade != e.IsUpgrade {
		return false
	} else if filter.CanSlot >= 0 && filter.CanSlot != e.CanSlot {
		return false
	} else if filter.CanUpgrade >= 0 && filter.CanUpgrade != e.CanUpgrade {
		return false
	} else if len(filter.Name) > 0 && strings.Index(e.Name, filter.Name) < 0 {
		return false
	} else if len(filter.Buff) > 0 && strings.Index(e.Buff, filter.Buff) < 0 {
		return false
	} else if len(filter.RandomBuff) > 0 && strings.Index(e.RandomBuff, filter.RandomBuff) < 0 {
		return false
	} else if len(filter.Effect) > 0 && strings.Index(e.Effect, filter.Effect) < 0 {
		return false
	} else if filter.Job != nil && !e.matchAnyJob(filter.Job) {
		return false
	} else {
		return true
	}
}

func (e *Equip) matchAnyJob(filters *[]job.Job) bool {
	for _, filter := range *filters {
		if filter == job.Unlimited {
			return true
		} else {
			for _, j := range *e.Job {
				if j == job.Unlimited || j == filter {
					return true
				}
			}
		}
	}
	return false
}
