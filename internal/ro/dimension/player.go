package dimension

import "github.com/dbstarll/game/internal/ro/dimension/job"

type Player struct {
	Character
	_job job.Job
}
