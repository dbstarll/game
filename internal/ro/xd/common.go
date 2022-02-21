package xd

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
)

var EmptyArray = []byte("[]")

type BuffEffects struct {
	buffEffects
}
type buffEffects struct {
	Buff []int `json:"buff"`
}

func (b *BuffEffects) UnmarshalJSON(data []byte) error {
	if !bytes.Equal(data, EmptyArray) {
		if err := unmarshalDisallowUnknownFields(data, &b.buffEffects); err != nil {
			return err
		}
	}
	return nil
}

type Effects struct {
	effects
}

type effects struct {
	items *map[string]float64
}

func (e *Effects) UnmarshalJSON(data []byte) error {
	if !bytes.Equal(data, EmptyArray) {
		item := make(map[string]float64)
		if err := json.Unmarshal(data, &item); err != nil {
			return errors.WithStack(err)
		} else {
			e.items = &item
		}
	}
	return nil
}
