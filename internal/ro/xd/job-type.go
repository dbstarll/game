package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	JobTypes *jobTypes
)

type jobTypes struct {
	ids   map[int]*JobType
	names map[string]*JobType
}

type JobType struct {
	Id   int
	Jobs []int
	Name string
}

func loadJobTypes(data []byte) (*jobTypes, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	jobTypes := &jobTypes{
		ids:   make(map[int]*JobType),
		names: make(map[string]*JobType),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := jobTypes.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d jobTypes", jobTypes.Size())
	return jobTypes, nil
}

func (r *jobTypes) add(data []byte) error {
	jobType := &JobType{}
	if err := unmarshalDisallowUnknownFields(data, jobType); err != nil {
		return err
	} else {
		r.ids[jobType.Id] = jobType
		r.names[jobType.Name] = jobType
		return nil
	}
}

func (r *jobTypes) Size() int {
	return len(r.ids)
}

func (r *jobTypes) Id(id int) (*JobType, bool) {
	jobType, exist := r.ids[id]
	return jobType, exist
}

func (r *jobTypes) Name(name string) (*JobType, bool) {
	jobType, exist := r.names[name]
	return jobType, exist
}
