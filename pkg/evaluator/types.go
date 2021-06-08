package evaluator

import (
	"fmt"

	"github.com/accurics/terrascan/pkg/runtime"
	"github.com/patilpankaj212/terrascan-policies/pkg/config"
)

// TerrascanPolicyEvaluator is a policy evaluator and uses terrascan packages
type TerrascanPolicyEvaluator struct {
	config.PolicyDirConfig
}

// NewTerrascanPolicyEvaluator creates a new TerrascanPolicyEvaluator
func NewTerrascanPolicyEvaluator(c config.PolicyDirConfig) *TerrascanPolicyEvaluator {
	return &TerrascanPolicyEvaluator{c}
}

// EvalPositive implementation for TerrascanPolicyEvaluator
func (t *TerrascanPolicyEvaluator) EvalPositive() error {

	// return an error if the positive iac file is not present in the policy directory
	if t.PositiveIacFilePath == "" {
		return fmt.Errorf("positive iac file should be present for evaluation")
	}

	output, err := t.execute(t.PositiveIacFilePath)
	if err != nil {
		return err
	}

	// the positive iac file should violate the policies present in the policy directory
	if len(output.Violations.Violations) < 1 {
		return fmt.Errorf("violations should be present for the positive iac file: '%s'", t.PositiveIacFilePath)
	}

	// verify that all rules present in the policy directory are violated
	for _, metadata := range t.RuleMetadataList {
		if isPresent := t.isRuleIDPresent(metadata.ReferenceID, output); !isPresent {
			return fmt.Errorf("ruleID '%s' should be present in the violations", metadata.ReferenceID)
		}
	}
	return nil
}

// EvalNegative implementation for TerrascanPolicyEvaluator
func (t *TerrascanPolicyEvaluator) EvalNegative() error {

	// return an error if the negative iac file is not present in the policy directory
	if t.NegativeIacFilePath == "" {
		return fmt.Errorf("negative iac file should be present for evaluation")
	}

	output, err := t.execute(t.NegativeIacFilePath)
	if err != nil {
		return err
	}

	// the negative iac file shouldn't violate the policies present in the policy directory
	if len(output.Violations.Violations) > 0 {
		return fmt.Errorf("no violations should be present for negative iac file: '%s'", t.NegativeIacFilePath)
	}

	return nil
}

// execute creates a terrascan executor and returns the violation result
func (t *TerrascanPolicyEvaluator) execute(iacFilePath string) (*runtime.Output, error) {
	var output runtime.Output
	executor, err := runtime.NewExecutor(t.IacType, "", []string{}, iacFilePath, "", []string{t.DirPath}, nil, nil, nil, "", false)
	if err != nil {
		return nil, err
	}

	output, err = executor.Execute()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// isRuleIDPresent checks if a rule id is present in the violations
func (t *TerrascanPolicyEvaluator) isRuleIDPresent(ruleID string, o *runtime.Output) bool {
	for _, v := range o.Violations.Violations {
		if v.RuleID == ruleID {
			return true
		}
	}
	return false
}
