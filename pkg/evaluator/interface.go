package evaluator

// PolicyEvaluator defines methods that a policy evaluator will implement
type PolicyEvaluator interface {
	EvalPositive() error
	EvalNegative() error
}
