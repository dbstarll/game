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
	values    *Values
}

func (f *Formula) add(totalKey string, objs ...interface{}) *Formula {
	values := []*Formula{f}
	return f.values.add(totalKey, append(values, f.values.GetAll(objs...)...)...)
}

func (f *Formula) multiply(totalKey string, objs ...interface{}) *Formula {
	values := []*Formula{f}
	return f.values.multiply(totalKey, append(values, f.values.GetAll(objs...)...)...)
}

func (f *Formula) Algorithm() string {
	if len(f.algorithm) == 0 {
		return f.String()
	} else {
		return fmt.Sprintf("%s = %s", f, f.algorithm)
	}
}

func (f *Formula) String() string {
	return fmt.Sprintf("%s[%v]", f.key, f.value)
}

func (v *Values) set(formula *Formula) *Formula {
	v.values[formula.key] = formula
	formula.values = v
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

//func (v *Values) GetAll(keys ...string) []*Formula {
//	values := make([]*Formula, len(keys))
//	for idx, key := range keys {
//		if value, exist := v.Get(key); exist && value != nil {
//			values[idx] = value
//		} else {
//			values[idx] = v.Set(key, 0)
//		}
//	}
//	return values
//}

func (v *Values) GetAll(objs ...interface{}) []*Formula {
	values := make([]*Formula, 0)
	for _, obj := range objs {
		if formula, ok := obj.(*Formula); ok {
			values = append(values, formula)
		} else if key, ok := obj.(string); !ok {
			// ignore
		} else if value, exist := v.Get(key); exist && value != nil {
			values = append(values, value)
		} else {
			values = append(values, v.Set(key, 0))
		}
	}
	return values
}

func (v *Values) Add(totalKey string, objs ...interface{}) *Formula {
	return v.add(totalKey, v.GetAll(objs...)...)
}

func (v *Values) Multiply(totalKey string, objs ...interface{}) *Formula {
	return v.multiply(totalKey, v.GetAll(objs...)...)
}

func (v *Values) add(totalKey string, items ...*Formula) *Formula {
	totalValue, algorithms := 0.0, make([]string, len(items))
	for idx, item := range items {
		totalValue += item.value
		algorithms[idx] = item.String()
	}
	return v.set(&Formula{
		key:       totalKey,
		value:     totalValue,
		algorithm: strings.Join(algorithms, " + "),
		refs:      &items,
	})
}

func (v *Values) multiply(totalKey string, items ...*Formula) *Formula {
	totalValue, algorithms := 1.0, make([]string, len(items))
	for idx, item := range items {
		totalValue *= item.value
		algorithms[idx] = item.String()
	}
	return v.set(&Formula{
		key:       totalKey,
		value:     totalValue,
		algorithm: strings.Join(algorithms, " * "),
		refs:      &items,
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
