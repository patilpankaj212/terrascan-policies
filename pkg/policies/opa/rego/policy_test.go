package rego_test

import (
	"path/filepath"
	"testing"

	"github.com/patilpankaj212/terrascan-policies/pkg/config"
	"github.com/patilpankaj212/terrascan-policies/pkg/evaluator"
	"github.com/patilpankaj212/terrascan-policies/pkg/utils"
)

var filterFolders = []string{"test"}

func TestAWSPolicies(t *testing.T) {
	awsPolicyDir, err := filepath.Abs("aws")
	if err != nil {
		t.Errorf("failed to load aws policy directory")
	}

	policyDirs, err := utils.FindAllSubDirectories(awsPolicyDir, filterFolders)
	if err != nil {
		t.Errorf("failed to policy directories in aws policy dir")
	}

	commonTest(t, policyDirs, "terraform")
}

func TestAzurePolicies(t *testing.T) {
	azurePolicyDir, err := filepath.Abs("azure")
	if err != nil {
		t.Errorf("failed to load azure policy directory")
	}

	policyDirs, err := utils.FindAllSubDirectories(azurePolicyDir, filterFolders)
	if err != nil {
		t.Errorf("failed to policy directories in azure policy dir")
	}

	commonTest(t, policyDirs, "terraform")
}

func TestGCPPolicies(t *testing.T) {
	gcpPolicyDir, err := filepath.Abs("gcp")
	if err != nil {
		t.Errorf("failed to load gcp policy directory")
	}

	policyDirs, err := utils.FindAllSubDirectories(gcpPolicyDir, filterFolders)
	if err != nil {
		t.Errorf("failed to policy directories in gcp policy dir")
	}

	commonTest(t, policyDirs, "terraform")
}

func TestK8sPolicies(t *testing.T) {
	k8sPolicyDir, err := filepath.Abs("k8s")
	if err != nil {
		t.Errorf("failed to load k8s policy directory")
	}

	policyDirs, err := utils.FindAllSubDirectories(k8sPolicyDir, filterFolders)
	if err != nil {
		t.Errorf("failed to policy directories in k8s policy dir")
	}

	commonTest(t, policyDirs, "k8s")
}

func commonTest(t *testing.T, policyDirs []string, iacType string) {

	for _, policyDir := range policyDirs {
		pConfigDir, err := config.NewPolicyDirConfig(policyDir, iacType)
		if err != nil {
			t.Errorf("failed to load policy config dir '%s'", filepath.Base(policyDir))
		}

		policyEvaluator, err := evaluator.GetEvaluator(evaluator.TerrascanEvaluator, *pConfigDir)
		if err != nil {
			t.Error("error while getting the policy evaluator")
		}

		testCaseName := filepath.Base(policyDir)
		t.Run(testCaseName, func(t *testing.T) {

			err = policyEvaluator.EvalPositive()
			if err != nil {
				t.Error(err.Error())
				return
			}

			err = policyEvaluator.EvalNegative()
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}
