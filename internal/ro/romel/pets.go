package romel

import (
	"encoding/json"
	"github.com/dbstarll/game/internal/ro/dimension/accessWay"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var Pets *pets

type pets struct {
	ids   map[string]*Pet
	names map[string]*Pet
}

type Pet struct {
	Id            string              `json:"id"`
	Name          string              `json:"name"`
	Intro         string              `json:"intro"`
	Race          int                 `json:"race"`
	Nature        int                 `json:"nature"`
	Size          int                 `json:"size"`
	Star          int                 `json:"star"`
	AccessWay     accessWay.AccessWay `json:"accessway"`
	Hobby         *[]Hobby            `json:"hobby"`
	Equip         *[]Hobby            `json:"equip"`
	Skill         *[]Skill            `json:"skill"`
	Cost          *[]Cost             `json:"cost"`
	AdventureBuff string              `json:"adventureBuff"`
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

func (p *pets) add(item map[string]interface{}, data []byte) error {
	pet := &Pet{}
	if err := json.Unmarshal(data, pet); err != nil {
		return errors.WithStack(err)
	} else {
		p.ids[pet.Id] = pet
		p.names[pet.Name] = pet
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

func (p *pets) Size() int {
	return len(p.ids)
}

func (p *pets) Get(name string) *Pet {
	return p.names[name]
}

func (p *pets) Filter(fn func(*Pet) error, filterFn ...func(filter *Pet)) (int, error) {
	count, filters := 0, make([]*Pet, len(filterFn))
	for idx, f := range filterFn {
		filters[idx] = &Pet{AccessWay: -1}
		f(filters[idx])
	}
	for _, pet := range p.ids {
		if pet.matchAny(filters...) {
			count++
			if err := fn(pet); err != nil {
				return 0, err
			}
		}
	}
	return count, nil
}

func (p *Pet) matchAny(filters ...*Pet) bool {
	for _, filter := range filters {
		if p.match(filter) {
			return true
		}
	}
	return len(filters) == 0
}

func (p *Pet) match(filter *Pet) bool {
	if filter.Race > 0 && filter.Race != p.Race {
		return false
	} else if filter.Nature > 0 && filter.Nature != p.Nature {
		return false
	} else if filter.Size > 0 && filter.Size != p.Size {
		return false
	} else if filter.Star > 0 && filter.Star != p.Star {
		return false
	} else if filter.AccessWay >= 0 && filter.AccessWay != p.AccessWay {
		return false
	} else if len(filter.Name) > 0 && strings.Index(p.Name, filter.Name) < 0 {
		return false
	} else if len(filter.AdventureBuff) > 0 && strings.Index(p.AdventureBuff, filter.AdventureBuff) < 0 {
		return false
	} else if filter.Cost != nil && !p.matchAllCost(filter.Cost) {
		return false
	} else if filter.Skill != nil && !p.matchAllSkill(filter.Skill) {
		return false
	} else {
		return true
	}
}

func (p *Pet) matchAllCost(filters *[]Cost) bool {
	for _, filter := range *filters {
		if len(filter.Name) > 0 {
			if p.Cost == nil {
				return false
			} else {
				match := false
				for _, cost := range *p.Cost {
					if strings.Index(cost.Name, filter.Name) >= 0 {
						match = true
						break
					}
				}
				if !match {
					return false
				}
			}
		}
	}
	return true
}

func (p *Pet) matchAllSkill(filters *[]Skill) bool {
	if p.Skill == nil {
		return false
	} else {
		for _, skill := range *p.Skill {
			if skill.matchAll(filters) {
				return true
			}
		}
		return false
	}
}

func (s *Skill) matchAll(filters *[]Skill) bool {
	for _, filter := range *filters {
		if len(filter.Intro) > 0 {
			if strings.Index(s.Intro, filter.Intro) < 0 {
				return false
			}
		}
	}
	return true
}
