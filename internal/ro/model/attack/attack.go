package attack

import (
	"github.com/dbstarll/game/internal/ro/dimension/job"
	"github.com/dbstarll/game/internal/ro/dimension/nature"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
)

type Attack struct {
	weapon   weapon.Weapon
	magic    bool
	remote   bool
	skill    bool
	critical bool
	nature   nature.Nature
}

func UseWeapon(job job.Job, weapon weapon.Weapon) *Attack {
	return &Attack{
		weapon: weapon,
		magic:  weapon.IsMagic(job),
		remote: weapon.IsRemote(job),
		nature: nature.Neutral,
	}
}

func (a *Attack) WithSkill() *Attack {
	a.skill = true
	return a
}

func (a *Attack) WithCritical() *Attack {
	a.critical = true
	return a
}

func (a *Attack) WithNature(nature nature.Nature) *Attack {
	a.nature = nature
	return a
}

func (a *Attack) IsMagic() bool {
	return a.magic
}

func (a *Attack) IsRemote() bool {
	return a.magic || a.remote
}

func (a *Attack) IsOrdinary() bool {
	return !a.skill
}

func (a *Attack) IsCritical() bool {
	return a.critical
}

func (a *Attack) GetNature() nature.Nature {
	return a.nature
}

func (a *Attack) GetWeapon() weapon.Weapon {
	return a.weapon
}
