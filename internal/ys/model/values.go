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
	totalValue, algorithms := operator.base(), make([]string, len(items))
	for idx, item := range items {
		if formula, ok := item.(*Formula); ok {
			totalValue = operator.operate(totalValue, formula.value)
			algorithms[idx] = formula.String()
		} else if floatValue, ok := item.(float64); ok {
			totalValue = operator.operate(totalValue, floatValue)
			algorithms[idx] = fmt.Sprintf("%v", floatValue)
		} else if intValue, ok := item.(int); ok {
			totalValue = operator.operate(totalValue, float64(intValue))
			algorithms[idx] = fmt.Sprintf("%d", intValue)
		} else {
			zap.S().Warnf("unknown item: [%s]%+v", reflect.TypeOf(item), item)
			algorithms[idx] = fmt.Sprintf("%s", item)
		}
	}
	return v.set(&Formula{
		key:       totalKey,
		value:     totalValue,
		algorithm: strings.Join(algorithms, operator.separator()),
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
