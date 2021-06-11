package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/accurics/terrascan/pkg/policy"
	"github.com/patilpankaj212/terrascan-policies/pkg/utils"
)

// PolicyDirConfig represents a policy directory
type PolicyDirConfig struct {
	IacType             string
	DirPath             string
	PositiveIacFilePath string
	NegativeIacFilePath string
	RuleMetadataList    []policy.RegoMetadata
}

// NewPolicyDirConfig creates and initialises a PolicyDirConfig
func NewPolicyDirConfig(absDirPath, iacType string) (*PolicyDirConfig, error) {
	pConfig := &PolicyDirConfig{DirPath: absDirPath, IacType: iacType}

	err := pConfig.loadMetadaFiles()
	if err != nil {
		return nil, err
	}

	pConfig.loadIacFiles()

	return pConfig, nil
}

// loadMetadaFiles loads all the rego metadata files present in the directory
func (p *PolicyDirConfig) loadMetadaFiles() error {
	metadataFiles, err := utils.FindAllMetadaFilesInDir(p.DirPath)
	if err != nil {
		return err
	}

	for i := range metadataFiles {
		metadata, err := ioutil.ReadFile(metadataFiles[i])
		if err != nil {
			return err
		}

		regoMetadata := policy.RegoMetadata{}
		if err = json.Unmarshal(metadata, &regoMetadata); err != nil {
			return err
		}

		p.RuleMetadataList = append(p.RuleMetadataList, regoMetadata)
	}
	return nil
}

// loadIacFiles initialises the positive and negative iac file paths
func (p *PolicyDirConfig) loadIacFiles() {
	iacDir := filepath.Join(p.DirPath, "test")
	_, err := os.Stat(iacDir)
	if os.IsNotExist(err) {
		return
	}

	positiveIacFilePath, _ := utils.FileFileByPrefix(iacDir, "positive")
	negativeIacFilePath, _ := utils.FileFileByPrefix(iacDir, "negative")

	if positiveIacFilePath != "" {
		p.PositiveIacFilePath = filepath.Join(iacDir, positiveIacFilePath)
	}

	if negativeIacFilePath != "" {
		p.NegativeIacFilePath = filepath.Join(iacDir, negativeIacFilePath)
	}
}
