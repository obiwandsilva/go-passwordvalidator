// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // linters aggregator
	_ "github.com/rogpeppe/godef"                           // find symbol information in Go source
	_ "golang.org/x/tools/cmd/goimports"                    // updates imports and formats code
)