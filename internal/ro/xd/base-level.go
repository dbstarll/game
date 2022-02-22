package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	BaseLevels *baseLevels
)

type baseLevels struct {
	ids map[int]*BaseLevel
}

type BaseLevel struct {
	Id       int
	AddPoint int
}

func loadBaseLevels(data []byte) (*baseLevels, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	baseLevels := &baseLevels{
		ids: make(map[int]*BaseLevel),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := baseLevels.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d baseLevels", baseLevels.Size())
	return baseLevels, nil
}

func (r *baseLevels) add(data []byte) error {
	baseLevel := &BaseLevel{}
	if err := unmarshalDisallowUnknownFields(data, baseLevel); err != nil {
		return err
	} else {
		r.ids[baseLevel.Id] = baseLevel
		return nil
	}
}

func (r *baseLevels) Size() int {
	return len(r.ids)
}

func (r *baseLevels) Id(id int) (*BaseLevel, bool) {
	baseLevel, exist := r.ids[id]
	return baseLevel, exist
}
