// +build tools
// https://github.com/onsi/ginkgo/tree/v2#go-module-tools-package
package tools

import (
	_ "github.com/onsi/ginkgo/ginkgo"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.
