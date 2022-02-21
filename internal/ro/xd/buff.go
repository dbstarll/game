package xd

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Buffs *buffs
)

type buffs struct {
	ids map[int]*Buff
}

type Buff struct {
	Id         int
	BuffName   string
	BuffDesc   string
	Dsc        string
	IsOverlay  Bool
	Condition  *Condition
	BuffEffect *BuffEffect
	BuffTarget *BuffTarget
}

type Condition struct {
	Type  string `json:"type"`
	items *map[string]interface{}
}

type BuffEffect struct {
	Type  string `json:"type"`
	items *map[string]interface{}
}

type BuffTarget struct {
	buffTarget
}

type buffTarget struct {
	Type           int           `json:"type"`
	Distance       *int          `json:"distance"`
	DistanceCount  *int          `json:"distance_count"`
	DynamicRange   *DynamicRange `json:"dynamic_range"`
	EntryType      *int          `json:"entry_type"`
	InnerRange     *float64      `json:"inner_range"`
	MinhpperTarget *Bool         `json:"minhpper_target"`
	NoSelf         *Bool         `json:"no_self"`
	NumLimit       *int          `json:"num_limit"`
	Range          *float64      `json:"range"`
	RecRange       *RecRange     `json:"rec_range"`
	TeamCount      *Bool         `json:"team_count"`
}

type RecRange struct {
	Length float64 `json:"length"`
	Width  int     `json:"width"`
}

type DynamicRange struct {
	Type int      `json:"type"`
	A    int      `json:"a"`
	B    *float64 `json:"b"`
	C    *int     `json:"c"`
}

func loadBuffs(data []byte) (*buffs, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	buffs := &buffs{ids: make(map[int]*Buff)}
	for _, val := range obj {
		if jsonData, err := json.Marshal(val); err != nil {
			return nil, errors.WithStack(err)
		} else if err := buffs.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d buffs", buffs.Size())
	return buffs, nil
}

func (b *buffs) add(data []byte) error {
	buff := &Buff{}
	if err := unmarshalDisallowUnknownFields(data, buff); err != nil {
		return err
	} else {
		if buff.BuffTarget.Type == 0 {
			buff.BuffTarget = nil
		}
		if buff.BuffEffect.Type == "" {
			buff.BuffTarget = nil
		}
		if buff.Condition.Type == "" {
			buff.Condition = nil
		}
		b.ids[buff.Id] = buff
		return nil
	}
}

func (b *buffs) Size() int {
	return len(b.ids)
}

func (c *Condition) UnmarshalJSON(data []byte) error {
	if !bytes.Equal(data, EmptyArray) {
		item := make(map[string]interface{})
		if err := json.Unmarshal(data, &item); err != nil {
			return errors.WithStack(err)
		} else if Type, ok := item["type"]; ok {
			if typeStr, ok := Type.(string); ok {
				c.Type = typeStr
				delete(item, "type")
				c.items = &item
			}
		}
	}
	return nil
}

func (b *BuffEffect) UnmarshalJSON(data []byte) error {
	if !bytes.Equal(data, EmptyArray) {
		item := make(map[string]interface{})
		if err := json.Unmarshal(data, &item); err != nil {
			return errors.WithStack(err)
		} else if Type, ok := item["type"]; ok {
			if typeStr, ok := Type.(string); ok {
				b.Type = typeStr
				delete(item, "type")
				b.items = &item
			}
		} else if Type, ok := item["ype"]; ok {
			if typeStr, ok := Type.(string); ok {
				b.Type = typeStr
				delete(item, "ype")
				b.items = &item
			}
		} else if Type, ok := item["type2"]; ok {
			if typeStr, ok := Type.(string); ok {
				b.Type = typeStr
				delete(item, "type2")
				b.items = &item
			}
		}
	}
	return nil
}

func (b *BuffTarget) UnmarshalJSON(data []byte) error {
	if !bytes.Equal(data, EmptyArray) {
		if err := unmarshalDisallowUnknownFields(data, &b.buffTarget); err != nil {
			return err
		}
	}
	return nil
}
