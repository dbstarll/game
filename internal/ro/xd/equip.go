package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Equips *equips
)

type equips struct {
	ids   map[int]*Equip
	names map[string]*Equip
}

type Equip struct {
	Id            int
	CanEnchant    bool
	CanEquip      []int
	CanRefine     bool
	CanStrength   bool
	Effect        *Effects
	EffectAdd     *Effects
	EquipType     int
	Icon          string
	MaxCardSlot   int
	NameEn        string
	NameZh        string
	RefineBuff    []RefineBuff
	RefineEffect  *Effects
	RefineEffect2 *Effects
	RefineMaxlv   int
	SuitID        []int
	Type          string
	TypeId        int
	UniqueEffect  *BuffEffects
	Upgrade       Upgrade
}

type RefineBuff struct {
	Buff []int `json:"buff"`
	Lv   int   `json:"lv"`
}

type Upgrade struct {
	BuffID  []int
	Product int
}

func loadEquips(data []byte) (*equips, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	equips := &equips{
		ids:   make(map[int]*Equip),
		names: make(map[string]*Equip),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := equips.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d equips", equips.Size())
	return equips, nil
}

func (r *equips) add(data []byte) error {
	equip := &Equip{}
	if err := unmarshalDisallowUnknownFields(data, equip); err != nil {
		return err
	} else {
		if equip.Effect.items == nil {
			equip.Effect = nil
		}
		if equip.EffectAdd.items == nil {
			equip.EffectAdd = nil
		}
		if equip.RefineEffect.items == nil {
			equip.RefineEffect = nil
		}
		if equip.RefineEffect2.items == nil {
			equip.RefineEffect2 = nil
		}
		if len(equip.UniqueEffect.Buff) == 0 {
			equip.UniqueEffect = nil
		}
		r.ids[equip.Id] = equip
		r.names[equip.NameZh] = equip
		return nil
	}
}

func (r *equips) Size() int {
	return len(r.ids)
}

func (r *equips) Id(id int) (*Equip, bool) {
	equip, exist := r.ids[id]
	return equip, exist
}

func (r *equips) Name(name string) (*Equip, bool) {
	equip, exist := r.names[name]
	return equip, exist
}
