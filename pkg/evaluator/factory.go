package evaluator

import (
	"fmt"

	"github.com/patilpankaj212/terrascan-policies/pkg/config"
)

// EvaluatorType is type for policy evaluator types
type EvaluatorType int

const (
	TerrascanEvaluator EvaluatorType = iota
)

// GetEvaluator returns a new policy evaluator
func GetEvaluator(e EvaluatorType, c config.PolicyDirConfig) (PolicyEvaluator, error) {

	switch e {
	case TerrascanEvaluator:
		return NewTerrascanPolicyEvaluator(c), nil
	default:
		return nil, fmt.Errorf("unsupported Policy Evaluator")
	}
}
