package evaluator

import (
	"fmt"

	"github.com/patilpankaj212/terrascan-policies/pkg/config"
)

type EvaluatorType int

const (
	TerrascanEvaluator EvaluatorType = iota
)

func GetEvaluator(e EvaluatorType, c config.PolicyDirConfig) (PolicyEvaluator, error) {

	switch e {
	case TerrascanEvaluator:
		return NewTerrascanPolicyEvaluator(c), nil
	default:
		return nil, fmt.Errorf("unsupported Policy Evaluator")
	}
}
