package model

var (
	AddOperator      = &addOperator{}
	MultiplyOperator = &multiplyOperator{}
)

type Operator interface {
	base() float64
	operate(one, two float64) float64
	separator() string
}

type addOperator struct {
}

func (o *addOperator) base() float64 {
	return 0
}

func (o *addOperator) operate(one, two float64) float64 {
	return one + two
}

func (o *addOperator) separator() string {
	return " + "
}

type multiplyOperator struct {
}

func (o *multiplyOperator) base() float64 {
	return 1
}

func (o *multiplyOperator) operate(one, two float64) float64 {
	return one * two
}

func (o *multiplyOperator) separator() string {
	return " * "
}
