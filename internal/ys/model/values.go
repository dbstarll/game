package model

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type Values struct {
	values map[string]*Formula
}

func NewValues() *Values {
	return &Values{values: make(map[string]*Formula)}
}

type Formula struct {
	key       string
	value     float64
	algorithm string
	refs      *[]*Formula
}

func (f *Formula) String() string {
	if len(f.algorithm) == 0 {
		return fmt.Sprintf("%s[%v]", f.key, f.value)
	} else {
		return fmt.Sprintf("%s[%v] = %s", f.key, f.value, f.algorithm)
	}
}

func (v *Values) set(formula *Formula) *Formula {
	v.values[formula.key] = formula
	return formula
}

func (v *Values) Set(key string, value float64) *Formula {
	return v.set(&Formula{
		key:       key,
		value:     value,
		algorithm: "",
	})
}

func (v *Values) Get(key string) (*Formula, bool) {
	if value, exist := v.values[key]; exist {
		return value, true
	} else {
		zap.S().Warnf("key not exist: %s", key)
		return nil, false
	}
}

func (v *Values) Add(totalKey string, keys ...string) *Formula {
	totalValue, values, algorithms := 0.0, make([]*Formula, len(keys)), make([]string, len(keys))
	for idx, key := range keys {
		if value, exist := v.Get(key); exist && value != nil {
			values[idx] = value
			totalValue += value.value
		} else {
			values[idx] = v.Set(key, 0)
		}
		algorithms[idx] = fmt.Sprintf("%s[%v]", values[idx].key, values[idx].value)
	}
	return v.set(&Formula{
		key:       totalKey,
		value:     totalValue,
		algorithm: strings.Join(algorithms, " + "),
		refs:      &values,
	})
}

func (v *Values) String() string {
	values := make(map[string]float64)
	for key, val := range v.values {
		if val.value != 0 {
			values[key] = val.value
		}
	}
	return fmt.Sprintf("%v", values)
}
