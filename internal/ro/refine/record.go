package refine

type Result string

const (
	Succeed Result = "+"
	Failed  Result = "-"
	Broken  Result = "x"
)

func (r Result) increment() int {
	switch r {
	case Succeed:
		return 1
	case Failed:
		return -1
	case Broken:
		return -1
	default:
		return 0
	}
}

type Record struct {
	level  int
	result Result
}
