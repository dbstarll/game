package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

var Pets *pets

type pets struct {
	ids   map[string]*Pet
	names map[string]*Pet
}

type Pet struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Intro         string  `json:"intro"`
	Race          int     `json:"race"`
	Nature        int     `json:"nature"`
	Size          int     `json:"size"`
	Star          int     `json:"star"`
	AccessWay     int     `json:"accessway"`
	Hobby         []Hobby `json:"hobby"`
	Equip         []Hobby `json:"equip"`
	Skill         []Skill `json:"skill"`
	Cost          []Cost  `json:"cost"`
	AdventureBuff string  `json:"adventureBuff"`
}

type Hobby struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Skill struct {
	Name  string `json:"name"`
	Intro string `json:"intro"`
	Lv    int    `json:"lv"`
}

type Cost struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	if pets, err := loadPets(); err != nil {
		log.Fatalf("load pets failed: %+v", err)
	} else {
		Pets = pets
	}
}

func loadPets() (*pets, error) {
	root, pets := RootRomel+"pet", &pets{
		ids:   make(map[string]*Pet),
		names: make(map[string]*Pet),
	}
	if err := iterate(root, func(item map[string]interface{}, data []byte) error {
		return pets.add(item, data)
	}); err != nil {
		return nil, err
	} else {
		log.Printf("load %d pets from %s", pets.Size(), root)
		return pets, nil
	}
}

func (e *pets) add(item map[string]interface{}, data []byte) error {
	pet := &Pet{}
	if err := json.Unmarshal(data, pet); err != nil {
		return errors.WithStack(err)
	} else {
		e.ids[pet.Id] = pet
		e.names[pet.Name] = pet
		delete(item, "id")
		delete(item, "name")
		delete(item, "avatar")
		delete(item, "intro")
		delete(item, "race")
		delete(item, "nature")
		delete(item, "size")
		delete(item, "star")
		delete(item, "accessway")
		delete(item, "hobby")
		delete(item, "equip")
		delete(item, "skill")
		delete(item, "adventureBuff")
		delete(item, "cost")
		if len(item) > 0 {
			for k, v := range item {
				log.Printf("unknown pet property: %s=%+v", k, v)
			}
		}
		return nil
	}
}

func (e *pets) Size() int {
	return len(e.ids)
}

func (e *pets) Get(name string) *Pet {
	return e.names[name]
}
