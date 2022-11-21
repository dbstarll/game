package artifacts

import "github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"

type FloatEntries []struct {
	Entry entry.Entry
	Value float64
}

type IntEntries []struct {
	Entry entry.Entry
	Value int
}

func (e *FloatEntries) LoopEntries(looper func(entry entry.Entry, value interface{}) error) error {
	for _, item := range *e {
		if err := looper(item.Entry, item.Value); err != nil {
			return err
		}
	}
	return nil
}

func (e *IntEntries) LoopEntries(looper func(entry entry.Entry, value interface{}) error) error {
	for _, item := range *e {
		if err := looper(item.Entry, item.Value); err != nil {
			return err
		}
	}
	return nil
}
