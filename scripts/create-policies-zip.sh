#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

#delete all the iac files (all iac files are present in 'test' directory)
find pkg/policies/opa/rego/ -type d -name test -prune -exec rm -rf {} \;
#delete the go test file
find pkg/policies/opa/rego/ -name policy_test.go -prune -exec rm -rf {} \;
#create zip of policies
cd pkg/ && zip -r policies.zip policies/opa/rego