package calculator

import (
	"fmt"
)

type Formula struct {
	key       string
	value     float64
	algorithm string
	refs      *[]interface{}
	values    *Values
}

func (f *Formula) add(totalKey string, items ...interface{}) *Formula {
	values := []interface{}{f}
	return f.values.Add(totalKey, append(values, items...)...)
}

func (f *Formula) reduce(totalKey string, items ...interface{}) *Formula {
	values := []interface{}{f}
	return f.values.Reduce(totalKey, append(values, items...)...)
}

func (f *Formula) multiply(totalKey string, items ...interface{}) *Formula {
	values := []interface{}{f}
	return f.values.Multiply(totalKey, append(values, items...)...)
}

func (f *Formula) divide(totalKey string, items ...interface{}) *Formula {
	values := []interface{}{f}
	return f.values.Divide(totalKey, append(values, items...)...)
}

func (f *Formula) Algorithm() string {
	if f == nil {
		return "nil"
	} else if len(f.algorithm) == 0 {
		return f.String()
	} else {
		return fmt.Sprintf("%s = %s", f, f.algorithm)
	}
}

func (f *Formula) Value() float64 {
	return f.value
}

func (f *Formula) String() string {
	return fmt.Sprintf("%s[%v]", f.key, f.value)
}
