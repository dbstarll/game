package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Roles *roles
)

type roles struct {
	ids   map[int]*Role
	vars  map[string]*Role
	names map[string]*Role
}

type Role struct {
	Id        int
	Base      int
	IsPercent Bool
	PropName  string
	VarName   string
}

func loadRoles(data []byte) (*roles, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	roles := &roles{
		ids:   make(map[int]*Role),
		vars:  make(map[string]*Role),
		names: make(map[string]*Role),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := roles.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d roles", roles.Size())
	return roles, nil
}

func (r *roles) add(data []byte) error {
	role := &Role{}
	if err := unmarshalDisallowUnknownFields(data, role); err != nil {
		return err
	} else {
		r.ids[role.Id] = role
		r.vars[role.VarName] = role
		r.names[role.PropName] = role
		return nil
	}
}

func (r *roles) Size() int {
	return len(r.ids)
}

func (r *roles) Id(id int) (*Role, bool) {
	role, exist := r.ids[id]
	return role, exist
}

func (r *roles) PropName(name string) (*Role, bool) {
	role, exist := r.names[name]
	return role, exist
}

func (r *roles) VarName(name string) (*Role, bool) {
	role, exist := r.vars[name]
	return role, exist
}
