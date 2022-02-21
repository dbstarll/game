package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Jobs *jobs
)

type jobs struct {
	ids   map[int]*Job
	names map[string]*Job
}

type Job struct {
	Id                 int
	NameEn             string
	NameZh             string
	MinBaseLevel       int
	MaxBaseLevel       int
	MinJobLevel        int
	MaxJobLevel        int
	PeakPoints         int
	PrevJobId          int
	PrevStageJobPoints int
	Skill              []int
	Type               int
	TypeBranch         int
	Stage              int
	ExtraPoints        int
}

func loadJobs(data []byte) (*jobs, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	jobs := &jobs{
		ids:   make(map[int]*Job),
		names: make(map[string]*Job),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := jobs.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d jobs", jobs.Size())
	return jobs, nil
}

func (r *jobs) add(data []byte) error {
	job := &Job{}
	if err := unmarshalDisallowUnknownFields(data, job); err != nil {
		return err
	} else {
		r.ids[job.Id] = job
		r.names[job.NameZh] = job
		return nil
	}
}

func (r *jobs) Size() int {
	return len(r.ids)
}

func (r *jobs) Id(id int) (*Job, bool) {
	job, exist := r.ids[id]
	return job, exist
}

func (r *jobs) Name(name string) (*Job, bool) {
	job, exist := r.names[name]
	return job, exist
}
