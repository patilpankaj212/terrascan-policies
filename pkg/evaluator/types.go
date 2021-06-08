package evaluator

import (
	"fmt"

	"github.com/accurics/terrascan/pkg/runtime"
	"github.com/patilpankaj212/terrascan-policies/pkg/config"
)

type TerrascanPolicyEvaluator struct {
	config.PolicyDirConfig
}

func NewTerrascanPolicyEvaluator(c config.PolicyDirConfig) *TerrascanPolicyEvaluator {
	return &TerrascanPolicyEvaluator{c}
}

func (t *TerrascanPolicyEvaluator) EvalPositive() error {

	if t.PositiveIacFilePath == "" {
		return fmt.Errorf("positive iac file should be present for evaluation")
	}

	executor, err := runtime.NewExecutor("", "", []string{}, t.PositiveIacFilePath, "", []string{t.DirPath}, nil, nil, nil, "", false)
	if err != nil {
		return err
	}

	output, err := executor.Execute()
	if err != nil {
		return err
	}

	if len(output.Violations.Violations) < 1 {
		return fmt.Errorf("violations should be present for the positive iac file: '%s'", t.PositiveIacFilePath)
	}

	for _, metadata := range t.RuleMetadataList {
		if isPresent := t.isRuleIDPresent(metadata.ReferenceID, output); !isPresent {
			return fmt.Errorf("ruleID '%s' should be present in the violations", metadata.ReferenceID)
		}
	}
	return nil
}

func (t *TerrascanPolicyEvaluator) EvalNegative() error {
	if t.NegativeIacFilePath == "" {
		return fmt.Errorf("negative iac file should be present for evaluation")
	}

	executor, err := runtime.NewExecutor("", "", []string{}, t.NegativeIacFilePath, "", []string{t.DirPath}, nil, nil, nil, "", false)
	if err != nil {
		return err
	}

	output, err := executor.Execute()
	if err != nil {
		return err
	}

	if len(output.Violations.Violations) > 0 {
		return fmt.Errorf("no violations should be present for negative iac file: '%s'", t.NegativeIacFilePath)
	}

	return nil
}

func (t *TerrascanPolicyEvaluator) isRuleIDPresent(ruleID string, o runtime.Output) bool {
	for _, v := range o.Violations.Violations {
		if v.RuleID == ruleID {
			return true
		}
	}
	return false
}
