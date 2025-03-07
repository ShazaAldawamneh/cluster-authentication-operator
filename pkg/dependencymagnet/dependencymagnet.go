//go:build tools
// +build tools

// go mod won't pull in code that isn't depended upon, but we have some code we don't depend on from code that must be included
// for our build to work.
package dependencymagnet

import (
	_ "github.com/openshift/api/authorization/v1/zz_generated.crd-manifests"
	_ "github.com/openshift/api/operator/v1/zz_generated.crd-manifests"
	_ "github.com/openshift/build-machinery-go"
)
