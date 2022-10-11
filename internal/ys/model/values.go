package model

import (
	"fmt"
	"go.uber.org/zap"
	"math"
	"reflect"
	"strings"
)

type Values struct {
	values map[string]*Formula
}

func NewValues() *Values {
	return &Values{values: make(map[string]*Formula)}
}

func (v *Values) set(formula *Formula) *Formula {
	v.values[formula.key] = formula
	formula.values = v
	formula.value = math.Round(formula.value*10000000) / 10000000
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

func (v *Values) getAll(items ...interface{}) []interface{} {
	values := make([]interface{}, len(items))
	for idx, item := range items {
		if formula, ok := item.(*Formula); ok {
			values[idx] = formula
		} else if key, ok := item.(string); !ok {
			values[idx] = item
		} else if value, exist := v.Get(key); exist && value != nil {
			values[idx] = value
		} else {
			values[idx] = v.Set(key, 0)
		}
	}
	return values
}

func (v *Values) Add(totalKey string, items ...interface{}) *Formula {
	return v.multiOperate(totalKey, AddOperator, items...)
}

func (v *Values) Multiply(totalKey string, items ...interface{}) *Formula {
	return v.multiOperate(totalKey, MultiplyOperator, items...)
}

func (v *Values) multiOperate(totalKey string, operator Operator, items ...interface{}) *Formula {
	items = v.getAll(items...)
	totalValue, initialized, algorithms := 0.0, false, make([]string, len(items))
	for idx, item := range items {
		if value, algorithm, ok := v.parseItem(item); !ok {
			algorithms[idx] = algorithm
		} else if !initialized {
			totalValue = value
			algorithms[idx] = algorithm
			initialized = true
		} else {
			totalValue = operator.operate(totalValue, value)
			algorithms[idx] = algorithm
		}
	}
	return v.set(&Formula{
		key:       totalKey,
		value:     totalValue,
		algorithm: strings.Join(algorithms, operator.separator()),
		refs:      &items,
	})
}

func (v *Values) parseItem(item interface{}) (float64, string, bool) {
	if formula, ok := item.(*Formula); ok {
		return formula.value, formula.String(), true
	} else if floatValue, ok := item.(float64); ok {
		return floatValue, fmt.Sprintf("%v", floatValue), true
	} else if intValue, ok := item.(int); ok {
		return float64(intValue), fmt.Sprintf("%d", intValue), true
	} else {
		zap.S().Warnf("unknown item: [%s]%+v", reflect.TypeOf(item), item)
		return 0, fmt.Sprintf("%s", item), false
	}
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
