package xd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	Sites *sites
)

type sites struct {
	ids   map[int]*Site
	names map[string]*Site
}

type Site struct {
	Id   int
	List []SingleSite
	Name string
}

type SingleSite struct {
	Id   int
	Name string
}

func loadSites(data []byte) (*sites, error) {
	obj := make(map[string]map[string]interface{})
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, errors.WithStack(err)
	}

	sites := &sites{
		ids:   make(map[int]*Site),
		names: make(map[string]*Site),
	}
	for _, item := range obj {
		if jsonData, err := json.Marshal(item); err != nil {
			return nil, errors.WithStack(err)
		} else if err := sites.add(jsonData); err != nil {
			return nil, err
		}
	}
	zap.S().Infof("load %d sites", sites.Size())
	return sites, nil
}

func (r *sites) add(data []byte) error {
	site := &Site{}
	if err := unmarshalDisallowUnknownFields(data, site); err != nil {
		return err
	} else {
		r.ids[site.Id] = site
		r.names[site.Name] = site
		return nil
	}
}

func (r *sites) Size() int {
	return len(r.ids)
}

func (r *sites) Id(id int) (*Site, bool) {
	site, exist := r.ids[id]
	return site, exist
}

func (r *sites) Name(name string) (*Site, bool) {
	site, exist := r.names[name]
	return site, exist
}
