package model

var (
	AddOperator      Operator = &addOperator{}
	MultiplyOperator Operator = &multiplyOperator{}
	ReduceOperator   Operator = &reduceOperator{}
)

type Operator interface {
	operate(first, second float64) float64
	separator() string
}

type addOperator struct {
}

func (o *addOperator) operate(first, second float64) float64 {
	return first + second
}

func (o *addOperator) separator() string {
	return " + "
}

type multiplyOperator struct {
}

func (o *multiplyOperator) operate(first, second float64) float64 {
	return first * second
}

func (o *multiplyOperator) separator() string {
	return " * "
}

type reduceOperator struct {
}

func (o *reduceOperator) operate(first, second float64) float64 {
	return first - second
}

func (o *reduceOperator) separator() string {
	return " - "
}
