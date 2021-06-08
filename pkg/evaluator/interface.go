package evaluator

type PolicyEvaluator interface {
	EvalPositive() error
	EvalNegative() error
}
