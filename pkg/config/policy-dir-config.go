package config

import (
	"github.com/accurics/terrascan/pkg/policy/opa"
)

type PolicyDirConfig struct {
	DirPath             string
	PositiveIacFilePath string
	NegativeIacFilePath string
	RuleMetadataList    []opa.RegoMetadata
}

func NewPolicyDirConfig(absDirPath string) PolicyDirConfig {
	return PolicyDirConfig{}
}
