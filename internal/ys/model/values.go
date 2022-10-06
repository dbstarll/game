package model

import "log"

type Values struct {
	values map[string]float64
}

func NewValues() *Values {
	return &Values{values: make(map[string]float64)}
}

func (v *Values) Set(key string, value float64) (float64, bool) {
	if oldValue, exist := v.values[key]; exist {
		v.values[key] = value
		return oldValue, true
	} else {
		v.values[key] = value
		return 0, false
	}
}

func (v *Values) Get(key string) float64 {
	if value, exist := v.values[key]; exist {
		return value
	} else {
		log.Printf("key not exist: %s\n", key)
		return 0
	}
}
