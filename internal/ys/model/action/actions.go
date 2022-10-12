package action

import (
	"fmt"
	"github.com/dbstarll/game/internal/ys/dimension/attackMode"
	"strings"
)

type Actions struct {
	actionList []*Action
	actionMap  map[attackMode.AttackMode]map[string]*Action
}

type Iterator func(index int, action *Action) bool

func NewActions() *Actions {
	return &Actions{
		actionList: make([]*Action, 0),
		actionMap:  make(map[attackMode.AttackMode]map[string]*Action),
	}
}

func (a *Actions) Add(action *Action) {
	if action != nil {
		if oldMap, exist := a.actionMap[action.mode]; !exist {
			a.actionList = append(a.actionList, action)
			a.actionMap[action.mode] = map[string]*Action{action.name: action}
		} else if oldAction, exist := oldMap[action.name]; !exist {
			a.actionList = append(a.actionList, action)
			oldMap[action.name] = action
		} else {
			oldAction.dmg = action.dmg
			oldAction.elemental = action.elemental
		}
	}
}

func (a *Actions) AddAll(other *Actions) {
	if other != nil {
		for _, action := range other.actionList {
			a.Add(action)
		}
	}
}

func (a *Actions) Get(mode attackMode.AttackMode, name string) *Action {
	if nameMaps, exist := a.actionMap[mode]; exist {
		if action, exist := nameMaps[name]; exist {
			return action
		}
		for key, action := range nameMaps {
			if strings.HasSuffix(key, name) {
				return action
			}
		}
	}
	return nil
}

func (a *Actions) Loop(iterator Iterator) {
	for idx, action := range a.actionList {
		if iterator(idx, action) {
			break
		}
	}
}

func (a *Actions) String() string {
	return fmt.Sprintf("%+v", a.actionMap)
}
