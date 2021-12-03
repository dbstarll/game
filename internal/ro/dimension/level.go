package dimension

//等级
type Level struct {
	Base int
	Job  int
}

func (l *Level) Add(incr *Level) {
	if incr != nil {
		l.Base += incr.Base
		l.Job += incr.Job
	}
}
