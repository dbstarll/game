package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

var Monsters *monsters

type monsters struct {
	ids   map[string]*Monster
	names map[string]*Monster
}

type Monster struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Race      int    `json:"race"`
	Nature    int    `json:"nature"`
	Size      int    `json:"size"`
	Level     int    `json:"level"`
	Hp        int    `json:"hp"`
	BaseExp   int    `json:"baseExp"`
	JobExp    int    `json:"jobExp"`
	IsStar    int    `json:"isStar"`
	IsDead    int    `json:"isDead"`
	IsSpecial int    `json:"isSpecial"`
}

func init() {
	if monsters, err := loadMonsters(); err != nil {
		log.Fatalf("load monsters failed: %+v", err)
	} else {
		Monsters = monsters
	}
}

func loadMonsters() (*monsters, error) {
	root, monsters := RootRomel+"monster", &monsters{
		ids:   make(map[string]*Monster),
		names: make(map[string]*Monster),
	}
	if err := iterate(root, func(item map[string]interface{}, data []byte) error {
		return monsters.add(item, data)
	}); err != nil {
		return nil, err
	} else {
		log.Printf("load %d monsters from %s", monsters.Size(), root)
		return monsters, nil
	}
}

func (e *monsters) add(item map[string]interface{}, data []byte) error {
	monster := &Monster{}
	if err := json.Unmarshal(data, monster); err != nil {
		return errors.WithStack(err)
	} else {
		e.ids[monster.Id] = monster
		e.names[monster.Name] = monster
		delete(item, "id")
		delete(item, "name")
		delete(item, "avatar")
		delete(item, "type")
		delete(item, "race")
		delete(item, "nature")
		delete(item, "size")
		delete(item, "level")
		delete(item, "hp")
		delete(item, "baseExp")
		delete(item, "jobExp")
		delete(item, "isStar")
		delete(item, "isDead")
		delete(item, "isSpecial")
		if len(item) > 0 {
			for k, v := range item {
				log.Printf("unknown monster property: %s=%+v", k, v)
			}
		}
		return nil
	}
}

func (e *monsters) Size() int {
	return len(e.ids)
}

func (e *monsters) Get(name string) *Monster {
	return e.names[name]
}
