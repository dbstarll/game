package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Cards *cards
)

type cards struct {
	ids   map[int]*Card
	names map[string]*Card
}

type Card struct {
	Id         int
	Quality    int
	Name       string
	Monster    string
	Position   []int
	BuffEffect CardBuffEffect
}

type CardBuffEffect struct {
	Buff []int `json:"buff"`
}

func loadCards(data []byte) (*cards, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	cards := &cards{
		ids:   make(map[int]*Card),
		names: make(map[string]*Card),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := cards.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d cards", cards.Size())
	return cards, nil
}

func (r *cards) add(data []byte) error {
	card := &Card{}
	if err := unmarshalDisallowUnknownFields(data, card); err != nil {
		return err
	} else {
		r.ids[card.Id] = card
		r.names[card.Name] = card
		return nil
	}
}

func (r *cards) Size() int {
	return len(r.ids)
}

func (r *cards) Id(id int) (*Card, bool) {
	card, exist := r.ids[id]
	return card, exist
}

func (r *cards) PropName(name string) (*Card, bool) {
	card, exist := r.names[name]
	return card, exist
}
