package romel

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

const RootRomel = "configs/romel/"

func iterate(root string, fn func(item map[string]interface{}, data []byte) error) error {
	if files, err := ioutil.ReadDir(root); err != nil {
		return errors.WithStack(err)
	} else {
		for _, file := range files {
			result := &Result{}
			if data, err := ioutil.ReadFile(root + "/" + file.Name()); err != nil {
				return errors.WithStack(err)
			} else if err := json.Unmarshal(data, result); err != nil {
				return errors.WithStack(err)
			} else {
				for _, item := range result.Data.List {
					if jsonItem, err := json.Marshal(item); err != nil {
						return errors.WithStack(err)
					} else if err := fn(item, jsonItem); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
}
